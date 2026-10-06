package action_executions

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
	t.Run("ListRemoteActionsNilRequest", func(t *testing.T) {
		_, response, err := service.ListRemoteActions(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRemoteActionNilRequest", func(t *testing.T) {
		_, response, err := service.GetRemoteAction(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListRemoteActionsForQueryNilRequest", func(t *testing.T) {
		_, response, err := service.ListRemoteActionsForQuery(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ExecuteNilRequest", func(t *testing.T) {
		_, response, err := service.Execute(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDeviceHistoryNilRequest", func(t *testing.T) {
		_, response, err := service.GetDeviceHistory(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDeviceHistoryInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetDeviceHistory(ctx, "..", load[HistoryRequest](t, "GetDeviceHistory_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
