package support_insights

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
	t.Run("GetCrashesNilRequest", func(t *testing.T) {
		_, response, err := service.GetCrashes(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCrashesInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetCrashes(ctx, "..", load[InsightsRequest](t, "GetCrashes_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCPUUsageNilRequest", func(t *testing.T) {
		_, response, err := service.GetCPUUsage(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCPUUsageInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetCPUUsage(ctx, "..", load[InsightsRequest](t, "GetCPUUsage_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMemoryUsageNilRequest", func(t *testing.T) {
		_, response, err := service.GetMemoryUsage(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMemoryUsageInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetMemoryUsage(ctx, "..", load[InsightsRequest](t, "GetMemoryUsage_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
