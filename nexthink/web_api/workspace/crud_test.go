package workspace

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace/mocks"
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
	return []contract{{name: "ListConversations", method: "GET", path: "/apigateway/nlp/assist/api/v2/conversations", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListConversations(ctx, "") }},
		{name: "GetConversation", method: "GET", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id?limit=25", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetConversation(ctx, "fixture-id", &PageOptions{Limit: 25})
		}},
		{name: "GetSharedConversation", method: "GET", path: "/apigateway/nlp/assist/api/v2/share/fixture-id?limit=25", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetSharedConversation(ctx, "fixture-id", &PageOptions{Limit: 25})
		}},
		{name: "UpdateConversation", method: "PATCH", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateConversation(ctx, "fixture-id", load[ConversationUpdate](t, "UpdateConversation_request"))
		}},
		{name: "DeleteConversation", method: "DELETE", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.DeleteConversation(ctx, "fixture-id")
			return nil, r, e
		}},
		{name: "CancelConversation", method: "POST", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id/cancel", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.CancelConversation(ctx, "fixture-id") }},
		{name: "MarkConversationRead", method: "PATCH", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id/read", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.MarkConversationRead(ctx, "fixture-id") }},
		{name: "CreateConversationShare", method: "POST", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id/share", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateConversationShare(ctx, "fixture-id")
		}},
		{name: "UploadConversationFile", method: "POST", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id/files", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UploadConversationFile(ctx, "fixture-id", load[ConversationFileRequest](t, "UploadConversationFile_request"))
		}},
		{name: "DeleteConversationFile", method: "DELETE", path: "/apigateway/nlp/assist/api/v2/conversations/fixture-id/files/file-id", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.DeleteConversationFile(ctx, "fixture-id", "file-id")
			return nil, r, e
		}},
		{name: "MCPProxy", method: "POST", path: "/apigateway/nlp/assist/api/v2/mcp-proxy", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.MCPProxy(ctx, load[MCPRequest](t, "MCPProxy_request"))
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
					if tt.name == "UploadConversationFile" {
						status = 201
					}
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
