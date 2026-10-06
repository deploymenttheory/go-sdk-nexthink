# GetTimelineV2

`GET /apigateway/workflow-executions-insights/api/v2/workflows/{workflowID}/executions/{executionID}/execution-timeline`

This reads the UI API.

Set `NEXTHINK_WORKFLOW_ID`, `NEXTHINK_EXECUTION_ID`.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/workflow_executions/GetTimelineV2
```
