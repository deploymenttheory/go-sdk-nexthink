package support_insights

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResponsePreservesFeatureFields(t *testing.T) {
	var value GetCrashesResponse
	input := `{"futureField":{"items":[0,false,null,"fixture"]}}`
	require.NoError(t, json.Unmarshal([]byte(input), &value))
	actual, err := json.Marshal(value)
	require.NoError(t, err)
	assert.JSONEq(t, input, string(actual))
	require.Contains(t, value.AdditionalFields, "futureField")
}
