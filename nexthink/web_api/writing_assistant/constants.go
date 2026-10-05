package writing_assistant

import _ "embed"

const (
	Endpoint    = "/apigateway/api/adopt/writing-assistant/graphql"
	operationID = "graphql.writing_assistant"
)

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

const EndpointList = "/apigateway/content-administration/api/v2/contents/writing-assistants"
