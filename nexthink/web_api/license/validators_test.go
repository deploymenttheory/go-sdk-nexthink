package license

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestGetFeatureStatusValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetFeatureStatus(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
