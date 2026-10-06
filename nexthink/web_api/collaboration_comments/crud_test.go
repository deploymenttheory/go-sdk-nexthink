package collaboration_comments

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collaboration_comments/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var result T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &result))
	return &result
}

type contractCase struct {
	name, method, path string
	hasBody, empty     bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "ResolveIdentifier", method: "POST", path: "/apigateway/rtc/api/v1/identifier", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ResolveIdentifier(ctx, load[IdentifierRequest](t, "ResolveIdentifier_input"))
	}},
		{name: "ListComments", method: "GET", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListComments(ctx, "fixture-id") }},
		{name: "CreateComment", method: "POST", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.CreateComment(ctx, "fixture-id", load[CreateMessageRequest](t, "CreateComment_input"))
			return nil, response, err
		}},
		{name: "CreateReply", method: "POST", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id/replies", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.CreateReply(ctx, "fixture-id", "fixture-id", load[CreateMessageRequest](t, "CreateReply_input"))
			return nil, response, err
		}},
		{name: "ListUserMentions", method: "GET", path: "/apigateway/rtc/api/v1/user-mentions?u=fixture", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListUserMentions(ctx, "fixture") }},
		{name: "ArchiveComment", method: "PATCH", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.ArchiveComment(ctx, "fixture-id", "fixture-id")
			return nil, response, err
		}},
		{name: "UnarchiveComment", method: "PATCH", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UnarchiveComment(ctx, "fixture-id", "fixture-id")
			return nil, response, err
		}},
		{name: "UpdateComment", method: "PATCH", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateComment(ctx, "fixture-id", "fixture-id", load[EditMessageRequest](t, "UpdateComment_input"))
			return nil, response, err
		}},
		{name: "UpdateReply", method: "PATCH", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id/replies/fixture-id", hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateReply(ctx, "fixture-id", "fixture-id", "fixture-id", load[EditMessageRequest](t, "UpdateReply_input"))
			return nil, response, err
		}},
		{name: "DeleteComment", method: "DELETE", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id", hasBody: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.DeleteComment(ctx, "fixture-id", "fixture-id")
			return nil, response, err
		}},
		{name: "DeleteReply", method: "DELETE", path: "/apigateway/rtc/api/v1/documents/fixture-id/comments/fixture-id/replies/fixture-id", hasBody: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.DeleteReply(ctx, "fixture-id", "fixture-id", "fixture-id")
			return nil, response, err
		}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					b, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(b))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 1, mock.GetTotalCallCount())
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			if !tt.empty {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			} else {
				assert.Nil(t, result)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(response.Body))
			}
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, failure := range []struct {
				name      string
				responder httpmock.Responder
			}{{"HTTP400", mocks.Responder(400, "error_validation")}, {"HTTP401", mocks.Responder(401, "error_unauthorized")}, {"HTTP403", mocks.Responder(403, "error_forbidden")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(r *http.Request) (*http.Response, error) {
				x := httpmock.NewStringResponse(200, "{broken")
				x.Header.Set("Content-Type", "application/json")
				return x, nil
			}}} {
				t.Run(failure.name, func(t *testing.T) {
					if tt.empty && failure.name == "Malformed" {
						return
					}
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, failure.responder)
					result, response, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
					if failure.name != "Transport" {
						require.NotNil(t, response)
					}
				})
			}
		})
	}
}
