package support_checklists

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
	t.Run("ListPropertiesInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.ListProperties(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.List(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.Get(ctx, "..", "fixture")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInvalidchecklistID", func(t *testing.T) {
		_, response, err := service.Get(ctx, "fixture", "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
