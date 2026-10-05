package nql_queries

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_queries/mocks"
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
			name:    "List",
			method:  "GET",
			path:    EndpointList,
			fixture: "list_success",
			status:  200,
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.List(context.Background()) },
		},
		{
			name:    "Create",
			method:  "POST",
			path:    "/apigateway/nqlapi/store/api/v1",
			fixture: "create_success",
			request: "create_request",
			status:  201,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Create(context.Background(), queryRequest("create_request"))
			},
		},
		{
			name:    "Get",
			method:  "GET",
			path:    "/apigateway/nqlapi/store/api/v1/fixture-id",
			fixture: "get_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.Get(context.Background(), "fixture-id") },
		},
		{
			name:    "Update",
			method:  "PUT",
			path:    "/apigateway/nqlapi/store/api/v1/fixture-id",
			fixture: "update_success",
			request: "update_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Update(context.Background(), queryRequest("update_request"))
			},
		},
		{
			name:    "Delete",
			method:  "DELETE",
			path:    "/apigateway/nqlapi/store/api/v1/fixture-id",
			fixture: "",
			request: "",
			status:  204,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				resp, err := s.Delete(context.Background(), "fixture-id")
				return nil, resp, err
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
					assert.Equal(t, "application/json", r.Header.Get("Accept"))
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

func queryRequest(name string) *SaveQueryRequest {
	var req SaveQueryRequest
	if err := json.Unmarshal(mocks.Fixture(name), &req); err != nil {
		panic(err)
	}
	return &req
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
