package custom_fields

import _ "embed"

const (
	Endpoint     = "/apigateway/nedm/customfields/graphql"
	EndpointList = "/apigateway/content-administration/api/v2/contents/custom-fields"
	operationID  = "graphql.custom_fields"
)

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

//go:embed queries/Get.graphql
var queryGet string
