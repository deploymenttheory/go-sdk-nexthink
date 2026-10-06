# Quick Start

The SDK has one entry point, `nexthink.NewClient`, and two optional API families: `PublicAPI` and `WebAPI`. `NewClientFromEnv` constructs the same client from environment variables. Choose the family that exposes the operation you need; each has its own authentication.

## Install

Use the Go version specified in [go.mod](../../go.mod). In a new project:

```sh
mkdir nexthink-example
cd nexthink-example
go mod init example.com/nexthink-example
go get github.com/deploymenttheory/go-sdk-nexthink/nexthink
```

Save one of the complete programs below as `main.go`, configure its environment, and run `go run .`. The first calls are read-only; they do not require a device to be enrolled.

## First public API call

Create an API client in Nexthink with the permissions needed for the service you intend to call. This example needs access to the Remote Actions API. Provide credentials through your normal secret-management mechanism; the placeholders below describe the required variables.

```sh
export NEXTHINK_API=public
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_CLIENT_ID='your-client-id'
export NEXTHINK_CLIENT_SECRET='your-client-secret'
```

`NEXTHINK_INSTANCE` is the tenant name, without `https://` or a domain. Valid regions are `us`, `eu`, `pac`, and `meta`.

```go
package main

import (
    "context"
    "errors"
    "fmt"
    "log"
    "time"

    "github.com/deploymenttheory/go-sdk-nexthink/nexthink"
    "github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
)

func main() {
    c, err := nexthink.NewClientFromEnv(
        nexthink.WithTimeout(30*time.Second),
        nexthink.WithRetryCount(0),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    if c.PublicAPI == nil {
        log.Fatal("set NEXTHINK_API=public or both")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    actions, response, err := c.PublicAPI.RemoteActions.ListRemoteActions(ctx)
    if err != nil {
        var apiErr *client.APIError
        if errors.As(err, &apiErr) {
            log.Printf("Nexthink returned HTTP %d", apiErr.StatusCode)
        }
        log.Fatal(err)
    }
    fmt.Printf("HTTP %d: %d remote actions\n", response.StatusCode, len(actions))
}
```

An empty list is a valid result. Listing available actions does not execute them. The SDK obtains and renews public OAuth tokens using the client ID and secret.

## First Web API call

The web APIs provide management and analytics contracts used by the browser UI. They require the appropriate user identity and permissions, which differ from a public API client's credentials.

For unattended CI or local password-only login, install the browser runtime during runner preparation:

```sh
# macOS/Windows, or Linux with system dependencies already installed:
go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install chromium
# On a Linux runner that also needs system libraries, use instead:
# go run github.com/mxschmitt/playwright-go/cmd/playwright@v0.6201.1 install --with-deps chromium

export NEXTHINK_API=web
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_WEB_AUTH=password
export NEXTHINK_USERNAME=your-local-account
export NEXTHINK_PASSWORD=your-local-password
```

Inject the username and password from your secret manager. This mode launches its own headless Chromium and needs neither a display nor desktop Chrome. Only local password-only login is supported; SSO, MFA and password-change challenges return errors. The browser and driver version must match the SDK dependency; the SDK never downloads them during a request. Unset `DEBUGP` and `PWDEBUG`; any nonempty value (including `0`) is rejected before browser startup to prevent protocol credential logging or interactive debugging. See the [complete password example](../../examples/nexthink/_build_client/headless_password/README.md) for a Linux container and a GitHub Actions job.

For a caller-managed access token instead:

```sh
export NEXTHINK_API=web
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_WEB_AUTH=token
export NEXTHINK_ACCESS_TOKEN='your-current-user-access-token'
```

Alternatively, on macOS, sign in to `https://your-instance.eu.nexthink.cloud` in Chrome and keep that tab open. Enable **View → Developer → Allow JavaScript from Apple Events** in Chrome and permit macOS Automation access when prompted. Then select the opt-in provider:

```sh
export NEXTHINK_API=web
export NEXTHINK_INSTANCE=your-instance
export NEXTHINK_REGION=eu
export NEXTHINK_WEB_AUTH=chrome
```

Use this complete program to retrieve the first page of application configurations:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/deploymenttheory/go-sdk-nexthink/nexthink"
)

