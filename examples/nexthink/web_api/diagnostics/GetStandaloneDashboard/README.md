# GetStandaloneDashboard

Use the standalone issue definition from the UI route, such as `exec-crashes-1`, and its binary parameter. `crashes` is a diagnostic context ID and does not identify this standalone dashboard. `issueQueryParameter.value` is a JSON-encoded array containing an observed binary name.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/diagnostics/GetStandaloneDashboard
```
