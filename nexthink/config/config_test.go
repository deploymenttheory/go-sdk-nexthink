package config

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
)

func validConfig() *AuthConfig {
	return &AuthConfig{
		Instance:  "fixture-tenant",
		Region:    "eu",
		PublicAPI: &ClientCredentials{ClientID: "fixture-id", ClientSecret: "fixture-secret"},
	}
}

func TestValidateFamilies(t *testing.T) {
	for _, family := range []string{"public", "web", "both"} {
		t.Run(family, func(t *testing.T) {
			cfg := validConfig()
			if family != "public" {
				cfg.WebAPI = &BrowserCredentials{AccessToken: "fixture-access-token"}
			}
			if family == "web" {
				cfg.PublicAPI = nil
			}
			require.NoError(t, cfg.Validate())
		})
	}
}

func TestValidateInvalidConfigurations(t *testing.T) {
	var nilProvider auth.TokenProviderFunc
	cases := map[string]func(*AuthConfig){
		"missing instance":      func(c *AuthConfig) { c.Instance = "" },
		"URL instance":          func(c *AuthConfig) { c.Instance = "https://example.test" },
		"domain instance":       func(c *AuthConfig) { c.Instance = "example.test" },
		"path instance":         func(c *AuthConfig) { c.Instance = "../other" },
		"invalid region":        func(c *AuthConfig) { c.Region = "invalid" },
		"no family":             func(c *AuthConfig) { c.PublicAPI = nil },
		"missing client ID":     func(c *AuthConfig) { c.PublicAPI.ClientID = " " },
		"missing client secret": func(c *AuthConfig) { c.PublicAPI.ClientSecret = "" },
		"empty web":             func(c *AuthConfig) { c.WebAPI = &BrowserCredentials{} },
		"whitespace token":      func(c *AuthConfig) { c.WebAPI = &BrowserCredentials{AccessToken: "token\nheader"} },
		"ambiguous web": func(c *AuthConfig) {
			c.WebAPI = &BrowserCredentials{
				AccessToken:   "token",
				TokenProvider: auth.StaticToken("token", time.Time{}),
			}
		},
		"typed nil provider": func(c *AuthConfig) { c.WebAPI = &BrowserCredentials{TokenProvider: nilProvider} },
		"provider expiry": func(c *AuthConfig) {
			c.WebAPI = &BrowserCredentials{
				TokenProvider: auth.StaticToken("token", time.Time{}),
				ExpiresAt:     time.Now().Add(time.Hour),
			}
		},
	}
	for name, mutate := range cases {
		t.Run(
			name,
			func(t *testing.T) { cfg := validConfig(); mutate(cfg); require.Error(t, cfg.Validate()) },
		)
	}
	var cfg *AuthConfig
	require.Error(t, cfg.Validate())
	cfg = validConfig()
	cfg.WebAPI = &BrowserCredentials{
		AccessToken: "expired",
		ExpiresAt:   time.Now().Add(-time.Second),
	}
	require.ErrorIs(t, cfg.Validate(), auth.ErrSessionExpired)
	cfg.WebAPI = &BrowserCredentials{
		TokenProvider: auth.TokenProviderFunc(func(context.Context) (auth.Token, error) {
			t.Fatal("configuration validation must not read a browser token")
			return auth.Token{}, nil
		}),
	}
	require.NoError(t, cfg.Validate())
}

func TestConfigFromEnvironment(t *testing.T) {
	for _, mode := range []string{"", "public", "web", "both"} {
		t.Run(mode, func(t *testing.T) {
			t.Setenv("NEXTHINK_INSTANCE", "fixture")
			t.Setenv("NEXTHINK_REGION", "eu")
			t.Setenv("NEXTHINK_API", mode)
			t.Setenv("NEXTHINK_WEB_AUTH", "token")
			t.Setenv("NEXTHINK_CLIENT_ID", "id")
			t.Setenv("NEXTHINK_CLIENT_SECRET", "secret")
			t.Setenv("NEXTHINK_ACCESS_TOKEN", "fixture-token")
			cfg, err := AuthConfigFromEnv()
			require.NoError(t, err)
			assert.Equal(t, mode != "web", cfg.PublicAPI != nil)
			assert.Equal(t, mode == "web" || mode == "both", cfg.WebAPI != nil)
		})
	}
}

func TestEnvironmentValidatesOnlySelectedFamily(t *testing.T) {
	t.Setenv("NEXTHINK_INSTANCE", "fixture")
	t.Setenv("NEXTHINK_REGION", "eu")
	t.Setenv("NEXTHINK_API", "web")
	t.Setenv("NEXTHINK_WEB_AUTH", "token")
	t.Setenv("NEXTHINK_CLIENT_ID", "")
	t.Setenv("NEXTHINK_CLIENT_SECRET", "")
	t.Setenv("NEXTHINK_ACCESS_TOKEN", "fixture-token")
	_, err := AuthConfigFromEnv()
	require.NoError(t, err)
	t.Setenv("NEXTHINK_API", "both")
	_, err = AuthConfigFromEnv()
	require.ErrorContains(t, err, "PublicAPI")
	t.Setenv("NEXTHINK_API", "unknown")
	_, err = AuthConfigFromEnv()
	require.ErrorContains(t, err, "NEXTHINK_API")
	t.Setenv("NEXTHINK_API", "web")
	t.Setenv("NEXTHINK_WEB_AUTH", "unknown")
	_, err = AuthConfigFromEnv()
	require.ErrorContains(t, err, "NEXTHINK_WEB_AUTH")
	t.Setenv("NEXTHINK_WEB_AUTH", "chrome")
	cfg, err := AuthConfigFromEnv()
	require.NoError(t, err)
	require.NotNil(t, cfg.WebAPI.TokenProvider)
}
