package campaigns

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
	_, _, err := s.Get(ctx, "")
	require.Error(t, err)
	_, _, err = s.Delete(ctx, "")
	require.Error(t, err)
	_, _, err = s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, nil)
	require.Error(t, err)
	r := load[CreateRequest](t, "Create_input")
	r.Campaign.NQLID = "missing_hash"
	_, _, err = s.Create(ctx, r)
	require.Error(t, err)
	r = load[CreateRequest](t, "Create_input")
	r.Campaign.Questions = nil
	_, _, err = s.Create(ctx, r)
	require.Error(t, err)
	for _, mutate := range []func(*UpdateRequest){func(r *UpdateRequest) { r.Campaign.ContentID = "" }, func(r *UpdateRequest) { r.Campaign.BCSUID = "" }, func(r *UpdateRequest) { r.Campaign.Status = "" }, func(r *UpdateRequest) { r.Campaign.Name = "" }} {
		r := load[UpdateRequest](t, "Update_input")
		mutate(r)
		_, _, err = s.Update(ctx, r)
		require.Error(t, err)
	}
	negative := -1
	_, _, err = s.List(ctx, &ListOptions{PageNumber: &negative})
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
func TestBuiltinNQLIDOnUpdate(t *testing.T) {
	r := load[UpdateRequest](t, "Update_input")
	r.Campaign.NQLID = "builtin_without_hash"
	require.NoError(t, ValidateUpdateRequest(r))
}
