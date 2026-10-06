// Package password obtains Nexthink user tokens through an isolated headless
// browser. Browser installation is an explicit deployment step.
package password

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
)

const DefaultLoginTimeout = 90 * time.Second

var (
	ErrBrowserUnavailable     = errors.New("headless browser runtime unavailable; install the pinned Playwright driver and Chromium")
	ErrInvalidCredentials     = errors.New("Nexthink rejected the username or password")
	ErrMFARequired            = errors.New("Nexthink requires MFA; password-only authentication cannot continue")
	ErrSSOUnsupported         = errors.New("corporate SSO is not supported by password authentication")
	ErrPasswordChangeRequired = errors.New("Nexthink requires a password change")
	ErrAccountLocked          = errors.New("Nexthink account is locked")
	ErrLoginFailed            = errors.New("Nexthink local login failed")
	ErrRenewalFailed          = errors.New("Nexthink token renewal failed")
	ErrClosed                 = errors.New("Nexthink password provider is closed")
	errRefreshRejected        = errors.New("refresh credential rejected")
)

// Config contains local-account credentials. Authentication state is never
// persisted. BrowserProxy applies to browser login and direct token renewal;
// Chromium uses the runner's trust store, not Go transport TLS callbacks.
type Config struct {
	Username              string        `json:"-"`
	Password              string        `json:"-"`
	LoginTimeout          time.Duration `json:"-"`
	BrowserProxy          string        `json:"-"`
	BrowserExecutablePath string        `json:"-"`
}

func (Config) String() string   { return "PasswordConfig{credentials:redacted}" }
func (Config) GoString() string { return "PasswordConfig{credentials:redacted}" }

// Validate checks configuration without launching or installing a browser.
func (c Config) Validate() error {
	if strings.TrimSpace(c.Username) == "" || c.Password == "" {
		return errors.New("password authentication requires a username and password")
	}
	if c.LoginTimeout < 0 {
		return errors.New("login timeout must not be negative")
	}
	if c.BrowserProxy != "" {
		u, err := url.Parse(c.BrowserProxy)
		if err != nil || u.Hostname() == "" || (u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "socks5") || u.RawQuery != "" || u.Fragment != "" || (u.Path != "" && u.Path != "/") {
			return errors.New("browser proxy must be an HTTP, HTTPS or SOCKS5 proxy URL")
		}
	}
	return nil
}

type session struct {
	AccessToken, RefreshToken, ClientID, TokenURL string
	ExpiresAt                                     time.Time
}

type attempt struct {
	done chan struct{}
	err  error
}

// Provider implements auth.TokenProvider. Concurrent callers share acquisition;
// each waiter can cancel independently. Close cancels and joins active work.
type Provider struct {
	mu          sync.Mutex
	origin      string
	cfg         Config
	current     session
	renewBefore time.Time
	active      *attempt
	terminal    error
	closed      bool
	lifetime    context.Context
	cancel      context.CancelFunc
	login       func(context.Context, string, Config) (session, error)
	httpClient  *http.Client
}

