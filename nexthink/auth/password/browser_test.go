package password

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestBrowserRejectsProtocolDebug(t *testing.T) {
	for _, name := range []string{"DEBUGP", "PWDEBUG"} {
		t.Run(name, func(t *testing.T) {
			t.Setenv(name, "1")
			_, err := browserLoginAt(context.Background(), "http://127.0.0.1", "http://127.0.0.1", Config{})
			if !errors.Is(err, ErrBrowserUnavailable) {
				t.Fatalf("expected safe startup rejection: %v", err)
			}
		})
	}
}

func TestBrowserAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := browserLoginAt(ctx, "http://127.0.0.1", "http://127.0.0.1", Config{})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation: %v", err)
	}
}

func TestBrowserOrigin(t *testing.T) {
	got, err := loginOrigin("https://example.eu.nexthink.cloud")
	if err != nil || got != "https://example-login.eu.nexthink.cloud" {
		t.Fatalf("unexpected derived origin %q: %v", got, err)
	}
	for _, invalid := range []string{"http://example.nexthink.cloud", "https://example.nexthink.cloud.evil.test", "https://user@example.nexthink.cloud", "https://example.nexthink.cloud:443", "https://example.nexthink.cloud/path", "https://example.nexthink.cloud?x=y"} {
		if _, err := loginOrigin(invalid); err == nil {
			t.Errorf("accepted invalid origin %q", invalid)
		}
	}
	if sameOrigin("https://user@example.nexthink.cloud", "https://example.nexthink.cloud") || sameOrigin("https://example.nexthink.cloud.evil.test", "https://example.nexthink.cloud") {
		t.Fatal("accepted foreign or credential-bearing origin")
	}
}

func TestBrowserToken(t *testing.T) {
	now := time.Now()
	s, err := browserToken([]byte(`{"access_token":"fixture-access","refresh_token":"fixture-refresh","token_type":"Bearer","expires_in":3600}`), "client_id=product-shell&grant_type=authorization_code", "https://example-login.nexthink.cloud/oauth2/default/v1/token", now)
	if err != nil || s.AccessToken != "fixture-access" || s.RefreshToken != "fixture-refresh" || s.ClientID != "product-shell" || !s.ExpiresAt.Equal(now.Add(time.Hour)) {
		t.Fatalf("token decoding failed: %v", err)
	}
	for _, body := range []string{`{}`, `{"access_token":"secret","expires_in":0,"token_type":"Bearer"}`, `{"access_token":"secret","expires_in":3600,"token_type":"other"}`, `{"access_token":"secret\r\nheader","expires_in":3600,"token_type":"Bearer"}`, `{"access_token":"secret","expires_in":9223372036854775807,"token_type":"Bearer"}`} {
		if _, err := browserToken([]byte(body), "", "", now); !errors.Is(err, ErrLoginFailed) || strings.Contains(err.Error(), "secret") {
			t.Fatalf("invalid response not rejected safely: %v", err)
		}
	}
}

func TestBrowserChallenges(t *testing.T) {
	for _, tc := range []struct {
		text string
		want error
	}{
		{"Your account is locked", ErrAccountLocked},
		{"Please change your password", ErrPasswordChangeRequired},
		{"Enter your verification code", ErrMFARequired},
		{"Unable to sign in", ErrInvalidCredentials},
		{"Welcome to Nexthink", nil},
	} {
		if !errors.Is(loginChallenge(tc.text), tc.want) {
			t.Errorf("wrong classification for %q", tc.text)
		}
	}
}

