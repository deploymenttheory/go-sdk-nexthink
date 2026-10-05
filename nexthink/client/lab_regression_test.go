package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"go.uber.org/zap"
)

func TestRefreshAfterInvalidationIsConcurrentAndDoesNotReenterMiddleware(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/token" {
			id, secret, ok := r.BasicAuth()
			if !ok || id != "id" || secret != "secret" {
				t.Error("token request must use Basic authentication")
			}
			fmt.Fprintf(w, `{"access_token":"token-%d","token_type":"Bearer","expires_in":900}`, calls.Add(1))
			return
		}
		if r.Header.Get("Authorization") != "Bearer token-2" {
			t.Errorf("API used stale token: %q", r.Header.Get("Authorization"))
		}
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer server.Close()
	c, err := NewTransport("id", "secret", "test", "eu", WithCustomTokenURL(server.URL+"/token"), WithBaseURL(server.URL), WithLogger(zap.NewNop()))
	if err != nil {
		t.Fatal(err)
	}
	c.InvalidateToken()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var result map[string]any
			if _, err := c.Get(ctx, "/data", nil, nil, &result); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if calls.Load() != 2 {
		t.Fatalf("expected initial token plus one refresh, got %d", calls.Load())
	}
	if err := c.RefreshToken(); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 3 {
		t.Fatal("explicit refresh reused cached token")
	}
}

func TestProviderCancellationAndOriginIsolation(t *testing.T) {
	var calls atomic.Int32
	p := auth.TokenProviderFunc(func(ctx context.Context) (auth.Token, error) { calls.Add(1); return auth.Token{}, ctx.Err() })
	c, err := NewTransportWithTokenProvider("https://tenant.example", p, WithLogger(zap.NewNop()))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Get(context.Background(), "https://other.example/data", nil, nil, nil); err == nil {
		t.Fatal("allowed cross-origin token transmission")
	}
	if calls.Load() != 0 {
		t.Fatal("provider invoked before origin validation")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = c.Get(ctx, "/data", nil, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("lost cancellation: %v", err)
	}
}

func TestNumericAndWrappedAPIError(t *testing.T) {
	err := ParseErrorResponse([]byte(`{"code":123,"message":"bad query"}`), 400, "Bad Request", "POST", "/nql", zap.NewNop())
	if GetErrorCode(fmt.Errorf("operation: %w", err)) != "123" || !IsBadRequest(fmt.Errorf("operation: %w", err)) {
		t.Fatalf("lost error code or wrapping: %v", err)
	}
	err = ParseErrorResponse([]byte(`{"code":"INVALID","details":"workflow detail"}`), 400, "Bad Request", "POST", "/workflows", zap.NewNop())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.Details != "workflow detail" {
		t.Fatalf("lost details: %v", err)
	}
}

func FuzzErrorResponse(f *testing.F) {
	f.Add([]byte(`{"code":123,"message":"bad query"}`))
	f.Add([]byte(`{"error":{"code":"BAD","details":{"field":"name"}}}`))
	f.Add([]byte(""))
	f.Fuzz(func(t *testing.T, body []byte) {
		err := ParseErrorResponse(body, 400, "Bad Request", "GET", "/test", zap.NewNop())
		if !IsBadRequest(err) {
			t.Fatal("lost HTTP status")
		}
	})
}
