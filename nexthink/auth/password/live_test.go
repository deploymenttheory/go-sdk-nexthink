package password

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/applications"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// TestLivePasswordAuthentication is deliberately opt-in and read-only. It needs
// a password-only local account, curl, and the pinned browser runtime. CI fixture
// jobs never enable it or receive tenant credentials.
func TestLivePasswordAuthentication(t *testing.T) {
	if os.Getenv("NEXTHINK_LIVE_PASSWORD_TEST") != "1" {
		t.Skip("live tenant validation is opt-in")
	}
	origin := fmt.Sprintf("https://%s.%s.nexthink.cloud", os.Getenv("NEXTHINK_INSTANCE"), os.Getenv("NEXTHINK_REGION"))
	p, err := New(origin, Config{Username: os.Getenv("NEXTHINK_USERNAME"), Password: os.Getenv("NEXTHINK_PASSWORD")})
	require.NoError(t, err)
	defer p.Close()
	var logins, refreshes atomic.Int32
	login := p.login
	p.login = func(ctx context.Context, origin string, cfg Config) (session, error) {
		logins.Add(1)
		return login(ctx, origin, cfg)
	}
	base := p.httpClient.Transport
	p.httpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) { refreshes.Add(1); return base.RoundTrip(r) })
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Minute)
	defer cancel()
	initial, err := p.Token(ctx)
	require.NoError(t, err)
	t.Logf("headless local login succeeded; access token lifetime remaining %s; refresh credential present: %t", time.Until(initial.ExpiresAt).Round(time.Second), p.current.RefreshToken != "")
	transport, err := client.NewTransportWithTokenProvider(origin, p, client.WithLogger(zap.NewNop()), client.WithRetryCount(0))
	require.NoError(t, err)
	service := applications.NewService(transport)
	check := func() {
		_, response, err := service.List(ctx, nil)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		token, err := p.Token(ctx)
		require.NoError(t, err)
		cmd := exec.CommandContext(ctx, "curl", "--silent", "--show-error", "--fail", "--max-time", "30", "--config", "-")
		cmd.Stdin = strings.NewReader("url = " + strconv.Quote(origin+applications.Endpoint+"/list?page_number=0&page_size=100&sort_name=asc") + "\nheader = " + strconv.Quote("Authorization: Bearer "+token.Value) + "\nheader = \"Accept: application/json\"\n")
		raw, err := cmd.Output()
		require.NoError(t, err, "curl failed (response suppressed)")
		var sdk, curl any
		require.NoError(t, json.Unmarshal(response.Body, &sdk))
		require.NoError(t, json.Unmarshal(raw, &curl))
		// Avoid assertion helpers that dump tenant data on comparison failure.
		sdkJSON, _ := json.Marshal(sdk)
		curlJSON, _ := json.Marshal(curl)
		require.True(t, string(sdkJSON) == string(curlJSON), "curl/SDK response mismatch (contents suppressed)")
	}
	check()
	t.Log("SDK read matched curl")
	if os.Getenv("NEXTHINK_LIVE_WAIT_FOR_EXPIRY") != "1" {
		return
	}
	delay := time.Until(initial.ExpiresAt.Add(time.Second))
	if delay > 10*time.Minute {
		t.Fatal("token lifetime exceeds bounded expiry test; repeat with a longer explicitly managed observation")
	}
	t.Logf("waiting %s for actual server-issued expiry", delay.Round(time.Second))
	timer := time.NewTimer(max(delay, 0))
	defer timer.Stop()
	select {
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	case <-timer.C:
	}
	check()
	next, err := p.Token(ctx)
	require.NoError(t, err)
	require.True(t, next.Value != initial.Value, "access token did not change after expiry")
	t.Logf("post-expiry curl/SDK comparison passed; browser logins=%d refresh requests=%d", logins.Load(), refreshes.Load())
}
