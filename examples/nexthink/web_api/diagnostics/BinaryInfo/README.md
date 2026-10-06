# BinaryInfo

Select an executable first. The browser only requests binary information when `details.params` contains an `executables` parameter. Its `value` is a JSON-encoded array of binary names, for example `"[\"Safari\"]"`. Discover names with `binaries | list binary.name`; choose a name returned by your tenant.

Copy `request.example.json`, replace the fixture values, and set `NEXTHINK_REQUEST_FILE` to that file. Where a timeframe is required, select a recent interval within the tenant’s retention period and use the matching timezone and UTC offset.

Configure authentication as described in the [Web API examples](../../README.md), then run from the repository root:

```sh
go run ./examples/nexthink/web_api/diagnostics/BinaryInfo
```
