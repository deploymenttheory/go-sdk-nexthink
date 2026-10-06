// Package config validates credentials for each enabled Nexthink API family.
package config

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth/chrome"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth/password"
)

type ClientCredentials struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
}

// UsernamePasswordCredentials configures SDK-managed headless local-account login.
type UsernamePasswordCredentials = password.Config

// BrowserCredentials accepts exactly one access token, token provider, or local login.
// Chrome is opt-in through auth/chrome.New or NEXTHINK_WEB_AUTH=chrome.
type BrowserCredentials struct {
	AccessToken      string `json:"-"`
	ExpiresAt        time.Time
	TokenProvider    auth.TokenProvider           `json:"-"`
	UsernamePassword *UsernamePasswordCredentials `json:"-"`
}

// AuthConfig enables only the API families whose credentials are supplied.
type AuthConfig struct {
	Instance  string
	Region    string
	PublicAPI *ClientCredentials
	WebAPI    *BrowserCredentials
}

// String and GoString prevent accidental credential disclosure in formatted diagnostics.
func (ClientCredentials) String() string      { return "ClientCredentials{credentials:redacted}" }
func (c ClientCredentials) GoString() string  { return c.String() }
func (BrowserCredentials) String() string     { return "BrowserCredentials{credentials:redacted}" }
func (c BrowserCredentials) GoString() string { return c.String() }
func (AuthConfig) String() string             { return "AuthConfig{credentials:redacted}" }
func (c AuthConfig) GoString() string         { return c.String() }

var instanceName = regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,98}[a-zA-Z0-9])?$`)

func (c *AuthConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("authentication configuration is required")
	}
	if !instanceName.MatchString(c.Instance) {
		return fmt.Errorf("instance must be a valid tenant name, without a URL or domain")
	}
	switch c.Region {
	case "us", "eu", "pac", "meta":
	default:
		return fmt.Errorf("region must be us, eu, pac or meta")
	}
	if c.PublicAPI == nil && c.WebAPI == nil {
		return fmt.Errorf("configure PublicAPI credentials, WebAPI credentials, or both")
	}
	if p := c.PublicAPI; p != nil {
		if strings.TrimSpace(p.ClientID) == "" || strings.TrimSpace(p.ClientSecret) == "" {
			return fmt.Errorf("PublicAPI requires client ID and client secret")
		}
	}
	if w := c.WebAPI; w != nil {
		hasToken := strings.TrimSpace(w.AccessToken) != ""
		hasProvider := w.TokenProvider != nil
		if hasProvider {
			v := reflect.ValueOf(w.TokenProvider)
			switch v.Kind() {
			case reflect.Pointer,
				reflect.Func,
				reflect.Map,
				reflect.Slice,
				reflect.Interface,
				reflect.Chan:
				if v.IsNil() {
					return fmt.Errorf("WebAPI token provider cannot be nil")
				}
			}
		}
		methods := 0
		for _, configured := range []bool{hasToken, hasProvider, w.UsernamePassword != nil} {
			if configured {
				methods++
			}
		}
		if methods != 1 {
			return fmt.Errorf("WebAPI requires exactly one access token, token provider or username/password configuration")
		}
		if w.UsernamePassword != nil {
			if err := w.UsernamePassword.Validate(); err != nil {
				return fmt.Errorf("WebAPI password authentication: %w", err)
			}
		}
		if hasToken {
			if strings.ContainsAny(w.AccessToken, " \t\r\n") {
				return fmt.Errorf("WebAPI access token must not contain whitespace")
			}
			if err := (auth.Token{Value: w.AccessToken, ExpiresAt: w.ExpiresAt}).Validate(); err != nil {
				return fmt.Errorf("WebAPI token: %w", err)
			}
		}
		if !hasToken && !w.ExpiresAt.IsZero() {
			return fmt.Errorf("WebAPI token expiry must be managed by its provider")
		}
	}
	return nil
}

// AuthConfigFromEnv reads NEXTHINK_API=public|web|both (default public).
// Public uses NEXTHINK_CLIENT_ID/SECRET. Web uses NEXTHINK_WEB_AUTH=token|chrome|password
// (default token) and NEXTHINK_ACCESS_TOKEN for token authentication.
func AuthConfigFromEnv() (*AuthConfig, error) {
	c := &AuthConfig{Instance: os.Getenv("NEXTHINK_INSTANCE"), Region: os.Getenv("NEXTHINK_REGION")}
	api := os.Getenv("NEXTHINK_API")
	if api == "" {
		api = "public"
	}
	switch api {
	case "public", "both":
		c.PublicAPI = &ClientCredentials{
			ClientID:     os.Getenv("NEXTHINK_CLIENT_ID"),
			ClientSecret: os.Getenv("NEXTHINK_CLIENT_SECRET"),
		}
	case "web":
	default:
		return nil, fmt.Errorf("NEXTHINK_API must be public, web or both")
	}
	if api == "web" || api == "both" {
		c.WebAPI = &BrowserCredentials{}
		switch method := os.Getenv("NEXTHINK_WEB_AUTH"); method {
		case "", "token":
			c.WebAPI.AccessToken = os.Getenv("NEXTHINK_ACCESS_TOKEN")
		case "password":
			credentials := &UsernamePasswordCredentials{
				Username:              os.Getenv("NEXTHINK_USERNAME"),
				Password:              os.Getenv("NEXTHINK_PASSWORD"),
				LoginTimeout:          90 * time.Second,
				BrowserProxy:          os.Getenv("NEXTHINK_BROWSER_PROXY"),
				BrowserExecutablePath: os.Getenv("NEXTHINK_BROWSER_EXECUTABLE_PATH"),
			}
			if raw := os.Getenv("NEXTHINK_LOGIN_TIMEOUT"); raw != "" {
				duration, err := time.ParseDuration(raw)
				if err != nil || duration <= 0 {
					return nil, fmt.Errorf("NEXTHINK_LOGIN_TIMEOUT must be a positive Go duration")
				}
				credentials.LoginTimeout = duration
			}
			c.WebAPI.UsernamePassword = credentials
		case "chrome":
			p, err := chrome.New(fmt.Sprintf("https://%s.%s.nexthink.cloud", c.Instance, c.Region))
			if err != nil {
				return nil, fmt.Errorf("Chrome provider: %w", err)
			}
			c.WebAPI.TokenProvider = p
		default:
			return nil, fmt.Errorf("NEXTHINK_WEB_AUTH must be token, chrome or password")
		}
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
