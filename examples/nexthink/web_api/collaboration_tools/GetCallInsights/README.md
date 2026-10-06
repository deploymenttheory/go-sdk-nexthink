# GetCallInsights

`POST /apigateway/collaboration-tools/api/v1/device/{deviceID}/call-insights`

Set `NEXTHINK_COLLECTOR_ID` to the Collector UID used by Device View. This is different from the device UID returned by NQL. The older `NEXTHINK_DEVICE_ID` environment variable remains supported as a fallback.

Copy `request.example.json`, choose `teams` or `zoom`, set a relevant time range and point `NEXTHINK_REQUEST_FILE` to the copy. Dates use `YYYY-MM-DDTHH:MM` with the service's local time context; UTC is the default.

Configure browser authentication using the root quick start, then run:

```sh
export NEXTHINK_API=web
export NEXTHINK_COLLECTOR_ID="<collector UID>"
export NEXTHINK_REQUEST_FILE="<request JSON path>"
go run ./examples/nexthink/web_api/collaboration_tools/GetCallInsights
```

Curl and SDK replay succeeded for both Teams and Zoom in the lab. An empty telemetry window returns markdown explaining that call quality cannot be assessed; HTTP success does not establish that the tenant has call records.
