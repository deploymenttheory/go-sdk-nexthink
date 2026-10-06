package application_experience

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/application_experience/mocks"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var result T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &result))
	return &result
}

type contractCase struct {
	name, method, path string
	hasBody, graphQL   bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "GetApplicationsOverviewDesktopInvestigations", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetApplicationsOverviewDesktopInvestigations(ctx, load[GetApplicationsOverviewDesktopInvestigationsRequest](t, "GetApplicationsOverviewDesktopInvestigations_input"))
	}},
		{name: "GetAvgNetworkResponseTime", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetAvgNetworkResponseTime(ctx, load[GetAvgNetworkResponseTimeRequest](t, "GetAvgNetworkResponseTime_input"))
		}},
		{name: "GetBinarySuggestion", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetBinarySuggestion(ctx, load[GetBinarySuggestionRequest](t, "GetBinarySuggestion_input"))
		}},
		{name: "GetDeviceCentricMetricBreakdown", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDeviceCentricMetricBreakdown(ctx, load[GetDeviceCentricMetricBreakdownRequest](t, "GetDeviceCentricMetricBreakdown_input"))
		}},
		{name: "GetFailedConnectionsRatio", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetFailedConnectionsRatio(ctx, load[GetFailedConnectionsRatioRequest](t, "GetFailedConnectionsRatio_input"))
		}},
		{name: "GetInsights", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetInsights(ctx, load[GetInsightsRequest](t, "GetInsights_input"))
		}},
		{name: "GetMetricBreakdown", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetMetricBreakdown(ctx, load[GetMetricBreakdownRequest](t, "GetMetricBreakdown_input"))
		}},
		{name: "GetNumOfCrashesAndDevices", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfCrashesAndDevices(ctx, load[GetNumOfCrashesAndDevicesRequest](t, "GetNumOfCrashesAndDevices_input"))
		}},
		{name: "GetNumOfCrashesAndEmployees", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfCrashesAndEmployees(ctx, load[GetNumOfCrashesAndEmployeesRequest](t, "GetNumOfCrashesAndEmployees_input"))
		}},
		{name: "GetNumOfDevices", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfDevices(ctx, load[GetNumOfDevicesRequest](t, "GetNumOfDevices_input"))
		}},
		{name: "GetNumOfDevicesWithCrashes", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfDevicesWithCrashes(ctx, load[GetNumOfDevicesWithCrashesRequest](t, "GetNumOfDevicesWithCrashes_input"))
		}},
		{name: "GetNumOfEmployees", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfEmployees(ctx, load[GetNumOfEmployeesRequest](t, "GetNumOfEmployees_input"))
		}},
		{name: "GetNumOfEmployeesWithCrashes", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetNumOfEmployeesWithCrashes(ctx, load[GetNumOfEmployeesWithCrashesRequest](t, "GetNumOfEmployeesWithCrashes_input"))
		}},
		{name: "OverviewDesktopTooltips", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.OverviewDesktopTooltips(ctx, load[OverviewDesktopTooltipsRequest](t, "OverviewDesktopTooltips_input"))
		}},
		{name: "TilesDesktopCrashesPerEmployee", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesDesktopCrashesPerEmployee(ctx, load[TilesDesktopCrashesPerEmployeeRequest](t, "TilesDesktopCrashesPerEmployee_input"))
		}},
		{name: "TilesDesktopNumberOfEmployees", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesDesktopNumberOfEmployees(ctx, load[TilesDesktopNumberOfEmployeesRequest](t, "TilesDesktopNumberOfEmployees_input"))
		}},
		{name: "TilesErrorCount", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesErrorCount(ctx, load[TilesErrorCountRequest](t, "TilesErrorCount_input"))
		}},
		{name: "TilesFrustratingPageLoads", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesFrustratingPageLoads(ctx, load[TilesFrustratingPageLoadsRequest](t, "TilesFrustratingPageLoads_input"))
		}},
		{name: "TilesNumberOfEmployees", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesNumberOfEmployees(ctx, load[TilesNumberOfEmployeesRequest](t, "TilesNumberOfEmployees_input"))
		}},
		{name: "TilesPageLoadTime", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesPageLoadTime(ctx, load[TilesPageLoadTimeRequest](t, "TilesPageLoadTime_input"))
		}},
		{name: "TilesTransactionTime", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesTransactionTime(ctx, load[TilesTransactionTimeRequest](t, "TilesTransactionTime_input"))
		}},
		{name: "TilesUsageTime", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.TilesUsageTime(ctx, load[TilesUsageTimeRequest](t, "TilesUsageTime_input"))
		}},
		{name: "WebOverviewTooltips", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.WebOverviewTooltips(ctx, load[WebOverviewTooltipsRequest](t, "WebOverviewTooltips_input"))
		}}}

}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			expected := mocks.Fixture(tt.name + "_success")
			if tt.graphQL {
				var envelope struct {
					Data json.RawMessage `json:"data"`
				}
				require.NoError(t, json.Unmarshal(expected, &envelope))
				expected = envelope.Data
			}
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, failure := range []struct {
				name      string
				responder httpmock.Responder
			}{{"HTTP401", mocks.Responder(401, "error_unauthorized")}, {"HTTP403", mocks.Responder(403, "error_unauthorized")}, {"HTTP409", mocks.Responder(409, "error_unauthorized")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, "{broken")
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			}}} {
				t.Run(failure.name, func(t *testing.T) {
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, failure.responder)
					result, _, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
				})
			}
		})
	}
}
func TestGraphQLPartialData(t *testing.T) {
	for _, tt := range contractCases(t) {
		if !tt.graphQL {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(200, tt.name+"_partial"))
			result, response, err := tt.call(NewService(transport))
			var graphErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &graphErrors)
			require.NotNil(t, result)
			require.NotNil(t, response)
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
			assert.JSONEq(t, string(expected.Data), string(actual))
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(200, "errors_only"))
			result, _, err = tt.call(NewService(transport))
			require.ErrorAs(t, err, &graphErrors)
			assert.Nil(t, result)
		})
	}
}
