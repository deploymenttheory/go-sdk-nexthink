package custom_fields

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
	for _, args := range [][2]string{{"", "MANUAL"}, {"id", "RULE_BASED"}, {"id", ""}} {
		_, _, err := s.Get(ctx, args[0], args[1])
		require.Error(t, err)
	}
	_, _, err := s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Delete(ctx, nil)
	require.Error(t, err)
	for _, mutate := range []func(*CreateRequest){func(r *CreateRequest) { r.Name = "" }, func(r *CreateRequest) { r.NQLID = "missing_hash" }, func(r *CreateRequest) { r.Type = "RULE_BASED" }, func(r *CreateRequest) { r.FieldDataType = "" }, func(r *CreateRequest) { r.DataModelObject = "" }, func(r *CreateRequest) { r.Type = "COMPUTED"; r.NQLQuery = "" }} {
		r := load[CreateRequest](t, "Create_input")
		mutate(r)
		_, _, err := s.Create(ctx, r)
		require.Error(t, err)
	}
	u := load[UpdateRequest](t, "Update_input")
	u.Revision = 0
	_, _, err = s.Update(ctx, u)
	require.Error(t, err)
	d := load[DeleteRequest](t, "Delete_input")
	d.DataModelObject = "device/device"
	_, _, err = s.Delete(ctx, d)
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
