package action_executions

// EndpointListRemoteActions is used by the first-party browser UI.
const EndpointListRemoteActions = "/apigateway/act/api/v2/remote-action"

// EndpointGetRemoteAction is used by the first-party browser UI.
const EndpointGetRemoteAction = "/apigateway/act/api/v2/remote-action/details"

// EndpointListRemoteActionsForQuery is used by the first-party browser UI.
const EndpointListRemoteActionsForQuery = "/apigateway/act/api/v2/remote-action/large"

// EndpointExecute is used by the first-party browser UI.
const EndpointExecute = "/apigateway/act/api/v3/execute"

// EndpointListActions is used by the first-party browser UI.
const EndpointListActions = "/apigateway/atl/action-executions-be/api/v1/actions"

// EndpointGetDeviceActions is the Amplify extension's device-scoped action list.
const EndpointGetDeviceActions = "/apigateway/atl/action-executions-be/api/v1/device/{deviceID}/actions"

// EndpointGetDeviceHistory is used by the first-party browser UI.
const EndpointGetDeviceHistory = "/apigateway/atl/action-executions-be/api/v1/device/{deviceID}/actions/executions/history"
