package workspace_tasks

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace_tasks/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, n string) *T {
	t.Helper()
	var r T
	require.NoError(t, json.Unmarshal(mocks.Fixture(n), &r))
	return &r
}

type contract struct {
	name, method, path string
	body, raw, empty   bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	ctx := context.Background()
	return []contract{{name: "ListTasks", method: "GET", path: "/apigateway/nlp/automations/api/v1/automations?page=0&pageSize=100", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ListTasks(ctx, &TaskListOptions{Page: 0, PageSize: 100})
	}},
		{name: "GetTask", method: "GET", path: "/apigateway/nlp/automations/api/v1/automations/fixture-id", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetTask(ctx, "fixture-id") }},
		{name: "CreateTask", method: "POST", path: "/apigateway/nlp/automations/api/v1/automations", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateTask(ctx, load[TaskRequest](t, "CreateTask_request"))
		}},
		{name: "UpdateTask", method: "PUT", path: "/apigateway/nlp/automations/api/v1/automations/fixture-id", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateTask(ctx, "fixture-id", load[TaskRequest](t, "UpdateTask_request"))
		}},
		{name: "DeleteTask", method: "DELETE", path: "/apigateway/nlp/automations/api/v1/automations/fixture-id", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.DeleteTask(ctx, "fixture-id")
			return nil, r, e
		}},
		{name: "ReconcileTaskAgentAccess", method: "POST", path: "/apigateway/nlp/automations/api/v1/automations/reconcile-it-agent-access", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ReconcileTaskAgentAccess(ctx)
			return nil, r, e
		}}}
}
func TestContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					if tt.raw {
						var expected string
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_request"), &expected))
						assert.Equal(t, expected, string(body))
						assert.Equal(t, "application/octet-stream", r.Header.Get("Content-Type"))
					} else {
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					}
				}
				status, body := 204, ""
				if !tt.empty {
					status = 200
					body = string(mocks.Fixture(tt.name + "_success"))
				}
				response := httpmock.NewStringResponse(status, body)
				response.Header.Set("Content-Type", "application/json")
				response.Header.Set("X-Request-ID", "fixture-request")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			if tt.empty {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				expected := mocks.Fixture(tt.name + "_success")
				actual, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(expected), string(actual))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}

func TestHTTPErrorsPreserveResponse(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, httpmock.NewStringResponder(403, string(mocks.Fixture("error"))))
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 403, response.StatusCode)
			assert.JSONEq(t, string(mocks.Fixture("error")), string(response.Body))
		})
	}
}
func TestMalformedJSONResponses(t *testing.T) {
	for _, tt := range contracts(t) {
		if tt.empty {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, "{broken")
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			})
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
		})
	}
}
