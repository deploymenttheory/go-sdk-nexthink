# NQL.ExportToJSON

Starts a saved NQL query export, polls until completion, downloads its results, and writes a new local file with mode `0600`. The output path must not already exist.

Configure `NEXTHINK_API=public`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, `NEXTHINK_CLIENT_ID`, and `NEXTHINK_CLIENT_SECRET`. The API client requires NQL API permissions.

Set `NEXTHINK_OUTPUT_FILE` to the destination path. Set `NEXTHINK_QUERY_ID` to an existing saved NQL API query ID such as `#your_saved_query`.

The service exports CSV; this helper converts it locally into JSON objects with string values.

The helper uses its default polling settings with an eleven-minute example context deadline.

```sh
go run ./examples/nexthink/public_api/nql/ExportToJSON
```

This example was compile-checked; no live export was started during this change.
