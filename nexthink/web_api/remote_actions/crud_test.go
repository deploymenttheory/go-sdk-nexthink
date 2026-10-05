package remote_actions

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/remote_actions/mocks"
)

type contractCase struct {
	name string
	call func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	return []contractCase{
		{name: "Create", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Create(context.Background(), writeRequest(t, "Create"))
		}},
		{name: "Update", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Update(context.Background(), writeRequest(t, "Update"))
		}},
		{name: "Delete", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Delete(context.Background(), "#fixture_action")
		}},

		{name: "Get", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Get(context.Background(), "fixture-uuid")
		}},
		{name: "GetForView", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetForView(context.Background(), "fixture-uuid")
		}},
		{
			name: "GetContentVolume",
			call: func(s *Service) (any, *interfaces.Response, error) { return s.GetContentVolume(context.Background()) },
		},
		{name: "InspectBashScript", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.InspectBashScript(
				context.Background(),
				scriptFixture(t, "InspectBashScript.tar.gz"),
			)
		}},
		{
			name: "InspectPowerShellScript",
			call: func(s *Service) (any, *interfaces.Response, error) {
				return s.InspectPowerShellScript(
					context.Background(),
					scriptFixture(t, "InspectPowerShellScript.ps1"),
				)
			},
		},
		{name: "GetPowerShellSignature", call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetPowerShellSignature(
				context.Background(),
				scriptFixture(t, "GetPowerShellSignature.ps1"),
			)
		}},
	}
}

func scriptFixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile("mocks/" + name)
	require.NoError(t, err)
	return data
}

func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				"POST",
				testutil.BaseURL+Endpoint,
				func(r *http.Request) (*http.Response, error) {
					assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					assert.Empty(t, r.URL.RawQuery)
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					return mocks.Responder(200, tt.name+"_success")(r)
				},
			)
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(actual))
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
				graphQL   bool
			}{
				{"HTTP401", mocks.Responder(401, "error_unauthorized"), false},
				{"HTTP403", mocks.Responder(403, "error_unauthorized"), false},
				{"GraphQL", mocks.Responder(200, "errors_only"), true},
				{"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF), false},
				{"Malformed", func(_ *http.Request) (*http.Response, error) {
					r := httpmock.NewStringResponse(200, "{broken")
					r.Header.Set("Content-Type", "application/json")
					return r, nil
				}, false},
			} {
				t.Run(failure.name, func(t *testing.T) {
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, failure.responder)
					result, _, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
					if failure.graphQL {
						var errors graphql.GraphQLErrors
						require.ErrorAs(t, err, &errors)
					}
				})
			}
		})
	}
}

func TestPartialData(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, "partial_error"))
	result, response, err := NewService(transport).Get(context.Background(), "fixture-uuid")
	var errors graphql.GraphQLErrors
	require.ErrorAs(t, err, &errors)
	require.NotNil(t, result)
	require.NotNil(t, response)
	assert.Equal(t, 200, response.StatusCode)
	data, marshalErr := json.Marshal(result)
	require.NoError(t, marshalErr)
	var expected struct {
		Data json.RawMessage `json:"data"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture("partial_error"), &expected))
	assert.JSONEq(t, string(expected.Data), string(data))
}

func writeRequest(t *testing.T, method string) *RemoteActionInput {
	t.Helper()
	var request struct {
		Variables map[string]json.RawMessage `json:"variables"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture(method+"_request"), &request))
	var result RemoteActionInput
	for _, data := range request.Variables {
		require.NoError(t, json.Unmarshal(data, &result))
	}
	return &result
}

func TestMutationPartialData(t *testing.T) {
	for _, tt := range contractCases(t) {
		if tt.name != "Create" && tt.name != "Update" && tt.name != "Delete" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(
				"POST",
				testutil.BaseURL+Endpoint,
				mocks.Responder(200, tt.name+"_partial_error"),
			)
			result, response, err := tt.call(NewService(transport))
			var gqlErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &gqlErrors)
			require.NotNil(t, response)
			require.NotNil(t, result)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &expected))
			actual, marshalErr := json.Marshal(result)
			require.NoError(t, marshalErr)
			assert.JSONEq(t, string(expected.Data), string(actual))
		})
	}
}

func TestList(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder(
		"GET",
		testutil.BaseURL+EndpointList,
		func(r *http.Request) (*http.Response, error) {
			assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
			assert.Empty(t, r.URL.RawQuery)
			return mocks.Responder(200, "List_success")(r)
		},
	)
	result, response, err := NewService(transport).List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Len(t, result.Rows, 1)
	assert.Equal(t, "fixture-content", result.Rows[0].ContentID)
	assert.Equal(t, "#fixture_action", result.Rows[0].NQLID)
	data, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(mocks.Fixture("List_success")), string(data))
}

func TestListFailures(t *testing.T) {
	for _, status := range []int{400, 401, 403} {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder(
			"GET",
			testutil.BaseURL+EndpointList,
			mocks.Responder(status, "error_unauthorized"),
		)
		result, response, err := NewService(transport).List(context.Background())
		require.Error(t, err)
		require.NotNil(t, response)
		assert.Nil(t, result)
	}
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder(
		"GET",
		testutil.BaseURL+EndpointList,
		httpmock.NewErrorResponder(io.ErrUnexpectedEOF),
	)
	result, _, err := NewService(transport).List(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
}

func TestListMalformedResponse(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder(
		"GET",
		testutil.BaseURL+EndpointList,
		func(_ *http.Request) (*http.Response, error) {
			r := httpmock.NewStringResponse(200, "{broken")
			r.Header.Set("Content-Type", "application/json")
			return r, nil
		},
	)
	result, response, err := NewService(transport).List(context.Background())
	require.Error(t, err)
	require.NotNil(t, response)
	assert.Nil(t, result)
}
