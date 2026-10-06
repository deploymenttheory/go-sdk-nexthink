package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

type redactionRoundTripper func(*http.Request) (*http.Response, error)

func (f redactionRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func TestRequestErrorRedactsQueryCredentials(t *testing.T) {
	const secret = "fixture-query-credential"
	networkErr := errors.New("fixture network failure")
	core, logs := observer.New(zap.DebugLevel)
	rt := redactionRoundTripper(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, secret, r.URL.Query().Get("dd-api-key"))
		return nil, networkErr
	})
	c, err := NewTransportWithTokenProvider("https://fixture.invalid", auth.StaticToken("fixture-bearer", time.Time{}), WithLogger(zap.New(core)), WithTransport(rt), WithRetryCount(0))
	require.NoError(t, err)
	_, err = c.PostWithQuery(context.Background(), "/observability", map[string]string{"dd-api-key": secret}, map[string]string{"fixture": "event"}, nil, nil)
	require.Error(t, err)
	assert.NotContains(t, err.Error(), secret)
	var urlErr *url.Error
	require.ErrorAs(t, err, &urlErr)
	assert.Equal(t, "https://fixture.invalid/observability", urlErr.URL)
	require.ErrorIs(t, err, networkErr)
	for _, entry := range logs.All() {
		assert.NotContains(t, fmt.Sprint(entry.ContextMap()), secret)
	}
}
func TestRequestPathRedactsOnlyDiagnostics(t *testing.T) {
	assert.Equal(t, "/fixture", requestLogPath("/fixture?dd-api-key=secret&harmless=value"))
	cause := errors.New("cause")
	wrapped := fmt.Errorf("outer: %w", &url.Error{Op: "Post", URL: "https://fixture.invalid/path?token=secret", Err: cause})
	redacted := redactRequestError(wrapped)
	assert.NotContains(t, redacted.Error(), "secret")
	require.ErrorIs(t, redacted, cause)
	var urlErr *url.Error
	require.ErrorAs(t, redacted, &urlErr)
	assert.Equal(t, "https://fixture.invalid/path", urlErr.URL)
}

func TestAPIErrorKeepsStatusAndRedactsEndpointQuery(t *testing.T) {
	const secret = "fixture-sensitive-query"
	core, logs := observer.New(zap.DebugLevel)
	rt := redactionRoundTripper(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, secret, r.URL.Query().Get("dd-api-key"))
		return &http.Response{StatusCode: 403, Status: "403 Forbidden", Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"code":"FORBIDDEN","message":"fixture denied"}`)), Request: r}, nil
	})
	c, err := NewTransportWithTokenProvider("https://fixture.invalid", auth.StaticToken("fixture-bearer", time.Time{}), WithLogger(zap.New(core)), WithTransport(rt), WithRetryCount(0))
	require.NoError(t, err)
	_, err = c.Post(context.Background(), "/observability?dd-api-key="+secret, nil, nil, nil)
	require.Error(t, err)
	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, 403, apiErr.StatusCode)
	assert.Equal(t, "FORBIDDEN", apiErr.Code)
	assert.Equal(t, "/observability", apiErr.Endpoint)
	assert.NotContains(t, err.Error(), secret)
	for _, entry := range logs.All() {
		assert.NotContains(t, fmt.Sprint(entry.ContextMap()), secret)
	}
}
