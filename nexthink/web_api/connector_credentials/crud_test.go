package connector_credentials

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/connector_credentials/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var v T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &v))
	return &v
}

type contractCase struct {
	name, verb, path string
	status           int
	body, empty      bool
	call             func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{
		{"List", "GET", EndpointList + "?enabled=true&type=conn_cr", 200, false, false, func(s *Service) (any, *interfaces.Response, error) { return s.List(ctx) }},
		{"ListIDs", "GET", Endpoint + "/config/conn_cr/all", 200, false, false, func(s *Service) (any, *interfaces.Response, error) { return s.ListIDs(ctx) }},
		{"Get", "GET", Endpoint + "/config/conn_cr-42", 200, false, false, func(s *Service) (any, *interfaces.Response, error) { return s.Get(ctx, "conn_cr-42") }},
		{"Create", "POST", Endpoint + "/credential/conn_cr-42", 201, true, false, func(s *Service) (any, *interfaces.Response, error) {
			return s.Create(ctx, "conn_cr-42", load[CredentialInput](t, "Create_request"))
		}},
		{"Update", "POST", Endpoint + "/credential/conn_cr-42", 200, true, false, func(s *Service) (any, *interfaces.Response, error) {
			return s.Update(ctx, "conn_cr-42", load[CredentialInput](t, "Update_request"))
		}},
		{"Delete", "POST", Endpoint + "/credential/conn_cr-42", 201, true, false, func(s *Service) (any, *interfaces.Response, error) { return s.Delete(ctx, "conn_cr-42") }},
	}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				if tt.empty {
					return mocks.Responder(tt.status, "")(r)
				}
				return mocks.Responder(tt.status, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, tt.status, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if tt.empty {
				assert.Nil(t, result)
			} else {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			}
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			for code, name := range map[int]string{401: "unauthorized", 403: "forbidden", 400: "validation"} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "error_"+name))
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
			if !tt.empty {
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
