package software_metering

import (
	"fmt"
	"strings"
)

func validateAutoConfigureMetering(r *AutoConfigureMeteringRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ApplicationID) == "" {
		return fmt.Errorf("applicationId is required")
	}
	return nil
}
func validateGetApplications(r *GetApplicationsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateGetConfigurationByApplicationUUID(r *GetConfigurationByApplicationUUIDRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ApplicationUUID) == "" {
		return fmt.Errorf("applicationUuid is required")
	}
	return nil
}
func validateGetConfigurationDetails(r *GetConfigurationDetailsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	return nil
}
func validateGetEmployeesTable(r *GetEmployeesTableRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.ApplicationUUID) == "" {
		return fmt.Errorf("applicationUuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetPackages(r *GetPackagesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateGetUsageBreakdown(r *GetUsageBreakdownRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.ApplicationUUID) == "" {
		return fmt.Errorf("applicationUuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetUsageByLicenseEndpoint(r *GetUsageByLicenseEndpointRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetUsageByLicenseEndpointCount(r *GetUsageByLicenseEndpointCountRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ConfigurationUUID) == "" {
		return fmt.Errorf("configurationUuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetUsageDistribution(r *GetUsageDistributionRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetUsageDistributionByCategory(r *GetUsageDistributionByCategoryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.CategoryKey) == "" {
		return fmt.Errorf("categoryKey is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetUsageOverview(r *GetUsageOverviewRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" {
		return fmt.Errorf("uuid is required")
	}
	if strings.TrimSpace(r.ApplicationUUID) == "" {
		return fmt.Errorf("applicationUuid is required")
	}
	if strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("timeRange is required")
	}
	return nil
}
func validateGetConfigurationUsageOverview(r *GetConfigurationUsageOverviewRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.UUID) == "" || strings.TrimSpace(r.TimeRange) == "" {
		return fmt.Errorf("uuid and timeRange are required")
	}
	return nil
}
