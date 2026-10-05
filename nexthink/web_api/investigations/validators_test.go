package investigations

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
	_, _, err := s.Create(ctx, nil)
	require.Error(t, err)
	_, _, err = s.Update(ctx, "fixture-id", &InvestigationInput{Name: "x"})
	require.Error(t, err)
	_, _, err = s.Get(ctx, "../x")
	require.Error(t, err)
	_, _, err = s.Export(ctx, "")
	require.Error(t, err)
	_, _, err = s.Import(ctx, &ExportDocument{Name: "x"})
	require.Error(t, err)
	_, err = s.Delete(ctx, "..")
	require.Error(t, err)
	require.Zero(t, mock.GetTotalCallCount())
}
