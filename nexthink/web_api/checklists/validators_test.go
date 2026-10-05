package checklists

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestInvalidRequestsDoNotReachTransport(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	cases := []func() error{func() error { _, _, err := s.Get(ctx, ".."); return err }, func() error { _, _, err := s.Create(ctx, nil, nil); return err }, func() error {
		r := load[ChecklistInput](t, "Create_input")
		r.Categories = nil
		_, _, err := s.Create(ctx, r, nil)
		return err
	}, func() error {
		r := load[ChecklistInput](t, "Create_input")
		r.Platforms = "unknown"
		_, _, err := s.Create(ctx, r, nil)
		return err
	}, func() error {
		_, _, err := s.Create(ctx, load[ChecklistInput](t, "Create_input"), &CreateOptions{LibraryUUID: "../x"})
		return err
	}, func() error {
		_, _, err := s.Update(ctx, "fixture-id", 0, load[ChecklistInput](t, "Update_input"))
		return err
	}, func() error { _, _, err := s.Update(ctx, "fixture-id", 2, nil); return err }, func() error { _, err := s.Delete(ctx, "fixture-id", 0); return err }}
	for i, call := range cases {
		require.Error(t, call(), "case %d", i)
	}
	require.Zero(t, mock.GetTotalCallCount())
}
