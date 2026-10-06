package password

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func newTestProvider(t *testing.T) *Provider {
	t.Helper()
	p, err := New("https://fixture.eu.nexthink.cloud", Config{Username: "fixture-user", Password: "private-password", LoginTimeout: time.Second})
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, p.Close()) })
	return p
}

func validSession(value string) session {
	return session{AccessToken: value, ExpiresAt: time.Now().Add(time.Hour), TokenURL: tokenURL("https://fixture.eu.nexthink.cloud"), ClientID: "product-shell"}
}

func response(status int, raw string) *http.Response {
	return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(raw)), Header: make(http.Header)}
}

func authFixture(t *testing.T, name string) string {
	t.Helper()
	raw, err := os.ReadFile("testdata/" + name + ".json")
	require.NoError(t, err)
	return string(raw)
}

func TestConfigAndRedaction(t *testing.T) {
	for _, cfg := range []Config{{}, {Username: "a"}, {Password: "b"}, {Username: "a", Password: "b", LoginTimeout: -1}, {Username: "a", Password: "b", BrowserProxy: "https://host/?secret=1"}} {
		require.Error(t, cfg.Validate())
	}
	cfg := Config{Username: "private-user", Password: "private-password", BrowserProxy: "http://secret:proxy@localhost:8080"}
	require.NoError(t, cfg.Validate())
	raw, err := json.Marshal(cfg)
	require.NoError(t, err)
	require.JSONEq(t, `{}`, string(raw))
	for _, value := range []string{fmt.Sprintf("%v", cfg), fmt.Sprintf("%+v", &cfg), fmt.Sprintf("%#v", cfg)} {
		require.NotContains(t, value, "private")
		require.NotContains(t, value, "secret")
	}
	for _, origin := range []string{"http://fixture.eu.nexthink.cloud", "https://fixture.eu.nexthink.cloud.evil.test", "https://user@fixture.eu.nexthink.cloud", "https://fixture.eu.nexthink.cloud/path", "https://fixture.eu.nexthink.cloud:443", "https://fixture.xx.nexthink.cloud"} {
		_, err := New(origin, cfg)
		require.Error(t, err)
	}
}

func TestConcurrentLoginCachedAndStaleInvalidation(t *testing.T) {
	p := newTestProvider(t)
	var calls atomic.Int32
	p.login = func(context.Context, string, Config) (session, error) {
		calls.Add(1)
		time.Sleep(10 * time.Millisecond)
		return validSession(fmt.Sprintf("token-%d", calls.Load())), nil
	}
	var wg sync.WaitGroup
	for range 20 {
		wg.Go(func() {
			v, err := p.Token(context.Background())
			require.NoError(t, err)
			require.Equal(t, "token-1", v.Value)
		})
	}
	wg.Wait()
	require.EqualValues(t, 1, calls.Load())
	p.InvalidateTokenIfCurrent("unrelated-token")
	_, err := p.Token(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, calls.Load())
	p.InvalidateTokenIfCurrent("token-1")
	v, err := p.Token(context.Background())
	require.NoError(t, err)
	require.Equal(t, "token-2", v.Value)
	p.InvalidateTokenIfCurrent("token-1")
	_, err = p.Token(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 2, calls.Load())
}

func TestPermanentChallengesDoNotResubmitCredentials(t *testing.T) {
	for _, problem := range []error{ErrInvalidCredentials, ErrMFARequired, ErrAccountLocked, ErrSSOUnsupported, ErrPasswordChangeRequired} {
		t.Run(problem.Error(), func(t *testing.T) {
			p := newTestProvider(t)
			calls := 0
			p.login = func(context.Context, string, Config) (session, error) { calls++; return session{}, problem }
			for range 3 {
				_, err := p.Token(context.Background())
				require.ErrorIs(t, err, problem)
			}
			require.Equal(t, 1, calls)
		})
	}
}

