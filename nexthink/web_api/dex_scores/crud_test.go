package dex_scores

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dex_scores/mocks"
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
	return []contractCase{{name: "GetCampaign", call: func(s *Service) (any, *interfaces.Response, error) {
		var r GetCampaignRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("GetCampaign_input"), &r))
		return s.GetCampaign(context.Background(), &r)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetCampaign(context.Background(), nil) }},
		{name: "GetDeviceExperience", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDeviceExperienceRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDeviceExperience_input"), &r))
			return s.GetDeviceExperience(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDeviceExperience(context.Background(), nil)
		}},
		{name: "GetDimensionBreakdowns", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDimensionBreakdownsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDimensionBreakdowns_input"), &r))
			return s.GetDimensionBreakdowns(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDimensionBreakdowns(context.Background(), nil)
		}},
		{name: "GetDimensionsV2", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDimensionsV2Request
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDimensionsV2_input"), &r))
			return s.GetDimensionsV2(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDimensionsV2(context.Background(), nil)
		}},
		{name: "GetInvestigationUrl", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetInvestigationUrlRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetInvestigationUrl_input"), &r))
			return s.GetInvestigationUrl(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetInvestigationUrl(context.Background(), nil)
		}},
		{name: "GetLeaves", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetLeavesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetLeaves_input"), &r))
			return s.GetLeaves(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetLeaves(context.Background(), nil) }},
		{name: "GetMetricThreshold", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetMetricThresholdRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetMetricThreshold_input"), &r))
			return s.GetMetricThreshold(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetMetricThreshold(context.Background(), nil)
		}},
		{name: "GetScores", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetScoresRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetScores_input"), &r))
			return s.GetScores(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetScores(context.Background(), nil) }},
		{name: "GetTrend", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrend_input"), &r))
			return s.GetTrend(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetTrend(context.Background(), nil) }},
		{name: "GetTrendDevices", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendDevicesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendDevices_input"), &r))
			return s.GetTrendDevices(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetTrendDevices(context.Background(), nil)
		}},
		{name: "GetTrendEmployeesWithIssues", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendEmployeesWithIssuesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendEmployeesWithIssues_input"), &r))
			return s.GetTrendEmployeesWithIssues(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetTrendEmployeesWithIssues(context.Background(), nil)
		}},
		{name: "GetTrendImprovement", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendImprovementRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendImprovement_input"), &r))
			return s.GetTrendImprovement(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetTrendImprovement(context.Background(), nil)
		}},
		{name: "GetTrendScore", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendScoreRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendScore_input"), &r))
			return s.GetTrendScore(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetTrendScore(context.Background(), nil) }},
		{name: "GetTrendTimeLost", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendTimeLostRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendTimeLost_input"), &r))
			return s.GetTrendTimeLost(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetTrendTimeLost(context.Background(), nil)
		}},
		{name: "GetTrendWithRange", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetTrendWithRangeRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetTrendWithRange_input"), &r))
			return s.GetTrendWithRange(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetTrendWithRange(context.Background(), nil)
		}},
		{name: "GetWhatsChanged", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetWhatsChangedRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetWhatsChanged_input"), &r))
			return s.GetWhatsChanged(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetWhatsChanged(context.Background(), nil)
		}}}
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