var tenantHost = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]*\.(us|eu|pac|meta)\.nexthink\.cloud$`)

// New constructs a lazy provider for an exact Nexthink tenant origin.
func New(origin string, cfg Config) (*Provider, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	u, err := url.Parse(origin)
	if err != nil || u.Scheme != "https" || u.User != nil || u.Port() != "" || !tenantHost.MatchString(u.Hostname()) || (u.Path != "" && u.Path != "/") || u.RawQuery != "" || u.Fragment != "" {
		return nil, errors.New("password authentication requires an HTTPS Nexthink tenant origin")
	}
	if cfg.LoginTimeout == 0 {
		cfg.LoginTimeout = DefaultLoginTimeout
	}
	cfg.Username = strings.TrimSpace(cfg.Username)
	transport := defaultRefreshTransport()
	if cfg.BrowserProxy != "" {
		proxyURL, _ := url.Parse(cfg.BrowserProxy)
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Provider{
		origin: "https://" + strings.ToLower(u.Host), cfg: cfg,
		lifetime: ctx, cancel: cancel, login: browserLogin,
		httpClient: &http.Client{Transport: transport, Timeout: 30 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }},
	}, nil
}

// defaultRefreshTransport owns its connection pool even when an application has
// replaced http.DefaultTransport with an instrumentation wrapper.
func defaultRefreshTransport() *http.Transport {
	if transport, ok := http.DefaultTransport.(*http.Transport); ok && transport != nil {
		return transport.Clone()
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: time.Second,
	}
}

func (*Provider) String() string   { return "PasswordProvider{credentials:redacted}" }
func (*Provider) GoString() string { return "PasswordProvider{credentials:redacted}" }

func (p *Provider) Token(ctx context.Context) (auth.Token, error) {
	p.mu.Lock()
	timeout := p.cfg.LoginTimeout
	p.mu.Unlock()
	// Bound time spent waiting for another caller as well as acquisition itself.
	ctx, cancelCall := context.WithTimeout(ctx, timeout)
	defer cancelCall()
	for {
		if err := ctx.Err(); err != nil {
			return auth.Token{}, err
		}
		p.mu.Lock()
		if p.closed {
			p.mu.Unlock()
			return auth.Token{}, ErrClosed
		}
		if p.terminal != nil {
			err := p.terminal
			p.mu.Unlock()
			return auth.Token{}, err
		}
		if p.current.AccessToken != "" && time.Now().Before(p.renewBefore) {
			token := auth.Token{Value: p.current.AccessToken, ExpiresAt: p.current.ExpiresAt}
			p.mu.Unlock()
			return token, nil
		}
		if a := p.active; a != nil {
			p.mu.Unlock()
			select {
			case <-ctx.Done():
				return auth.Token{}, ctx.Err()
			case <-a.done:
				if a.err != nil {
					return auth.Token{}, a.err
				}
				continue
			}
		}
		a := &attempt{done: make(chan struct{})}
		p.active = a
		previous, cfg := p.current, p.cfg
		p.mu.Unlock()

		work, cancel := context.WithCancel(ctx)
		stop := context.AfterFunc(p.lifetime, cancel)
		next, err := p.acquire(work, previous, cfg)
		// Preserve a successfully received rotation even if the initiating caller
		// was canceled while the browser or response body was being cleaned up.
		if err != nil && work.Err() != nil && !permanent(err) {
			err = work.Err()
		}
		stop()
		cancel()

		p.mu.Lock()
		if p.closed {
			err = ErrClosed
		}
		if err == nil {
			p.current = next
			buffer := min(30*time.Second, time.Until(next.ExpiresAt)/10)
			p.renewBefore = next.ExpiresAt.Add(-buffer)
		} else if permanent(err) {
			p.terminal = err
		}
		a.err = err
		p.active = nil
		close(a.done)
		p.mu.Unlock()
		if err != nil {
			return auth.Token{}, err
		}
	}
}

func permanent(err error) bool {
	return errors.Is(err, ErrInvalidCredentials) || errors.Is(err, ErrAccountLocked) || errors.Is(err, ErrMFARequired) || errors.Is(err, ErrSSOUnsupported) || errors.Is(err, ErrPasswordChangeRequired)
}

func (p *Provider) acquire(ctx context.Context, previous session, cfg Config) (session, error) {
	if previous.RefreshToken != "" {
		next, err := p.refresh(ctx, previous)
		if err == nil {
			return next, validateSession(next)
		}
		if !errors.Is(err, errRefreshRejected) {
			return session{}, err
		}
	}
	next, err := p.login(ctx, p.origin, cfg)
	if err != nil {
		return session{}, err
	}
	if err := validateSession(next); err != nil {
		return session{}, err
	}
	if next.RefreshToken != "" && (next.TokenURL != tokenURL(p.origin) || next.ClientID == "" || strings.ContainsAny(next.ClientID, "\r\n")) {
		return session{}, ErrLoginFailed
	}
	return next, nil
}

func validateSession(s session) error {
	if s.AccessToken == "" || strings.ContainsAny(s.AccessToken, " \t\r\n") || !time.Now().Before(s.ExpiresAt) {
		return ErrLoginFailed
	}
	return nil
}

func tokenURL(origin string) string {
	u, _ := url.Parse(origin)
	parts := strings.SplitN(u.Host, ".", 2)
	return fmt.Sprintf("https://%s-login.%s/oauth2/default/v1/token", parts[0], parts[1])
}

// InvalidateToken preserves the refresh credential but discards the cached access token.
func (p *Provider) InvalidateToken() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.current.AccessToken = ""
	p.renewBefore = time.Time{}
}

// InvalidateTokenIfCurrent ignores a late 401 for an already replaced token.
func (p *Provider) InvalidateTokenIfCurrent(value string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if value != "" && value == p.current.AccessToken {
		p.current.AccessToken = ""
		p.renewBefore = time.Time{}
	}
}

func (p *Provider) Close() error {
	p.mu.Lock()
	p.closed = true
	p.cancel()
	a := p.active
	p.current = session{}
	p.cfg.Password = ""
	p.mu.Unlock()
	if a != nil {
		<-a.done
	}
	p.httpClient.CloseIdleConnections()
	return nil
}

var _ auth.TokenProvider = (*Provider)(nil)