func TestBrowserIntegration(t *testing.T) {
	if os.Getenv("NEXTHINK_BROWSER_TEST") != "1" {
		t.Skip("set NEXTHINK_BROWSER_TEST=1 after installing the pinned Playwright Chromium runtime")
	}
	for _, tc := range []struct {
		name, result string
		want         error
	}{
		{"combined", "success", nil}, {"staged", "success", nil}, {"telemetry", "success", nil}, {"iframe", "success", nil},
		{"invalid", "Unable to sign in", ErrInvalidCredentials},
		{"mfa", "Enter your verification code", ErrMFARequired},
		{"password_change", "Please change your password", ErrPasswordChangeRequired},
		{"locked", "Your account is locked", ErrAccountLocked},
		{"sso", "sso", ErrSSOUnsupported},
		{"cancel", "pending", context.DeadlineExceeded},
		{"changed", "changed", context.DeadlineExceeded},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var credentialCalls atomic.Int32
			mux := http.NewServeMux()
			mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				if tc.name == "changed" {
					_, _ = w.Write([]byte(`<h1>Unrecognized login page</h1>`))
					return
				}
				if tc.name == "iframe" {
					_, _ = w.Write([]byte(`<iframe src="https://login.okta.com/discovery/iframe.html"></iframe><h1 style="display:none">Enter verification code</h1>`))
				}
				if tc.name == "telemetry" {
					_, _ = w.Write([]byte(`<script>fetch("https://telemetry.example.invalid/events",{method:"POST",body:"fixture"}).catch(()=>{});</script>`))
				}
				passwordField := `<input type="password" name="password">`
				if tc.name == "staged" {
					passwordField = ""
				}
				_, _ = fmt.Fprintf(w, `<form><input name="username">%s<button type="submit">Sign in</button></form><script>
document.querySelector('form').onsubmit=async(e)=>{e.preventDefault();let p=document.querySelector('input[type=password]');
if(!p){document.querySelector('form').insertAdjacentHTML('afterbegin','<input type="password" name="password">');return;}
let r=await fetch('/credentials',{method:'POST',body:new URLSearchParams(new FormData(e.target))});
let result=await r.text(); if(result==='success'){await fetch('/oauth2/default/v1/token',{method:'POST',body:'client_id=product-shell&grant_type=authorization_code'});}
else if(result==='sso'){location.href='https://idp.example.invalid/login';}
else if(result!=='pending'){document.body.innerHTML='<h1>'+result+'</h1>';}}
</script>`, passwordField)
			})
			mux.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
				credentialCalls.Add(1)
				if err := r.ParseForm(); err != nil || r.Form.Get("username") != "fixture-user" || r.Form.Get("password") != "fixture-password" {
					http.Error(w, "invalid fixture", 400)
					return
				}
				_, _ = w.Write([]byte(tc.result))
			})
			mux.HandleFunc("/oauth2/default/v1/token", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"access_token":"fixture-access","refresh_token":"fixture-refresh","token_type":"Bearer","expires_in":3600}`))
			})
			server := httptest.NewServer(mux)
			defer server.Close()
			timeout := 30 * time.Second
			if tc.name == "cancel" || tc.name == "changed" {
				timeout = 6 * time.Second
			}
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			s, err := browserLoginAt(ctx, server.URL, server.URL, Config{Username: "fixture-user", Password: "fixture-password"})
			if !errors.Is(err, tc.want) {
				t.Fatalf("wanted %v, got %v", tc.want, err)
			}
			if tc.want == nil && s.AccessToken != "fixture-access" {
				t.Fatal("missing captured token")
			}
			expectedCalls := int32(1)
			if tc.name == "changed" {
				expectedCalls = 0
			}
			if credentialCalls.Load() != expectedCalls {
				t.Fatalf("expected %d submissions, got %d", expectedCalls, credentialCalls.Load())
			}
			if err != nil && strings.Contains(err.Error(), "fixture-password") {
				t.Fatal("password exposed")
			}
		})
	}
}

func TestBrowserGate(t *testing.T) {
	ctx := context.Background()
	if err := acquireBrowser(ctx); err != nil {
		t.Fatal(err)
	}
	waiting, cancel := context.WithTimeout(ctx, 20*time.Millisecond)
	defer cancel()
	if err := acquireBrowser(waiting); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("queued cancellation: %v", err)
	}
	<-browserGate
	if err := acquireBrowser(ctx); err != nil {
		t.Fatal(err)
	}
	<-browserGate
}

func TestBrowserParallelAcquisition(t *testing.T) {
	if os.Getenv("NEXTHINK_BROWSER_TEST") != "1" {
		t.Skip("requires installed Playwright runtime")
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/default/v1/token" {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"fixture-access","token_type":"Bearer","expires_in":3600}`))
			return
		}
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<script>fetch('/oauth2/default/v1/token',{method:'POST',body:'client_id=product-shell'});</script>`))
	}))
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	done := make(chan error, 2)
	for range 2 {
		go func() { _, err := browserLoginAt(ctx, server.URL, server.URL, Config{}); done <- err }()
	}
	for range 2 {
		if err := <-done; err != nil {
			t.Errorf("concurrent login: %v", err)
		}
	}
}
