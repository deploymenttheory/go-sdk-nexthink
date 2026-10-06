package workflows

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflows/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type managementCase struct {
	name     string
	endpoint string
	call     func(*Service) (any, *interfaces.Response, error)
}

func managementCases(t *testing.T) []managementCase {
	t.Helper()
	return []managementCase{{name: "GetFromLibrary", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request GetFromLibraryRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementGetFromLibrary_input"), &request))
		return service.GetFromLibrary(context.Background(), &request)
	}}, {name: "SetActive", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request SetActiveRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementSetActive_input"), &request))
		return service.SetActive(context.Background(), &request)
	}}, {name: "Import", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request ImportRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementImport_input"), &request))
		return service.Import(context.Background(), &request)
	}}}
}
func TestManagementWireContracts(t *testing.T) {
	for _, tt := range managementCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+tt.endpoint, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				assert.Empty(t, r.URL.RawQuery)
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Management"+tt.name+"_request")), string(b))
				return mocks.Responder(200, "Management"+tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture("Management"+tt.name+"_success"), &expected))
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(actual))
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestManagementFailures(t *testing.T) {
	for _, tt := range managementCases(t) {
		for _, failure := range []struct {
			name      string
			responder httpmock.Responder
		}{
			{"Unauthorized", mocks.Responder(401, "error_unauthorized")}, {"Forbidden", mocks.Responder(403, "error_unauthorized")}, {"GraphQL", mocks.Responder(200, "errors_only")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(_ *http.Request) (*http.Response, error) {
				r := httpmock.NewStringResponse(200, "{broken")
				r.Header.Set("Content-Type", "application/json")
				return r, nil
			}},
		} {
			t.Run(tt.name+failure.name, func(t *testing.T) {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+tt.endpoint, failure.responder)
				result, _, err := tt.call(NewService(transport))
				require.Error(t, err)
				assert.Nil(t, result)
			})
		}
		t.Run(tt.name+"PartialData", func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+tt.endpoint, mocks.Responder(200, "Management"+tt.name+"_partial"))
			result, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, result)
			require.NotNil(t, response)
			var gqlErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &gqlErrors)
		})
	}
}
func TestManagementNilRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	{
		result, response, err := service.GetFromLibrary(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.SetActive(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.Import(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	assert.Equal(t, 0, mock.GetTotalCallCount())
}
