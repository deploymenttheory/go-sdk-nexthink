package application_experience

import _ "embed"

const (
	Endpoint    = "/apigateway/bus/appexgw/graphql"
	operationID = "graphql.application_experience"
)

//go:embed queries/GetApplicationsOverviewDesktopInvestigations.graphql
var queryGetApplicationsOverviewDesktopInvestigations string

//go:embed queries/GetAvgNetworkResponseTime.graphql
var queryGetAvgNetworkResponseTime string

//go:embed queries/GetBinarySuggestion.graphql
var queryGetBinarySuggestion string

//go:embed queries/GetDeviceCentricMetricBreakdown.graphql
var queryGetDeviceCentricMetricBreakdown string

//go:embed queries/GetFailedConnectionsRatio.graphql
var queryGetFailedConnectionsRatio string

//go:embed queries/GetInsights.graphql
var queryGetInsights string

//go:embed queries/GetMetricBreakdown.graphql
var queryGetMetricBreakdown string

//go:embed queries/GetNumOfCrashesAndDevices.graphql
var queryGetNumOfCrashesAndDevices string

//go:embed queries/GetNumOfCrashesAndEmployees.graphql
var queryGetNumOfCrashesAndEmployees string

//go:embed queries/GetNumOfDevices.graphql
var queryGetNumOfDevices string

//go:embed queries/GetNumOfDevicesWithCrashes.graphql
var queryGetNumOfDevicesWithCrashes string

//go:embed queries/GetNumOfEmployees.graphql
var queryGetNumOfEmployees string

//go:embed queries/GetNumOfEmployeesWithCrashes.graphql
var queryGetNumOfEmployeesWithCrashes string

//go:embed queries/OverviewDesktopTooltips.graphql
var queryOverviewDesktopTooltips string

//go:embed queries/TilesDesktopCrashesPerEmployee.graphql
var queryTilesDesktopCrashesPerEmployee string

//go:embed queries/TilesDesktopNumberOfEmployees.graphql
var queryTilesDesktopNumberOfEmployees string

//go:embed queries/TilesErrorCount.graphql
var queryTilesErrorCount string

//go:embed queries/TilesFrustratingPageLoads.graphql
var queryTilesFrustratingPageLoads string

//go:embed queries/TilesNumberOfEmployees.graphql
var queryTilesNumberOfEmployees string

//go:embed queries/TilesPageLoadTime.graphql
var queryTilesPageLoadTime string

//go:embed queries/TilesTransactionTime.graphql
var queryTilesTransactionTime string

//go:embed queries/TilesUsageTime.graphql
var queryTilesUsageTime string

//go:embed queries/WebOverviewTooltips.graphql
var queryWebOverviewTooltips string
