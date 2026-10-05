package collector_management

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestValidateUpdateConfiguration(t *testing.T) {
	for _, value := range []string{"", "null", "[]", "{}", `{"configRevision":null}`, "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).SetUpdateConfiguration(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
	require.NoError(t, ValidateUpdateConfiguration(json.RawMessage(`{"configRevision":0}`)))
}
