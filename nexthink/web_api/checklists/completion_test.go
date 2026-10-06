package checklists

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/checklists/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func completionLoad[T any](t *testing.T, name string) *T {
	t.Helper()
	var value T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &value))
	return &value
}

type completionCase struct {
	name, verb, path, mode string
	status                 int
	hasBody, raw, graph    bool
	call                   func(*Service) (any, *interfaces.Response, error)
}

func completionCases(t *testing.T) []completionCase {
	t.Helper()
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	executionID := "22222222-2222-4333-8444-555555555555"
	_ = id
	_ = executionID
	return []completionCase{{name: "CompletionGetFromLibrary", verb: "GET", path: "/apigateway/atl/support-checklist-config-be/api/v1/builtincontent/library/11111111-2222-4333-8444-555555555555", status: 200, mode: "json", hasBody: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetFromLibrary(ctx, id) }}}
}

func TestCompletionWireContracts(t *testing.T) {
	for _, tt := range completionCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					if tt.raw {
						var expected string
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_request"), &expected))
						assert.Equal(t, expected, string(body))
						assert.Equal(t, "application/octet-stream", r.Header.Get("Content-Type"))
					} else {
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					}
				}
				if tt.mode != "json" {
					body := ""
					if tt.mode == "text" {
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &body))
					}
					res := httpmock.NewStringResponse(tt.status, body)
					res.Header.Set("Content-Type", "text/plain")
					res.Header.Set("X-Request-ID", "fixture-request")
					return res, nil
				}
				return mocks.Responder(tt.status, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, tt.status, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if tt.mode == "empty" {
				assert.Nil(t, result)
				return
			}
			expected := mocks.Fixture(tt.name + "_success")
			if tt.mode == "text" {
				var message string
				require.NoError(t, json.Unmarshal(expected, &message))
				expected, _ = json.Marshal(map[string]string{"message": message})
			}
			if tt.graph {
				var envelope struct {
					Data json.RawMessage `json:"data"`
				}
				require.NoError(t, json.Unmarshal(expected, &envelope))
				expected = envelope.Data
			}
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
func TestCompletionFailures(t *testing.T) {
	for _, tt := range completionCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for code, name := range map[int]string{400: "validation", 401: "unauthorized", 403: "forbidden"} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "completion_error_"+name))
				_, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, code, response.StatusCode)
				assert.JSONEq(t, string(mocks.Fixture("completion_error_"+name)), string(response.Body))
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			if tt.mode == "json" {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
					res := httpmock.NewStringResponse(200, "{broken")
					res.Header.Set("Content-Type", "application/json")
					return res, nil
				})
				_, _, err = tt.call(NewService(transport))
				require.Error(t, err)
			}
			if tt.graph {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewJsonResponderOrPanic(200, map[string]any{"data": map[string]any{}, "errors": []map[string]any{{"message": "fixture conflict"}}}))
				result, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				require.NotNil(t, result)
				assert.Equal(t, 200, response.StatusCode)
			}
		})
	}
}
func TestCompletionValidationBeforeTransport(t *testing.T) {
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	_ = ctx
	_ = id
	cases := []struct {
		name string
		call func(*Service) (any, *interfaces.Response, error)
	}{{name: "GetFromLibrary", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetFromLibrary(ctx, "../bad") }}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, response)
			assert.Zero(t, mock.GetTotalCallCount())
		})
	}
}
