package workspace

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPreserveResponseFieldsAndTypedEdits(t *testing.T) {
	raw := `{"id":"fixture-id","future_field":{"enabled":false},"title":null,"enabled":null,"revision":null,"name":"agent"}`
	var value Conversation
	require.NoError(t, json.Unmarshal([]byte(raw), &value))
	assert.Equal(t, "fixture-id", value.ID)
	out, err := json.Marshal(value)
	require.NoError(t, err)
	assert.JSONEq(t, raw, string(out))
	value.ID = "updated-id"
	out, err = json.Marshal(value)
	require.NoError(t, err)
	var decoded map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(out, &decoded))
	assert.JSONEq(t, `"updated-id"`, string(decoded["id"]))
	assert.JSONEq(t, `{"enabled":false}`, string(decoded["future_field"]))
}
func TestRejectMalformedTypedResponse(t *testing.T) {
	var value Conversation
	require.Error(t, json.Unmarshal([]byte(`{"id":42}`), &value))
}

func TestExplicitClearingRequest(t *testing.T) {
	raw := `{"tags":[]}`
	var request ConversationUpdate
	require.NoError(t, json.Unmarshal([]byte(raw), &request))
	encoded, err := json.Marshal(request)
	require.NoError(t, err)
	assert.JSONEq(t, raw, string(encoded))
}

func TestChatMessageEmptyTextAndSteps(t *testing.T) {
	raw := `{"author":"agent","content":[{"type":"text","text":""}],"steps":[],"artifacts":[{"category":"test"}],"feedback":{"rating":"NONE"}}`
	var message ChatMessage
	require.NoError(t, json.Unmarshal([]byte(raw), &message))
	encoded, err := json.Marshal(message)
	require.NoError(t, err)
	assert.JSONEq(t, raw, string(encoded))
}
