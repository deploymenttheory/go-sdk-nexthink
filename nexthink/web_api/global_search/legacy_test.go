package global_search

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/global_search/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func legacyFixture(t *testing.T) *LegacyDashboardSearchRequest {
	t.Helper()
	var r LegacyDashboardSearchRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("LegacySearch_request"), &r))
	return &r
}
func legacyCases(t *testing.T) []struct {
	name, query string
	call        func(*Service) (any, *interfaces.Response, error)
} {
	session := &PortalSession{Cookie: "portal=fixture", XAuthToken: "fixture-portal-token"}
	return []struct {
		name, query string
		call        func(*Service) (any, *interfaces.Response, error)
	}{{"LegacyAuth", "getAuthToken", func(s *Service) (any, *interfaces.Response, error) {
		return s.GetLegacyAuthToken(context.Background(), session)
	}}, {"LegacySearch", "searchCustomDashboards", func(s *Service) (any, *interfaces.Response, error) {
		return s.SearchLegacyDashboards(context.Background(), session, legacyFixture(t))
	}}}
}
func TestLegacyWireContracts(t *testing.T) {
	for _, tt := range legacyCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/PortalServlet", func(r *http.Request) (*http.Response, error) {
				assert.Empty(t, r.Header.Get("Authorization"))
				assert.Equal(t, "portal=fixture", r.Header.Get("Cookie"))
				assert.Equal(t, "fixture-portal-token", r.Header.Get("X-Auth-Token"))
				assert.True(t, strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded"))
				b, e := io.ReadAll(r.Body)
				require.NoError(t, e)
				form, e := url.ParseQuery(string(b))
				require.NoError(t, e)
				assert.Equal(t, tt.query, form.Get("query"))
				var expected map[string]string
				fixture := "LegacyAuth_request"
				if tt.name == "LegacySearch" {
					fixture = "LegacySearch_form"
				}
				require.NoError(t, json.Unmarshal(mocks.Fixture(fixture), &expected))
				assert.Len(t, form, len(expected))
				for key, value := range expected {
					assert.Equal(t, []string{value}, form[key])
				}
				if tt.name == "LegacyAuth" {
					assert.Len(t, form, 1)
				} else {
					assert.Len(t, form, 5)
					assert.Equal(t, "fixture", form.Get("search"))
					assert.Equal(t, "6", form.Get("maxRoleBasedResults"))
				}
				response := httpmock.NewStringResponse(200, string(mocks.Fixture(tt.name+"_success")))
				response.Header.Set("Content-Type", "application/json")
				response.Header.Set("X-Auth-Token", "rotated-fixture")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			assert.Equal(t, "rotated-fixture", response.Headers.Get("X-Auth-Token"))
			encoded, e := json.Marshal(result)
			require.NoError(t, e)
			assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(encoded))
		})
	}
}
func TestLegacyFailuresAndInBandErrors(t *testing.T) {
	for _, tt := range legacyCases(t) {
		for _, status := range []int{200, 401, 403, 500} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/PortalServlet", func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(status, string(mocks.Fixture("Legacy_error")))
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			if status == 200 {
				require.NotNil(t, result)
				var portalErr *PortalStatus
				require.ErrorAs(t, err, &portalErr)
				assert.Equal(t, 401, portalErr.Code)
			} else {
				assert.Nil(t, result)
			}
		}
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+"/PortalServlet", httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
		_, _, err := tt.call(NewService(transport))
		require.Error(t, err)
	}
}
func TestLegacyValidationBeforeRequest(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	for _, session := range []*PortalSession{nil, {}, {XAuthToken: "bad\r\nheader"}, {Cookie: "bad\x00cookie"}} {
		_, _, err := s.GetLegacyAuthToken(context.Background(), session)
		require.Error(t, err)
	}
	session := &PortalSession{Cookie: "portal=fixture"}
	_, _, err := s.SearchLegacyDashboards(context.Background(), session, nil)
	require.Error(t, err)
	_, _, err = s.SearchLegacyDashboards(context.Background(), session, &LegacyDashboardSearchRequest{Search: "fixture"})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
	encoded, err := json.Marshal(session)
	require.NoError(t, err)
	assert.NotContains(t, string(encoded), "fixture")
}
func TestLegacyInvalidEnvelope(t *testing.T) {
	for _, body := range []string{`{}`, `{"resultStatus":{"code":0},"result":null}`, `{"resultStatus":{"code":0},"result":{}}`, `{broken`} {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+"/PortalServlet", func(r *http.Request) (*http.Response, error) {
			response := httpmock.NewStringResponse(200, body)
			response.Header.Set("Content-Type", "application/json")
			return response, nil
		})
		_, _, err := NewService(transport).GetLegacyAuthToken(context.Background(), &PortalSession{Cookie: "portal=fixture"})
		require.Error(t, err)
	}
}
