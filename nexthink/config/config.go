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
)

type ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

// BrowserCredentials accepts exactly one caller-managed access token or token provider.
// Chrome is opt-in through auth/chrome.New or NEXTHINK_WEB_AUTH=chrome.
type BrowserCredentials struct {
	AccessToken   string
	ExpiresAt     time.Time
	TokenProvider auth.TokenProvider
}

// AuthConfig enables only the API families whose credentials are supplied.
type AuthConfig struct {
	Instance  string
	Region    string
	PublicAPI *ClientCredentials
	WebAPI    *BrowserCredentials
}

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
		if hasToken == hasProvider {
			return fmt.Errorf("WebAPI requires exactly one access token or token provider")
		}
		if hasToken {
			if strings.ContainsAny(w.AccessToken, " \t\r\n") {
				return fmt.Errorf("WebAPI access token must not contain whitespace")
			}
			if err := (auth.Token{Value: w.AccessToken, ExpiresAt: w.ExpiresAt}).Validate(); err != nil {
				return fmt.Errorf("WebAPI token: %w", err)
			}
		}
		if hasProvider && !w.ExpiresAt.IsZero() {
			return fmt.Errorf("WebAPI token expiry must be managed by its provider")
		}
	}
	return nil
}

// AuthConfigFromEnv reads NEXTHINK_API=public|web|both (default public).
// Public uses NEXTHINK_CLIENT_ID/SECRET. Web uses NEXTHINK_WEB_AUTH=token|chrome
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
		case "chrome":
			p, err := chrome.New(fmt.Sprintf("https://%s.%s.nexthink.cloud", c.Instance, c.Region))
			if err != nil {
				return nil, fmt.Errorf("Chrome provider: %w", err)
			}
			c.WebAPI.TokenProvider = p
		default:
			return nil, fmt.Errorf("NEXTHINK_WEB_AUTH must be token or chrome")
		}
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}
