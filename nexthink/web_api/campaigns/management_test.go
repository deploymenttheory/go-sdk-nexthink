package campaigns

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/campaigns/mocks"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
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
	return []managementCase{{name: "GetBranding", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		return service.GetBranding(context.Background())
	}}, {name: "UpdateBranding", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request UpdateBrandingRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementUpdateBranding_input"), &request))
		return service.UpdateBranding(context.Background(), &request)
	}}, {name: "SetStatus", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request SetStatusRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementSetStatus_input"), &request))
		return service.SetStatus(context.Background(), &request)
	}}, {name: "GetByNQLID", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request GetByNQLIDRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementGetByNQLID_input"), &request))
		return service.GetByNQLID(context.Background(), &request)
	}}, {name: "GetFromLibrary", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request GetFromLibraryRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementGetFromLibrary_input"), &request))
		return service.GetFromLibrary(context.Background(), &request)
	}}, {name: "GetWithV6", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
		var request GetWithV6Request
		require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementGetWithV6_input"), &request))
		return service.GetWithV6(context.Background(), &request)
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
		result, response, err := service.UpdateBranding(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.SetStatus(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.GetByNQLID(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.GetFromLibrary(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.GetWithV6(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	assert.Equal(t, 0, mock.GetTotalCallCount())
}

func TestRetireUnpublishedCampaignRetainsBusinessError(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, "ManagementSetStatus_rejected"))
	id := "fixture-campaign"
	result, response, err := NewService(transport).SetStatus(context.Background(), &SetStatusRequest{ContentID: &id, Status: "RETIRED"})
	require.NotNil(t, result)
	assert.Nil(t, result.Campaign)
	require.NotNil(t, response)
	var gqlErrors graphql.GraphQLErrors
	require.ErrorAs(t, err, &gqlErrors)
	require.Len(t, gqlErrors, 1)
	assert.Equal(t, "CANNOT_RETIRE_UNPUBLISHED_CAMPAIGN", gqlErrors[0].Extensions["code"])
}
