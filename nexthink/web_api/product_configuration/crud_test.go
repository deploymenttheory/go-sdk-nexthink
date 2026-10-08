package product_configuration

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/product_configuration/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var value T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &value))
	return &value
}

type contract struct {
	name, verb, path                 string
	body, multipart, download, empty bool
	call                             func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	ctx := context.Background()
	return []contract{{name: "GetInstance", verb: "GET", path: "/apigateway/entity-manager/api/v1/configuration/instances/platform.assist.web-search", body: false, multipart: false, download: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetInstance(ctx, "platform.assist.web-search")
	}},
		{name: "CreateInstance", verb: "POST", path: "/apigateway/entity-manager/api/v1/configuration/instances", body: true, multipart: false, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.CreateInstance(ctx, *load[InstanceConfiguration](t, "CreateInstance_request"))
			return nil, r, e
		}},
		{name: "UpdateInstance", verb: "PUT", path: "/apigateway/entity-manager/api/v1/configuration/instances/platform.assist.web-search", body: true, multipart: false, download: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.UpdateInstance(ctx, "platform.assist.web-search", *load[InstanceConfiguration](t, "UpdateInstance_request"))
			return nil, r, e
		}}}
}
func TestContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					if tt.multipart {
						checkMultipart(t, r)
					} else {
						body, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					}
				}
				status, body, contentType := 200, "", "application/json"
				if tt.empty {
					status = 204
				} else if tt.download {
					require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &body))
					contentType = "text/csv"
				} else {
					body = string(mocks.Fixture(tt.name + "_success"))
				}
				response := httpmock.NewStringResponse(status, body)
				response.Header.Set("Content-Type", contentType)
				response.Header.Set("X-Request-ID", "fixture-request")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if tt.empty {
				assert.Nil(t, result)
				assert.Equal(t, 204, response.StatusCode)
			} else if tt.download {
				var expected string
				require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
				assert.Equal(t, []byte(expected), result)
			} else {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			}
		})
	}
}
func TestHTTPErrors(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewStringResponder(404, string(mocks.Fixture("error"))))
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 404, response.StatusCode)
			assert.JSONEq(t, string(mocks.Fixture("error")), string(response.Body))
		})
	}
}

func checkMultipart(t *testing.T, r *http.Request) {
	t.Helper()
	t.Fatal("unexpected multipart request")
}

func TestValidationDoesNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	_, _, err := service.GetInstance(context.Background(), "")
	require.Error(t, err)
	_, err = service.CreateInstance(context.Background(), nil)
	require.Error(t, err)
	_, err = service.CreateInstance(context.Background(), InstanceConfiguration{"key": json.RawMessage("invalid")})
	require.Error(t, err)
	_, err = service.UpdateInstance(context.Background(), "key", InstanceConfiguration{"other": json.RawMessage("false")})
	require.Error(t, err)
	assert.Equal(t, 0, mock.GetTotalCallCount())
}

func TestMalformedJSON(t *testing.T) {
	for _, tt := range contracts(t) {
		if tt.empty || tt.download {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
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
