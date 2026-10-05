package device_configuration

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestSetProfilesValidation(t *testing.T) {
	for _, value := range []string{"", "null", "[]", "true", "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).SetProfiles(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestSetSettingsValidation(t *testing.T) {
	for _, value := range []string{"", "null", "[]", "true", "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).SetSettings(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
