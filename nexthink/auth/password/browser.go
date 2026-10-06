package password

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mxschmitt/playwright-go"
)

// Playwright's Run configuration updates a package-level logger read by active
// event handlers. Serialize browser acquisition across SDK providers; token
// refresh and API requests remain concurrent.
var browserGate = make(chan struct{}, 1)

func acquireBrowser(ctx context.Context) error {
	select {
	case browserGate <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func browserLogin(ctx context.Context, origin string, cfg Config) (session, error) {
	login, err := loginOrigin(origin)
	if err != nil {
		return session{}, err
	}
	return browserLoginAt(ctx, origin, login, cfg)
}

func loginOrigin(origin string) (string, error) {
	u, err := url.Parse(origin)
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
		return "", ErrLoginFailed
	}
	parts := strings.Split(u.Hostname(), ".")
	if len(parts) < 3 || !strings.HasSuffix(u.Hostname(), ".nexthink.cloud") {
		return "", ErrLoginFailed
	}
	parts[0] += "-login"
	return "https://" + strings.Join(parts, "."), nil
}

func sameOrigin(rawURL, origin string) bool {
	u, err := url.Parse(rawURL)
	o, originErr := url.Parse(origin)
	return err == nil && originErr == nil && u.User == nil && u.Scheme == o.Scheme && strings.EqualFold(u.Host, o.Host)
}

// browserLoginAt is private so only controlled tests can substitute a local origin.
func browserLoginAt(ctx context.Context, origin, login string, cfg Config) (session, error) {
	if err := ctx.Err(); err != nil {
		return session{}, err
	}
	if os.Getenv("DEBUGP") != "" || os.Getenv("PWDEBUG") != "" {
		return session{}, fmt.Errorf("%w: disable DEBUGP and PWDEBUG for password authentication", ErrBrowserUnavailable)
	}
	if err := acquireBrowser(ctx); err != nil {
		return session{}, err
	}
	defer func() { <-browserGate }()
	pw, err := playwright.Run(&playwright.RunOptions{Stdout: io.Discard, Stderr: io.Discard, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))})
	if err != nil {
		return session{}, ErrBrowserUnavailable
	}
	var stopOnce sync.Once
	stopDriver := func() { stopOnce.Do(func() { _ = pw.Stop() }) }
	defer stopDriver()
	stopCancellation := context.AfterFunc(ctx, stopDriver)
	defer stopCancellation()
	options := playwright.BrowserTypeLaunchOptions{Headless: playwright.Bool(true), Timeout: playwright.Float(browserTimeout(ctx))}
	if cfg.BrowserExecutablePath != "" {
		options.ExecutablePath = playwright.String(cfg.BrowserExecutablePath)
	}
	if cfg.BrowserProxy != "" {
		proxyURL, parseErr := url.Parse(cfg.BrowserProxy)
		if parseErr != nil {
			return session{}, ErrLoginFailed
		}
		options.Proxy = &playwright.Proxy{}
		if proxyURL.User != nil {
			options.Proxy.Username = playwright.String(proxyURL.User.Username())
			if password, exists := proxyURL.User.Password(); exists {
				options.Proxy.Password = playwright.String(password)
			}
			proxyURL.User = nil
		}
		options.Proxy.Server = proxyURL.String()
	}
	browser, err := pw.Chromium.Launch(options)
	if err != nil {
		return session{}, browserError(ctx, ErrBrowserUnavailable)
	}
	defer func() { _ = browser.Close() }()
	bc, err := browser.NewContext(playwright.BrowserNewContextOptions{AcceptDownloads: playwright.Bool(false), ServiceWorkers: playwright.ServiceWorkerPolicyBlock})
	if err != nil {
		return session{}, browserError(ctx, ErrLoginFailed)
	}
	defer func() { _ = bc.Close() }()
	bc.SetDefaultTimeout(2000)
	blocked := make(chan error, 1)
	report := func(err error) {
		select {
		case blocked <- err:
		default:
		}
	}
	// Block foreign navigation and submissions before the request leaves the browser.
	// Static resources may use Nexthink's CDN but never receive a form submission.
	if err := bc.Route("**/*", func(route playwright.Route) {
		r := route.Request()
		allowed := sameOrigin(r.URL(), origin) || sameOrigin(r.URL(), login)
		if !allowed && (r.IsNavigationRequest() || r.Method() != "GET") {
			if r.IsNavigationRequest() {
				frame := r.Frame()
				if frame == nil || frame.ParentFrame() == nil {
					report(ErrSSOUnsupported)
				}
			}
			_ = route.Abort()
			return
		}
		_ = route.Continue()
	}); err != nil {
		return session{}, ErrLoginFailed
	}
	page, err := bc.NewPage()
	if err != nil {
		return session{}, ErrLoginFailed
	}
	tokens := make(chan session, 1)
	page.OnResponse(func(response playwright.Response) {
		u, parseErr := url.Parse(response.URL())
		if parseErr != nil || !sameOrigin(response.URL(), login) || u.Path != "/oauth2/default/v1/token" || response.Request().Method() != "POST" || response.Status() != 200 {
			return
		}
		body, bodyErr := response.Body()
		form, formErr := response.Request().PostData()
		if bodyErr != nil || formErr != nil {
			return
		}
		s, decodeErr := browserToken(body, form, login+"/oauth2/default/v1/token", time.Now())
		if decodeErr != nil {
			report(decodeErr)
			return
		}
		select {
		case tokens <- s:
		default:
		}
	})
	_, err = page.Goto(strings.TrimRight(origin, "/")+"/login", playwright.PageGotoOptions{WaitUntil: playwright.WaitUntilStateDomcontentloaded, Timeout: playwright.Float(browserTimeout(ctx))})
	if err != nil {
		select {
		case blockedErr := <-blocked:
			return session{}, blockedErr
		default:
			return session{}, browserError(ctx, ErrLoginFailed)
		}
	}
	return driveLogin(ctx, page, cfg, origin, login, tokens, blocked)
}

