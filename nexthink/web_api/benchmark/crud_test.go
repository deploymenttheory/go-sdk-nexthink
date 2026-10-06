package benchmark

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/benchmark/mocks"
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
	return []contractCase{{name: "GetBinaryProductMaps", call: func(s *Service) (any, *interfaces.Response, error) {
		var r GetBinaryProductMapsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("GetBinaryProductMaps_input"), &r))
		return s.GetBinaryProductMaps(context.Background(), &r)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetBinaryProductMaps(context.Background(), nil)
	}},
		{name: "GetProfile", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetProfileRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetProfile_input"), &r))
			return s.GetProfile(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetProfile(context.Background(), nil) }},
		{name: "GetProfileSearchItems", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetProfileSearchItemsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetProfileSearchItems_input"), &r))
			return s.GetProfileSearchItems(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetProfileSearchItems(context.Background(), nil)
		}},
		{name: "GetProfileVersions", call: func(s *Service) (any, *interfaces.Response, error) {
			var r GetProfileVersionsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetProfileVersions_input"), &r))
			return s.GetProfileVersions(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetProfileVersions(context.Background(), nil)
		}},
		{name: "LookupBenchmark", call: func(s *Service) (any, *interfaces.Response, error) {
			var r LookupBenchmarkRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("LookupBenchmark_input"), &r))
			return s.LookupBenchmark(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.LookupBenchmark(context.Background(), nil)
		}},
		{name: "ProductProperties", call: func(s *Service) (any, *interfaces.Response, error) {
			var r ProductPropertiesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ProductProperties_input"), &r))
			return s.ProductProperties(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ProductProperties(context.Background(), nil)
		}},
		{name: "ProfileProperties", call: func(s *Service) (any, *interfaces.Response, error) {
			var r ProfilePropertiesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ProfileProperties_input"), &r))
			return s.ProfileProperties(context.Background(), &r)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ProfileProperties(context.Background(), nil)
		}}}
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
