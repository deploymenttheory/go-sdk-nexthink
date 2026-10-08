package workspace_assignments

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace_assignments/mocks"
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
	return []contract{{name: "ListAssignments", method: "GET", path: "/infinity/assignments/api/v1/assignments?sort=priority&sources=workspace&sources=test", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ListAssignments(ctx, &AssignmentOptions{Sources: []string{"workspace", "test"}, Sort: "priority"})
	}},
		{name: "GetAssignment", method: "GET", path: "/infinity/assignments/api/v1/assignments/fixture-id", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetAssignment(ctx, "fixture-id", nil) }},
		{name: "UpdateAssignment", method: "PATCH", path: "/infinity/assignments/api/v1/assignments/fixture-id", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateAssignment(ctx, "fixture-id", load[AssignmentUpdate](t, "UpdateAssignment_request"))
		}},
		{name: "MarkAssignmentRead", method: "POST", path: "/infinity/assignments/api/v1/assignments/fixture-id/read", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.MarkAssignmentRead(ctx, "fixture-id")
			return nil, r, e
		}},
		{name: "GetUnreadAssignmentCount", method: "GET", path: "/infinity/assignments/api/v1/assignments/summary/unread-count", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetUnreadAssignmentCount(ctx, nil) }},
		{name: "ListAssignees", method: "GET", path: "/infinity/assignments/api/v1/assignees?assignmentId=fixture-id", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListAssignees(ctx, "fixture-id") }}}
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

func TestAssignmentNotFoundEnvelope(t *testing.T) {
	for _, method := range []string{"GET", "PATCH"} {
		t.Run(method, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(method, testutil.BaseURL+AssignmentsEndpoint+"/fixture-id", func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, string(mocks.Fixture("assignment_not_found")))
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			})
			service := NewService(transport)
			var response *interfaces.Response
			var err error
			if method == "GET" {
				_, response, err = service.GetAssignment(context.Background(), "fixture-id", nil)
			} else {
				_, response, err = service.UpdateAssignment(context.Background(), "fixture-id", &AssignmentUpdate{State: "done", Revision: 1})
			}
			require.ErrorContains(t, err, "not found")
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
		})
	}
}
