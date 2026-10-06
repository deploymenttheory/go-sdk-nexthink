# ListActivities

`GET /apigateway/workflow-executions-insights/api/v1/workflows/{workflowID}/executions/{executionID}/activities-history`

This reads the UI API.

Set `NEXTHINK_WORKFLOW_ID`, `NEXTHINK_EXECUTION_ID`.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/workflow_executions/ListActivities
```
