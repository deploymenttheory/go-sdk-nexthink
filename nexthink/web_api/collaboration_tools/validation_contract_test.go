package collaboration_tools

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
	t.Run("GetCallInsightsNilRequest", func(t *testing.T) {
		_, response, err := service.GetCallInsights(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCallInsightsInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetCallInsights(ctx, "..", load[CallInsightsRequest](t, "GetCallInsights_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
