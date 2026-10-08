package data_exploration

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/data_exploration/mocks"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
	"time"
)

type contractCase struct {
	name    string
	call    func(*Service, *TimeContext) (any, *interfaces.Response, error)
	invalid func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contractCase {
	t.Helper()
	return []contractCase{{name: "Query", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
		var request QueryRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Query_input"), &request))
		return s.Query(context.Background(), &request, options)
	}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Query(context.Background(), nil, nil) }},
		{name: "Inspect", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request QueryInput
			require.NoError(t, json.Unmarshal(mocks.Fixture("Inspect_input"), &request))
			return s.Inspect(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Inspect(context.Background(), nil, nil) }},
		{name: "GetFilterValues", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request FilterValuesRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetFilterValues_input"), &request))
			return s.GetFilterValues(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetFilterValues(context.Background(), nil, nil)
		}},
		{name: "ListFields", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request FieldsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ListFields_input"), &request))
			return s.ListFields(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListFields(context.Background(), nil, nil)
		}},
		{name: "ListSystemRatings", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request SystemRatingsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ListSystemRatings_input"), &request))
			return s.ListSystemRatings(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListSystemRatings(context.Background(), nil, nil)
		}},
		{name: "ListOrganisationFields", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			return s.ListOrganisationFields(context.Background(), options)
		}, invalid: nil},
		{name: "GetMenu", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request MenuRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetMenu_input"), &request))
			return s.GetMenu(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetMenu(context.Background(), nil, nil) }},
		{name: "ListBreakdownFields", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request BreakdownFieldsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ListBreakdownFields_input"), &request))
			return s.ListBreakdownFields(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListBreakdownFields(context.Background(), nil, nil)
		}},
		{name: "GetBreakdownInsights", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request BreakdownInsightsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetBreakdownInsights_input"), &request))
			return s.GetBreakdownInsights(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetBreakdownInsights(context.Background(), nil, nil)
		}},
		{name: "ListByDurations", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request ByDurationsRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("ListByDurations_input"), &request))
			return s.ListByDurations(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListByDurations(context.Background(), nil, nil)
		}},
		{name: "GetOrganisation", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request OrganisationRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetOrganisation_input"), &request))
			return s.GetOrganisation(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetOrganisation(context.Background(), nil, nil)
		}},
		{name: "GetItemMeta", call: func(s *Service, options *TimeContext) (any, *interfaces.Response, error) {
			var request ItemMetaRequest
			require.NoError(t, json.Unmarshal(mocks.Fixture("GetItemMeta_input"), &request))
			return s.GetItemMeta(context.Background(), &request, options)
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetItemMeta(context.Background(), nil, nil)
		}}}
}
func fixedTimeContext() *TimeContext {
	return &TimeContext{TimeZone: "Europe/London", UTCOffset: -60, ISODateTime: "2026-10-06T05:30:00Z", AppName: "sdk-fixture"}
}
func TestWireContracts(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tc.name+"_request")), string(body))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				assert.Equal(t, "Europe/London", r.Header.Get("x-nxt-waas-timezone"))
				assert.Equal(t, "-60", r.Header.Get("x-nxt-waas-utc-offset"))
				assert.Equal(t, "2026-10-06T05:30:00Z", r.Header.Get("x-nxt-waas-iso-date-time"))
				assert.Equal(t, "sdk-fixture", r.Header.Get("x-nxt-waas-app-name"))
				return mocks.Responder(200, tc.name+"_success")(r)
			})
			result, response, err := tc.call(NewService(transport), fixedTimeContext())
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &expected))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(data))
		})
	}
}

