package application_experience

import (
	"fmt"
	"strings"
)

func validateGetApplicationsOverviewDesktopInvestigations(r *GetApplicationsOverviewDesktopInvestigationsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetAvgNetworkResponseTime(r *GetAvgNetworkResponseTimeRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetBinarySuggestion(r *GetBinarySuggestionRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateGetDeviceCentricMetricBreakdown(r *GetDeviceCentricMetricBreakdownRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.DesktopAppExperienceID) == "" {
		return fmt.Errorf("desktopAppExperienceId is required")
	}
	if strings.TrimSpace(r.DimensionToGroupBy) == "" {
		return fmt.Errorf("dimensionToGroupBy is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	if r.Limit <= 0 {
		return fmt.Errorf("limit must be positive")
	}
	return nil
}
func validateGetFailedConnectionsRatio(r *GetFailedConnectionsRatioRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetInsights(r *GetInsightsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ApplicationID) == "" {
		return fmt.Errorf("applicationId is required")
	}
	if strings.TrimSpace(r.ApplicationName) == "" {
		return fmt.Errorf("applicationName is required")
	}
	if strings.TrimSpace(r.InsightType) == "" {
		return fmt.Errorf("insightType is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetMetricBreakdown(r *GetMetricBreakdownRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.DesktopAppExperienceID) == "" {
		return fmt.Errorf("desktopAppExperienceId is required")
	}
	if strings.TrimSpace(r.DimensionToGroupBy) == "" {
		return fmt.Errorf("dimensionToGroupBy is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	if r.Limit <= 0 {
		return fmt.Errorf("limit must be positive")
	}
	return nil
}
func validateGetNumOfCrashesAndDevices(r *GetNumOfCrashesAndDevicesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetNumOfCrashesAndEmployees(r *GetNumOfCrashesAndEmployeesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetNumOfDevices(r *GetNumOfDevicesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetNumOfDevicesWithCrashes(r *GetNumOfDevicesWithCrashesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetNumOfEmployees(r *GetNumOfEmployeesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateGetNumOfEmployeesWithCrashes(r *GetNumOfEmployeesWithCrashesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateOverviewDesktopTooltips(r *OverviewDesktopTooltipsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.DesktopTimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("desktopTimeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.DesktopTimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("desktopTimeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.DesktopTimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("desktopTimeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.DesktopTimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("desktopTimeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesDesktopCrashesPerEmployee(r *TilesDesktopCrashesPerEmployeeRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesDesktopNumberOfEmployees(r *TilesDesktopNumberOfEmployeesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesErrorCount(r *TilesErrorCountRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesFrustratingPageLoads(r *TilesFrustratingPageLoadsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesNumberOfEmployees(r *TilesNumberOfEmployeesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesPageLoadTime(r *TilesPageLoadTimeRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesTransactionTime(r *TilesTransactionTimeRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateTilesUsageTime(r *TilesUsageTimeRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
func validateWebOverviewTooltips(r *WebOverviewTooltipsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalFrom) == "" {
		return fmt.Errorf("timeFrameFilter.LocalFrom is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalTo) == "" {
		return fmt.Errorf("timeFrameFilter.LocalTo is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.LocalNow) == "" {
		return fmt.Errorf("timeFrameFilter.LocalNow is required")
	}
	if strings.TrimSpace(r.TimeFrameFilter.TimeZone) == "" {
		return fmt.Errorf("timeFrameFilter.TimeZone is required")
	}
	return nil
}
