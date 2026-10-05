package ratings

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
	cases := []func() error{func() error { _, _, err := s.Get(ctx, ".."); return err }, func() error { _, _, err := s.Create(ctx, nil); return err }, func() error {
		r := load[RatingInput](t, "Create_input")
		r.Conditions = map[string]string{"4": "devices"}
		_, _, err := s.Create(ctx, r)
		return err
	}, func() error {
		r := load[RatingInput](t, "Create_input")
		r.Conditions = map[string]string{"1": " "}
		_, _, err := s.Create(ctx, r)
		return err
	}, func() error { _, _, err := s.Update(ctx, nil); return err }, func() error {
		r := load[UpdateRequest](t, "Update_input")
		r.Revision = 0
		_, _, err := s.Update(ctx, r)
		return err
	}, func() error { _, _, err := s.Delete(ctx, "fixture-id", -1); return err }, func() error { _, _, err := s.Delete(ctx, "../x", 1); return err }}
	for i, call := range cases {
		require.Error(t, call(), "case %d", i)
	}
	require.Zero(t, mock.GetTotalCallCount())
}
