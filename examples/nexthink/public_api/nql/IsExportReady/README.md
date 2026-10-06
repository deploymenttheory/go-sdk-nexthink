# NQL.IsExportReady

Checks whether an existing export has completed. This example does not start an export or download its results.

Configure `NEXTHINK_API=public`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, `NEXTHINK_CLIENT_ID`, and `NEXTHINK_CLIENT_SECRET`. The API client requires NQL API permissions.

Set `NEXTHINK_EXPORT_ID` to an export returned by `StartNQLExport`.

```sh
go run ./examples/nexthink/public_api/nql/IsExportReady
```

This example was compile-checked; no live export was started during this change.
