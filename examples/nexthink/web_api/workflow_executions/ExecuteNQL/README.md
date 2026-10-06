# ExecuteNQL

`POST /apigateway/workflows/api/v2/execute/nql`

This submits a real execution. Set `NEXTHINK_ALLOW_EXECUTION=true` only after choosing approved targets.

Copy `request.example.json`, replace fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Dates use `YYYY-MM-DDTHH:MM` for Support/VDI local time ranges.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/workflow_executions/ExecuteNQL
```
