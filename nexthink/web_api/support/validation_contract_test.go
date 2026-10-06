package support

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
	t.Run("GetProfileInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetProfile(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetPlatformInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetPlatform(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("ListUsersInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.ListUsers(ctx, "..")
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("SearchNilRequest", func(t *testing.T) {
		_, response, err := service.Search(ctx, nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
