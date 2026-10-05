# Lab validation migration

This alpha revision intentionally corrects incompatible wire models and invalid examples.

- Import `github.com/deploymenttheory/go-sdk-nexthink`, replacing the previous `go-api-sdk-nexthink` module path.
- Replace `ExecuteRequest.Platform` and `ExportRequest.Platform` with `Parameters`. Parameters work only when declared by the saved query.
- `ExportRequest.Format` controls local conversion in `ExportWorkflow`; it is not sent to Nexthink. The server exports CSV. Local JSON conversion keeps CSV values as strings. `Compression` is the server-side option.
- Failed exports now return an error from wait helpers. Cancellation and deadline errors remain discoverable with `errors.Is`.
- `Workflow.TriggerMethods` is an object with boolean fields. Use `.Enabled()` for a display list. Versions contain `Version`, `Status`, `Definition`, `Parameters`, `Valid`, `HasDevice`, and `HasUser`.
- Workflow detail requests use `nqlId`. Remote-action details use `nql-id`. Both accept IDs returned by their list endpoints, including IDs without a leading `#`. Saved NQL execution IDs still require `#`.
- Query-builder filters and event computations preserve call order. A filter before `Compute` applies within the event relation; a filter after it applies to the computed result. Adjacent computations are combined. String helpers escape their arguments as literals; use `Where` for explicit NQL expressions or enumeration comparisons.
- Corrected templates filter event fields before computing, filter page counts before summarizing, compare enumerations without string quotes, and calculate memory usage from used memory and installed memory.
- Enrichment accepts up to 10,000 operations, matching the published schema.
- Authenticated transports reject cross-origin URLs and redirects. Base URLs cannot include embedded credentials, query strings, or fragments. Custom transport/proxy/TLS settings also apply to the separate token HTTP client.
- Examples that write data require `NEXTHINK_REQUEST_FILE`; read examples use environment-provided query, workflow, remote-action, and export IDs.

Undocumented services are opt-in through `nexthink/experimental`. Browser and client-credentials identities have different permissions. Endpoint reachability does not establish permission to perform every operation or GraphQL mutation.
