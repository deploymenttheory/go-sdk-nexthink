package rule_based_custom_fields

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidationPreventsRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	ctx := context.Background()
	for _, id := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		_, _, err := s.Get(ctx, id)
		require.Error(t, err)
		_, _, err = s.Update(ctx, id, load[FieldInput](t, "Update_input"))
		require.Error(t, err)
		_, err = s.Delete(ctx, id, load[DeleteRequest](t, "Delete_input"))
		require.Error(t, err)
	}
	_, _, err := s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, "id", load[FieldInput](t, "Create_input"))
	require.Error(t, err)
	_, err = s.Delete(ctx, "id", nil)
	require.Error(t, err)
	d := load[DeleteRequest](t, "Delete_input")
	d.InventoryObject = "device/device"
	_, err = s.Delete(ctx, "id", d)
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
