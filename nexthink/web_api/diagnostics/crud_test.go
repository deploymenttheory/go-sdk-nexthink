package diagnostics

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/diagnostics/mocks"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type contractCase struct {
	name    string
	call    func(*Service) (any, *interfaces.Response, error)
	invalid func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contractCase {
	t.Helper()
	return []contractCase{{name: "BinaryInfo", call: func(s *Service) (any, *interfaces.Response, error) {
		var r BinaryInfoRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("BinaryInfo_input"), &r))
		return s.BinaryInfo(context.Background(), &r)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.BinaryInfo(context.Background(), nil) }},
		{name: "GetDiagnosticContexts", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDiagnosticContextsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDiagnosticContexts_input"), &r))
			return s.GetDiagnosticContexts(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDiagnosticContexts(context.Background(), nil)
		}},
		{name: "GetDiagnosticOverview", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDiagnosticOverviewRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDiagnosticOverview_input"), &r))
			return s.GetDiagnosticOverview(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDiagnosticOverview(context.Background(), nil)
		}},
		{name: "GetImpactedAndTotalObjects", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetImpactedAndTotalObjectsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetImpactedAndTotalObjects_input"), &r))
			return s.GetImpactedAndTotalObjects(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetImpactedAndTotalObjects(context.Background(), nil)
		}},
		{name: "GetStandaloneDashboard", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetStandaloneDashboardRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetStandaloneDashboard_input"), &r))
			return s.GetStandaloneDashboard(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetStandaloneDashboard(context.Background(), nil)
		}},
		{name: "HierarchyBreakdownDimensions", call: func(s *Service) (any, *interfaces.Response, error) {
			var r HierarchyBreakdownDimensionsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("HierarchyBreakdownDimensions_input"), &r))
			return s.HierarchyBreakdownDimensions(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.HierarchyBreakdownDimensions(context.Background(), nil)
		}},
		{name: "IssueTimeseries", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssueTimeseriesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssueTimeseries_input"), &r))
			return s.IssueTimeseries(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssueTimeseries(context.Background(), nil)
		}},
		{name: "TroubleshootingInsights", call: func(s *Service) (any, *interfaces.Response, error) {
			var r TroubleshootingInsightsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("TroubleshootingInsights_input"), &r))
			return s.TroubleshootingInsights(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.TroubleshootingInsights(context.Background(), nil)
		}},
		{name: "GetConfiguration", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetConfigurationRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetConfiguration_input"), &r))
			return s.GetConfiguration(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetConfiguration(context.Background(), nil)
		}},
		{name: "IssueEventsAndAssociatedObjectsByHierarchyBreakdown", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssueEventsAndAssociatedObjectsByHierarchyBreakdownRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssueEventsAndAssociatedObjectsByHierarchyBreakdown_input"), &r))
			return s.IssueEventsAndAssociatedObjectsByHierarchyBreakdown(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssueEventsAndAssociatedObjectsByHierarchyBreakdown(context.Background(), nil)
		}},
		{name: "IssueEventsAndAssociatedObjectsByLocationBreakdown", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssueEventsAndAssociatedObjectsByLocationBreakdownRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssueEventsAndAssociatedObjectsByLocationBreakdown_input"), &r))
			return s.IssueEventsAndAssociatedObjectsByLocationBreakdown(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssueEventsAndAssociatedObjectsByLocationBreakdown(context.Background(), nil)
		}},
		{name: "IssueEventsAndAssociatedObjectsByTechnicalBreakdown", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssueEventsAndAssociatedObjectsByTechnicalBreakdownRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssueEventsAndAssociatedObjectsByTechnicalBreakdown_input"), &r))
			return s.IssueEventsAndAssociatedObjectsByTechnicalBreakdown(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssueEventsAndAssociatedObjectsByTechnicalBreakdown(context.Background(), nil)
		}},
		{name: "GetDashboard", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDashboardRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDashboard_input"), &r))
			return s.GetDashboard(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetDashboard(context.Background(), nil) }}}
}
func TestWireContracts(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
				b, e := io.ReadAll(r.Body)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture(tc.name+"_request")), string(b))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				return mocks.Responder(200, tc.name+"_success")(r)
			})
			result, response, err := tc.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			var envelope struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &envelope))
			got, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(envelope.Data), string(got))
		})
	}
}
func TestGraphQLErrorsRetainPartialData(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			var body map[string]any
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &body))
			body["errors"] = []any{map[string]any{"message": "fixture failure"}}
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewJsonResponderOrPanic(200, body))
			result, response, err := tc.call(NewService(transport))
			require.Error(t, err)
			var graphErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &graphErrors)
			require.Equal(t, "fixture failure", graphErrors[0].Message)
			require.NotNil(t, response)
			require.NotNil(t, result)
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			for _, status := range []int{400, 401, 403, 500} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(status, `{"message":"fixture failure"}`))
				_, response, err := tc.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, status, response.StatusCode)
			}
			for _, body := range []string{`{`, `{"data":null}`, `{"data":[]}`} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(200, body))
				_, _, err := tc.call(NewService(transport))
				require.Error(t, err)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewErrorResponder(errors.New("fixture transport failure")))
			_, _, err := tc.call(NewService(transport))
			require.Error(t, err)
		})
	}
}
func TestNilRequestsMakeNoHTTPCall(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			_, response, err := tc.invalid(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, response)
			assert.Zero(t, mock.GetTotalCallCount())
		})
	}
}

func TestNullableRootResults(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, tc.name+"_null"))
			result, response, err := tc.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			var envelope struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_null"), &envelope))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(envelope.Data), string(data))
		})
	}
}
