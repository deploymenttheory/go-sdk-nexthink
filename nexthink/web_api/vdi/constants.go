package vdi

// EndpointGetSessionTimeline is used by the first-party browser UI.
const EndpointGetSessionTimeline = "/apigateway/vdi/vdi-service/api/v2/session/{sessionID}/timeline/details"

// EndpointGetHypervisorTimeline is used by the first-party browser UI.
const EndpointGetHypervisorTimeline = "/apigateway/vdi/vdi-service/api/v1/session/{sessionID}/timeline/hypervisor"

// EndpointGetGlobalHealth is used by the first-party browser UI.
const EndpointGetGlobalHealth = "/apigateway/vdi/vdi-service/api/v1/session/{sessionID}/insight/global-health"

// EndpointValidateHostname is used by the first-party browser UI.
const EndpointValidateHostname = "/apigateway/vdi/vdi-service/api/v1/hypervisor/hostname/validate"
