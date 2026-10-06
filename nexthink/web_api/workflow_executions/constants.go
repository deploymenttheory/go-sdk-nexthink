package workflow_executions

// EndpointGetTimeline is used by the first-party browser UI.
const EndpointGetTimeline = "/apigateway/workflow-executions-insights/api/v3/workflows/{workflowID}/executions/{executionID}/execution-timeline"

// EndpointGetTimelineV2 is used by the first-party browser UI.
const EndpointGetTimelineV2 = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/execution-timeline"

// EndpointListActivities is used by the first-party browser UI.
const EndpointListActivities = "/apigateway/workflow-executions-insights/api/v1/workflows/{workflowID}/executions/{executionID}/activities-history"

// EndpointGetRemoteActionDetails is used by the first-party browser UI.
const EndpointGetRemoteActionDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/remote-action/{thinkletID}"

// EndpointGetCustomFieldsDetails is used by the first-party browser UI.
const EndpointGetCustomFieldsDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/updatecustomfields/{thinkletID}"

// EndpointGetCampaignDetails is used by the first-party browser UI.
const EndpointGetCampaignDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/campaign/{thinkletID}"

// EndpointGetFunctionDetails is used by the first-party browser UI.
const EndpointGetFunctionDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/function/{thinkletID}"

// EndpointGetMessageDetails is used by the first-party browser UI.
const EndpointGetMessageDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/message/{thinkletID}"

// EndpointGetSAPIDetails is used by the first-party browser UI.
const EndpointGetSAPIDetails = "/apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/sapi/{thinkletID}"

// EndpointListWorkflows is used by the first-party browser UI.
const EndpointListWorkflows = "/apigateway/workflows/api/v1/workflows"

// EndpointGetWorkflow is used by the first-party browser UI.
const EndpointGetWorkflow = "/apigateway/workflows/api/v2/workflows/{workflowID}"

// EndpointGet is used by the first-party browser UI.
const EndpointGet = "/apigateway/workflows/api/v1/workflows/{workflowID}/executions/{executionID}"

// EndpointGetHistory is used by the first-party browser UI.
const EndpointGetHistory = "/apigateway/workflows/api/v1/workflows/{workflowID}/executions/{executionID}/history"

// EndpointExecute is used by the first-party browser UI.
const EndpointExecute = "/apigateway/workflows/api/v2/execute"

// EndpointExecuteNQL is used by the first-party browser UI.
const EndpointExecuteNQL = "/apigateway/workflows/api/v2/execute/nql"
