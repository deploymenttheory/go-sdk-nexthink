package amplify_ai

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/amplify_ai/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInvalidRequestsDoNotReachTransport(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	ctx := context.Background()
	_, _, err := s.GenerateOrFetchAnalysis(ctx, &AnalysisRequest{DeviceID: "device"})
	require.Error(t, err)
	_, _, err = s.SubmitFeedback(ctx, &FeedbackRequest{ResolutionPlanID: "plan", Feedback: Feedback{Rating: 6}})
	require.Error(t, err)
	_, err = s.PostMetric(ctx, &MetricRequest{})
	require.Error(t, err)
	_, err = s.ReportTicketRetrievalDuration(ctx, &TicketRetrievalDurationRequest{DurationSeconds: -1, Status: "success"})
	require.Error(t, err)
	action := ExecuteActionRequest{"resolutionPlanId": json.RawMessage(`123`), "resolutionStepId": json.RawMessage(`"step"`)}
	_, err = s.ExecuteAction(ctx, &action)
	require.Error(t, err)
	_, err = s.ExecuteUserAction(ctx, &ExecuteUserActionRequest{ResolutionPlanID: "plan"})
	require.Error(t, err)
	_, _, err = s.RefreshResolutionStep(ctx, &RefreshResolutionStepRequest{ID: "step"})
	require.Error(t, err)
	_, _, err = s.UpdateResolutionStep(ctx, &UpdateResolutionStepRequest{ID: "step"})
	require.Error(t, err)
	_, err = s.ResolveTicket(ctx, &ResolveTicketRequest{})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}

func TestQueryIdentifierIsEscaped(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+EndpointGetResolutionPlan+"?resolution_plan_id=plan%26other%3Dvalue", mocks.Responder(200, "GetResolutionPlan_success"))
	_, _, err := NewService(transport).GetResolutionPlan(context.Background(), "plan&other=value")
	require.NoError(t, err)
}

func TestAnalysisAllowsTicketIDWithoutOnTheFly(t *testing.T) {
	require.NoError(t, validateGenerateOrFetchAnalysis(&AnalysisRequest{DeviceID: "device", TicketID: "ticket"}))
}

func TestExecuteActionPreservesWidgetVariant(t *testing.T) {
	request := ExecuteActionRequest{
		"resolutionPlanId": json.RawMessage(`"plan"`), "resolutionStepId": json.RawMessage(`"step"`),
		"remoteActionId": json.RawMessage(`"#example"`), "params": json.RawMessage(`{"input":"value"}`),
		"targets": json.RawMessage(`{"devices":{"collectorUids":["collector"]}}`),
	}
	require.NoError(t, validateExecuteAction(&request))
	data, err := json.Marshal(request)
	require.NoError(t, err)
	assert.JSONEq(t, `{"resolutionPlanId":"plan","resolutionStepId":"step","remoteActionId":"#example","params":{"input":"value"},"targets":{"devices":{"collectorUids":["collector"]}}}`, string(data))
}
