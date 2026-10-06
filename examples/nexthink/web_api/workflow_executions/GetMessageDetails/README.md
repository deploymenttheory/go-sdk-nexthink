# GetMessageDetails

`GET /apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/thinklet/message/{thinkletID}`

This reads the UI API.

Set `NEXTHINK_WORKFLOW_ID`, `NEXTHINK_EXECUTION_ID`, `NEXTHINK_THINKLET_ID`.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/workflow_executions/GetMessageDetails
```
