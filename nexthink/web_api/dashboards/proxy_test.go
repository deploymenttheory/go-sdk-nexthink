package dashboards

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dashboards/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAsyncProxyPreservesOriginalService(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	original := NewService(transport)
	proxy := original.WithAsyncProxy().WithAsyncProxy()
	for _, path := range []string{Endpoint, EndpointAsyncProxy} {
		path := path
		mock.RegisterResponder("POST", testutil.BaseURL+path, func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
			if path == EndpointAsyncProxy {
				assert.Equal(t, "true", r.Header.Get("x-nxt-waas-allow-long-running"))
			} else {
				assert.Empty(t, r.Header.Get("x-nxt-waas-allow-long-running"))
			}
			assert.NotEmpty(t, r.Header.Get("x-nxt-waas-iso-date-time"))
			assert.Equal(t, "UTC", r.Header.Get("x-nxt-waas-timezone"))
			var body struct {
				OperationName string `json:"operationName"`
			}
			require.NoError(t, json.NewDecoder(r.Body).Decode(&body))
			assert.Equal(t, "Collections", body.OperationName)
			return mocks.Responder(200, "CompletionListCollections_success")(r)
		})
	}
	for _, s := range []*Service{proxy, original} {
		result, response, err := s.ListCollections(context.Background())
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, response.StatusCode)
	}
	assert.Equal(t, 2, mock.GetTotalCallCount())
}
func TestProductShellMenuAndProxyREST(t *testing.T) {
	for _, status := range []int{200, 403} {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("GET", testutil.BaseURL+EndpointProductShellMenu, func(r *http.Request) (*http.Response, error) {
			assert.Empty(t, r.Header.Get("x-nxt-waas-allow-long-running"))
			fixture := "GetProductShellMenu_success"
			if status != 200 {
				fixture = "extension_error_forbidden"
			}
			return mocks.Responder(status, fixture)(r)
		})
		result, response, err := NewService(transport).WithAsyncProxy().GetProductShellMenu(context.Background())
		require.NotNil(t, response)
		assert.Equal(t, status, response.StatusCode)
		if status == 200 {
			require.NoError(t, err)
			require.NotNil(t, result.Items)
			assert.Empty(t, result.Items)
		} else {
			require.Error(t, err)
			assert.Nil(t, result)
		}
	}
}
func TestAsyncProxyRetainsGraphQLErrors(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+EndpointAsyncProxy, mocks.Responder(200, "errors_only"))
	_, response, err := NewService(transport).WithAsyncProxy().ListCollections(context.Background())
	require.Error(t, err)
	require.NotNil(t, response)
	assert.Equal(t, 200, response.StatusCode)
}
