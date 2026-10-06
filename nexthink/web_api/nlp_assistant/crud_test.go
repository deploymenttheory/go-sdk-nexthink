package nlp_assistant

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nlp_assistant/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func fixtureRequest(t *testing.T) *ChatRequest {
	t.Helper()
	var r ChatRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("request"), &r))
	return &r
}
func TestWireContract(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	var fixture struct {
		Stream string `json:"stream"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture("success"), &fixture))
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
		assert.Equal(t, "text/event-stream", r.Header.Get("Accept"))
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		assert.JSONEq(t, string(mocks.Fixture("request")), string(body))
		response := httpmock.NewStringResponse(200, fixture.Stream)
		response.Header.Set("Content-Type", "text/event-stream")
		response.Header.Set("X-Request-ID", "fixture-request")
		return response, nil
	})
	result, response, err := NewService(transport).Chat(context.Background(), fixtureRequest(t))
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Len(t, result.Events, 2)
	require.NotNil(t, response)
	assert.Equal(t, fixture.Stream, string(response.Body))
	assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
}
func TestHTTPAndTransportErrors(t *testing.T) {
	for _, status := range []int{401, 403, 429, 500} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewStringResponder(status, string(mocks.Fixture("error"))))
			result, response, err := NewService(transport).Chat(context.Background(), fixtureRequest(t))
			require.Error(t, err)
			assert.Nil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
		})
	}
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
	_, _, err := NewService(transport).Chat(context.Background(), fixtureRequest(t))
	require.Error(t, err)
}
func TestNilRequestDoesNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	_, response, err := NewService(transport).Chat(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	assert.Zero(t, mock.GetTotalCallCount())
}
