package nql_editor

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_editor/mocks"
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
			name:    "Validate",
			method:  "POST",
			path:    "/apigateway/nql-editor/nql-ls/validate",
			fixture: "validate_success",
			request: "validate_request",
			status:  200,
			query:   "nqlMode=fixture+mode",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Validate(context.Background(), validationRequest(), "fixture mode")
			},
		},
		{
			name:    "Complete",
			method:  "POST",
			path:    "/apigateway/nql-editor/nql-ls/complete",
			fixture: "complete_success",
			request: "position_request",
			status:  200,
			query:   "nqlMode=fixture+mode",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Complete(context.Background(), positionRequest(), "fixture mode")
			},
		},
		{
			name:    "Hover",
			method:  "POST",
			path:    "/apigateway/nql-editor/nql-ls/hover",
			fixture: "hover_success",
			request: "position_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Hover(context.Background(), positionRequest())
			},
		},
		{
			name:    "Resolve",
			method:  "POST",
			path:    "/apigateway/nql-editor/nql-ls/resolve",
			fixture: "resolve_success",
			request: "resolve_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.Resolve(context.Background(), completionItem())
			},
		},
		{
			name:    "GetHighlighting",
			method:  "GET",
			path:    "/apigateway/nql-editor/highlighting",
			fixture: "highlighting_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetHighlighting(context.Background()) },
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

func validationRequest() *ValidationRequest {
	var req ValidationRequest
	if err := json.Unmarshal(mocks.Fixture("validate_request"), &req); err != nil {
		panic(err)
	}
	return &req
}

func positionRequest() *PositionRequest {
	var req PositionRequest
	if err := json.Unmarshal(mocks.Fixture("position_request"), &req); err != nil {
		panic(err)
	}
	return &req
}

func completionItem() CompletionItem {
	var req CompletionItem
	if err := json.Unmarshal(mocks.Fixture("resolve_request"), &req); err != nil {
		panic(err)
	}
	return req
}

func TestEmptyDiagnosticsAndNullHover(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	mock.RegisterResponder(
		"POST",
		testutil.BaseURL+EndpointValidate,
		mocks.Responder(200, "empty_diagnostics"),
	)
	result, resp, err := s.Validate(context.Background(), validationRequest(), "")
	require.NoError(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	mock.RegisterResponder(
		"POST",
		testutil.BaseURL+EndpointHover,
		mocks.Responder(200, "null_hover"),
	)
	hover, resp, err := s.Hover(context.Background(), positionRequest())
	require.NoError(t, err)
	assert.JSONEq(t, "null", string(hover))
	assert.Equal(t, 200, resp.StatusCode)
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
