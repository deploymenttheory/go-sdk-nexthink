package monitors

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
	_, _, err = s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Delete(ctx, nil)
	require.Error(t, err)
	r := load[MonitorInput](t, "Create_input")
	id := "existing"
	r.UUID = &id
	_, _, err = s.Create(ctx, r)
	require.Error(t, err)
	for _, mutate := range []func(*UpdateRequest){func(r *UpdateRequest) { r.DocUUID = "" }, func(r *UpdateRequest) { r.Revision = 0 }, func(r *UpdateRequest) { r.Monitor.UUID = nil }, func(r *UpdateRequest) { r.Monitor.Name = "" }} {
		r := load[UpdateRequest](t, "Update_input")
		mutate(r)
		_, _, err = s.Update(ctx, r)
		require.Error(t, err)
	}
	d := load[DeleteRequest](t, "Delete_input")
	d.MonitorType = "unknown"
	_, _, err = s.Delete(ctx, d)
	require.Error(t, err)
	assert.Zero(t, mock.GetTotalCallCount())
}
