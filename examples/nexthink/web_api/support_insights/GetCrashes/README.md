# GetCrashes

`POST /apigateway/atl/support-views-insights-be/api/v1/device/{deviceID}/insights/crashes`

This reads the UI API.

Set `NEXTHINK_DEVICE_ID`.

Copy `request.example.json`, replace fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Dates use `YYYY-MM-DDTHH:MM` for Support/VDI local time ranges.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/support_insights/GetCrashes
```
