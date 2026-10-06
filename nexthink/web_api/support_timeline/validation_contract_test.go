package support_timeline

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidationPreventsHTTP(t *testing.T) {
	ctx := context.Background()
	transport, mock := testutil.NewTransport(t)
	service := NewService(transport)
	t.Run("GetAlertsAndErrorsNilRequest", func(t *testing.T) {
		_, response, err := service.GetAlertsAndErrors(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetAlertsAndErrorsInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetAlertsAndErrors(ctx, "..", load[TimeRange](t, "GetAlertsAndErrors_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetPerformanceNilRequest", func(t *testing.T) {
		_, response, err := service.GetPerformance(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetPerformanceInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetPerformance(ctx, "..", load[TimeRange](t, "GetPerformance_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetConnectivityNilRequest", func(t *testing.T) {
		_, response, err := service.GetConnectivity(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetConnectivityInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetConnectivity(ctx, "..", load[TimeRange](t, "GetConnectivity_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetApplicationConnectivityNilRequest", func(t *testing.T) {
		_, response, err := service.GetApplicationConnectivity(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetApplicationConnectivityInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetApplicationConnectivity(ctx, "..", load[TimeRange](t, "GetApplicationConnectivity_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetActivityNilRequest", func(t *testing.T) {
		_, response, err := service.GetActivity(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetActivityInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetActivity(ctx, "..", load[TimeRange](t, "GetActivity_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetApplicationsNilRequest", func(t *testing.T) {
		_, response, err := service.GetApplications(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetApplicationsInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetApplications(ctx, "..", load[TimeRange](t, "GetApplications_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetUserInteractionsNilRequest", func(t *testing.T) {
		_, response, err := service.GetUserInteractions(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetUserInteractionsInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetUserInteractions(ctx, "..", load[TimeRange](t, "GetUserInteractions_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCollaborationNilRequest", func(t *testing.T) {
		_, response, err := service.GetCollaboration(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCollaborationInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetCollaboration(ctx, "..", load[TimeRange](t, "GetCollaboration_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetErrorsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetErrorsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetErrorsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetErrorsDrilldown(ctx, "..", load[TimeRange](t, "GetErrorsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetFreezesDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetFreezesDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetFreezesDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetFreezesDrilldown(ctx, "..", load[TimeRange](t, "GetFreezesDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetAlertsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetAlertsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetAlertsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetAlertsDrilldown(ctx, "..", load[AlertsRequest](t, "GetAlertsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetActionsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetActionsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetActionsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetActionsDrilldown(ctx, "..", load[TimeRange](t, "GetActionsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSystemBootsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetSystemBootsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSystemBootsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetSystemBootsDrilldown(ctx, "..", load[TimeRange](t, "GetSystemBootsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSystemBootsAndSuspendsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetSystemBootsAndSuspendsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetSystemBootsAndSuspendsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetSystemBootsAndSuspendsDrilldown(ctx, "..", load[TimeRange](t, "GetSystemBootsAndSuspendsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTeamsCallsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetTeamsCallsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetTeamsCallsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetTeamsCallsDrilldown(ctx, "..", load[TimeRange](t, "GetTeamsCallsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetZoomCallsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetZoomCallsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetZoomCallsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetZoomCallsDrilldown(ctx, "..", load[TimeRange](t, "GetZoomCallsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDesktopConnectivityDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetDesktopConnectivityDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDesktopConnectivityDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetDesktopConnectivityDrilldown(ctx, "..", load[ApplicationRequest](t, "GetDesktopConnectivityDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWebConnectivityDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetWebConnectivityDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWebConnectivityDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetWebConnectivityDrilldown(ctx, "..", load[ApplicationRequest](t, "GetWebConnectivityDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDesktopApplicationsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetDesktopApplicationsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDesktopApplicationsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetDesktopApplicationsDrilldown(ctx, "..", load[ApplicationRequest](t, "GetDesktopApplicationsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWebApplicationsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetWebApplicationsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWebApplicationsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetWebApplicationsDrilldown(ctx, "..", load[ApplicationRequest](t, "GetWebApplicationsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInstallationsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetInstallationsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetInstallationsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetInstallationsDrilldown(ctx, "..", load[TimeRange](t, "GetInstallationsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetEthernetDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetEthernetDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetEthernetDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetEthernetDrilldown(ctx, "..", load[TimeRange](t, "GetEthernetDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWiFiDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetWiFiDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetWiFiDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetWiFiDrilldown(ctx, "..", load[TimeRange](t, "GetWiFiDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetConnectionsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetConnectionsDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetConnectionsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetConnectionsDrilldown(ctx, "..", load[TimeRange](t, "GetConnectionsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNetworkApplicationDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetNetworkApplicationDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNetworkApplicationDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetNetworkApplicationDrilldown(ctx, "..", load[ApplicationRequest](t, "GetNetworkApplicationDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCPUDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetCPUDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetCPUDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetCPUDrilldown(ctx, "..", load[TimeRange](t, "GetCPUDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMemoryDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetMemoryDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetMemoryDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetMemoryDrilldown(ctx, "..", load[TimeRange](t, "GetMemoryDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDiskPerformanceDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetDiskPerformanceDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDiskPerformanceDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetDiskPerformanceDrilldown(ctx, "..", load[TimeRange](t, "GetDiskPerformanceDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDriveSpaceDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetDriveSpaceDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetDriveSpaceDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetDriveSpaceDrilldown(ctx, "..", load[TimeRange](t, "GetDriveSpaceDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetGPUDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetGPUDrilldown(ctx, "fixture", "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetGPUDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetGPUDrilldown(ctx, "..", "fixture", load[TimeRange](t, "GetGPUDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetGPUDrilldownInvalidslot", func(t *testing.T) {
		_, response, err := service.GetGPUDrilldown(ctx, "fixture", "..", load[TimeRange](t, "GetGPUDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNPUDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetNPUDrilldown(ctx, "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNPUDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetNPUDrilldown(ctx, "..", load[TimeRange](t, "GetNPUDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetUserInteractionsDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetUserInteractionsDrilldown(ctx, "fixture", "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetUserInteractionsDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetUserInteractionsDrilldown(ctx, "..", "fixture", load[TimeRange](t, "GetUserInteractionsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetUserInteractionsDrilldownInvaliduserID", func(t *testing.T) {
		_, response, err := service.GetUserInteractionsDrilldown(ctx, "fixture", "..", load[TimeRange](t, "GetUserInteractionsDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRoundTripTimeDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetRoundTripTimeDrilldown(ctx, "fixture", "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRoundTripTimeDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetRoundTripTimeDrilldown(ctx, "..", "fixture", load[TimeRange](t, "GetRoundTripTimeDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetRoundTripTimeDrilldownInvaliduserID", func(t *testing.T) {
		_, response, err := service.GetRoundTripTimeDrilldown(ctx, "fixture", "..", load[TimeRange](t, "GetRoundTripTimeDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNetworkLatencyDrilldownNilRequest", func(t *testing.T) {
		_, response, err := service.GetNetworkLatencyDrilldown(ctx, "fixture", "fixture", nil)
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNetworkLatencyDrilldownInvaliddeviceID", func(t *testing.T) {
		_, response, err := service.GetNetworkLatencyDrilldown(ctx, "..", "fixture", load[TimeRange](t, "GetNetworkLatencyDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	t.Run("GetNetworkLatencyDrilldownInvaliduserID", func(t *testing.T) {
		_, response, err := service.GetNetworkLatencyDrilldown(ctx, "fixture", "..", load[TimeRange](t, "GetNetworkLatencyDrilldown_request"))
		require.Error(t, err)
		assert.Nil(t, response)
	})
	assert.Equal(t, 0, mock.GetTotalCallCount())
	assert.Equal(t, "UTC", service.headers()["time-zone"])
	assert.Equal(t, "0", service.headers()["utc-offset"])
}
