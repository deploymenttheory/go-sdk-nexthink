package nexthink

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	shellmocks "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/product_shell/mocks"
)

func tokenResponder(t *testing.T) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		t.Helper()
		id, secret, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "fixture-id", id)
		assert.Equal(t, "fixture-secret", secret)
		response := httpmock.NewStringResponse(
			200,
			`{"access_token":"public-token","token_type":"Bearer","expires_in":3600}`,
		)
		response.Header.Set("Content-Type", "application/json")
		return response, nil
	}
}

func TestOneClientRoutesBothFamilies(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder(
		"POST",
		"https://fixture-login.eu.nexthink.cloud/oauth2/default/v1/token",
		tokenResponder(t),
	)
	mock.RegisterResponder(
		"GET",
		"https://fixture.api.eu.nexthink.cloud/api/v1/workflows",
		func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer public-token", r.Header.Get("Authorization"))
			response := httpmock.NewStringResponse(200, `[]`)
			response.Header.Set("Content-Type", "application/json")
			return response, nil
		},
	)
	mock.RegisterResponder(
		"POST",
		"https://fixture.eu.nexthink.cloud/apigateway/api/v1/product-shell/menu",
		func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer web-token", r.Header.Get("Authorization"))
			return shellmocks.Responder(200, "menu_success")(r)
		},
	)
	c, err := NewClient(
		&AuthConfig{
			Instance:  "fixture",
			Region:    "eu",
			PublicAPI: &ClientCredentials{ClientID: "fixture-id", ClientSecret: "fixture-secret"},
			WebAPI:    &BrowserCredentials{AccessToken: "web-token"},
		},
		WithLogger(zap.NewNop()),
		WithTransport(mock),
		WithRetryCount(0),
	)
	require.NoError(t, err)
	require.NotNil(t, c.PublicAPI)
	require.NotNil(t, c.WebAPI)
	_, _, err = c.PublicAPI.Workflows.ListWorkflows(context.Background())
	require.NoError(t, err)
	_, _, err = c.WebAPI.ProductShell.GetMenu(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 3, mock.GetTotalCallCount())
	require.NotNil(t, c.PublicAPI.GetTokenManager())
	require.NotNil(t, c.GetLogger())
}

func TestWebOnlyRequiresNoOAuthCredentials(t *testing.T) {
	mock := httpmock.NewMockTransport()
	c, err := NewClient(
		&AuthConfig{
			Instance: "fixture",
			Region:   "eu",
			WebAPI:   &BrowserCredentials{AccessToken: "web-token"},
		},
		WithLogger(zap.NewNop()),
		WithTransport(mock),
	)
	require.NoError(t, err)
	assert.Nil(t, c.PublicAPI)
	require.NotNil(t, c.WebAPI)
	assert.Zero(t, mock.GetTotalCallCount())
}

func TestInvalidWebConfigurationPreventsOAuthRequest(t *testing.T) {
	mock := httpmock.NewMockTransport()
	c, err := NewClient(
		&AuthConfig{
			Instance:  "fixture",
			Region:    "eu",
			PublicAPI: &ClientCredentials{ClientID: "fixture-id", ClientSecret: "fixture-secret"},
			WebAPI:    &BrowserCredentials{},
		},
		WithTransport(mock),
	)
	require.ErrorContains(t, err, "WebAPI")
	assert.Nil(t, c)
	assert.Zero(t, mock.GetTotalCallCount())
}

func TestScopedTransportOptions(t *testing.T) {
	mock := httpmock.NewMockTransport()
	mock.RegisterResponder("POST", "https://login.example.test/token", tokenResponder(t))
	c, err := NewClient(
		&AuthConfig{
			Instance:  "fixture",
			Region:    "eu",
			PublicAPI: &ClientCredentials{ClientID: "fixture-id", ClientSecret: "fixture-secret"},
			WebAPI:    &BrowserCredentials{AccessToken: "web-token"},
		},
		WithLogger(zap.NewNop()),
		WithTransport(mock),
		WithPublicAPIOptions(
			client.WithBaseURL("https://public.example.test"),
			client.WithCustomTokenURL("https://login.example.test/token"),
		),
		WithWebAPIOptions(client.WithBaseURL("https://web.example.test")),
	)
	require.NoError(t, err)
	assert.Equal(t, "https://public.example.test", c.PublicAPI.transport.BaseURL)
	assert.Equal(t, "https://web.example.test", c.WebAPI.transport.BaseURL)
}

func TestExpiredProviderAndCanceledRequestDoNotSend(t *testing.T) {
	mock := httpmock.NewMockTransport()
	c, err := NewClient(
		&AuthConfig{
			Instance: "fixture",
			Region:   "eu",
			WebAPI: &BrowserCredentials{
				TokenProvider: auth.StaticToken("expired", time.Now().Add(-time.Second)),
			},
		},
		WithLogger(zap.NewNop()),
		WithTransport(mock),
		WithRetryCount(0),
	)
	require.NoError(t, err)
	_, _, err = c.WebAPI.ProductShell.GetMenu(context.Background())
	require.ErrorIs(t, err, auth.ErrSessionExpired)
	assert.Zero(t, mock.GetTotalCallCount())
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, _, err = c.WebAPI.ProductShell.GetMenu(ctx)
	require.ErrorIs(t, err, context.Canceled)
	assert.Zero(t, mock.GetTotalCallCount())
}
