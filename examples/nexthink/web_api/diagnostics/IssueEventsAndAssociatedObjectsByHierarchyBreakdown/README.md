# IssueEventsAndAssociatedObjectsByHierarchyBreakdown

Obtain `hierarchyId` and `levelDepth` from `HierarchyBreakdownDimensions`. If that call returns no hierarchy dimensions, this example has no valid fixture; do not submit the placeholder hierarchy ID.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/diagnostics/IssueEventsAndAssociatedObjectsByHierarchyBreakdown
```
