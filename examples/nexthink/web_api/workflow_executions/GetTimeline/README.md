# GetTimeline

`GET /apigateway/workflow-executions-insights/api/v3/workflows/{workflowID}/executions/{executionID}/execution-timeline`

This reads the UI API.

Set `NEXTHINK_WORKFLOW_ID`, `NEXTHINK_EXECUTION_ID`.

Copy `request.example.json`, replace fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Dates use `YYYY-MM-DDTHH:MM` for Support/VDI local time ranges.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/workflow_executions/GetTimeline
```
