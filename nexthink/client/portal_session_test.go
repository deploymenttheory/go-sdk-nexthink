package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/portalsession"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"resty.dev/v3"
)

func TestPortalSessionDoesNotForwardOrPersistOtherCredentials(t *testing.T) {
	var providerCalls, portalCalls atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if portalsession.Allowed(r.URL.Path) {
			assert.Empty(t, r.Header.Get("Authorization"))
			assert.Equal(t, "portal-fixture", r.Header.Get("X-Auth-Token"))
			assert.Equal(t, "portal=explicit", r.Header.Get("Cookie"))
			http.SetCookie(w, &http.Cookie{Name: "portal", Value: "response-secret", Path: "/"})
			if portalCalls.Add(1) == 1 {
				w.WriteHeader(503)
				fmt.Fprint(w, `{"message":"retry fixture"}`)
				return
			}
		} else {
			assert.Equal(t, "Bearer browser-fixture", r.Header.Get("Authorization"))
			assert.Empty(t, r.Header.Get("Cookie"))
			assert.Empty(t, r.Header.Get("X-Auth-Token"))
		}
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer server.Close()
	provider := auth.TokenProviderFunc(func(context.Context) (auth.Token, error) {
		providerCalls.Add(1)
		return auth.Token{Value: "browser-fixture"}, nil
	})
	c, err := NewTransportWithTokenProvider(server.URL, provider, WithTransport(server.Client().Transport), WithLogger(zap.NewNop()), WithRetryCount(1), WithRetryWaitTime(time.Millisecond))
	require.NoError(t, err)
	require.Nil(t, c.client.CookieJar())
	c.client.SetAuthToken("stale-client-default")
	c.globalHeaders["Authorization"] = "Bearer stale-global-default"
	// Explicitly enable retry in this test to exercise middleware on both attempts.
	c.client.SetRetryAllowNonIdempotent(true)
	c.client.AddRetryConditions(func(response *resty.Response, err error) bool { return response != nil && response.StatusCode() == 503 })
	var result map[string]bool
	_, err = c.PostForm(portalsession.Context(context.Background()), portalsession.Endpoint, map[string]string{"query": "searchCustomDashboards"}, map[string]string{"Cookie": "portal=explicit", "X-Auth-Token": "portal-fixture"}, &result)
	require.NoError(t, err)
	assert.EqualValues(t, 2, portalCalls.Load())
	assert.Zero(t, providerCalls.Load())
	for _, path := range []string{portalsession.GetAsset, portalsession.SaveAsset} {
		_, err = c.Post(portalsession.Context(context.Background()), path, map[string]string{"name": "menu-logo"}, map[string]string{"Cookie": "portal=explicit", "X-Auth-Token": "portal-fixture"}, &result)
		require.NoError(t, err)
	}
	assert.Zero(t, providerCalls.Load())
	_, err = c.Get(context.Background(), "/normal", nil, nil, &result)
	require.NoError(t, err)
	assert.EqualValues(t, 1, providerCalls.Load())
}
func TestPortalSessionMarkerCannotDisableAuthOnOtherRoutes(t *testing.T) {
	var providerCalls atomic.Int32
	c, err := NewTransportWithTokenProvider("https://fixture.invalid", auth.TokenProviderFunc(func(context.Context) (auth.Token, error) {
		providerCalls.Add(1)
		return auth.Token{Value: "fixture"}, nil
	}), WithLogger(zap.NewNop()), WithRetryCount(0))
	require.NoError(t, err)
	for _, path := range []string{"/other", "/PortalServlet?query=other", "/PortalServlet/", "/%50ortalServlet", "/PortalApiServlet/nxportalbranding/other", "/PortalApiServlet/nxportalbranding/getAsset?x=1"} {
		_, err = c.PostForm(portalsession.Context(context.Background()), path, nil, nil, nil)
		require.ErrorContains(t, err, "restricted")
	}
	_, _, err = c.GetBytes(portalsession.Context(context.Background()), "/other", nil, nil)
	require.ErrorContains(t, err, "restricted")
	assert.Zero(t, providerCalls.Load())
}
