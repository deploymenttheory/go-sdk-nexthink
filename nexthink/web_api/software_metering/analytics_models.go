package software_metering

import "encoding/json"

// JSONValue preserves polymorphic server values without coercion.
type JSONValue = json.RawMessage

// Pagination selects the offset and size of a metering table.
type Pagination struct {
	Offset int `json:"offset"`
	Size   int `json:"size"`
}
type OrderInput struct {
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
}

// Filter selects the usage and category dimensions observed in the metering UI.
type Filter struct {
	UsageTypes                  []string                  `json:"usageTypes,omitempty"`
	UsageDistribution           []UsageDistributionFilter `json:"usageDistribution,omitempty"`
	UsageDistributionByCategory []UsageCategoryFilter     `json:"usageDistributionByCategory,omitempty"`
}
type UsageDistributionFilter struct {
	Type  string   `json:"type"`
	Usage []string `json:"usage"`
}
type UsageCategoryFilter struct {
	CategoryKey string               `json:"categoryKey"`
	Values      []UsageCategoryValue `json:"values"`
}
type UsageCategoryValue struct {
	Name       string   `json:"name"`
	UsageTypes []string `json:"usageTypes,omitempty"`
}

// AutoConfigureMeteringRequest contains the variables for autoConfigureMetering.
type AutoConfigureMeteringRequest struct {
	ApplicationID string `json:"applicationId"`
}
type AutoConfigureMeteringResponseAutoConfigureMetering struct {
	ConfigurationUUID *string `json:"configurationUuid"`
	NQLID             *string `json:"nqlId"`
	Name              *string `json:"name"`
	ApplicationID     *string `json:"applicationId"`
	Status            *string `json:"status"`
	Message           *string `json:"message"`
}

type AutoConfigureMeteringResponse struct {
	AutoConfigureMetering *AutoConfigureMeteringResponseAutoConfigureMetering `json:"autoConfigureMetering"`
}

