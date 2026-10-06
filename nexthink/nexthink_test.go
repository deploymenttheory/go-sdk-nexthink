package nexthink

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	shellmocks "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/product_shell/mocks"
	actionmocks "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/remote_actions/mocks"
	workflowmocks "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflows/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
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
	mock.RegisterResponder(
		"POST",
		"https://fixture.eu.nexthink.cloud/apigateway/workflows/manage/graphql",
		func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer web-token", r.Header.Get("Authorization"))
			return workflowmocks.Responder(200, "List_success")(r)
		},
	)
	mock.RegisterResponder(
		"POST",
		"https://fixture.eu.nexthink.cloud/apigateway/act/manage/graphql",
		func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer web-token", r.Header.Get("Authorization"))
			return actionmocks.Responder(200, "GetContentVolume_success")(r)
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
	require.NotNil(t, c.WebAPI.Workflows)
	require.NotNil(t, c.WebAPI.RemoteActions)
	require.NotNil(t, c.WebAPI.Applications)
	require.NotNil(t, c.WebAPI.WritingAssistant)
	require.NotNil(t, c.WebAPI.SoftwareMetering)
	require.NotNil(t, c.WebAPI.CustomFields)
	require.NotNil(t, c.WebAPI.RuleBasedCustomFields)
	require.NotNil(t, c.WebAPI.Monitors)
	require.NotNil(t, c.WebAPI.Campaigns)
	require.NotNil(t, c.WebAPI.Assets)
	require.NotNil(t, c.WebAPI.Checklists)
	require.NotNil(t, c.WebAPI.Dashboards)
	require.NotNil(t, c.WebAPI.Ratings)
	require.NotNil(t, c.WebAPI.Investigations)
	require.NotNil(t, c.WebAPI.Connectors)
	require.NotNil(t, c.WebAPI.KnowledgeBases)
	require.NotNil(t, c.WebAPI.LegacyConnectors)
	require.NotNil(t, c.WebAPI.Webhooks)
	require.NotNil(t, c.WebAPI.DataExporters)
	require.NotNil(t, c.WebAPI.ConnectorCredentials)
	require.NotNil(t, c.WebAPI.ApplicationExperience)
	require.NotNil(t, c.WebAPI.DexConfiguration)
	require.NotNil(t, c.WebAPI.AlertHub)
	require.NotNil(t, c.WebAPI.Diagnostics)
	require.NotNil(t, c.WebAPI.Benchmark)
	require.NotNil(t, c.WebAPI.DexScores)
	require.NotNil(t, c.WebAPI.CCIInsights)
	require.NotNil(t, c.WebAPI.NetworkInsights)
	require.NotNil(t, c.WebAPI.DataExploration)
	require.NotNil(t, c.WebAPI.Library)
	require.NotNil(t, c.WebAPI.Recommendations)
	require.NotNil(t, c.WebAPI.CollaborationComments)
	require.NotNil(t, c.WebAPI.Support)
	require.NotNil(t, c.WebAPI.SupportChecklists)
	require.NotNil(t, c.WebAPI.SupportTimeline)
	require.NotNil(t, c.WebAPI.SupportInsights)
	require.NotNil(t, c.WebAPI.CollaborationTools)
	require.NotNil(t, c.WebAPI.VDI)
	require.NotNil(t, c.WebAPI.WorkflowExecutions)
	require.NotNil(t, c.WebAPI.ActionExecutions)
	require.NotNil(t, c.WebAPI.AccessManagement)
	require.NotNil(t, c.WebAPI.LegacyAccessManagement)
	require.NotNil(t, c.WebAPI.ContentSharing)
	require.NotNil(t, c.WebAPI.DataExport)
	require.NotNil(t, c.WebAPI.VisualEditor)
	require.NotNil(t, c.WebAPI.CustomFieldValues)
	require.NotNil(t, c.WebAPI.GlobalSearch)
	require.NotNil(t, c.WebAPI.NLPAssistant)
	require.NotNil(t, c.WebAPI.Autopilot)
	_, _, err = c.PublicAPI.Workflows.ListWorkflows(context.Background())
	require.NoError(t, err)
	_, _, err = c.WebAPI.ProductShell.GetMenu(context.Background())
	require.NoError(t, err)
	_, _, err = c.WebAPI.Workflows.List(context.Background())
	require.NoError(t, err)
	_, _, err = c.WebAPI.RemoteActions.GetContentVolume(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 5, mock.GetTotalCallCount())
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
