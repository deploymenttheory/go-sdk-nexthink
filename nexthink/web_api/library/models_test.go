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

func TestPackPartialResponseRoundTrip(t *testing.T) {
	for _, fixture := range []string{"InstallPack_success", "Pack_partial"} {
		t.Run(fixture, func(t *testing.T) {
			data := mocks.Fixture(fixture)
			var pack Pack
			require.NoError(t, json.Unmarshal(data, &pack))
			require.Equal(t, "Fixture pack", pack.Name)
			encoded, err := json.Marshal(pack)
			require.NoError(t, err)
			require.JSONEq(t, string(data), string(encoded))
		})
	}
}

func TestPackPartialResponseEdits(t *testing.T) {
	var pack Pack
	require.NoError(t, json.Unmarshal(mocks.Fixture("Pack_partial"), &pack))
	pack.InstallationState = "INSTALLED"
	pack.PackUUID = "newly-added-id"
	encoded, err := json.Marshal(pack)
	require.NoError(t, err)
	var fields map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(encoded, &fields))
	require.JSONEq(t, `"INSTALLED"`, string(fields["installationState"]))
	require.JSONEq(t, `"newly-added-id"`, string(fields["packUuid"]))
	require.JSONEq(t, `{"value":null}`, string(fields["futureMetadata"]))
	require.NotContains(t, fields, "isCustomPack")
}

func TestPackRequestRoundTrip(t *testing.T) {
	// Installation reuses the selected catalog pack as its request body. A
	// response codec must not drop fields from that subsequent write.
	data := mocks.Fixture("InstallPack_request")
	var pack Pack
	require.NoError(t, json.Unmarshal(data, &pack))
	encoded, err := json.Marshal(pack)
	require.NoError(t, err)
	require.JSONEq(t, string(data), string(encoded))
}

func TestPackUnmarshalReplacesPreviousFields(t *testing.T) {
	var pack Pack
	require.NoError(t, json.Unmarshal(mocks.Fixture("Pack_partial"), &pack))
	require.NoError(t, json.Unmarshal(mocks.Fixture("InstallPack_success"), &pack))
	encoded, err := json.Marshal(pack)
	require.NoError(t, err)
	require.JSONEq(t, string(mocks.Fixture("InstallPack_success")), string(encoded))
}
