package amplify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/amplify/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var r T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &r))
	return &r
}
func TestDiscoveryContracts(t *testing.T) {
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	cases := []struct {
		name, verb, path string
		body, empty      bool
		call             func(*Service) (any, *interfaces.Response, error)
		invalid          func(*Service) (any, *interfaces.Response, error)
	}{{name: "CreateConfiguration", verb: "POST", path: EndpointCreateConfiguration, body: true, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.CreateConfiguration(ctx, load[ConfigurationRequest](t, "CreateConfiguration_request"))
	}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.CreateConfiguration(ctx, nil) }},
		{name: "UpdateConfiguration", verb: "PUT", path: fmt.Sprintf(EndpointUpdateConfiguration, id) + "?revisionNumber=1", body: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateConfiguration(ctx, id, 1, load[ConfigurationRequest](t, "UpdateConfiguration_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateConfiguration(ctx, id, 1, nil) }},
		{name: "Search", verb: "POST", path: EndpointSearch, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Search(ctx, load[SearchRequest](t, "Search_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Search(ctx, nil) }},
		{name: "SearchDevices", verb: "POST", path: EndpointSearchDevices, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.SearchDevices(ctx, load[SearchRequest](t, "SearchDevices_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.SearchDevices(ctx, nil) }},
		{name: "GetDeviceProperties", verb: "GET", path: fmt.Sprintf(EndpointGetDeviceProperties, id), body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetDeviceProperties(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetDeviceProperties(ctx, "") }},
		{name: "GetDeviceUsers", verb: "GET", path: fmt.Sprintf(EndpointGetDeviceUsers, id), body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetDeviceUsers(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetDeviceUsers(ctx, "") }},
		{name: "GetDevicePackages", verb: "GET", path: fmt.Sprintf(EndpointGetDevicePackages, id), body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetDevicePackages(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetDevicePackages(ctx, "") }},
		{name: "GetUserProperties", verb: "GET", path: fmt.Sprintf(EndpointGetUserProperties, id), body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetUserProperties(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetUserProperties(ctx, "") }},
		{name: "GetUserDevices", verb: "GET", path: fmt.Sprintf(EndpointGetUserDevices, id), body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetUserDevices(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetUserDevices(ctx, "") }},
		{name: "GetConfiguration", verb: "GET", path: EndpointGetConfiguration, body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetConfiguration(ctx) }},
		{name: "PostInsights", verb: "POST", path: EndpointPostInsights, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.PostInsights(ctx, load[InsightsRequest](t, "PostInsights_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.PostInsights(ctx, nil)
			return nil, r, e
		}}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				}
				if tt.empty {
					return httpmock.NewStringResponse(200, ""), nil
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, resp, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode)
			if !tt.empty {
				encoded, err := json.Marshal(result)
				require.NoError(t, err)
				expected := mocks.Fixture(tt.name + "_success")
				assert.JSONEq(t, string(expected), string(encoded))
			}
			for _, status := range []int{400, 401, 403, 409, 422, 500} {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(status, "error"))
				_, resp, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, status, resp.StatusCode)
			}
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err = tt.call(NewService(transport))
			require.Error(t, err)
			if !tt.empty {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
					r := httpmock.NewStringResponse(200, "{broken")
					r.Header.Set("Content-Type", "application/json")
					return r, nil
				})
				_, _, err = tt.call(NewService(transport))
				require.Error(t, err)
			}
			if tt.invalid != nil {
				before := mock.GetTotalCallCount()
				_, resp, err := tt.invalid(NewService(transport))
				require.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, before, mock.GetTotalCallCount())
			}
		})
	}
}

func TestIdentifiersAreEscaped(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/ast/assist-extension-be/api/v2/user/a%2Fb%3Fc/properties", mocks.Responder(200, "GetUserProperties_success"))
	_, _, err := NewService(transport).GetUserProperties(context.Background(), "a/b?c")
	require.NoError(t, err)
}
func TestEmptyKeywordAndAction(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	_, _, err := s.Search(context.Background(), &SearchRequest{Keyword: "  "})
	require.Error(t, err)
	_, _, err = s.SearchDevices(context.Background(), &SearchRequest{})
	require.Error(t, err)
	_, err = s.PostInsights(context.Background(), &InsightsRequest{Action: " "})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}

func TestConfigurationClearAndRevision(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	ctx := context.Background()
	request := &ConfigurationRequest{ITSMConfigList: []WebApplication{}, EnableUsageDataReporting: false}
	mock.RegisterResponder("PUT", testutil.BaseURL+"/apigateway/ast/assist-admin-be/api/v1/admin-conf/config?revisionNumber=0", func(r *http.Request) (*http.Response, error) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		assert.JSONEq(t, `{"itsmConfigList":[],"enableUsageDataReporting":false}`, string(body))
		return mocks.Responder(200, "UpdateConfiguration_success")(r)
	})
	_, _, err := s.UpdateConfiguration(ctx, "config", 0, request)
	require.NoError(t, err)
	before := mock.GetTotalCallCount()
	_, _, err = s.UpdateConfiguration(ctx, "config", -1, request)
	require.Error(t, err)
	_, _, err = s.UpdateConfiguration(ctx, "", 0, request)
	require.Error(t, err)
	_, _, err = s.CreateConfiguration(ctx, &ConfigurationRequest{})
	require.Error(t, err)
	_, _, err = s.CreateConfiguration(ctx, &ConfigurationRequest{ITSMConfigList: []WebApplication{{ITSMURL: " "}}})
	require.Error(t, err)
	_, _, err = s.CreateConfiguration(ctx, &ConfigurationRequest{ITSMConfigList: []WebApplication{{ITSMURL: "https://example.invalid/*", Substitution: "replacement"}}})
	require.Error(t, err)
	assert.Equal(t, before, mock.GetTotalCallCount())
}

func TestExistingConfigurationPreservesServerFields(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+EndpointGetConfiguration, mocks.Responder(200, "GetConfiguration_existing_success"))
	result, _, err := NewService(transport).GetConfiguration(context.Background())
	require.NoError(t, err)
	require.Len(t, result, 1)
	actual, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(mocks.Fixture("GetConfiguration_existing_success")), string(actual))
	assert.JSONEq(t, `[]`, string(result[0]["itsmConfigList"]))
	assert.JSONEq(t, `false`, string(result[0]["enableUsageDataReporting"]))
}
