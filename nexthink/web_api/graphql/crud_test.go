package graphql

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql/mocks"
)

type contractCase struct {
	name, method, path, fixture, request, query string
	status                                      int
	headers                                     map[string]string
	call                                        func(*Service) (any, *interfaces.Response, error)
}

func contractCases() []contractCase {
	return []contractCase{
		{
			name:    "graphql.writing_assistant",
			method:  "POST",
			path:    "/apigateway/api/adopt/writing-assistant/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.writing_assistant",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.benchmark",
			method:  "POST",
			path:    "/apigateway/cci/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.benchmark",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.dex_configuration",
			method:  "POST",
			path:    "/apigateway/dex-ec/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.dex_configuration",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.dashboards",
			method:  "POST",
			path:    "/apigateway/dash/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.dashboards",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.workflows",
			method:  "POST",
			path:    "/apigateway/workflows/manage/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.workflows",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.campaigns",
			method:  "POST",
			path:    "/apigateway/euf-gateway/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.campaigns",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.visual_editor",
			method:  "POST",
			path:    "/apigateway/visual-editor/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.visual_editor",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.value_provider",
			method:  "POST",
			path:    "/apigateway/value-provider/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.value_provider",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.monitors",
			method:  "POST",
			path:    "/apigateway/mnt/alert/config/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.monitors",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.remote_actions",
			method:  "POST",
			path:    "/apigateway/act/manage/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.remote_actions",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
		{
			name:    "graphql.software_metering",
			method:  "POST",
			path:    "/apigateway/dex-metering/graphql",
			fixture: "execute_success",
			request: "execute_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Execute(
					context.Background(),
					"graphql.software_metering",
					GraphQLRequest{
						Query:         "query Fixture($limit: Int) { __typename }",
						Variables:     map[string]any{"limit": 1},
						OperationName: "Fixture",
					},
				)
			},
		},
	}
}

func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases() {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				tt.method,
				testutil.BaseURL+tt.path,
				func(r *http.Request) (*http.Response, error) {
					assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
					assert.Equal(t, "*/*", r.Header.Get("Accept"))
					assert.Equal(t, tt.query, r.URL.RawQuery)
					for k, v := range tt.headers {
						assert.Equal(t, v, r.Header.Get(k))
					}
					if tt.request != "" {
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
						body, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						assert.JSONEq(t, string(mocks.Fixture(tt.request)), string(body))
					} else if r.Body != nil {
						body, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						assert.Empty(t, body)
					}
					return mocks.Responder(tt.status, tt.fixture)(r)
				},
			)
			result, resp, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tt.status, resp.StatusCode)
			assert.Equal(t, "fixture-request", resp.Headers.Get("X-Request-ID"))
			if tt.fixture != "" {
				data, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.fixture)), string(data))
			} else {
				assert.Nil(t, result)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}

func TestAPIErrors(t *testing.T) {
	for _, tt := range contractCases() {
		t.Run(tt.name, func(t *testing.T) {
			for _, failure := range []struct {
				name   string
				status int
			}{{"error_validation", 400}, {"error_unauthorized", 401}, {"error_forbidden", 403}} {
				t.Run(failure.name, func(t *testing.T) {
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder(
						tt.method,
						testutil.BaseURL+tt.path,
						mocks.Responder(failure.status, failure.name),
					)
					result, resp, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
					require.NotNil(t, resp)
					assert.Equal(t, failure.status, resp.StatusCode)
					assert.Equal(t, "fixture-request", resp.Headers.Get("X-Request-ID"))
				})
			}
		})
	}
}

func TestTransportFailure(t *testing.T) {
	for _, tt := range contractCases() {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				tt.method,
				testutil.BaseURL+tt.path,
				httpmock.NewErrorResponder(io.ErrUnexpectedEOF),
			)
			result, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, result)
		})
	}
}

func TestGraphQLErrorsPreserveData(t *testing.T) {
	for _, fixture := range []string{"partial_error", "errors_only"} {
		t.Run(fixture, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				"POST",
				testutil.BaseURL+"/apigateway/workflows/manage/graphql",
				mocks.Responder(200, fixture),
			)
			result, resp, err := NewService(
				transport,
			).Execute(context.Background(), "graphql.workflows", GraphQLRequest{Query: "{ sample }"})
			var gqlErrors GraphQLErrors
			require.ErrorAs(t, err, &gqlErrors)
			require.NotEmpty(t, gqlErrors)
			require.NotNil(t, result)
			assert.Equal(t, 200, resp.StatusCode)
			data, marshalErr := json.Marshal(result)
			require.NoError(t, marshalErr)
			assert.JSONEq(t, string(mocks.Fixture(fixture)), string(data))
		})
	}
}

func TestMalformedResponse(t *testing.T) {
	for _, tt := range contractCases() {
		if tt.fixture == "" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				tt.method,
				testutil.BaseURL+tt.path,
				func(_ *http.Request) (*http.Response, error) {
					response := httpmock.NewStringResponse(tt.status, "{malformed")
					response.Header.Set("Content-Type", "application/json")
					return response, nil
				},
			)
			result, resp, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, result)
			require.NotNil(t, resp)
			assert.Equal(t, tt.status, resp.StatusCode)
		})
	}
}
