package library

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/library/mocks"
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
	name, method, path              string
	hasBody, graphQL, empty, caller bool
	call                            func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "ListContents", method: "GET", path: "/apigateway/library-service/v1/builtincontent", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListContents(ctx) }},
		{name: "ListPacks", method: "GET", path: "/apigateway/library-service/v1/packs", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListPacks(ctx) }},
		{name: "GetContent", method: "GET", path: "/apigateway/library-service/v1/builtincontent/fixture.json", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetContent(ctx, "fixture.json") }},
		{name: "GetCustomContent", method: "GET", path: "/apigateway/library-service/v1/builtincontent/contents/fixture-id/resourceName/fixture-resource", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetCustomContent(ctx, "fixture-id", "fixture-resource")
		}},
		{name: "GetCreateCopyInfo", method: "GET", path: "/apigateway/coad/v1/contents/act/config", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetCreateCopyInfo(ctx, "act") }},
		{name: "InstallContent", method: "POST", path: "/apigateway/library-service/v1/builtincontent/import", hasBody: true, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.InstallContent(ctx, load[ContentInstallationRequest](t, "InstallContent_input"))
		}},
		{name: "InstallCustomContent", method: "POST", path: "/apigateway/library-service/v1/builtincontent/contents/fixture-id/resourceName/fixture-resource/install", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.InstallCustomContent(ctx, "fixture-id", "fixture-resource")
		}},
		{name: "InstallPack", method: "POST", path: "/apigateway/library-service/v1/packs", hasBody: true, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.InstallPack(ctx, load[Pack](t, "InstallPack_input"))
		}},
		{name: "InstallCustomPack", method: "POST", path: "/apigateway/library-service/v1/packs/fixture-pack/install", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.InstallCustomPack(ctx, "fixture-pack") }},
		{name: "UpdateContent", method: "PUT", path: "/apigateway/library-service/v1/builtincontent/update", hasBody: true, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateContent(ctx, load[ContentUpdateRequest](t, "UpdateContent_input"))
		}},
		{name: "UpdatePack", method: "POST", path: "/apigateway/library-service/v1/packs/fixture-pack/install/latest", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdatePack(ctx, "fixture-pack") }},
		{name: "UpdateCustomContent", method: "POST", path: "/apigateway/library-service/v1/builtincontent/contents/fixture-id/resourceName/fixture-resource/install/latest", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateCustomContent(ctx, "fixture-id", "fixture-resource")
		}},
		{name: "GetDependenciesStatus", method: "POST", path: "/apigateway/library-service/v1/builtincontent/dependency-status", hasBody: true, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDependenciesStatus(ctx, load[DependenciesRequest](t, "GetDependenciesStatus_input"))
		}},
		{name: "InstallDependencies", method: "POST", path: "/apigateway/library-service/v1/builtincontent/multi-import", hasBody: true, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.InstallDependencies(ctx, load[DependenciesRequest](t, "InstallDependencies_input"))
		}},
		{name: "GetPack", method: "GET", path: "/apigateway/library-service/v1/packs/fixture.json?packUuid=fixture-pack", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetPack(ctx, "fixture.json", "fixture-pack")
		}},
		{name: "ImportCustomPack", method: "POST", path: "/apigateway/library-service/v1/packs/fixture-id/import", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.ImportCustomPack(ctx, "fixture-id") }},
		{name: "GetPackInstallationStatus", method: "GET", path: "/apigateway/library-service/v1/packs/fixture-id/status", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetPackInstallationStatus(ctx, "fixture-id")
		}},
		{name: "GetContentInstallationStatus", method: "GET", path: "/apigateway/library-service/v1/builtincontent/contents/fixture-id/resourceName/fixture-resource/status", hasBody: false, empty: false, caller: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetContentInstallationStatus(ctx, "fixture-id", "fixture-resource")
		}},
		{name: "DeleteCustomPack", method: "DELETE", path: "/apigateway/library-service/v1/packs/fixture-id", hasBody: false, empty: true, caller: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.DeleteCustomPack(ctx, "fixture-id")
			return nil, response, err
		}},
		{name: "GetLocale", method: "GET", path: "/apigateway/library-service/v1/locale", hasBody: false, empty: false, caller: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLocale(ctx) }}}

}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.caller {
					assert.Equal(t, CallerService, r.Header.Get("nx-caller-service"))
				}
				if !tt.hasBody && r.Body != nil {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.Empty(t, string(body))
				}
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				if tt.empty {
					return mocks.Responder(204, "")(r)
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			if !tt.empty {
				require.NotNil(t, result)
			}
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			if tt.empty {
				assert.Nil(t, result)
				return
			}
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
					if tt.empty && failure.name == "Malformed" {
						t.Skip("empty response is not decoded")
					}
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

func TestReferenceAndQueryEscaping(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/builtincontent/contents/fixture%2Fid/resourceName/resource%3Fname", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, Endpoint+"/builtincontent/contents/fixture%2Fid/resourceName/resource%3Fname", r.URL.EscapedPath())
		return mocks.Responder(200, "GetCustomContent_success")(r)
	})
	_, _, err := NewService(transport).GetCustomContent(context.Background(), "fixture/id", "resource?name")
	require.NoError(t, err)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/packs/fixture.json?packUuid=pack%26other%3Dvalue", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "pack&other=value", r.URL.Query().Get("packUuid"))
		assert.Len(t, r.URL.Query(), 1)
		return mocks.Responder(200, "GetPack_success")(r)
	})
	_, _, err = NewService(transport).GetPack(context.Background(), "fixture.json", "pack&other=value")
	require.NoError(t, err)
}
