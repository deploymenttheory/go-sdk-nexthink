package appearance

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/appearance/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func legacySaveFixture(t *testing.T) *SaveLegacyAssetRequest {
	var r SaveLegacyAssetRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("LegacySave_request"), &r))
	return &r
}
func TestLegacyAppearanceContracts(t *testing.T) {
	session := &auth.PortalSession{Cookie: "portal=fixture", XAuthToken: "portal-token"}
	for _, tt := range []struct {
		name, path string
		call       func(*Service) (any, *interfaces.Response, error)
	}{
		{"LegacyGet", EndpointLegacyGetAsset, func(s *Service) (any, *interfaces.Response, error) {
			return s.GetLegacyAsset(context.Background(), session, MenuLogo)
		}},
		{"LegacySave", EndpointLegacySaveAsset, func(s *Service) (any, *interfaces.Response, error) {
			return s.SaveLegacyAsset(context.Background(), session, legacySaveFixture(t))
		}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			s := NewService(transport)
			mock.RegisterResponder("POST", testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Empty(t, r.Header.Get("Authorization"))
				assert.Equal(t, "portal=fixture", r.Header.Get("Cookie"))
				assert.Equal(t, "portal-token", r.Header.Get("X-Auth-Token"))
				assert.Contains(t, r.Header.Get("Content-Type"), "application/json")
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, resp, err := tt.call(s)
			require.NoError(t, err)
			require.NotNil(t, resp)
			body, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(body))
			for _, status := range []int{400, 401, 403, 500} {
				mock.RegisterResponder("POST", testutil.BaseURL+tt.path, mocks.Responder(status, "error"))
				_, resp, err := tt.call(s)
				require.Error(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, status, resp.StatusCode)
			}
			mock.RegisterResponder("POST", testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err = tt.call(s)
			require.Error(t, err)
			mock.RegisterResponder("POST", testutil.BaseURL+tt.path, mocks.Responder(200, "Legacy_status_failure"))
			result, resp, err = tt.call(s)
			require.Error(t, err)
			require.NotNil(t, result)
			require.NotNil(t, resp)
		})
	}
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	mock.RegisterResponder("POST", testutil.BaseURL+EndpointLegacySaveAsset, mocks.Responder(200, "LegacySave_failure"))
	result, resp, err := s.SaveLegacyAsset(context.Background(), session, legacySaveFixture(t))
	require.ErrorContains(t, err, "invalid_svg")
	require.NotNil(t, result.Result.Error)
	require.NotNil(t, resp)
	before := mock.GetTotalCallCount()
	_, _, err = s.GetLegacyAsset(context.Background(), nil, MenuLogo)
	require.Error(t, err)
	_, _, err = s.GetLegacyAsset(context.Background(), session, "../other")
	require.Error(t, err)
	for _, r := range []*SaveLegacyAssetRequest{nil, {}, {Name: MenuLogo, ID: json.RawMessage(`{}`), Version: json.RawMessage(`1`)}, {Name: MenuLogo, ID: json.RawMessage(`1`), Version: json.RawMessage(`1`), Filename: "fixture.png", Blob: "data:image/png;base64,eA=="}} {
		_, _, err = s.SaveLegacyAsset(context.Background(), session, r)
		require.Error(t, err)
	}
	assert.Equal(t, before, mock.GetTotalCallCount())
}
func TestLegacyErrorCodeTruthiness(t *testing.T) {
	for _, value := range []string{`null`, `false`, `0`, `""`} {
		assert.False(t, legacyCodeTruthy(json.RawMessage(value)))
	}
	for _, value := range []string{`1`, `"error"`, `true`} {
		assert.True(t, legacyCodeTruthy(json.RawMessage(value)))
	}
}