func TestWaiterCancellationAndClose(t *testing.T) {
	p := newTestProvider(t)
	started := make(chan struct{})
	finished := make(chan error, 1)
	p.login = func(ctx context.Context, _ string, _ Config) (session, error) {
		close(started)
		<-ctx.Done()
		return session{}, ctx.Err()
	}
	go func() { _, err := p.Token(context.Background()); finished <- err }()
	<-started
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := p.Token(ctx)
	require.ErrorIs(t, err, context.Canceled)
	require.NoError(t, p.Close())
	require.ErrorIs(t, <-finished, ErrClosed)
	_, err = p.Token(context.Background())
	require.ErrorIs(t, err, ErrClosed)
}

func TestLoginCancellationCanRecover(t *testing.T) {
	p := newTestProvider(t)
	p.login = func(ctx context.Context, _ string, _ Config) (session, error) {
		<-ctx.Done()
		return session{}, ctx.Err()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := p.Token(ctx)
	require.ErrorIs(t, err, context.DeadlineExceeded)
	p.login = func(context.Context, string, Config) (session, error) { return validSession("recovered"), nil }
	v, err := p.Token(context.Background())
	require.NoError(t, err)
	require.Equal(t, "recovered", v.Value)
}

func TestRefreshRotationAndWireContract(t *testing.T) {
	p := newTestProvider(t)
	logins := 0
	p.login = func(context.Context, string, Config) (session, error) {
		logins++
		s := validSession("initial")
		s.RefreshToken = "refresh-1"
		return s, nil
	}
	_, err := p.Token(context.Background())
	require.NoError(t, err)
	var grants []string
	p.httpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		require.Equal(t, "POST", r.Method)
		require.Equal(t, tokenURL(p.origin), r.URL.String())
		require.Empty(t, r.Header.Get("Authorization"))
		require.NoError(t, r.ParseForm())
		require.Equal(t, "refresh_token", r.Form.Get("grant_type"))
		require.Equal(t, "product-shell", r.Form.Get("client_id"))
		require.Empty(t, r.Form.Get("password"))
		grants = append(grants, r.Form.Get("refresh_token"))
		return response(200, authFixture(t, "token_success")), nil
	})
	for range 2 {
		p.InvalidateToken()
		v, err := p.Token(context.Background())
		require.NoError(t, err)
		require.Equal(t, "renewed", v.Value)
	}
	require.Equal(t, []string{"refresh-1", "refresh-2"}, grants)
	require.Equal(t, 1, logins)
}

func TestRefreshFailurePolicy(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		status     int
		wantLogin  bool
	}{
		{"revoked", authFixture(t, "refresh_rejected"), 400, true},
		{"server error", `{"error":"server_error"}`, 503, false},
		{"client error", `{"error":"invalid_client"}`, 400, false},
		{"redirect", `{}`, 302, false},
		{"malformed", `private-secret`, 200, false},
		{"bad type", `{"token_type":"Basic","access_token":"private-secret","expires_in":300}`, 200, false},
		{"expired", `{"token_type":"Bearer","access_token":"private-secret","expires_in":0}`, 200, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := newTestProvider(t)
			logins := 0
			p.login = func(context.Context, string, Config) (session, error) {
				logins++
				s := validSession("initial")
				s.RefreshToken = "refresh-1"
				return s, nil
			}
			_, err := p.Token(context.Background())
			require.NoError(t, err)
			p.InvalidateToken()
			p.httpClient.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return response(tc.status, tc.body), nil })
			_, err = p.Token(context.Background())
			if tc.wantLogin {
				require.NoError(t, err)
				require.Equal(t, 2, logins)
			} else {
				require.ErrorIs(t, err, ErrRenewalFailed)
				require.NotContains(t, err.Error(), "private-secret")
				require.Equal(t, 1, logins)
			}
		})
	}
}

func TestNoRefreshTokenAndExpiry(t *testing.T) {
	p := newTestProvider(t)
	calls := 0
	p.login = func(context.Context, string, Config) (session, error) { calls++; return validSession("fresh"), nil }
	_, err := p.Token(context.Background())
	require.NoError(t, err)
	p.mu.Lock()
	p.renewBefore = time.Now().Add(-time.Second)
	p.mu.Unlock()
	_, err = p.Token(context.Background())
	require.NoError(t, err)
	require.Equal(t, 2, calls)
}

