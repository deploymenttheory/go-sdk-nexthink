package assets

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
	cases := []func() error{func() error { _, _, err := s.GetSignedURL(ctx, ".."); return err }, func() error { _, _, err := s.Create(ctx, nil); return err }, func() error {
		_, _, err := s.Create(ctx, &UploadRequest{Name: "a\r\nb", MediaType: "image/png", Data: []byte("x")})
		return err
	}, func() error {
		_, _, err := s.Create(ctx, &UploadRequest{Name: "a", MediaType: "invalid", Data: []byte("x")})
		return err
	}, func() error { _, err := s.Update(ctx, "../x", load[UploadRequest](t, "Update_input")); return err }, func() error { _, err := s.Update(ctx, "fixture-id", nil); return err }, func() error { _, err := s.Delete(ctx, ""); return err }}
	for i, call := range cases {
		require.Error(t, call(), "case %d", i)
	}
	require.Zero(t, mock.GetTotalCallCount())
}
