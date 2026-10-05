package ratings

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ratings/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var value T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &value))
	return &value
}

type contractCase struct {
	name, verb, path              string
	status                        int
	hasBody, raw, metadata, graph bool
	call                          func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "List", verb: "GET", path: EndpointList, status: 200, hasBody: false, raw: false, metadata: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.List(ctx) }},
		{name: "Get", verb: "GET", path: Endpoint + "/fixture-id", status: 200, hasBody: false, raw: false, metadata: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.Get(ctx, "fixture-id") }},
		{name: "Create", verb: "POST", path: Endpoint, status: 200, hasBody: true, raw: false, metadata: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Create(ctx, load[RatingInput](t, "Create_input"))
		}},
		{name: "Update", verb: "PUT", path: Endpoint + "/fixture-id?revision=2", status: 200, hasBody: true, raw: false, metadata: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Update(ctx, load[UpdateRequest](t, "Update_input"))
		}},
		{name: "Delete", verb: "DELETE", path: Endpoint + "/fixture-id?revision=2", status: 200, hasBody: true, raw: false, metadata: false, graph: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.Delete(ctx, "fixture-id", 2) }}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contracts(t) {
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
						assert.Equal(t, "text/plain", r.Header.Get("Content-Type"))
						assert.Equal(t, "fixture.svg", r.Header.Get("x-nxt-asset-name"))
					} else {
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					}
				}
				if tt.metadata {
					res := httpmock.NewStringResponse(tt.status, "")
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
			if tt.metadata {
				assert.Nil(t, result)
				return
			}
			require.NotNil(t, result)
			expected := mocks.Fixture(tt.name + "_success")
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
func TestFailures(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, code := range []int{401, 403, 409} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "error_unauthorized"))
				result, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, code, response.StatusCode)
				assert.Nil(t, result)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			result, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, result)
			if !tt.metadata {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
					res := httpmock.NewStringResponse(200, "{broken")
					res.Header.Set("Content-Type", "application/json")
					return res, nil
				})
				result, _, err = tt.call(NewService(transport))
				require.Error(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
