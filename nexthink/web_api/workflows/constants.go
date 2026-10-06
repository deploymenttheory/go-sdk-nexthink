package workflows

import _ "embed"

const (
	Endpoint    = "/apigateway/workflows/manage/graphql"
	operationID = "graphql.workflows"
)

//go:embed queries/List.graphql
var queryList string

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/Export.graphql
var queryExport string

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

const (
	EndpointConnectorDefinitions = "/apigateway/workflows/manage/api/externals/third-party-connectors/v1/connectors"
	EndpointConnectorCredentials = "/apigateway/workflows/manage/api/externals/connectors-credentials/v1/connectors"
)

//go:embed queries/GetFromLibrary.graphql
var queryGetFromLibrary string

//go:embed queries/SetActive.graphql
var querySetActive string

//go:embed queries/Import.graphql
var queryImport string
