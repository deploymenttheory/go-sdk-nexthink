package content_administration

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration/mocks"
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
			name:    "GetConfiguration",
			method:  "GET",
			path:    "/apigateway/content-administration/api/v2/config/sample%20%3F%23%25",
			fixture: "configuration_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.GetConfiguration(context.Background(), "sample ?#%")
			},
		},
		{
			name:    "List",
			method:  "GET",
			path:    "/apigateway/content-administration/api/v2/contents/sample%20%3F%23%25",
			fixture: "list_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.List(context.Background(), "sample ?#%") },
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

// System-owned content returns JSON null for ownership/audit fields.
func TestListPreservesNullOwnership(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder(
		"GET",
		testutil.BaseURL+"/apigateway/content-administration/api/v2/contents/remoteactions",
		mocks.Responder(200, "list_nullable_success"),
	)
	result, _, err := NewService(transport).List(context.Background(), "remoteactions")
	require.NoError(t, err)
	require.NotEmpty(t, result.Rows)
	assert.Nil(t, result.Rows[0].ContentOwner)
	assert.Nil(t, result.Rows[0].CreatedBy)
	assert.Nil(t, result.Rows[0].UpdatedBy)
	data, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(mocks.Fixture("list_nullable_success")), string(data))
}

// Workflow listings include metadata absent from other content families.
func TestWorkflowListPreservesMetadata(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/content-administration/api/v2/contents/workflows", mocks.Responder(200, "list_workflows_success"))
	result, _, err := NewService(transport).List(context.Background(), "workflows")
	require.NoError(t, err)
	require.Len(t, result.Rows, 1)
	require.NotNil(t, result.Rows[0].NQLID)
	assert.Equal(t, "#sdk_fixture_workflow", *result.Rows[0].NQLID)
	data, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(mocks.Fixture("list_workflows_success")), string(data))
}
