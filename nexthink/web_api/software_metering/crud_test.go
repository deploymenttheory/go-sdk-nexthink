package software_metering

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/software_metering/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var result T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &result))
	return &result
}

type contractCase struct {
	name, method, path string
	hasBody, graphQL   bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "List", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.List(ctx) }},
		{name: "Get", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.Get(ctx, "fixture-id") }},
		{name: "Create", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Create(ctx, load[ConfigurationInput](t, "Create_input"))
		}},
		{name: "Update", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Update(ctx, "fixture-id", load[ConfigurationInput](t, "Update_input"))
		}},
		{name: "Delete", method: "POST", path: Endpoint, hasBody: true, graphQL: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.Delete(ctx, "fixture-id") }}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			expected := mocks.Fixture(tt.name + "_success")
			if tt.graphQL {
				var envelope struct {
					Data json.RawMessage `json:"data"`
				}
				require.NoError(t, json.Unmarshal(expected, &envelope))
				expected = envelope.Data
			}
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, failure := range []struct {
				name      string
				responder httpmock.Responder
			}{{"HTTP401", mocks.Responder(401, "error_unauthorized")}, {"HTTP403", mocks.Responder(403, "error_unauthorized")}, {"HTTP409", mocks.Responder(409, "error_unauthorized")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, "{broken")
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			}}} {
				t.Run(failure.name, func(t *testing.T) {
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, failure.responder)
					result, _, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
				})
			}
		})
	}
}
func TestGraphQLPartialData(t *testing.T) {
	for _, tt := range contractCases(t) {
		if !tt.graphQL {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(200, tt.name+"_partial"))
			result, response, err := tt.call(NewService(transport))
			var graphErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &graphErrors)
			require.NotNil(t, result)
			require.NotNil(t, response)
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
			assert.JSONEq(t, string(expected.Data), string(actual))
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(200, "errors_only"))
			result, _, err = tt.call(NewService(transport))
			require.ErrorAs(t, err, &graphErrors)
			assert.Nil(t, result)
		})
	}
}
