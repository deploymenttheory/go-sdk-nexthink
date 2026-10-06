package support

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResponsePreservesFeatureFields(t *testing.T) {
	var value GetProfileResponseName
	input := `{"futureField":{"items":[0,false,null,"fixture"]}}`
	require.NoError(t, json.Unmarshal([]byte(input), &value))
	actual, err := json.Marshal(value)
	require.NoError(t, err)
	assert.JSONEq(t, input, string(actual))
	require.Contains(t, value.AdditionalFields, "futureField")
}

func TestResponsePreservesKnownNull(t *testing.T) {
	var result GetPlatformResponse
	require.NoError(t, json.Unmarshal([]byte(`{"platform":null}`), &result))
	data, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, `{"platform":null}`, string(data))
	platform := "mac"
	result.Platform = &platform
	data, err = json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, `{"platform":"mac"}`, string(data))
}
