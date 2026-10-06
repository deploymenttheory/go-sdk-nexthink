package software_metering

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnalyticsValidationBeforeTransport(t *testing.T) {
	t.Run("AutoConfigureMetering", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).AutoConfigureMetering(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetApplications", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetApplications(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetConfigurationByApplicationUUID", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetConfigurationByApplicationUUID(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetConfigurationDetails", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetConfigurationDetails(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetEmployeesTable", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetEmployeesTable(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetPackages", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetPackages(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageBreakdown", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageBreakdown(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageByLicenseEndpoint", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageByLicenseEndpoint(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageByLicenseEndpointCount", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageByLicenseEndpointCount(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageDistribution", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageDistribution(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageDistributionByCategory", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageDistributionByCategory(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
	t.Run("GetUsageOverview", func(t *testing.T) {
		c, m := testutil.NewTransport(t)
		_, response, err := NewService(c).GetUsageOverview(context.Background(), nil)
		require.Error(t, err)
		require.Nil(t, response)
		require.Zero(t, m.GetTotalCallCount())
	})
}

func TestPackageEmptySearchAccepted(t *testing.T) {
	require.NoError(t, validateGetPackages(&GetPackagesRequest{}))
}
