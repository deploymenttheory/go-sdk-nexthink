# Get

`GET /apigateway/atl/support-checklist-values-be/api/v1/device/{deviceID}/checklist/{checklistID}`

This reads the UI API.

Set `NEXTHINK_DEVICE_ID`, `NEXTHINK_CHECKLIST_ID`.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/support_checklists/Get
```
