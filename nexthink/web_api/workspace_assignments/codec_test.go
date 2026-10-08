package workspace_assignments

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPreserveResponseFieldsAndTypedEdits(t *testing.T) {
	raw := `{"id":"fixture-id","future_field":{"enabled":false},"title":null,"enabled":null,"revision":null,"name":"agent"}`
	var value Assignment
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
	var value Assignment
	require.Error(t, json.Unmarshal([]byte(`{"id":42}`), &value))
}
