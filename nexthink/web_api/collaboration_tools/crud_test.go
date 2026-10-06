package collaboration_tools

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collaboration_tools/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var v T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &v))
	return &v
}

type contract struct {
	name, method, path, query string
	call                      func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	t.Helper()
	ctx := context.Background()
	return []contract{{name: "GetCallInsights", method: "POST", path: "/apigateway/collaboration-tools/api/v1/device/fixture-deviceID/call-insights", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetCallInsights(ctx, "fixture-deviceID", load[CallInsightsRequest](t, "GetCallInsights_request"))
	}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "Europe/London", r.Header.Get("time-zone"))
				assert.Equal(t, "-60", r.Header.Get("utc-offset"))
				var query map[string]string
				require.NoError(t, json.Unmarshal([]byte(tt.query), &query))
				actual := map[string]string{}
				for k := range r.URL.Query() {
					actual[k] = r.URL.Query().Get(k)
				}
				assert.Equal(t, query, actual)
				if tt.method == "POST" {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport, WithTimeZone("Europe/London", -60)))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestErrors(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(403, "error"))
			result, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, 403, response.StatusCode)
		})
	}
}
func TestMalformedSuccess(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
				r := httpmock.NewStringResponse(200, "{broken")
				r.Header.Set("Content-Type", "application/json")
				return r, nil
			})
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
		})
	}
}
func TestInvalidIdentifier(t *testing.T) {
	for _, id := range []string{"", " ", "..", "a/b", "a?b", "a#b", "a\nb"} {
		assert.Error(t, validateID(id))
	}
	assert.NoError(t, validateID("fixture-id"))
}
