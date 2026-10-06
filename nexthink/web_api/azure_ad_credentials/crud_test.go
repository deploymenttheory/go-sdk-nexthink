package azure_ad_credentials

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/azure_ad_credentials/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type contractCase struct {
	name, verb, path string
	form, body, ack  bool
	call             func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	return []contractCase{{name: "CheckCredentials", verb: "POST", path: Endpoint + "/check-credentials", form: true, body: true, ack: true, call: func(s *Service) (any, *interfaces.Response, error) {
		var request CheckCredentialsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("CheckCredentials_request"), &request))
		response, err := s.CheckCredentials(context.Background(), &request)
		return nil, response, err
	}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					if tt.form {
						assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
						require.NoError(t, r.ParseForm())
						var expected map[string]string
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_request"), &expected))
						assert.Len(t, r.PostForm, len(expected))
						for k, v := range expected {
							assert.Equal(t, []string{v}, r.PostForm[k])
						}
					} else {
						body, err := io.ReadAll(r.Body)
						require.NoError(t, err)
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					}
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if !tt.ack {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			}
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for code, name := range map[int]string{400: "validation", 401: "unauthorized", 403: "forbidden", 404: "not_found"} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "contract_error_"+name))
				_, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, code, response.StatusCode)
				assert.JSONEq(t, string(mocks.Fixture("contract_error_"+name)), string(response.Body))
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			if !tt.ack {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(*http.Request) (*http.Response, error) {
					response := httpmock.NewStringResponse(200, "{broken")
					response.Header.Set("Content-Type", "application/json")
					return response, nil
				})
				_, _, err := tt.call(NewService(transport))
				require.Error(t, err)
			}
		})
	}
}

func TestCheckWithSavedCredentials(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint+"/check-credentials", func(r *http.Request) (*http.Response, error) {
		require.NoError(t, r.ParseForm())
		assert.Equal(t, "tenant", r.PostForm.Get("tenant_id"))
		assert.Equal(t, "client", r.PostForm.Get("client_id"))
		assert.Equal(t, "US_L4", r.PostForm.Get("ms_national_cloud"))
		_, present := r.PostForm["client_secret"]
		assert.False(t, present)
		return httpmock.NewStringResponse(200, ""), nil
	})
	_, err := NewService(transport).CheckCredentials(context.Background(), &CheckCredentialsRequest{TenantID: "tenant", ClientID: "client", NationalCloud: "US_L4"})
	require.NoError(t, err)
}