func main() {
    c, err := nexthink.NewClientFromEnv(
        nexthink.WithTimeout(30*time.Second),
        nexthink.WithRetryCount(0),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    if c.WebAPI == nil {
        log.Fatal("set NEXTHINK_API=web or both")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
    defer cancel()
    result, response, err := c.WebAPI.Applications.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("HTTP %d: %d applications on this page, %d total\n",
        response.StatusCode, len(result.Items), result.Total)
}
```

The request context includes authentication, so this example allows two minutes for first login plus the API call. `WithTimeout` configures the HTTP transport separately. `Applications.List` retrieves one page. Use its `ListOptions` and returned `Links.Next`/`Total` to request subsequent pages. The [resource guide](../../examples/nexthink/web_api/applications/README.md) covers additional operations. Browser endpoints can change, and a tenant's permissions or feature flags may restrict access; consult the [coverage inventory](../web-api-coverage.md) for established contracts and live-validation limits.

## Authentication and token lifetime

| Setting | Meaning |
| --- | --- |
| `NEXTHINK_API` | `public` (default), `web`, or `both` |
| `NEXTHINK_INSTANCE` | Tenant name, without a URL or domain |
| `NEXTHINK_REGION` | `us`, `eu`, `pac`, or `meta` |
| `NEXTHINK_CLIENT_ID`, `NEXTHINK_CLIENT_SECRET` | Required when public APIs are enabled |
| `NEXTHINK_WEB_AUTH` | `token` (default), `password`, or `chrome`, when web APIs are enabled |
| `NEXTHINK_USERNAME`, `NEXTHINK_PASSWORD` | Local account credentials required for `password` authentication |
| `NEXTHINK_LOGIN_TIMEOUT` | Password login timeout as a Go duration; default `90s` |
| `NEXTHINK_BROWSER_PROXY` | Optional proxy for browser login and token renewal, separate from API HTTP proxy options |
| `NEXTHINK_BROWSER_EXECUTABLE_PATH` | Optional compatible Chromium executable; the Playwright driver is still required |
| `NEXTHINK_ACCESS_TOKEN` | Required for web `token` authentication; supply only the token, without the `Bearer ` prefix |

Set `NEXTHINK_API=both` and supply both credential sets to enable both families. A public-only client has a nil `WebAPI`; a web-only client has a nil `PublicAPI`. Constructor validation checks every enabled family before creating transports. Creating a public client obtains its initial OAuth token. Browser providers retrieve tokens at request time, and each endpoint evaluates authorization when called.

For configuration loaded from your own secret manager, pass `AuthConfig` directly. This complete example constructs both families with an opt-in Chrome provider; it obtains the initial public OAuth token but makes no resource requests:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/deploymenttheory/go-sdk-nexthink/nexthink"
    "github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth/chrome"
)

func main() {
    instance, region := os.Getenv("NEXTHINK_INSTANCE"), os.Getenv("NEXTHINK_REGION")
    provider, err := chrome.New(fmt.Sprintf("https://%s.%s.nexthink.cloud", instance, region))
    if err != nil {
        log.Fatal(err)
    }
    c, err := nexthink.NewClient(&nexthink.AuthConfig{
        Instance: instance,
        Region:   region,
        PublicAPI: &nexthink.ClientCredentials{
            ClientID:     os.Getenv("NEXTHINK_CLIENT_ID"),
            ClientSecret: os.Getenv("NEXTHINK_CLIENT_SECRET"),
        },
        WebAPI: &nexthink.BrowserCredentials{TokenProvider: provider},
    })
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()
    fmt.Printf("Public configured: %t; web configured: %t\n", c.PublicAPI != nil, c.WebAPI != nil)
}
```

Omit the `PublicAPI` or `WebAPI` credential field to disable that family. Web credentials accept exactly one `AccessToken`, `TokenProvider`, or `UsernamePassword`. Set `defer c.Close()` after successful construction. It closes only SDK-owned authentication resources; application-provided token providers remain caller-owned:

- **Headless password:** the SDK obtains tokens through local password login in an isolated Chromium context, then closes the browser. Tokens remain in memory. It caches valid tokens and renews through a supported refresh exchange when available or fresh headless login otherwise. Concurrent requests share authentication work. Invalid credentials or unsupported challenges stop automatic credential submissions for that client. This mode does not save reusable session files or automate SSO/MFA. Controlled browser fixtures do not establish your tenant's login or refresh policy.
- **Public OAuth:** the SDK acquires and refreshes tokens from client credentials.
- **Chrome provider:** the SDK reads the current access token and expiry from the matching signed-in tab. Chrome owns login and renewal. The SDK does not read browser cookies or refresh tokens, or automate interactive login. An expired session returns `auth.ErrSessionExpired`; sign in again in Chrome.
- **Static web token:** renewal is the caller's responsibility. Programmatic `BrowserCredentials.ExpiresAt` enables expiry validation. The environment configuration supplies no expiry; it does not infer or renew the token's lifetime.
- **Application-owned provider:** implement `auth.TokenProvider` or use `auth.TokenProviderFunc`. The provider supplies current tokens, manages renewal, observes context cancellation, and must be safe for concurrent calls. With a provider, its returned token owns expiry; do not also set `BrowserCredentials.ExpiresAt`.

For explicit password configuration, use `WebAPI: &nexthink.BrowserCredentials{UsernamePassword: &nexthink.UsernamePasswordCredentials{Username: username, Password: password}}`. Optional fields are `LoginTimeout`, `BrowserProxy`, and `BrowserExecutablePath`. The [complete Go example](../../examples/nexthink/_build_client/headless_password/main.go) includes cleanup, context deadlines and a read-only call.

The public integration host and browser tenant host use separate transports and token sources. Public credentials are never silently substituted for a browser session.

## Configuration and errors

Pass `nexthink.WithTimeout`, `WithRetryCount`, `WithLogger`, `WithProxy`, `WithTLSClientConfig`, or `WithTransport` to either constructor. See [root options](../../nexthink/options.go) and the [transport package reference](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-nexthink/nexthink/client) for signatures. Scope advanced `client.ClientOption` values with `nexthink.WithPublicAPIOptions(...)` or `nexthink.WithWebAPIOptions(...)`; for example, a base URL override should apply only to the intended API family. `WithDebug` omits headers and bodies from method/status logging.

Browser login uses Chromium's certificate trust configuration. Go HTTP TLS options and custom round trippers do not configure Chromium; provision runner/browser trust and `BrowserProxy` separately. Browser TLS verification is never silently disabled. A Web API HTTP 401 invalidates the password provider token for the next call but does not replay the failed operation; a 403 remains a permission error. Authentication errors omit credential values.

Most resource methods return a typed result, `*interfaces.Response`, and an error. Methods with no decoded acknowledgment return response metadata and error. Handle errors before dereferencing results; validation or transport failures may have no HTTP response. HTTP failures retain metadata and the response body. The public example above shows `errors.As` with `*client.APIError`.

Some responses need additional checks:

- GraphQL can return HTTP 200 with errors and partial data. The GraphQL helper returns those errors while preserving the available data and response metadata.
- Enrichment HTTP 207 is a successful transport response whose payload can contain individual failures.
- Some browser APIs return business status inside a successful HTTP envelope. Inspect the typed status/result fields described by that resource's guide.
- A successful device deletion request schedules work; it does not prove deletion has finished. Spark handoff requires a configured Teams identity.

Authenticated requests reject cross-origin URLs and do not follow redirects. Export downloads use a separate unauthenticated HTTP client for signed result URLs; keep those URLs out of logs.

## Next calls and examples

From a repository checkout, the [public examples](../../examples/nexthink/public_api) and [web example index](../../examples/nexthink/web_api/README.md) provide runnable programs and resource-specific request formats. Depending on the example, inputs include:

| Variable | Used for |
| --- | --- |
| `NEXTHINK_QUERY_ID` / `NEXTHINK_EXPORT_QUERY_ID` | Saved NQL query IDs |
| `NEXTHINK_EXPORT_ID` | An existing export |
| `NEXTHINK_REMOTE_ACTION_ID` / `NEXTHINK_WORKFLOW_ID` | Public detail lookups |
| `NEXTHINK_REQUEST_FILE` | Explicit JSON request bodies, including write/trigger examples |
| `NEXTHINK_CONTENT_ID` | Resource identifiers in examples that use this variable |
| `NEXTHINK_USER_UPN` / `NEXTHINK_TIMEZONE` | Spark recipient and optional timezone |

These variables belong to the examples, not the client's authentication configuration. Inspect each example's guide for the required variables and side effects before executing it. Examples can print tenant data; export examples can write files in the current directory.

Public NQL execution accepts a saved query ID, not arbitrary query text. Query builders and templates generate text to save separately; `ExecuteQueryBuilder` still executes the server's saved query. Query parameters must match the saved query's declarations. [NQL guides](nql-query-building.md) explain the workflow.

NQL exports are CSV. `ExportToJSON` converts downloaded CSV to JSON string values locally. `StartNQLExport` supports `NONE`, `GZIP`, and `ZSTD` compression for raw downloads; see the [export guide](nql-export-workflow.md).

The generic [web request example](../../examples/nexthink/web_api/Request) accepts `NEXTHINK_WEB_OPERATION` and optional `NEXTHINK_WEB_REQUEST_FILE` with `PathParams`, `Query`, and `Body`. Prefer typed resource methods when available, especially for uploads and operations with special headers or body encoding.

## Developing the SDK

Run checks from the repository root:

```sh
go test ./...
go test -race ./...
go vet ./...
go build -o /tmp/nexthink-basic-client ./examples/nexthink/_build_client/new_client
go build -o /tmp/nexthink-logged-client ./examples/nexthink/_build_client/new_client_with_logger
go build -o /tmp/nexthink-password-client ./examples/nexthink/_build_client/headless_password
```

Go's `./...` pattern skips directories beginning with `_`, so client-construction examples need explicit builds. After installing the pinned browser runtime, run `NEXTHINK_BROWSER_TEST=1 go test -race ./nexthink/auth/password -run TestBrowser -timeout 180s` for local controlled browser fixtures. The separate browser-auth workflow runs those fixtures on Linux, macOS and Windows without lab secrets. See [migration notes](../migration-lab-validation.md), [live validation evidence](../lab-validation.md), and [CONTRIBUTING.md](../../CONTRIBUTING.md) for further context.
