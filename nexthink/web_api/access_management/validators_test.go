package access_management

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
		_, _, err := s.ListUsers(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.CreateUser(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateUser(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetUser(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.DeleteUser(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateMappings(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetViewDomains(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateSSOConfiguration(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetAPICredential(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.CreateAPICredential(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateAPICredential(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.DeleteAPICredential(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.CreateRole(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateRole(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetRole(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetRoleSummary(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.DeleteLegacyProfile(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.DeleteRole(ctx, "")
		require.Error(t, err)
	}
	{
		_, _, err := s.GetFeatureFlags(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetSharedContents(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ListContents(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateAccount(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ResetUserMFA(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ResetUserPassword(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ResendActivationEmail(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UnlockUser(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.ChangePassword(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.GetSupportAccess(ctx, "")
		require.Error(t, err)
	}
	{
		_, _, err := s.CreateSupportAccess(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateSupportAccess(ctx, "fixture-id", nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.UpdateSupportAccess(ctx, "", load[SupportAccess](t, "UpdateSupportAccess_input"))
		require.Error(t, err)
	}
	{
		_, _, err := s.DeleteSupportAccess(ctx, "")
		require.Error(t, err)
	}
	{
		_, _, err := s.GetRolePermissions(ctx, "")
		require.Error(t, err)
	}
	{
		_, _, err := s.GrantRoleContentPermissions(ctx, nil)
		require.Error(t, err)
	}
	{
		_, _, err := s.RevokeRoleContentPermissions(ctx, nil)
		require.Error(t, err)
	}
	require.Equal(t, 0, mock.GetTotalCallCount())
}