// NQL API logs can have a collection URI without a display collection name.
// Preserve the server's null instead of manufacturing an empty string.
func TestQueryPreservesNullCollectionName(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, "Query_null_collection_success"))
	result, response, err := NewService(transport).Query(context.Background(), &QueryRequest{
		QueryInput: QueryInput{Query: "platform.nql_api_logs during past 7d | summarize total = count()"},
	}, fixedTimeContext())
	require.NoError(t, err)
	require.NotNil(t, response)
	var expected struct {
		Data json.RawMessage `json:"data"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture("Query_null_collection_success"), &expected))
	encoded, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(expected.Data), string(encoded))
}
func TestPartialGraphQLErrors(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, tc.name+"_partial"))
			result, response, err := tc.call(NewService(transport), fixedTimeContext())
			require.Error(t, err)
			var graphqlErrors graphql.GraphQLErrors
			require.ErrorAs(t, err, &graphqlErrors)
			assert.Equal(t, "Fixture GraphQL failure", graphqlErrors[0].Message)
			require.NotNil(t, response)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_partial"), &expected))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(data))
		})
	}
}
func TestNullResults(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(200, tc.name+"_null"))
			result, response, err := tc.call(NewService(transport), nil)
			require.NoError(t, err)
			require.NotNil(t, response)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_null"), &expected))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(data))
		})
	}
}
func TestHTTPAndTransportFailures(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			for _, status := range []int{400, 401, 403, 500} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(status, "http_error"))
				_, response, err := tc.call(NewService(transport), nil)
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, status, response.StatusCode)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewErrorResponder(errors.New("fixture network failure")))
			_, _, err := tc.call(NewService(transport), nil)
			require.Error(t, err)
			for _, body := range []string{`{`, `{"data":null}`, `{"data":[]}`} {
				mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(200, body))
				_, _, err := tc.call(NewService(transport), nil)
				require.Error(t, err)
			}
		})
	}
}
func TestValidationPreventsHTTP(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			s := NewService(transport)
			if tc.invalid != nil {
				_, response, err := tc.invalid(s)
				require.Error(t, err)
				require.Nil(t, response)
			}
			_, response, err := tc.call(s, &TimeContext{ISODateTime: "not-a-date"})
			require.Error(t, err)
			require.Nil(t, response)
			assert.Zero(t, mock.GetTotalCallCount())
		})
	}
}
func TestDefaultTimeHeaders(t *testing.T) {
	for _, tc := range contracts(t) {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			start := time.Now().UTC().Add(-time.Second)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "UTC", r.Header.Get("x-nxt-waas-timezone"))
				assert.Equal(t, "0", r.Header.Get("x-nxt-waas-utc-offset"))
				assert.Equal(t, "sdk", r.Header.Get("x-nxt-waas-app-name"))
				value, err := time.Parse(time.RFC3339, r.Header.Get("x-nxt-waas-iso-date-time"))
				require.NoError(t, err)
				assert.True(t, !value.Before(start))
				assert.True(t, !value.After(time.Now().UTC()))
				return mocks.Responder(200, tc.name+"_success")(r)
			})
			_, _, err := tc.call(NewService(transport), nil)
			require.NoError(t, err)
		})
	}
}

func TestRichQueryAndInspectMetadata(t *testing.T) {
	cases := []contractCase{{name: "Query_rich", call: func(s *Service, o *TimeContext) (any, *interfaces.Response, error) {
		var request QueryRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Query_rich_input"), &request))
		return s.Query(context.Background(), &request, o)
	}}, {name: "Inspect_rich", call: func(s *Service, o *TimeContext) (any, *interfaces.Response, error) {
		var request QueryInput
		require.NoError(t, json.Unmarshal(mocks.Fixture("Inspect_rich_input"), &request))
		return s.Inspect(context.Background(), &request, o)
	}}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
				data, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tc.name+"_request")), string(data))
				return mocks.Responder(200, tc.name+"_success")(r)
			})
			result, _, err := tc.call(NewService(transport), fixedTimeContext())
			require.NoError(t, err)
			var expected struct {
				Data json.RawMessage `json:"data"`
			}
			require.NoError(t, json.Unmarshal(mocks.Fixture(tc.name+"_success"), &expected))
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected.Data), string(data))
		})
	}
}
