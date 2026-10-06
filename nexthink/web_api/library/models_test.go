package library

import (
	"encoding/json"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/library/mocks"
	"github.com/stretchr/testify/require"
)

func TestDependencyOptionalFilenameRoundTrip(t *testing.T) {
	// Campaign dependencies can omit fileName; action dependencies supply it.
	data := mocks.Fixture("Dependencies_optional_filename")
	var dependencies map[string][]Dependency
	require.NoError(t, json.Unmarshal(data, &dependencies))
	require.Empty(t, dependencies["engage"][0].FileName)
	require.Equal(t, "fixture-action.json", dependencies["act"][0].FileName)

	encoded, err := json.Marshal(dependencies)
	require.NoError(t, err)
	require.JSONEq(t, string(data), string(encoded))
}
