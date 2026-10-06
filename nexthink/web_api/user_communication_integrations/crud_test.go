package user_communication_integrations

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/user_communication_integrations/mocks"
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
	return []contractCase{{name: "List", verb: "GET", path: Endpoint + "", form: false, body: false, ack: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.List(context.Background()) }},
		{name: "Get", verb: "GET", path: Endpoint + "/fixture-channel", form: false, body: false, ack: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Get(context.Background(), "fixture-channel")
		}},
		{name: "Create", verb: "POST", path: Endpoint + "", form: false, body: true, ack: false, call: func(s *Service) (any, *interfaces.Response, error) {
			var request IntegrationInput
			require.NoError(t, json.Unmarshal(mocks.Fixture("Create_request"), &request))
			return s.Create(context.Background(), &request)
		}},
		{name: "Update", verb: "PUT", path: Endpoint + "/fixture-channel", form: false, body: true, ack: false, call: func(s *Service) (any, *interfaces.Response, error) {
			var request IntegrationInput
			require.NoError(t, json.Unmarshal(mocks.Fixture("Update_request"), &request))
			return s.Update(context.Background(), "fixture-channel", &request)
		}},
		{name: "Delete", verb: "DELETE", path: Endpoint + "/fixture-channel", form: false, body: false, ack: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.Delete(context.Background(), "fixture-channel")
			return nil, response, err
		}},
		{name: "ListAzureConnectors", verb: "GET", path: Endpoint + "/externals/azureconnectors", form: false, body: false, ack: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListAzureConnectors(context.Background())
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
