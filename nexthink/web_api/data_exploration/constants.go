package data_exploration

import _ "embed"

const Endpoint = "/apigateway/dash/graphql"
const operationID = "graphql.dashboards"

//go:embed queries/Query.graphql
var queryQuery string

//go:embed queries/Inspect.graphql
var queryInspect string

//go:embed queries/GetFilterValues.graphql
var queryGetFilterValues string

//go:embed queries/ListFields.graphql
var queryListFields string

//go:embed queries/ListSystemRatings.graphql
var queryListSystemRatings string

//go:embed queries/ListOrganisationFields.graphql
var queryListOrganisationFields string

//go:embed queries/GetMenu.graphql
var queryGetMenu string

//go:embed queries/ListBreakdownFields.graphql
var queryListBreakdownFields string

//go:embed queries/GetBreakdownInsights.graphql
var queryGetBreakdownInsights string

//go:embed queries/ListByDurations.graphql
var queryListByDurations string

//go:embed queries/GetOrganisation.graphql
var queryGetOrganisation string

//go:embed queries/GetItemMeta.graphql
var queryGetItemMeta string
