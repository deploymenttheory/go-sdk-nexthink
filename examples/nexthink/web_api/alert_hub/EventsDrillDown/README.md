# EventsDrillDown

Use the actual monitor name. For an existing alert, copy `startedAt` from its `lastTriggerDateTime` and copy the matching context values. The time is an epoch timestamp in seconds. A zero timestamp is only a placeholder.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/alert_hub/EventsDrillDown
```
