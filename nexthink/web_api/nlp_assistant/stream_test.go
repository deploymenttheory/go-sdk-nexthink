package nlp_assistant

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSSEMultilineCommentsAndErrors(t *testing.T) {
	raw := ":keepalive\r\nid: fixture-id\r\nevent: status\r\ndata: {\r\ndata: \"content\": \"working\"}\r\n\r\nevent: error\ndata: {\"message\":\"denied\",\"status\":403}\n\n"
	r, err := decode([]byte(raw))
	require.Error(t, err)
	require.Len(t, r.Events, 2)
	assert.Equal(t, "fixture-id", r.Events[0].ID)
	data, e := r.Events[0].Decode()
	require.NoError(t, e)
	assert.Equal(t, "working", data.Content)
	var server *StreamError
	require.ErrorAs(t, err, &server)
	assert.Equal(t, 403, server.Status)
}
func TestSSELargeJSONAndTruncatedTail(t *testing.T) {
	body, _ := json.Marshal(EventData{Content: strings.Repeat("x", 128*1024)})
	raw := "data: " + string(body) + "\n\nevent: final-response\ndata: {\"content\":\"unfinished\"}"
	r, err := decode([]byte(raw))
	require.ErrorContains(t, err, "truncated")
	require.Len(t, r.Events, 1)
	data, e := r.Events[0].Decode()
	require.NoError(t, e)
	assert.Len(t, data.Content, 128*1024)
}
func TestInvalidEventJSON(t *testing.T) {
	r, err := decode([]byte("data: {bad}\n\n"))
	require.Error(t, err)
	assert.Empty(t, r.Events)
}
func TestValidation(t *testing.T) {
	r := fixtureRequest(t)
	require.NoError(t, validateRequest(r))
	r.ID = ""
	require.Error(t, validateRequest(r))
	r = fixtureRequest(t)
	r.Source = "other"
	require.Error(t, validateRequest(r))
	r = fixtureRequest(t)
	r.ConversationMode = "other"
	require.Error(t, validateRequest(r))
	r = fixtureRequest(t)
	r.Messages = nil
	require.Error(t, validateRequest(r))
	r = fixtureRequest(t)
	r.Messages[0].Author = "other"
	require.Error(t, validateRequest(r))
	r = fixtureRequest(t)
	r.Messages[0].Content = " "
	require.Error(t, validateRequest(r))
}

func TestSSECarriageReturnDelimiters(t *testing.T) {
	r, err := decode([]byte("data: {\"content\":\"fixture\"}\r\r"))
	require.NoError(t, err)
	require.Len(t, r.Events, 1)
}

func TestUnexpectedJSONResponseIsNotSilentlyAccepted(t *testing.T) {
	_, err := decode([]byte(`{"error":"server returned JSON instead of SSE"}`))
	require.ErrorContains(t, err, "no SSE data events")
}
