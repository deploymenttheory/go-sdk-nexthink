package library

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidationBeforeTransport(t *testing.T) {
	t.Run("GetContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetContent(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetCustomContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetCustomContent(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetCreateCopyInfo", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetCreateCopyInfo(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("InstallContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).InstallContent(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("InstallCustomContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).InstallCustomContent(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("InstallPack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).InstallPack(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("InstallCustomPack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).InstallCustomPack(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("UpdateContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).UpdateContent(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("UpdatePack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).UpdatePack(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("UpdateCustomContent", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).UpdateCustomContent(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetDependenciesStatus", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetDependenciesStatus(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("InstallDependencies", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).InstallDependencies(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetPack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetPack(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("ImportCustomPack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).ImportCustomPack(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetPackInstallationStatus", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetPackInstallationStatus(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetContentInstallationStatus", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetContentInstallationStatus(context.Background(), "", "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("DeleteCustomPack", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		response, err := NewService(c).DeleteCustomPack(context.Background(), "")
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
}
