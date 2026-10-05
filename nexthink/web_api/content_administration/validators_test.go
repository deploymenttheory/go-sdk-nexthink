package content_administration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestGetConfigurationValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetConfiguration(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestListValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).List(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
