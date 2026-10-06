# GetMetricThreshold

Set `metricId` to a leaf ID returned by `DEXScores.GetLeaves`, preserving its integer value as a string. Aggregate score IDs are accepted by other score operations but can fail this threshold query. Do not reuse an arbitrary parent score ID.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/dex_scores/GetMetricThreshold
```
