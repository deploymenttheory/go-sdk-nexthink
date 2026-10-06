package autopilot

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvalidRequestsDoNotSend(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	ctx := context.Background()
	var err error
	_, err = s.UpdateWebSearch(ctx, nil)
	require.Error(t, err)
	_, err = s.UpdateFilesystemTool(ctx, nil)
	require.Error(t, err)
	_, err = s.UpdateAgentName(ctx, nil)
	require.Error(t, err)
	_, _, err = s.SaveSettings(ctx, nil)
	require.Error(t, err)
	_, _, err = s.ReplaceWebSearchDomains(ctx, nil)
	require.Error(t, err)
	_, _, err = s.GetApproval(ctx, " ")
	require.Error(t, err)
	_, _, err = s.CreateApproval(ctx, nil)
	require.Error(t, err)
	_, _, err = s.UpdateApproval(ctx, "fixture-id", nil)
	require.Error(t, err)
	_, _, err = s.UpdateApproval(ctx, " ", load[ApprovalInput](t, "UpdateApproval_request"))
	require.Error(t, err)
	_, _, err = s.CreateTicket(ctx, nil)
	require.Error(t, err)
	_, _, err = s.GetConversation(ctx, " ")
	require.Error(t, err)
	_, _, err = s.GetRecommendationConversationIDs(ctx, " ")
	require.Error(t, err)
	_, err = s.UpdateAgentActionInputs(ctx, "fixture-id", nil)
	require.Error(t, err)
	_, err = s.UpdateAgentActionInputs(ctx, " ", load[AgentActionInputsRequest](t, "UpdateAgentActionInputs_request"))
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
func TestExplicitZeroAndEmptyValues(t *testing.T) {
	require.NoError(t, validateUpdateWebSearch(&BooleanValue{Value: false}))
	require.NoError(t, validateReplaceWebSearchDomains(&WebSearchDomainsRequest{Domains: []WebSearchDomainInput{}}))
	require.Error(t, validateReplaceWebSearchDomains(&WebSearchDomainsRequest{}))
	require.NoError(t, validateUpdateAgentActionInputs(&AgentActionInputsRequest{Inputs: []AgentActionInput{}}))
	require.Error(t, validateUpdateAgentActionInputs(&AgentActionInputsRequest{}))
	require.Error(t, validateUpdateAgentName(&StringValue{}))
	require.Error(t, validateCreateApproval(&Approval{}))
	require.Error(t, validateCreateTicket(&TicketRequest{}))
	require.Error(t, validateSaveSettings(&Settings{Revision: -1}))
}
