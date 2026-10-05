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

Both API families now use `nexthink.NewClient(&nexthink.AuthConfig{...}, options...)` or `nexthink.NewClientFromEnv()`. The positional four-string constructor has been replaced by a typed configuration. Supply `PublicAPI: &nexthink.ClientCredentials{ClientID: ..., ClientSecret: ...}` and/or `WebAPI: &nexthink.BrowserCredentials{AccessToken: ...}` (or `TokenProvider`).

- Change `client.NQL`, `client.Workflows`, etc. to `client.PublicAPI.NQL`, `client.PublicAPI.Workflows`, etc.
- Move imports from `nexthink/services/<resource>` to `nexthink/public_api/<resource>`.
- Replace the separate `experimental.NewClient` with the main constructor and use `client.WebAPI.<Resource>`. Models live in `nexthink/web_api/<resource>`; the `experimental` package has been removed.
- Use `client.WebAPI.GraphQL.Execute` for management GraphQL requests and `web_api.Operations()` for endpoint evidence.
- Constructor options now come from `nexthink`. Scope low-level `client` options with `WithPublicAPIOptions` or `WithWebAPIOptions`. `NEXTHINK_BASE_URL` is replaced by these explicit per-family overrides.
- Token-manager operations now live under `client.PublicAPI`.
- Examples are grouped under `examples/nexthink/public_api` and `examples/nexthink/web_api`. The web example uses `NEXTHINK_WEB_OPERATION` and `NEXTHINK_WEB_REQUEST_FILE`.

Browser and client-credentials identities have different permissions. Endpoint reachability does not establish permission for every operation or GraphQL mutation.

## Typed web helper signatures

`DeviceConfiguration.SetProfiles` now takes `*device_configuration.SaveProfilesRequest` and returns `*ProfilesResponse`. Build `Settings` from explicit `{profileId,name,newValue}` changes; the service updates settings on existing profiles rather than replacing the collection. `ProductShell.ValidateClaims` now takes `*product_shell.ClaimsRequest` and returns `*ClaimsResponse`; inspect `Result.Result` for the boolean. `Claims` is a pointer to a slice so an explicitly empty array differs from an omitted field.

The added web management resources are initialized by the existing root constructor. No additional clients or authentication implementations are needed. Product-specific IDs and revisions are represented separately in their request types.
