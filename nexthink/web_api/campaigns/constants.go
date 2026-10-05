package campaigns

import _ "embed"

const (
	Endpoint    = "/apigateway/euf-gateway/graphql"
	operationID = "graphql.campaigns"
)

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/List.graphql
var queryList string
