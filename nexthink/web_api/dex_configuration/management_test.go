package dex_configuration

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dex_configuration/mocks"
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
	return []managementCase{
		{name: "GetApplications", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.GetApplications(context.Background())
		}}, {name: "GetAccount", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.GetAccount(context.Background())
		}}, {name: "UpdateApplications", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			var request UpdateApplicationsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementUpdateApplications_input"), &request))
			return service.UpdateApplications(context.Background(), &request)
		}}, {name: "GetScoreMetrics", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			var request GetScoreMetricsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementGetScoreMetrics_input"), &request))
			return service.GetScoreMetrics(context.Background(), &request)
		}}, {name: "UpdateScoreMetrics", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			var request UpdateScoreMetricsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementUpdateScoreMetrics_input"), &request))
			return service.UpdateScoreMetrics(context.Background(), &request)
		}}, {name: "GetVDIOptIn", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.GetVDIOptIn(context.Background())
		}}, {name: "OptInVDI", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.OptInVDI(context.Background())
		}}, {name: "GetMemoryMetricsOptIn", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.GetMemoryMetricsOptIn(context.Background())
		}}, {name: "OptInMemoryMetrics", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.OptInMemoryMetrics(context.Background())
		}}, {name: "GetCampaign", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			return service.GetCampaign(context.Background())
		}}, {name: "EnableCampaign", endpoint: Endpoint, call: func(service *Service) (any, *interfaces.Response, error) {
			var request EnableCampaignRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ManagementEnableCampaign_input"), &request))
			return service.EnableCampaign(context.Background(), &request)
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
		result, response, err := service.UpdateApplications(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.GetScoreMetrics(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.UpdateScoreMetrics(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	{
		result, response, err := service.EnableCampaign(context.Background(), nil)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
	}
	assert.Equal(t, 0, mock.GetTotalCallCount())
}

func TestScoreMetricPatchPreservesNullAndFalse(t *testing.T) {
	disabled := false
	patch := ScoreMetricInput{ID: 10203000000, Average: json.RawMessage("null"), AverageEnabled: &disabled, Selected: &disabled}
	b, err := json.Marshal(patch)
	require.NoError(t, err)
	assert.JSONEq(t, `{"id":10203000000,"avg":null,"avgEnabled":false,"selected":false}`, string(b))
}
func TestInvalidMetricThresholdsDoNotSend(t *testing.T) {
	for _, threshold := range []string{`"12"`, `true`, `{}`, `[12]`, `broken`} {
		transport, mock := testutil.NewTransport(t)
		result, response, err := NewService(transport).UpdateScoreMetrics(context.Background(), &UpdateScoreMetricsRequest{Metrics: []ScoreMetricInput{{ID: 1, Average: json.RawMessage(threshold)}}})
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
func TestNilCollectionsDoNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	_, _, err := service.UpdateApplications(context.Background(), &UpdateApplicationsRequest{})
	require.Error(t, err)
	_, _, err = service.UpdateScoreMetrics(context.Background(), &UpdateScoreMetricsRequest{})
	require.Error(t, err)
	_, _, err = service.EnableCampaign(context.Background(), &EnableCampaignRequest{})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
