package workflow_executions

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidationPreventsHTTP(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	t.Run("GetTimelineNilRequest", func(t *testing.T) {
		_, response, err := service.GetTimeline(ctx, "fixture", "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTimelineInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetTimeline(ctx, "..", "fixture", load[TimelineOptions](t, "GetTimeline_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTimelineInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetTimeline(ctx, "fixture", "..", load[TimelineOptions](t, "GetTimeline_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTimelineV2InvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetTimelineV2(ctx, "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTimelineV2InvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetTimelineV2(ctx, "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListActivitiesInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.ListActivities(ctx, "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListActivitiesInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.ListActivities(ctx, "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRemoteActionDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetRemoteActionDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRemoteActionDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetRemoteActionDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRemoteActionDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetRemoteActionDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCustomFieldsDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetCustomFieldsDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCustomFieldsDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetCustomFieldsDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCustomFieldsDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetCustomFieldsDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCampaignDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetCampaignDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCampaignDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetCampaignDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCampaignDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetCampaignDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetFunctionDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetFunctionDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetFunctionDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetFunctionDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetFunctionDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetFunctionDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMessageDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetMessageDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMessageDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetMessageDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMessageDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetMessageDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSAPIDetailsInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetSAPIDetails(ctx, "..", "fixture", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSAPIDetailsInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetSAPIDetails(ctx, "fixture", "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSAPIDetailsInvalidthinkletID", func(t *testing.T) {
		_, response, err := service.GetSAPIDetails(ctx, "fixture", "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListWorkflowsNilRequest", func(t *testing.T) {
		_, response, err := service.ListWorkflows(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWorkflowInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetWorkflow(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.Get(ctx, "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.Get(ctx, "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetHistoryInvalidworkflowID", func(t *testing.T) {
		_, response, err := service.GetHistory(ctx, "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetHistoryInvalidexecutionID", func(t *testing.T) {
		_, response, err := service.GetHistory(ctx, "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ExecuteNilRequest", func(t *testing.T) {
		_, response, err := service.Execute(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ExecuteNQLNilRequest", func(t *testing.T) {
		_, response, err := service.ExecuteNQL(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
