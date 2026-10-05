package collector_management

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collector_management/mocks"
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
			name:    "GetDownloadLinks",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/links",
			fixture: "download_links_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetDownloadLinks(context.Background()) },
		},
		{
			name:    "GetUpdateConfiguration",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/config",
			fixture: "update_configuration_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.GetUpdateConfiguration(context.Background())
			},
		},
		{
			name:    "GetPlatforms",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/query/platforms",
			fixture: "platforms_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetPlatforms(context.Background()) },
		},
		{
			name:    "GetVersions",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/query/versions",
			fixture: "versions_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetVersions(context.Background()) },
		},
		{
			name:    "GetGroups",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/query/groups",
			fixture: "groups_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetGroups(context.Background()) },
		},
		{
			name:    "GetTargetVersions",
			method:  "GET",
			path:    "/apigateway/api/v1/clt-updater/query/target-versions",
			fixture: "target_versions_success",
			request: "",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call:    func(s *Service) (any, *interfaces.Response, error) { return s.GetTargetVersions(context.Background()) },
		},
		{
			name:    "SetUpdateConfiguration",
			method:  "POST",
			path:    "/apigateway/api/v1/clt-updater/config",
			fixture: "set_update_configuration_success",
			request: "configuration_request",
			status:  200,
			query:   "",
			headers: map[string]string{},
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.SetUpdateConfiguration(
					context.Background(),
					json.RawMessage(mocks.Fixture("configuration_request")),
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
