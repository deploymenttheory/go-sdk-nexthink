# MonitorConfigView

Set `monitorConfigViewInput.alertName` to the actual monitor name from the monitor definition or alert row. This field takes the name, not the monitor UUID or NQL ID.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/alert_hub/MonitorConfigView
```
