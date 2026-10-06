package workflow_executions

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflow_executions/mocks"
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
	return []contract{{name: "GetTimeline", method: "GET", path: "/apigateway/workflow-executions-insights/api/v3/workflows/fixture-workflowID/executions/fixture-executionID/execution-timeline", query: "{\"executionStatus\": \"COMPLETED\", \"terminationDate\": \"2026-01-02T00:00:00Z\"}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetTimeline(ctx, "fixture-workflowID", "fixture-executionID", load[TimelineOptions](t, "GetTimeline_request"))
	}}, {name: "GetTimelineV2", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/execution-timeline", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetTimelineV2(ctx, "fixture-workflowID", "fixture-executionID")
	}}, {name: "ListActivities", method: "GET", path: "/apigateway/workflow-executions-insights/api/v1/workflows/fixture-workflowID/executions/fixture-executionID/activities-history", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ListActivities(ctx, "fixture-workflowID", "fixture-executionID")
	}}, {name: "GetRemoteActionDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/remote-action/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetRemoteActionDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "GetCustomFieldsDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/updatecustomfields/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetCustomFieldsDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "GetCampaignDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/campaign/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetCampaignDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "GetFunctionDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/function/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetFunctionDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "GetMessageDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/message/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetMessageDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "GetSAPIDetails", method: "GET", path: "/apigateway/workflow-executions-insights/api/v2/workflows/fixture-workflowID/executions/fixture-executionID/thinklet/sapi/fixture-thinkletID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetSAPIDetails(ctx, "fixture-workflowID", "fixture-executionID", "fixture-thinkletID")
	}}, {name: "ListWorkflows", method: "GET", path: "/apigateway/workflows/api/v1/workflows", query: "{\"fetchOnlyActiveWorkflows\": \"true\", \"fetchWorkflowContent\": \"false\", \"target\": \"DEVICE\"}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ListWorkflows(ctx, load[ListOptions](t, "ListWorkflows_request"))
	}}, {name: "GetWorkflow", method: "GET", path: "/apigateway/workflows/api/v2/workflows/fixture-workflowID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetWorkflow(ctx, "fixture-workflowID") }}, {name: "Get", method: "GET", path: "/apigateway/workflows/api/v1/workflows/fixture-workflowID/executions/fixture-executionID", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.Get(ctx, "fixture-workflowID", "fixture-executionID")
	}}, {name: "GetHistory", method: "GET", path: "/apigateway/workflows/api/v1/workflows/fixture-workflowID/executions/fixture-executionID/history", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetHistory(ctx, "fixture-workflowID", "fixture-executionID")
	}}, {name: "Execute", method: "POST", path: "/apigateway/workflows/api/v2/execute", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.Execute(ctx, load[ExecuteRequest](t, "Execute_request"))
	}}, {name: "ExecuteNQL", method: "POST", path: "/apigateway/workflows/api/v2/execute/nql", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ExecuteNQL(ctx, load[ExecuteNQLRequest](t, "ExecuteNQL_request"))
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
