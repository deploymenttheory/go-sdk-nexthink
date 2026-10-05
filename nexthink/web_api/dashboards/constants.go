package dashboards

import _ "embed"

const (
	Endpoint     = "/apigateway/dash/graphql"
	EndpointList = "/apigateway/content-administration/api/v2/contents/dashboards"
)
const operationID = "graphql.dashboards"

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string
