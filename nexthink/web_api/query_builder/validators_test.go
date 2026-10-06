package query_builder

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDrilldownConditions(t *testing.T) {
	// The UI sends null conditions when the query has no context identifiers.
	require.NoError(t, validateTransformDrilldown(&DrilldownRequest{Query: "devices", Conditions: nil}))
	for _, conditions := range []*Conditions{{}, {Identifiers: []ContextIdentifier{{ObjectURI: "device/device/collector/uid", Identifier: "device.collector.uid"}}, Values: [][]json.RawMessage{{}}}, {Identifiers: []ContextIdentifier{{ObjectURI: "device/device/collector/uid", Identifier: "device.collector.uid"}}, Values: [][]json.RawMessage{{json.RawMessage(`broken`)}}}} {
		require.Error(t, validateTransformDrilldown(&DrilldownRequest{Query: "devices", Conditions: conditions}))
	}
	require.NoError(t, validateTransformDrilldown(&DrilldownRequest{Query: "devices", Conditions: &Conditions{Identifiers: []ContextIdentifier{{ObjectURI: "device/device/collector/uid", Identifier: "device.collector.uid"}}, Values: [][]json.RawMessage{{json.RawMessage(`null`)}}}}))
}
func TestTransformExtensions(t *testing.T) {
	var transform Transform
	const data = `{"objectURI":"user/user","futureOption":{"keep":true}}`
	require.NoError(t, json.Unmarshal([]byte(data), &transform))
	require.Equal(t, "user/user", transform.ObjectURI)
	body, err := json.Marshal(transform)
	require.NoError(t, err)
	require.JSONEq(t, data, string(body))
	transform.AdditionalFields["objectURI"] = json.RawMessage(`"override"`)
	body, err = json.Marshal(transform)
	require.NoError(t, err)
	require.JSONEq(t, data, string(body))
}
