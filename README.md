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
    "github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/nql"
)

c, err := nexthink.NewClientFromEnv()
if err != nil {
    log.Fatal(err)
}
result, response, err := c.PublicAPI.NQL.ExecuteNQLV2(context.Background(), &nql.ExecuteRequest{
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

## One client, two API families

Like the [Jamf Pro SDK](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/blob/main/jamfpro/jamfpro.go), all resources are wired through one entry point, [`nexthink.NewClient`](nexthink/nexthink.go).

```go
provider, err := chrome.New("https://your-instance.eu.nexthink.cloud")
// Handle err. Chrome access is explicitly opt-in.
c, err := nexthink.NewClient(&nexthink.AuthConfig{
    Instance: "your-instance",
    Region: "eu",
    PublicAPI: &nexthink.ClientCredentials{
        ClientID: os.Getenv("NEXTHINK_CLIENT_ID"),
        ClientSecret: os.Getenv("NEXTHINK_CLIENT_SECRET"),
    },
    WebAPI: &nexthink.BrowserCredentials{TokenProvider: provider},
})
// Handle err before accessing either family.
links, response, err := c.WebAPI.CollectorManagement.GetDownloadLinks(ctx)
```

Omit `PublicAPI` or `WebAPI` credentials to disable that family. The constructor validates all enabled families before making requests. Public APIs require a client ID and secret. Web APIs require exactly one access token or token provider. A web-only client does not require OAuth client credentials.

`PublicAPI` targets the public integration host; `WebAPI` targets the tenant's browser host. Both use the same transport implementation, with separate token sources and destinations. Resources live under `nexthink/public_api/<resource>` and `nexthink/web_api/<resource>`, with models, CRUD methods, validators, tests and JSON fixtures.

[Web API documentation](docs/web-api.md) describes the observed browser endpoints. Their contracts and permissions may change; the [catalog](nexthink/web_api/operations.json) records live evidence separately from routes found only in frontend code. Chrome owns session login and renewal. Caller-managed access tokens and `auth.TokenProvider` implementations are also supported.

## Examples

[Examples](examples/nexthink) use credentials from the environment. The latest lab run passed 19 public read examples and nine web examples. Read-only fixtures use:

- `NEXTHINK_QUERY_ID` and optionally `NEXTHINK_EXPORT_QUERY_ID`
- `NEXTHINK_EXPORT_ID` for an existing export
- `NEXTHINK_REMOTE_ACTION_ID` / `NEXTHINK_WORKFLOW_ID` for detail lookups

Examples that trigger actions, campaigns, workflows, enrichment, or deletion read an explicit JSON body from `NEXTHINK_REQUEST_FILE`. Inspect its targets before running it. Spark also uses `NEXTHINK_USER_UPN` and optional `NEXTHINK_TIMEZONE`.

```sh
go run ./examples/nexthink/public_api/nql/ExecuteNQLV2
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/Request
```

The web API example accepts `NEXTHINK_WEB_OPERATION` and an optional `NEXTHINK_WEB_REQUEST_FILE` containing `PathParams`, `Query`, and `Body`. Example output can contain tenant data; export examples write files in the current directory.

## Configuration and errors

Pass options from `nexthink`, such as `WithTimeout`, `WithLogger`, `WithProxy`, `WithRetryCount`, and `WithTLSClientConfig`, to the constructor. Scope advanced transport options with `nexthink.WithPublicAPIOptions(client.WithBaseURL(...))` or `nexthink.WithWebAPIOptions(...)`; each applies only to its API family. `WithDebug` prints HTTP method/status while omitting headers and bodies. Authenticated requests reject cross-origin URLs and do not follow redirects. Export downloads use a separate unauthenticated HTTP client for signed download URLs.

Web management LCRUD methods are available through `c.WebAPI.Workflows` and `c.WebAPI.RemoteActions`. See the [web API coverage inventory](docs/web-api-coverage.md) for validated methods and discovered APIs still awaiting implementation.

Service methods return result, response metadata, and error. Inspect `client.APIError` with `errors.As`; status helpers also support wrapped errors. HTTP 207 is a successful transport response with enrichment errors in its payload. `c.WebAPI.GraphQL.Execute` reports GraphQL errors even when HTTP status is 200, retaining partial data.

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
