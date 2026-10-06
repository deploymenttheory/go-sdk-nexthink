# GetGlobalHealth

`GET /apigateway/vdi/vdi-service/api/v1/session/{sessionID}/insight/global-health`

This reads the UI API.

Set `NEXTHINK_SESSION_ID`.

Copy `request.example.json`, replace fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Dates use `YYYY-MM-DDTHH:MM:SS` for Support/VDI local time ranges.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/vdi/GetGlobalHealth
```
