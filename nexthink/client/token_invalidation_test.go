package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
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

type invalidationFixture struct {
	mu            sync.Mutex
	value         string
	invalidated   []string
	unconditional atomic.Int32
}

func (f *invalidationFixture) Token(context.Context) (auth.Token, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.value == "" {
		f.value = "replacement"
	}
	return auth.Token{Value: f.value}, nil
}
func (f *invalidationFixture) InvalidateToken() { f.unconditional.Add(1) }
func (f *invalidationFixture) InvalidateTokenIfCurrent(value string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.invalidated = append(f.invalidated, value)
	if f.value == value {
		f.value = ""
	}
}

func TestUnauthorizedInvalidatesWithoutReplay(t *testing.T) {
	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE", "bytes"} {
		t.Run(method, func(t *testing.T) {
			var requests atomic.Int32
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				w.Header().Set("Content-Type", "application/json")
				if r.Header.Get("Authorization") == "Bearer rejected" {
					w.WriteHeader(401)
					fmt.Fprint(w, `{"message":"expired"}`)
					return
				}
				assert.Equal(t, "Bearer replacement", r.Header.Get("Authorization"))
				fmt.Fprint(w, `{}`)
			}))
			defer server.Close()
			provider := &invalidationFixture{value: "rejected"}
			transport, err := NewTransportWithTokenProvider(server.URL, provider, WithTransport(server.Client().Transport), WithLogger(zap.NewNop()), WithRetryCount(3), WithRetryWaitTime(time.Millisecond))
			require.NoError(t, err)
			transport.client.SetRetryAllowNonIdempotent(true)
			transport.client.AddRetryConditions(func(*resty.Response, error) bool { return true })
			if method == "bytes" {
				_, _, err = transport.GetBytes(context.Background(), "/fixture", nil, nil)
			} else {
				_, err = transport.executeRequest(transport.client.R().SetContext(context.Background()), method, "/fixture")
			}
			require.Error(t, err)
			assert.EqualValues(t, 1, requests.Load(), "401 must not replay even if custom retry policy requests it")
			assert.Equal(t, []string{"rejected"}, provider.invalidated)
			assert.Zero(t, provider.unconditional.Load())
			transport.client.SetRetryCount(0)
			_, err = transport.Get(context.Background(), "/fixture", nil, nil, nil)
			require.NoError(t, err)
			assert.EqualValues(t, 2, requests.Load())
		})
	}
}

func TestForbiddenAndPortalDoNotInvalidateBearer(t *testing.T) {
	for _, portal := range []bool{false, true} {
		t.Run(fmt.Sprint(portal), func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if portal {
					assert.Empty(t, r.Header.Get("Authorization"))
					w.WriteHeader(401)
				} else {
					w.WriteHeader(403)
				}
				fmt.Fprint(w, `{"message":"denied"}`)
			}))
			defer server.Close()
			provider := &invalidationFixture{value: "valid"}
			transport, err := NewTransportWithTokenProvider(server.URL, provider, WithTransport(server.Client().Transport), WithLogger(zap.NewNop()), WithRetryCount(0))
			require.NoError(t, err)
			if portal {
				_, err = transport.PostForm(portalsession.Context(context.Background()), portalsession.Endpoint, nil, map[string]string{"Cookie": "fixture=session"}, nil)
			} else {
				_, err = transport.Get(context.Background(), "/fixture", nil, nil, nil)
			}
			require.Error(t, err)
			assert.Empty(t, provider.invalidated)
			assert.Equal(t, "valid", provider.value)
		})
	}
}

func TestDelayedUnauthorizedPassesRejectedTokenNotCurrentToken(t *testing.T) {
	rejectedStarted, releaseRejected := make(chan struct{}), make(chan struct{})
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/delayed" {
			close(rejectedStarted)
			<-releaseRejected
			w.WriteHeader(401)
		}
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	provider := &invalidationFixture{value: "old"}
	transport, err := NewTransportWithTokenProvider(server.URL, provider, WithTransport(server.Client().Transport), WithLogger(zap.NewNop()), WithRetryCount(0))
	require.NoError(t, err)
	done := make(chan error, 1)
	go func() {
		_, requestErr := transport.Get(context.Background(), "/delayed", nil, nil, nil)
		done <- requestErr
	}()
	<-rejectedStarted
	provider.mu.Lock()
	provider.value = "new"
	provider.mu.Unlock()
	_, err = transport.Get(context.Background(), "/current", nil, nil, nil)
	require.NoError(t, err)
	close(releaseRejected)
	require.Error(t, <-done)
	assert.Equal(t, []string{"old"}, provider.invalidated)
	assert.Equal(t, "new", provider.value)
}

func TestUnauthorizedHandlingPreservesServerRetryPolicy(t *testing.T) {
	var requests atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if requests.Add(1) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	provider := &invalidationFixture{value: "valid"}
	transport, err := NewTransportWithTokenProvider(server.URL, provider, WithTransport(server.Client().Transport), WithLogger(zap.NewNop()), WithRetryCount(1), WithRetryWaitTime(time.Millisecond))
	require.NoError(t, err)
	transport.client.AddRetryConditions(func(resp *resty.Response, _ error) bool {
		return resp != nil && resp.StatusCode() == http.StatusServiceUnavailable
	})
	_, err = transport.Get(context.Background(), "/fixture", nil, nil, nil)
	require.NoError(t, err)
	assert.EqualValues(t, 2, requests.Load())
	assert.Empty(t, provider.invalidated)
}