func browserTimeout(ctx context.Context) float64 {
	if deadline, ok := ctx.Deadline(); ok {
		return max(1, float64(time.Until(deadline).Milliseconds()))
	}
	return 90000
}

func browserError(ctx context.Context, fallback error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return fallback
}

func browserToken(body []byte, form, tokenURL string, now time.Time) (session, error) {
	var token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int64  `json:"expires_in"`
	}
	values, err := url.ParseQuery(form)
	if err != nil || json.Unmarshal(body, &token) != nil || token.AccessToken == "" || token.ExpiresIn <= 0 || token.ExpiresIn > 31536000 || !strings.EqualFold(token.TokenType, "Bearer") || strings.ContainsAny(token.AccessToken, "\r\n") {
		return session{}, ErrLoginFailed
	}
	return session{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken, ClientID: values.Get("client_id"), TokenURL: tokenURL, ExpiresAt: now.Add(time.Duration(token.ExpiresIn) * time.Second)}, nil
}

func driveLogin(ctx context.Context, page playwright.Page, cfg Config, origin, login string, tokens <-chan session, blocked <-chan error) (session, error) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	usernameSubmitted, passwordSubmitted := false, false
	for {
		select {
		case <-ctx.Done():
			return session{}, ctx.Err()
		case err := <-blocked:
			return session{}, err
		case s := <-tokens:
			return s, nil
		case <-ticker.C:
		}
		if !sameOrigin(page.URL(), origin) && !sameOrigin(page.URL(), login) {
			return session{}, ErrSSOUnsupported
		}
		// Inspect challenge headings and alerts, not arbitrary help/footer text.
		text, _ := page.Locator("h1:visible, h2:visible, [role=alert]:visible, .o-form-error-container:visible, .infobox-error:visible").AllTextContents()
		if challenge := loginChallenge(strings.Join(text, " ")); challenge != nil {
			return session{}, challenge
		}
		mfa := page.Locator(`input[autocomplete="one-time-code"], input[name="passCode"], input[name="otp"]`).First()
		if visible, _ := mfa.IsVisible(); visible {
			return session{}, ErrMFARequired
		}
		user := page.Locator(`input[name="username"]:visible, input[name="identifier"]:visible, input[type="email"]:visible, input[autocomplete="username"]:visible, #okta-signin-username:visible`).First()
		pass := page.Locator(`input[type="password"]:visible`).First()
		userVisible, _ := user.IsVisible()
		passVisible, _ := pass.IsVisible()
		if passwordSubmitted || (!userVisible && !passVisible) {
			continue
		}
		if passVisible {
			if count, _ := page.Locator(`input[type="password"]:visible`).Count(); count > 1 {
				return session{}, ErrPasswordChangeRequired
			}
			if userVisible && !usernameSubmitted {
				if err := user.Fill(cfg.Username); err != nil {
					return session{}, browserError(ctx, ErrLoginFailed)
				}
			}
			if err := pass.Fill(cfg.Password); err != nil {
				return session{}, browserError(ctx, ErrLoginFailed)
			}
			passwordSubmitted = true
		} else {
			if usernameSubmitted {
				continue
			}
			if err := user.Fill(cfg.Username); err != nil {
				return session{}, browserError(ctx, ErrLoginFailed)
			}
			usernameSubmitted = true
		}
		button := page.Locator(`button[type="submit"]:visible, input[type="submit"]:visible, button:has-text("Sign in"):visible, button:has-text("Next"):visible`).First()
		if err := button.Click(); err != nil {
			return session{}, browserError(ctx, ErrLoginFailed)
		}
	}
}

func loginChallenge(text string) error {
	text = strings.ToLower(text)
	for _, rule := range []struct {
		phrases []string
		err     error
	}{
		{[]string{"account is locked", "account locked", "too many attempts"}, ErrAccountLocked},
		{[]string{"change your password", "password expired", "reset your password", "set a new password"}, ErrPasswordChangeRequired},
		{[]string{"verify your identity", "verification code", "multi-factor", "multifactor", "authenticator", "security code", "approve the request"}, ErrMFARequired},
		{[]string{"unable to sign in", "invalid credentials", "incorrect password", "invalid username", "authentication failed", "sign in failed"}, ErrInvalidCredentials},
	} {
		for _, phrase := range rule.phrases {
			if strings.Contains(text, phrase) {
				return rule.err
			}
		}
	}
	return nil
}
