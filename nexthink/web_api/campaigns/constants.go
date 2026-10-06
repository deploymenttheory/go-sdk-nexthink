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

//go:embed queries/GetBranding.graphql
var queryGetBranding string

//go:embed queries/UpdateBranding.graphql
var queryUpdateBranding string

//go:embed queries/SetStatus.graphql
var querySetStatus string

//go:embed queries/GetByNQLID.graphql
var queryGetByNQLID string

//go:embed queries/GetFromLibrary.graphql
var queryGetFromLibrary string

//go:embed queries/GetWithV6.graphql
var queryGetWithV6 string

const EndpointFeatures = "/apigateway/api/v1/euf/features"
