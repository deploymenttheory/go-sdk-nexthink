# Go SDK for Nexthink

An unofficial, alpha Go client for Nexthink Infinity. Requires the Go version in [go.mod](go.mod).

```sh
go get github.com/deploymenttheory/go-sdk-nexthink/nexthink
```

## Public APIs

Set `NEXTHINK_CLIENT_ID`, `NEXTHINK_CLIENT_SECRET`, `NEXTHINK_INSTANCE`, and `NEXTHINK_REGION` (`us`, `eu`, `pac`, or `meta`). The client obtains and refreshes OAuth client-credentials tokens.

```go
import (
    "context"
    "log"

    "github.com/deploymenttheory/go-sdk-nexthink/nexthink"
    "github.com/deploymenttheory/go-sdk-nexthink/nexthink/services/nql"
)

c, err := nexthink.NewClientFromEnv()
if err != nil {
    log.Fatal(err)
}
result, response, err := c.NQL.ExecuteNQLV2(context.Background(), &nql.ExecuteRequest{
    QueryID: "#my_saved_query",
    Parameters: map[string]string{"device_name": "lab-mac"},
})
// Handle err before using result or response.
```

The saved query must declare `$device_name`. The public API executes saved query IDs; it does not accept arbitrary query text. Query builders and templates generate text to save separately. `ExecuteQueryBuilder` validates the local builder but still executes the server's saved query.

| Service | Operations |
| --- | --- |
| `NQL` | Execute v1/v2, start export, poll status, download, result helpers |
| `RemoteActions` | List, details, trigger |
| `Workflows` | List, details, trigger v1/v2 |
| `Enrichment` | Enrich fields, including partial-success responses |
| `Campaigns` | Trigger a campaign |
| `DataManagement` | Schedule device deletions and inspect per-device outcomes |
| `Spark` | Hand off a conversation to a user's Teams account |

A successful deletion request schedules work; it does not confirm completed deletion. Spark requires a configured Teams identity. Permissions alone do not establish working device, workflow, campaign, or user fixtures.

## Undocumented APIs

[Experimental APIs](docs/experimental.md) use a separate client on the tenant's browser host. Their contracts and permission requirements may change. [The operation catalog](nexthink/experimental/operations.json) records the evidence and authentication results for each route.

```go
provider, err := chrome.New("https://your-instance.eu.nexthink.cloud")
// Handle err. Creating this provider explicitly opts into Chrome session access.
c, err := experimental.NewClient("your-instance", "eu", provider)
// Handle err.
links, response, err := c.CollectorManagement.GetDownloadLinks(ctx)
```

Imports are `nexthink/auth/chrome` and `nexthink/experimental` under this module. Chrome session access is macOS-only and reads the current access token from an already signed-in tab. Chrome owns login and renewal; the SDK does not extract a refresh token. Callers can instead supply `auth.StaticToken`, their own `auth.TokenProvider`, or a public client's `GetTokenManager()` where the endpoint accepts client credentials.

## Examples

[Examples](examples/nexthink) use credentials from the environment. Read-only fixtures use:

- `NEXTHINK_QUERY_ID` and optionally `NEXTHINK_EXPORT_QUERY_ID`
- `NEXTHINK_EXPORT_ID` for an existing export
- `NEXTHINK_REMOTE_ACTION_ID` / `NEXTHINK_WORKFLOW_ID` for detail lookups

Examples that trigger actions, campaigns, workflows, enrichment, or deletion read an explicit JSON body from `NEXTHINK_REQUEST_FILE`. Inspect its targets before running it. Spark also uses `NEXTHINK_USER_UPN` and optional `NEXTHINK_TIMEZONE`.

```sh
go run ./examples/nexthink/nql/ExecuteNQLV2
NEXTHINK_AUTH=chrome go run ./examples/nexthink/experimental/Request
```

The experimental example accepts `NEXTHINK_EXPERIMENTAL_OPERATION` and an optional `NEXTHINK_EXPERIMENTAL_REQUEST_FILE` containing `PathParams`, `Query`, and `Body`. Example output can contain tenant data; export examples write files in the current directory.

## Configuration and errors

Pass options from `nexthink/client`, such as `WithTimeout`, `WithLogger`, `WithProxy`, `WithRetryCount`, and `WithTLSClientConfig`, to the constructor. `WithDebug` prints HTTP method/status while omitting headers and bodies. Authenticated requests reject cross-origin URLs and do not follow redirects. Export downloads use a separate unauthenticated HTTP client for signed download URLs.

Service methods return result, response metadata, and error. Inspect `client.APIError` with `errors.As`; status helpers also support wrapped errors. HTTP 207 is a successful transport response with enrichment errors in its payload. `experimental.GraphQL` reports GraphQL errors even when HTTP status is 200, retaining partial data.

Exports are CSV. `ExportToJSON` converts CSV locally into JSON string values; it does not request server-side JSON. Use `Compression` (`NONE`, `GZIP`, `ZSTD`) with `StartNQLExport` for raw compressed downloads. Keep signed result URLs out of logs.

## Validation and migration

See [live validation evidence and remaining gaps](docs/lab-validation.md), [breaking changes](docs/migration-lab-validation.md), and the [NQL guides](docs/guides).

```sh
go test ./...
go test -race ./...
go vet ./...
```

Go's `./...` pattern skips directories beginning with `_`; the two client-construction examples under `examples/nexthink/_build_client` need explicit builds.

API reference: [Nexthink developer documentation](https://docs.nexthink.com/api). This project is not affiliated with or endorsed by Nexthink. Licensed under [MIT](LICENSE).
