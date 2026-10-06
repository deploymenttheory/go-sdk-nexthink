package software_metering

import _ "embed"

//go:embed queries/AutoConfigureMetering.graphql
var queryAutoConfigureMetering string

//go:embed queries/GetApplications.graphql
var queryGetApplications string

//go:embed queries/GetConfigurationByApplicationUUID.graphql
var queryGetConfigurationByApplicationUUID string

//go:embed queries/GetConfigurationDetails.graphql
var queryGetConfigurationDetails string

//go:embed queries/GetEmployeesTable.graphql
var queryGetEmployeesTable string

//go:embed queries/GetPackages.graphql
var queryGetPackages string

//go:embed queries/GetUsageBreakdown.graphql
var queryGetUsageBreakdown string

//go:embed queries/GetUsageByLicenseEndpoint.graphql
var queryGetUsageByLicenseEndpoint string

//go:embed queries/GetUsageByLicenseEndpointCount.graphql
var queryGetUsageByLicenseEndpointCount string

//go:embed queries/GetUsageDistribution.graphql
var queryGetUsageDistribution string

//go:embed queries/GetUsageDistributionByCategory.graphql
var queryGetUsageDistributionByCategory string

//go:embed queries/GetUsageOverview.graphql
var queryGetUsageOverview string

//go:embed queries/GetConfigurationUsageOverview.graphql
var queryGetConfigurationUsageOverview string
