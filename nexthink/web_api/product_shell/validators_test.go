package product_shell

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestGetFlagValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetFlag(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestGetDynamicMenuValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetDynamicMenu(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestValidateClaimsValidation(t *testing.T) {
	for _, value := range []string{"", "null", "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).ValidateClaims(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestPostTelemetryValidation(t *testing.T) {
	for _, value := range []string{"", "null", "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).PostTelemetry(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestOpaqueRequestShapes(t *testing.T) {
	for _, raw := range []string{`{}`, `[]`, `["fixture-claim"]`} {
		require.NoError(t, ValidateRequest(json.RawMessage(raw)))
	}
}
