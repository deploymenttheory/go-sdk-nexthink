package writing_assistant

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
	for _, id := range []string{"", " "} {
		_, _, err := s.Get(ctx, id)
		require.Error(t, err)
		_, _, err = s.Delete(ctx, id)
		require.Error(t, err)
	}
	for _, req := range []*CreateRequest{nil, {}, {Name: "fixture", Instructions: "Check spelling", Tool: "REVIEW"}} {
		_, _, err := s.Create(ctx, req)
		require.Error(t, err)
	}
	for _, req := range []*UpdateRequest{nil, {}, {ID: "id", Name: "fixture", Instructions: "Check spelling", Revision: 0}, {Name: "fixture", Instructions: "Check spelling", Revision: 1}} {
		_, _, err := s.Update(ctx, req)
		require.Error(t, err)
	}
	assert.Zero(t, mock.GetTotalCallCount())
}