// GetApplicationsRequest contains the variables for getApplications.
type GetApplicationsRequest struct {
	UUID *string `json:"uuid,omitempty"`
}
type GetApplicationsResponseApplications struct {
	UUID *string `json:"uuid"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type GetApplicationsResponse struct {
	Applications []GetApplicationsResponseApplications `json:"applications"`
}

// GetConfigurationByApplicationUUIDRequest contains the variables for getConfigurationByApplicationUuid.
type GetConfigurationByApplicationUUIDRequest struct {
	ApplicationUUID string `json:"applicationUuid"`
}
type GetConfigurationByApplicationUUIDResponseConfigurationByApplicationUUID struct {
	UUID          *string `json:"uuid"`
	NQLID         *string `json:"nqlId"`
	CreatedAt     *string `json:"createdAt"`
	DataStartTime *string `json:"dataStartTime"`
}

type GetConfigurationByApplicationUUIDResponse struct {
	ConfigurationByApplicationUUID *GetConfigurationByApplicationUUIDResponseConfigurationByApplicationUUID `json:"configurationByApplicationUuid"`
}

// GetConfigurationDetailsRequest contains the variables for getConfigurationDetails.
type GetConfigurationDetailsRequest struct {
	UUID string `json:"uuid"`
}
type GetConfigurationDetailsResponseConfigurationThresholds struct {
	AppType    *string  `json:"appType"`
	Value      *float64 `json:"value"`
	TimeUnit   *string  `json:"timeUnit"`
	MetricType *string  `json:"metricType"`
}

type GetConfigurationDetailsResponseConfigurationPackages struct {
	AppUUID *string `json:"appUuid"`
}

type GetConfigurationDetailsResponseConfigurationApplications struct {
	UUID *string `json:"uuid"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

type GetConfigurationDetailsResponseConfiguration struct {
	UUID          *string                                                    `json:"uuid"`
	Name          *string                                                    `json:"name"`
	LicenseType   *string                                                    `json:"licenseType"`
	Thresholds    []GetConfigurationDetailsResponseConfigurationThresholds   `json:"thresholds"`
	Packages      []GetConfigurationDetailsResponseConfigurationPackages     `json:"packages"`
	Applications  []GetConfigurationDetailsResponseConfigurationApplications `json:"applications"`
	DataStartTime *string                                                    `json:"dataStartTime"`
}

type GetConfigurationDetailsResponse struct {
	Configuration *GetConfigurationDetailsResponseConfiguration `json:"configuration"`
}

// GetEmployeesTableRequest contains the variables for getEmployeesTable.
type GetEmployeesTableRequest struct {
	UUID            string     `json:"uuid"`
	ApplicationUUID string     `json:"applicationUuid"`
	Pagination      Pagination `json:"pagination"`
	OrderInput      OrderInput `json:"orderInput"`
	TimeRange       string     `json:"timeRange"`
}
type GetEmployeesTableResponseUsageByLicenseEndpointLicenseEndpointUsages struct {
	Name            *string  `json:"name"`
	UsageSecondsWeb *float64 `json:"usageSecondsWeb"`
	LastUsage       *string  `json:"lastUsage"`
}

type GetEmployeesTableResponseUsageByLicenseEndpoint struct {
	TotalSize             *int64                                                                 `json:"totalSize"`
	NQLQuery              *string                                                                `json:"nqlQuery"`
	LicenseEndpointUsages []GetEmployeesTableResponseUsageByLicenseEndpointLicenseEndpointUsages `json:"licenseEndpointUsages"`
}

type GetEmployeesTableResponse struct {
	UsageByLicenseEndpoint *GetEmployeesTableResponseUsageByLicenseEndpoint `json:"usageByLicenseEndpoint"`
}

// GetPackagesRequest contains the variables for getPackages.
type GetPackagesRequest struct {
	Search string `json:"search"`
}
type GetPackagesResponsePackages struct {
	Name *string `json:"name"`
}

type GetPackagesResponse struct {
	Packages []GetPackagesResponsePackages `json:"packages"`
}

// GetUsageBreakdownRequest contains the variables for getUsageBreakdown.
type GetUsageBreakdownRequest struct {
	UUID            string `json:"uuid"`
	ApplicationUUID string `json:"applicationUuid"`
	TimeRange       string `json:"timeRange"`
}
type GetUsageBreakdownResponseUsageDistributionByCategoryUsage struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
}

type GetUsageBreakdownResponseUsageDistributionByCategory struct {
	Name  *string                                                    `json:"name"`
	Usage *GetUsageBreakdownResponseUsageDistributionByCategoryUsage `json:"usage"`
}

type GetUsageBreakdownResponse struct {
	UsageDistributionByCategory []GetUsageBreakdownResponseUsageDistributionByCategory `json:"usageDistributionByCategory"`
}

// GetUsageByLicenseEndpointRequest contains the variables for getUsageByLicenseEndpoint.
type GetUsageByLicenseEndpointRequest struct {
	UUID       string     `json:"uuid"`
	Pagination Pagination `json:"pagination"`
	OrderInput OrderInput `json:"orderInput"`
	TimeRange  string     `json:"timeRange"`
	Filter     *Filter    `json:"filter,omitempty"`
}
type GetUsageByLicenseEndpointResponseUsageByLicenseEndpointLicenseEndpointUsages struct {
	Name                *string  `json:"name"`
	UsageSecondsDesktop *float64 `json:"usageSecondsDesktop"`
	UsageSecondsWeb     *float64 `json:"usageSecondsWeb"`
	LastUsage           *string  `json:"lastUsage"`
}

type GetUsageByLicenseEndpointResponseUsageByLicenseEndpoint struct {
	TotalSize             *int64                                                                         `json:"totalSize"`
	LicenseEndpointUsages []GetUsageByLicenseEndpointResponseUsageByLicenseEndpointLicenseEndpointUsages `json:"licenseEndpointUsages"`
	InvestigationURL      *string                                                                        `json:"investigationUrl"`
	NQLQuery              *string                                                                        `json:"nqlQuery"`
}

