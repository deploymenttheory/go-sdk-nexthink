# NQL.ExportWorkflow

Starts a saved NQL query export, polls until completion, downloads its results, and writes a new local file with mode `0600`. The output path must not already exist.

Configure `NEXTHINK_API=public`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, `NEXTHINK_CLIENT_ID`, and `NEXTHINK_CLIENT_SECRET`. The API client requires NQL API permissions.

Set `NEXTHINK_OUTPUT_FILE` to the destination path. Copy [request.example.json](request.example.json), replace `queryId` with an existing saved NQL API query ID, add its required parameters, and set `NEXTHINK_REQUEST_FILE` to that file.

`NEXTHINK_EXPORT_FORMAT` accepts `csv` (default) or `json`. JSON conversion runs locally on downloaded CSV; JSON values remain strings. Status transitions are written to stderr.

The example polls every five seconds with a ten-minute workflow timeout and an eleven-minute context deadline.

```sh
go run ./examples/nexthink/public_api/nql/ExportWorkflow
```

This example was compile-checked; no live export was started during this change.