func TestRefreshNetworkErrorRedaction(t *testing.T) {
	p := newTestProvider(t)
	p.httpClient.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) { return nil, errors.New("private-network-secret") })
	s := validSession("token")
	s.RefreshToken = "private-refresh"
	_, err := p.refresh(context.Background(), s)
	require.ErrorIs(t, err, ErrRenewalFailed)
	require.NotContains(t, err.Error(), "private")
}

func TestConcurrentWaitersShareTimeoutFailure(t *testing.T) {
	p := newTestProvider(t)
	p.cfg.LoginTimeout = 250 * time.Millisecond
	var calls atomic.Int32
	started := make(chan struct{})
	p.login = func(ctx context.Context, _ string, _ Config) (session, error) {
		if calls.Add(1) == 1 {
			close(started)
		}
		<-ctx.Done()
		return session{}, ctx.Err()
	}
	results := make(chan error, 20)
	go func() { _, err := p.Token(context.Background()); results <- err }()
	<-started
	for range 19 {
		go func() { _, err := p.Token(context.Background()); results <- err }()
	}
	for range 20 {
		require.ErrorIs(t, <-results, context.DeadlineExceeded)
	}
	require.EqualValues(t, 1, calls.Load(), "waiters must not resubmit credentials after a shared timeout")
}

func TestNewWithWrappedDefaultTransport(t *testing.T) {
	original := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		t.Fatal("provider must own its refresh transport")
		return nil, errors.New("unexpected roundtrip")
	})
	defer func() { http.DefaultTransport = original }()
	p := newTestProvider(t)
	owned, ok := p.httpClient.Transport.(*http.Transport)
	require.True(t, ok)
	require.NotNil(t, owned.DialContext)
	require.Equal(t, 10*time.Second, owned.TLSHandshakeTimeout)
	require.True(t, owned.ForceAttemptHTTP2)
}

type cancelOnCloseBody struct {
	io.Reader
	cancel context.CancelFunc
}

func (b cancelOnCloseBody) Close() error { b.cancel(); return nil }

func TestCancellationAfterRefreshResponsePreservesRotation(t *testing.T) {
	p := newTestProvider(t)
	p.login = func(context.Context, string, Config) (session, error) {
		s := validSession("initial")
		s.RefreshToken = "refresh-1"
		return s, nil
	}
	_, err := p.Token(context.Background())
	require.NoError(t, err)
	p.InvalidateToken()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var refreshes atomic.Int32
	p.httpClient.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		refreshes.Add(1)
		return &http.Response{StatusCode: 200, Header: make(http.Header), Body: cancelOnCloseBody{Reader: strings.NewReader(`{"token_type":"Bearer","access_token":"renewed","refresh_token":"refresh-2","expires_in":300}`), cancel: cancel}}, nil
	})
	_, err = p.Token(ctx)
	require.ErrorIs(t, err, context.Canceled)
	value, err := p.Token(context.Background())
	require.NoError(t, err)
	require.Equal(t, "renewed", value.Value)
	require.EqualValues(t, 1, refreshes.Load())
	p.mu.Lock()
	currentRefresh := p.current.RefreshToken
	p.mu.Unlock()
	require.Equal(t, "refresh-2", currentRefresh)
}

func TestCancellationDoesNotDiscardPermanentFailure(t *testing.T) {
	p := newTestProvider(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	calls := 0
	p.login = func(context.Context, string, Config) (session, error) {
		calls++
		cancel()
		return session{}, ErrInvalidCredentials
	}
	_, err := p.Token(ctx)
	require.ErrorIs(t, err, ErrInvalidCredentials)
	_, err = p.Token(context.Background())
	require.ErrorIs(t, err, ErrInvalidCredentials)
	require.Equal(t, 1, calls)
}