type GetUsageByLicenseEndpointResponse struct {
	UsageByLicenseEndpoint *GetUsageByLicenseEndpointResponseUsageByLicenseEndpoint `json:"usageByLicenseEndpoint"`
}

// GetUsageByLicenseEndpointCountRequest contains the variables for getUsageByLicenseEndpointCount.
type GetUsageByLicenseEndpointCountRequest struct {
	ConfigurationUUID string  `json:"configurationUuid"`
	TimeRange         string  `json:"timeRange"`
	Filter            *Filter `json:"filter,omitempty"`
}
type GetUsageByLicenseEndpointCountResponse struct {
	UsageByLicenseEndpointCount *int64 `json:"usageByLicenseEndpointCount"`
}

// GetUsageDistributionRequest contains the variables for getUsageDistribution.
type GetUsageDistributionRequest struct {
	UUID      string  `json:"uuid"`
	TimeRange string  `json:"timeRange"`
	Filter    *Filter `json:"filter,omitempty"`
}
type GetUsageDistributionResponseUsageDistributionWeb struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
	NotUsing   *int64 `json:"notUsing"`
}

type GetUsageDistributionResponseUsageDistributionDesktop struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
	NotUsing   *int64 `json:"notUsing"`
}

type GetUsageDistributionResponseUsageDistribution struct {
	Web     *GetUsageDistributionResponseUsageDistributionWeb     `json:"web"`
	Desktop *GetUsageDistributionResponseUsageDistributionDesktop `json:"desktop"`
}

type GetUsageDistributionResponse struct {
	UsageDistribution *GetUsageDistributionResponseUsageDistribution `json:"usageDistribution"`
}

// GetUsageDistributionByCategoryRequest contains the variables for getUsageDistributionByCategory.
type GetUsageDistributionByCategoryRequest struct {
	UUID        string  `json:"uuid"`
	CategoryKey string  `json:"categoryKey"`
	TimeRange   string  `json:"timeRange"`
	Filter      *Filter `json:"filter,omitempty"`
}
type GetUsageDistributionByCategoryResponseUsageDistributionByCategoryUsage struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
	NotUsing   *int64 `json:"notUsing"`
}

type GetUsageDistributionByCategoryResponseUsageDistributionByCategory struct {
	Name  *string                                                                 `json:"name"`
	Usage *GetUsageDistributionByCategoryResponseUsageDistributionByCategoryUsage `json:"usage"`
}

type GetUsageDistributionByCategoryResponse struct {
	UsageDistributionByCategory []GetUsageDistributionByCategoryResponseUsageDistributionByCategory `json:"usageDistributionByCategory"`
}

// GetUsageOverviewRequest contains the variables for getUsageOverview.
type GetUsageOverviewRequest struct {
	UUID            string `json:"uuid"`
	ApplicationUUID string `json:"applicationUuid"`
	TimeRange       string `json:"timeRange"`
}
type GetUsageOverviewResponseUsageOverview struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
}

type GetUsageOverviewResponse struct {
	UsageOverview *GetUsageOverviewResponseUsageOverview `json:"usageOverview"`
}

// GetConfigurationUsageOverviewRequest covers an entire metering configuration, including optional filters.
type GetConfigurationUsageOverviewRequest struct {
	UUID      string  `json:"uuid"`
	TimeRange string  `json:"timeRange"`
	Filter    *Filter `json:"filter,omitempty"`
}
type GetConfigurationUsageOverviewResponse struct {
	UsageOverview *ConfigurationUsageOverview `json:"usageOverview"`
}
type ConfigurationUsageOverview struct {
	Using      *int64 `json:"using"`
	UnderUsing *int64 `json:"underUsing"`
	NotUsing   *int64 `json:"notUsing"`
}
