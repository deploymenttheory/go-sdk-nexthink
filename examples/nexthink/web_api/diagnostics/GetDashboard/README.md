# GetDashboard

Set `issueReference.alertId` to an actual alert occurrence UUID returned by AlertHub. A monitor definition UUID or a diagnostic context ID cannot supply that alert fixture. Use `GetStandaloneDashboard` for a standalone issue definition.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/diagnostics/GetDashboard
```
