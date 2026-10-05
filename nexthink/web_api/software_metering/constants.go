package software_metering

import _ "embed"

const (
	Endpoint    = "/apigateway/dex-metering/graphql"
	operationID = "graphql.software_metering"
)

//go:embed queries/List.graphql
var queryList string

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string
