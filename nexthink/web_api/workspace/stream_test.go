package workspace

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestChatContract(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+ChatEndpoint, func(r *http.Request) (*http.Response, error) {
		body, e := io.ReadAll(r.Body)
		require.NoError(t, e)
		assert.JSONEq(t, string(mocks.Fixture("Chat_request")), string(body))
		assert.Equal(t, "text/event-stream", r.Header.Get("Accept"))
		assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
		var stream string
		require.NoError(t, json.Unmarshal(mocks.Fixture("Chat_success"), &stream))
		response := httpmock.NewStringResponse(200, stream)
		response.Header.Set("Content-Type", "text/event-stream")
		return response, nil
	})
	result, response, err := NewService(transport).Chat(context.Background(), load[ChatRequest](t, "Chat_request"))
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Len(t, result.Events, 3)
	assert.Contains(t, string(result.Events[1].Data), "SDK")
	assert.Equal(t, 1, mock.GetTotalCallCount())
}
func TestDecodeEvents(t *testing.T) {
	cases := []struct {
		name, body string
		count      int
		failure    bool
	}{
		{"multiline CRLF and comments", "\ufeff: ping\r\nid: cursor\r\nevent: message\r\ndata: {\r\ndata: \"type\":\"finish\"}\r\n\r\n", 1, false},
		{"done terminator", "data: {\"type\":\"finish\"}\n\ndata: [DONE]\n\n", 1, false},
		{"partial then error", "data: {\"type\":\"text-delta\",\"delta\":\"SDK\"}\n\ndata: {\"type\":\"error\",\"errorText\":\"fixture failure\"}\n\n", 2, true},
		{"explicit SSE error", "event: error\ndata: {\"message\":\"fixture failure\"}\n\n", 1, true},
		{"truncated", "data: {\"type\":\"finish\"}\n", 0, true},
		{"malformed", "data: broken\n\n", 0, true},
		{"empty success body", "", 0, true},
		{"non-SSE response", "<html>Login</html>", 0, true},
		{"large event", "data: {\"delta\":\"" + strings.Repeat("x", 128*1024) + "\"}\n\n", 1, false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result, err := decodeEvents([]byte(tt.body))
			assert.Len(t, result.Events, tt.count)
			if tt.failure {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
func TestChatCancellationAndHTTPError(t *testing.T) {
	t.Run("cancellation propagates", func(t *testing.T) {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+ChatEndpoint, func(r *http.Request) (*http.Response, error) { return nil, r.Context().Err() })
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, _, err := NewService(transport).Chat(ctx, load[ChatRequest](t, "Chat_request"))
		require.Error(t, err)
	})
	t.Run("HTTP errors retain metadata", func(t *testing.T) {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+ChatEndpoint, httpmock.NewStringResponder(403, string(mocks.Fixture("error"))))
		result, response, err := NewService(transport).Chat(context.Background(), load[ChatRequest](t, "Chat_request"))
		require.Error(t, err)
		require.Nil(t, result)
		require.NotNil(t, response)
		assert.Equal(t, 403, response.StatusCode)
	})
}

type chatTransportStub struct {
	interfaces.HTTPClient
	response *interfaces.Response
	err      error
}

func (s chatTransportStub) Post(context.Context, string, any, map[string]string, any) (*interfaces.Response, error) {
	return s.response, s.err
}
func TestChatMissingResponseAndPartialTransportError(t *testing.T) {
	request := load[ChatRequest](t, "Chat_request")
	_, _, err := NewService(chatTransportStub{}).Chat(context.Background(), request)
	require.ErrorContains(t, err, "missing chat response")
	response := &interfaces.Response{StatusCode: 200, Body: []byte("data: {\"type\":\"text-delta\",\"delta\":\"SDK\"}\n\ndata: {\"unfinished\":")}
	result, metadata, err := NewService(chatTransportStub{response: response, err: context.Canceled}).Chat(context.Background(), request)
	require.True(t, errors.Is(err, context.Canceled))
	require.ErrorContains(t, err, "truncated")
	require.Same(t, response, metadata)
	require.Len(t, result.Events, 1)
}
