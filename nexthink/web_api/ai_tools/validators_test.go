package ai_tools

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ai_tools/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCustomCreationRequiresHashPrefixedNQLID(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	var request ToolRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("Create_request"), &request))
	request.NQLID = "custom_tool_without_hash"
	_, response, err := NewService(transport).Create(context.Background(), &request)
	require.ErrorContains(t, err, "start with #")
	assert.Nil(t, response)
	assert.Zero(t, mock.GetTotalCallCount())
	// Library tools have unprefixed IDs and remain valid for updates.
	require.NoError(t, validateToolRequest(&request))
}

func TestGoalDatesUseServerTimestamps(t *testing.T) {
	var request GoalRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("CreateGoal_request"), &request))
	require.NoError(t, validateGoalRequest(&request))
	request.Timeline.StartDate = "2099-01-01"
	require.ErrorContains(t, validateGoalRequest(&request), "RFC3339")
	request.Timeline.StartDate = "2099-01-01T00:00:00Z"
	request.Timeline.EndDate = json.RawMessage(`"not-a-date"`)
	require.ErrorContains(t, validateGoalRequest(&request), "endDate")
}

func TestGoalProgressSupportsClearingManualStatus(t *testing.T) {
	var request GoalRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("UpdateGoal_request"), &request))
	request.Progress = json.RawMessage(`null`)
	data, err := json.Marshal(request)
	require.NoError(t, err)
	var document map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(data, &document))
	assert.Equal(t, "null", string(document["progress"]))
	request.Timeline.EndDate = json.RawMessage(`null`)
	require.NoError(t, validateGoalRequest(&request))
	request.Progress = nil
	data, err = json.Marshal(request)
	require.NoError(t, err)
	document = nil
	require.NoError(t, json.Unmarshal(data, &document))
	assert.NotContains(t, document, "progress")
}

func TestValidators(t *testing.T) {
	assert.Error(t, validateRevision(-1))
	assert.NoError(t, validateRevision(0))
	for _, id := range []string{"", " ", ".", "..", "a/b", "a?x", "a#x", "a\\b"} {
		assert.Error(t, validateID(id))
	}
	for _, lang := range []string{"", "en", "ja"} {
		assert.NoError(t, validateLanguage(lang))
	}
	assert.Error(t, validateLanguage("en-US"))
	assert.Error(t, validateToolRequest(&ToolRequest{}))
	assert.Error(t, validateToolRequest(&ToolRequest{Name: "Tool", NQLID: "INVALID"}))
	assert.Error(t, validateCopilotRequest(&CopilotRequest{}))
	assert.Error(t, validateCredentialsRequest(&CredentialsRequest{}))
	assert.Error(t, validateToolInsightsRequest(&ToolInsightsRequest{}))
	assert.Error(t, validateGoalRequest(&GoalRequest{}))
}

func TestGoalScopePresence(t *testing.T) {
	for _, wire := range []string{
		`{"individualTools":[{"toolId":"11111111-1111-4111-8111-111111111111"}]}`,
		`{"policyGroups":["POLICY_GROUP_ALL"]}`,
		`{"policyGroups":[],"individualTools":[]}`,
	} {
		var scope GoalToolScope
		require.NoError(t, json.Unmarshal([]byte(wire), &scope))
		result, err := json.Marshal(scope)
		require.NoError(t, err)
		assert.JSONEq(t, wire, string(result))
	}
}
