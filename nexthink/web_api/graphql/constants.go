package graphql

const (
	EndpointWritingAssistant = "/apigateway/api/adopt/writing-assistant/graphql"
	EndpointBenchmark        = "/apigateway/cci/graphql"
	EndpointDexConfiguration = "/apigateway/dex-ec/graphql"
	EndpointDashboards       = "/apigateway/dash/graphql"
	EndpointWorkflows        = "/apigateway/workflows/manage/graphql"
	EndpointCampaigns        = "/apigateway/euf-gateway/graphql"
	EndpointVisualEditor     = "/apigateway/visual-editor/graphql"
	EndpointValueProvider    = "/apigateway/value-provider/graphql"
	EndpointMonitors         = "/apigateway/mnt/alert/config/graphql"
	EndpointRemoteActions    = "/apigateway/act/manage/graphql"
	EndpointSoftwareMetering = "/apigateway/dex-metering/graphql"
)

var endpoints = map[string]string{
	"graphql.writing_assistant": EndpointWritingAssistant,
	"graphql.benchmark":         EndpointBenchmark,
	"graphql.dex_configuration": EndpointDexConfiguration,
	"graphql.dashboards":        EndpointDashboards,
	"graphql.workflows":         EndpointWorkflows,
	"graphql.campaigns":         EndpointCampaigns,
	"graphql.visual_editor":     EndpointVisualEditor,
	"graphql.value_provider":    EndpointValueProvider,
	"graphql.monitors":          EndpointMonitors,
	"graphql.remote_actions":    EndpointRemoteActions,
	"graphql.software_metering": EndpointSoftwareMetering,
}
