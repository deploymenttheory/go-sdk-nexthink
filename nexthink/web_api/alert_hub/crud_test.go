package alert_hub

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/alert_hub/mocks"
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
	return []contractCase{{name: "AlertImpactAssessment", call: func(s *Service) (any, *interfaces.Response, error) {
		var r AlertImpactAssessmentRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("AlertImpactAssessment_input"), &r))
		return s.AlertImpactAssessment(context.Background(), &r)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) {
		return s.AlertImpactAssessment(context.Background(), nil)
	}},
		{name: "EventsDrillDown", call: func(s *Service) (any, *interfaces.Response, error) {
			var r EventsDrillDownRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("EventsDrillDown_input"), &r))
			return s.EventsDrillDown(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.EventsDrillDown(context.Background(), nil)
		}},
		{name: "Issues", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssuesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("Issues_input"), &r))
			return s.Issues(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Issues(context.Background(), nil) }},
		{name: "IssuesSelectedPeriod", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssuesSelectedPeriodRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssuesSelectedPeriod_input"), &r))
			return s.IssuesSelectedPeriod(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssuesSelectedPeriod(context.Background(), nil)
		}},
		{name: "IssuesTimeline", call: func(s *Service) (any, *interfaces.Response, error) {
			var r IssuesTimelineRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("IssuesTimeline_input"), &r))
			return s.IssuesTimeline(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.IssuesTimeline(context.Background(), nil)
		}},
		{name: "MonitorConfigView", call: func(s *Service) (any, *interfaces.Response, error) {
			var r MonitorConfigViewRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("MonitorConfigView_input"), &r))
			return s.MonitorConfigView(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.MonitorConfigView(context.Background(), nil)
		}},
		{name: "AlertTriggerInfo", call: func(s *Service) (any, *interfaces.Response, error) {
			var r AlertTriggerInfoRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("AlertTriggerInfo_input"), &r))
			return s.AlertTriggerInfo(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.AlertTriggerInfo(context.Background(), nil)
		}},
		{name: "AlertsImpactedAssociationOverTime", call: func(s *Service) (any, *interfaces.Response, error) {
			var r AlertsImpactedAssociationOverTimeRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("AlertsImpactedAssociationOverTime_input"), &r))
			return s.AlertsImpactedAssociationOverTime(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.AlertsImpactedAssociationOverTime(context.Background(), nil)
		}},
		{name: "GetAlertOnChangeTimeSeries", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetAlertOnChangeTimeSeriesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetAlertOnChangeTimeSeries_input"), &r))
			return s.GetAlertOnChangeTimeSeries(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetAlertOnChangeTimeSeries(context.Background(), nil)
		}},
		{name: "GetIssuePeriods", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetIssuePeriodsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetIssuePeriods_input"), &r))
			return s.GetIssuePeriods(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetIssuePeriods(context.Background(), nil)
		}},
		{name: "GetStaticAlertTimeSeries", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetStaticAlertTimeSeriesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetStaticAlertTimeSeries_input"), &r))
			return s.GetStaticAlertTimeSeries(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetStaticAlertTimeSeries(context.Background(), nil)
		}},
		{name: "Tags", call: func(s *Service) (any, *interfaces.Response, error) {
			var r TagsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("Tags_input"), &r))
			return s.Tags(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Tags(context.Background(), nil) }}}
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
