package collaboration_comments

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvalidRequestsDoNotSend(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	{
		_, _, err := s.ResolveIdentifier(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ListComments(ctx, "")
		require.Error(t, err)
	}
	{
		_, err := s.CreateComment(ctx, "fixture-id", nil)
		require.Error(t, err)
	}
	{
		_, err := s.CreateComment(ctx, "", load[CreateMessageRequest](t, "CreateComment_input"))
		require.Error(t, err)
	}
	{
		_, err := s.CreateReply(ctx, "fixture-id", "fixture-id", nil)
		require.Error(t, err)
	}
	{
		_, err := s.CreateReply(ctx, "", "fixture-id", load[CreateMessageRequest](t, "CreateReply_input"))
		require.Error(t, err)
	}
	{
		_, err := s.CreateReply(ctx, "fixture-id", "", load[CreateMessageRequest](t, "CreateReply_input"))
		require.Error(t, err)
	}
	{
		_, err := s.ArchiveComment(ctx, "", "fixture-id")
		require.Error(t, err)
	}
	{
		_, err := s.ArchiveComment(ctx, "fixture-id", "")
		require.Error(t, err)
	}
	{
		_, err := s.UnarchiveComment(ctx, "", "fixture-id")
		require.Error(t, err)
	}
	{
		_, err := s.UnarchiveComment(ctx, "fixture-id", "")
		require.Error(t, err)
	}
	{
		_, err := s.UpdateComment(ctx, "fixture-id", "fixture-id", nil)
		require.Error(t, err)
	}
	{
		_, err := s.UpdateComment(ctx, "", "fixture-id", load[EditMessageRequest](t, "UpdateComment_input"))
		require.Error(t, err)
	}
	{
		_, err := s.UpdateComment(ctx, "fixture-id", "", load[EditMessageRequest](t, "UpdateComment_input"))
		require.Error(t, err)
	}
	{
		_, err := s.UpdateReply(ctx, "fixture-id", "fixture-id", "fixture-id", nil)
		require.Error(t, err)
	}
	{
		_, err := s.UpdateReply(ctx, "", "fixture-id", "fixture-id", load[EditMessageRequest](t, "UpdateReply_input"))
		require.Error(t, err)
	}
	{
		_, err := s.UpdateReply(ctx, "fixture-id", "", "fixture-id", load[EditMessageRequest](t, "UpdateReply_input"))
		require.Error(t, err)
	}
	{
		_, err := s.UpdateReply(ctx, "fixture-id", "fixture-id", "", load[EditMessageRequest](t, "UpdateReply_input"))
		require.Error(t, err)
	}
	{
		_, err := s.DeleteComment(ctx, "", "fixture-id")
		require.Error(t, err)
	}
	{
		_, err := s.DeleteComment(ctx, "fixture-id", "")
		require.Error(t, err)
	}
	{
		_, err := s.DeleteReply(ctx, "", "fixture-id", "fixture-id")
		require.Error(t, err)
	}
	{
		_, err := s.DeleteReply(ctx, "fixture-id", "", "fixture-id")
		require.Error(t, err)
	}
	{
		_, err := s.DeleteReply(ctx, "fixture-id", "fixture-id", "")
		require.Error(t, err)
	}
	require.Equal(t, 0, mock.GetTotalCallCount())
}
