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

//go:embed queries/CreateWidget.graphql
var queryCreateWidget string

//go:embed queries/UpdateWidget.graphql
var queryUpdateWidget string

//go:embed queries/DeleteWidget.graphql
var queryDeleteWidget string

//go:embed queries/CreateFilter.graphql
var queryCreateFilter string

//go:embed queries/UpdateFilter.graphql
var queryUpdateFilter string

//go:embed queries/DeleteFilter.graphql
var queryDeleteFilter string

//go:embed queries/CreateTab.graphql
var queryCreateTab string

//go:embed queries/UpdateTab.graphql
var queryUpdateTab string

//go:embed queries/UpdateTabs.graphql
var queryUpdateTabs string

//go:embed queries/DeleteTab.graphql
var queryDeleteTab string

//go:embed queries/UpdateLayout.graphql
var queryUpdateLayout string

//go:embed queries/Export.graphql
var queryExport string

//go:embed queries/Duplicate.graphql
var queryDuplicate string

//go:embed queries/Import.graphql
var queryImport string

//go:embed queries/GetConfiguration.graphql
var queryGetConfiguration string

//go:embed queries/ListCollections.graphql
var queryListCollections string

//go:embed queries/ListFields.graphql
var queryListFields string
