package cci_insights

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/cci_insights/mocks"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type contractCase struct {
	name    string
	call    func(*Service) (any, *interfaces.Response, error)
	invalid func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contractCase {
	t.Helper()
	return []contractCase{{name: "GetBinaryInsights", call: func(s *Service) (any, *interfaces.Response, error) {
		var r GetBinaryInsightsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("GetBinaryInsights_input"), &r))
		return s.GetBinaryInsights(context.Background(), &r)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetBinaryInsights(context.Background(), nil)
	}},
		{name: "GetDiagnosticBinaryInsights", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDiagnosticBinaryInsightsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDiagnosticBinaryInsights_input"), &r))
			return s.GetDiagnosticBinaryInsights(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetDiagnosticBinaryInsights(context.Background(), nil)
		}},
		{name: "GetDataset", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetDatasetRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetDataset_input"), &r))
			return s.GetDataset(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetDataset(context.Background(), nil) }}}
}
func TestWireContracts(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
				b, e := io.ReadAll(r.Body)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture(tc.name+"_request")), string(b))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				return mocks.Responder(200, tc.name+"_success")(r)
			})
			result, response, err := tc.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			var envelope struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &envelope))
			got, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(envelope.Data), string(got))
		})
	}
}
func TestGraphQLErrorsRetainPartialData(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			var body map[string]any
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &body))
			body["errors"] = []any{map[string]any{"message": "fixture failure"}}
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewJsonResponderOrPanic(200, body))
			result, response, err := tc.call(NewService(transport))
			require.Error(t, err)
			var graphErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &graphErrors)
			require.Equal(t, "fixture failure", graphErrors[0].Message)
			require.NotNil(t, response)
			require.NotNil(t, result)
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			for _, status := range []int{400, 401, 403, 500} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(status, `{"message":"fixture failure"}`))
				_, response, err := tc.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, status, response.StatusCode)
			}
			for _, body := range []string{`{`, `{"data":null}`, `{"data":[]}`} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(200, body))
				_, _, err := tc.call(NewService(transport))
				require.Error(t, err)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewErrorResponder(errors.New("fixture transport failure")))
			_, _, err := tc.call(NewService(transport))
			require.Error(t, err)
		})
	}
}
func TestNilRequestsMakeNoHTTPCall(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			_, response, err := tc.invalid(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, response)
			assert.Zero(t, mock.GetTotalCallCount())
		})
	}
}

func TestNullableRootResults(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, tc.name+"_null"))
			result, response, err := tc.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			var envelope struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_null"), &envelope))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(envelope.Data), string(data))
		})
	}
}
