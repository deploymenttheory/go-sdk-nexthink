package data_exploration

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequiredInputs(t *testing.T) {
	cases := []any{&QueryRequest{}, &QueryInput{}, &FieldsRequest{}, &SystemRatingsRequest{}, &FilterValuesRequest{}, &BreakdownFieldsRequest{}, &BreakdownInsightsRequest{}, &ByDurationsRequest{}, &OrganisationRequest{}, &ItemMetaRequest{}, &MenuRequest{}}
	for _, request := range cases {
		require.Error(t, validateRequest(request))
	}
}
func TestBreakdownVariablesExcludeUIFlags(t *testing.T) {
	var request BreakdownInsightsRequest
	require.NoError(t, json.Unmarshal([]byte(`{"nqlVariables":{"query":"devices | summarize count = count()","limit":10,"showDrilldown":true,"fetchTotalCount":true,"includeGlobalTime":true,"uniqueKey":"fixture"}}`), &request))
	data, err := json.Marshal(request)
	require.NoError(t, err)
	assert.JSONEq(t, `{"nqlVariables":{"query":"devices | summarize count = count()","limit":10}}`, string(data))
}
func TestInvalidEmbeddedJSON(t *testing.T) {
	_, err := variables(&QueryInput{Query: "devices", Filters: []json.RawMessage{json.RawMessage(`{`)}})
	require.Error(t, err)
}
func TestPartialTimeContextDefaults(t *testing.T) {
	headers, err := contextHeaders(&TimeContext{AppName: "fixture"})
	require.NoError(t, err)
	assert.Equal(t, "UTC", headers["x-nxt-waas-timezone"])
	assert.Equal(t, "0", headers["x-nxt-waas-utc-offset"])
	assert.Equal(t, "fixture", headers["x-nxt-waas-app-name"])
}
