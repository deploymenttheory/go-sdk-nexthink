package vdi

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
	t.Run("GetSessionTimelineNilRequest", func(t *testing.T) {
		_, response, err := service.GetSessionTimeline(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSessionTimelineInvalidsessionID", func(t *testing.T) {
		_, response, err := service.GetSessionTimeline(ctx, "..", load[TimelineRequest](t, "GetSessionTimeline_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetHypervisorTimelineNilRequest", func(t *testing.T) {
		_, response, err := service.GetHypervisorTimeline(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetHypervisorTimelineInvalidsessionID", func(t *testing.T) {
		_, response, err := service.GetHypervisorTimeline(ctx, "..", load[TimelineRequest](t, "GetHypervisorTimeline_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetGlobalHealthNilRequest", func(t *testing.T) {
		_, response, err := service.GetGlobalHealth(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetGlobalHealthInvalidsessionID", func(t *testing.T) {
		_, response, err := service.GetGlobalHealth(ctx, "..", load[HealthRequest](t, "GetGlobalHealth_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ValidateHostnameNilRequest", func(t *testing.T) {
		_, response, err := service.ValidateHostname(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
