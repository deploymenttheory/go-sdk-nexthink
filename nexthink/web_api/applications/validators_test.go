package applications

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
		_, _, err = s.Update(ctx, id, load[ApplicationInput](t, "Update_input"))
		require.Error(t, err)
		_, _, err = s.Delete(ctx, id, 1)
		require.Error(t, err)
	}
	for _, req := range []*ApplicationInput{nil, {}, {Name: "fixture", Category: "STANDARD"}, {Name: "fixture", Category: "STANDARD", Revision: -1, Desktop: &DesktopConfiguration{}}} {
		_, _, err := s.Create(ctx, req)
		require.Error(t, err)
	}
	_, _, err := s.Update(ctx, "id", load[ApplicationInput](t, "Create_input"))
	require.Error(t, err)
	_, _, err = s.Delete(ctx, "id", 0)
	require.Error(t, err)
	for _, options := range []*ListOptions{{PageNumber: -1}, {PageSize: -1}, {SortName: "invalid"}} {
		_, _, err = s.List(ctx, options)
		require.Error(t, err)
	}
	assert.Zero(t, mock.GetTotalCallCount())
}
