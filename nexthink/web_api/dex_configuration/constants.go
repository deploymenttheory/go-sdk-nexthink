package dex_configuration

import _ "embed"

const (
	Endpoint    = "/apigateway/dex-ec/graphql"
	operationID = "graphql.dex_configuration"
)

//go:embed queries/GetAccount.graphql
var queryGetAccount string

//go:embed queries/UpdateApplications.graphql
var queryUpdateApplications string

//go:embed queries/GetScoreMetrics.graphql
var queryGetScoreMetrics string

//go:embed queries/UpdateScoreMetrics.graphql
var queryUpdateScoreMetrics string

//go:embed queries/GetVDIOptIn.graphql
var queryGetVDIOptIn string

//go:embed queries/OptInVDI.graphql
var queryOptInVDI string

//go:embed queries/GetMemoryMetricsOptIn.graphql
var queryGetMemoryMetricsOptIn string

//go:embed queries/OptInMemoryMetrics.graphql
var queryOptInMemoryMetrics string

//go:embed queries/GetCampaign.graphql
var queryGetCampaign string

//go:embed queries/EnableCampaign.graphql
var queryEnableCampaign string

//go:embed queries/GetApplications.graphql
var queryGetApplications string
