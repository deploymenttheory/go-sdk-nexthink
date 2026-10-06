package config

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPasswordConfiguration(t *testing.T) {
	for name, mutate := range map[string]func(*BrowserCredentials){
		"valid":                 func(*BrowserCredentials) {},
		"missing username":      func(w *BrowserCredentials) { w.UsernamePassword.Username = " " },
		"missing password":      func(w *BrowserCredentials) { w.UsernamePassword.Password = "" },
		"negative timeout":      func(w *BrowserCredentials) { w.UsernamePassword.LoginTimeout = -time.Second },
		"token and password":    func(w *BrowserCredentials) { w.AccessToken = "fixture-token" },
		"provider and password": func(w *BrowserCredentials) { w.TokenProvider = auth.StaticToken("fixture-token", time.Time{}) },
		"password expiry":       func(w *BrowserCredentials) { w.ExpiresAt = time.Now().Add(time.Hour) },
		"bad proxy":             func(w *BrowserCredentials) { w.UsernamePassword.BrowserProxy = "ftp://fixture.invalid" },
	} {
		t.Run(name, func(t *testing.T) {
			cfg := &AuthConfig{Instance: "fixture", Region: "eu", WebAPI: &BrowserCredentials{UsernamePassword: &UsernamePasswordCredentials{Username: "local@example.invalid", Password: "fixture-password"}}}
			mutate(cfg.WebAPI)
			if name == "valid" {
				require.NoError(t, cfg.Validate())
			} else {
				require.Error(t, cfg.Validate())
			}
		})
	}
}

func TestPasswordFromEnvironment(t *testing.T) {
	t.Setenv("NEXTHINK_API", "web")
	t.Setenv("NEXTHINK_INSTANCE", "fixture")
	t.Setenv("NEXTHINK_REGION", "eu")
	t.Setenv("NEXTHINK_WEB_AUTH", "password")
	t.Setenv("NEXTHINK_USERNAME", "local@example.invalid")
	t.Setenv("NEXTHINK_PASSWORD", "fixture-password")
	t.Setenv("NEXTHINK_ACCESS_TOKEN", "unselected-token")
	t.Setenv("NEXTHINK_LOGIN_TIMEOUT", "")
	t.Setenv("NEXTHINK_BROWSER_PROXY", "http://proxy.example.invalid:8080")
	t.Setenv("NEXTHINK_BROWSER_EXECUTABLE_PATH", "/fixture/chromium")
	cfg, err := AuthConfigFromEnv()
	require.NoError(t, err)
	assert.Equal(t, 90*time.Second, cfg.WebAPI.UsernamePassword.LoginTimeout)
	assert.Equal(t, "local@example.invalid", cfg.WebAPI.UsernamePassword.Username)
	assert.Equal(t, "fixture-password", cfg.WebAPI.UsernamePassword.Password)
	assert.Equal(t, "http://proxy.example.invalid:8080", cfg.WebAPI.UsernamePassword.BrowserProxy)
	assert.Equal(t, "/fixture/chromium", cfg.WebAPI.UsernamePassword.BrowserExecutablePath)
	assert.Empty(t, cfg.WebAPI.AccessToken)
	for _, value := range []string{"0s", "-1s", "invalid"} {
		t.Setenv("NEXTHINK_LOGIN_TIMEOUT", value)
		_, err = AuthConfigFromEnv()
		require.ErrorContains(t, err, "NEXTHINK_LOGIN_TIMEOUT")
	}
	t.Setenv("NEXTHINK_LOGIN_TIMEOUT", "2m")
	cfg, err = AuthConfigFromEnv()
	require.NoError(t, err)
	assert.Equal(t, 2*time.Minute, cfg.WebAPI.UsernamePassword.LoginTimeout)
}

func TestConfigurationRedaction(t *testing.T) {
	cfg := AuthConfig{Instance: "fixture", Region: "eu", PublicAPI: &ClientCredentials{ClientID: "secret-client", ClientSecret: "secret-api"}, WebAPI: &BrowserCredentials{AccessToken: "secret-token", UsernamePassword: &UsernamePasswordCredentials{Username: "secret-user", Password: "secret-password"}}}
	for _, value := range []any{cfg, &cfg, *cfg.PublicAPI, cfg.PublicAPI, *cfg.WebAPI, cfg.WebAPI} {
		encoded, err := json.Marshal(value)
		require.NoError(t, err)
		for _, output := range []string{string(encoded), fmt.Sprintf("%v", value), fmt.Sprintf("%+v", value), fmt.Sprintf("%#v", value)} {
			for _, secret := range []string{"secret-client", "secret-api", "secret-token", "secret-user", "secret-password"} {
				assert.NotContains(t, output, secret)
			}
		}
	}
}
