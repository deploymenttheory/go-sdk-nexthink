package monitors

import _ "embed"

const (
	Endpoint     = "/apigateway/mnt/alert/config/graphql"
	EndpointList = "/apigateway/content-administration/api/v2/contents/alerts"
	operationID  = "graphql.monitors"
)

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/Export.graphql
var queryExport string

//go:embed queries/ExportLibrary.graphql
var queryExportLibrary string

//go:embed queries/Import.graphql
var queryImport string
