package remote_actions

import _ "embed"

const (
	Endpoint    = "/apigateway/act/manage/graphql"
	operationID = "graphql.remote_actions"
)

//go:embed queries/Get.graphql
var queryGet string

//go:embed queries/GetForView.graphql
var queryGetForView string

//go:embed queries/GetContentVolume.graphql
var queryGetContentVolume string

//go:embed queries/InspectBashScript.graphql
var queryInspectBashScript string

//go:embed queries/InspectPowerShellScript.graphql
var queryInspectPowerShellScript string

//go:embed queries/GetPowerShellSignature.graphql
var queryGetPowerShellSignature string

//go:embed queries/Create.graphql
var queryCreate string

//go:embed queries/Update.graphql
var queryUpdate string

//go:embed queries/Delete.graphql
var queryDelete string

const EndpointList = "/apigateway/content-administration/api/v2/contents/remoteactions"

//go:embed queries/Export.graphql
var queryExport string

//go:embed queries/Import.graphql
var queryImport string

//go:embed queries/GetFromLibrary.graphql
var queryGetFromLibrary string
