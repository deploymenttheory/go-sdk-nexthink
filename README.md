# Go SDK for Nexthink

[![Go Report Card](https://goreportcard.com/badge/github.com/deploymenttheory/go-sdk-nexthink)](https://goreportcard.com/report/github.com/deploymenttheory/go-sdk-nexthink)
[![GoDoc](https://pkg.go.dev/badge/github.com/deploymenttheory/go-sdk-nexthink/nexthink.svg)](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-nexthink/nexthink)
[![License](https://img.shields.io/github/license/deploymenttheory/go-sdk-nexthink)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/deploymenttheory/go-sdk-nexthink)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/deploymenttheory/go-sdk-nexthink)](https://github.com/deploymenttheory/go-sdk-nexthink/releases)
[![Tests](https://github.com/deploymenttheory/go-sdk-nexthink/actions/workflows/test.yml/badge.svg)](https://github.com/deploymenttheory/go-sdk-nexthink/actions/workflows/test.yml)
[![Lint](https://github.com/deploymenttheory/go-sdk-nexthink/actions/workflows/go-lint.yml/badge.svg)](https://github.com/deploymenttheory/go-sdk-nexthink/actions/workflows/go-lint.yml)
![Status: Alpha](https://img.shields.io/badge/status-alpha-yellow)

An unofficial Go client for Nexthink Infinity, supporting the public integration APIs and the APIs used by the Nexthink web UI. Both API families share one entry point, `nexthink.NewClient`, with typed resources, request validation, response metadata, and a shared HTTP transport.

The SDK is **alpha**. Public and undocumented web contracts have different stability guarantees. The [coverage inventory](docs/web-api-coverage.md) and [lab validation report](docs/lab-validation.md) distinguish live-tested operations from contracts established through frontend source and unit tests.

The current [systematic acceptance report](docs/acceptance/README.md) records **418 passed methods and 274 outstanding validation gaps across 692 exported methods** in both API families. Implemented coverage does not imply complete live validation.

The latest product-area expansion adds AI Tools/governance, Amplify configuration, Workspace, and product/device/user administration resources. Collaboration dashboards and call views reuse the existing dashboard and data-exploration services. See the [role-to-resource reconciliation](docs/web-api-coverage.md#product-coverage-reconciliation-role-permissions-6-october-2026) for inspected source coverage and remaining validation work; this is not an exhaustive inventory of every Nexthink endpoint.

## Why include the Web UI APIs?

Nexthink's public integration APIs primarily execute saved queries, read available actions and workflows, trigger operations, enrich data, and schedule device deletion. Managing the content behind those operations requires additional APIs: for example, creating or editing a workflow or remote action, configuring dashboards and integrations, or retrieving analytics used by the UI.

The SDK exposes those observed browser contracts through `client.WebAPI`, alongside `client.PublicAPI`. Management resources provide list, create, read, update, and delete operations where those operations exist, plus their observed auxiliary calls. The [web API guide](docs/web-api.md) explains REST and GraphQL coverage and discovery limits; the SDK does not assume every resource supports every operation.

| API family | Typical use | Authentication |
| --- | --- | --- |
| `client.PublicAPI` | Saved NQL execution/export, action and workflow triggers, enrichment | API client ID and secret; OAuth tokens acquired and refreshed by the SDK |
| `client.WebAPI` | Content management, configuration, analytics, browser administration | Local username/password through headless Chromium, a user access token, or a token provider |

The credential types are not interchangeable. An API client's permissions do not establish a signed-in browser identity. Web access also depends on the user's permissions, tenant features, and available fixtures. Undocumented endpoints can change without notice.

## Quick Start

```sh
go get github.com/deploymenttheory/go-sdk-nexthink/nexthink
```

Follow the **[Quick Start Guide](docs/guides/quick-start.md)** for installation, complete read-only programs for both API families, environment configuration, headless CI authentication, browser session setup, and error handling. Use the Go version required by [go.mod](go.mod).

## Examples

The [examples directory](examples/nexthink) contains runnable programs using the same SDK entry point:

- **[Public APIs](examples/nexthink/public_api):** NQL, remote actions, workflows, campaigns, enrichment, data management, and Spark.
- **[Web APIs](examples/nexthink/web_api/README.md):** resource guides and examples for management, analytics, integrations, identity, and support operations.
- **[Headless password authentication](examples/nexthink/_build_client/headless_password/README.md):** local-account login, explicit browser installation, a Linux container, and a GitHub Actions example.

From a checkout with the appropriate credentials configured:

```sh
NEXTHINK_API=public go run ./examples/nexthink/public_api/remote_actions/ListRemoteActions
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/applications/List
```

Mutation examples require explicit targets or request files. Each resource guide describes its inputs and effects; examples that execute actions, send messages, or change configuration perform those operations when run. Output can contain tenant data.

## HTTP Client Configuration

Both families use the same transport implementation with separate authentication and destinations. Pass common [client options](nexthink/options.go) to `NewClient` or `NewClientFromEnv`:

| Option | Purpose |
| --- | --- |
| `WithTimeout` | Set the HTTP request timeout; request contexts can impose a shorter deadline |
| `WithRetryCount` | Configure retry attempts; use zero when explicitly handling retries yourself |
| `WithLogger` | Supply a zap logger |
| `WithProxy` | Configure a proxy |
| `WithTLSClientConfig` | Supply TLS configuration |
| `WithTransport` | Supply an HTTP round tripper |
| `WithDebug` | Log request method/status without headers or bodies |
| `WithPublicAPIOptions` / `WithWebAPIOptions` | Apply advanced `client.ClientOption` settings to one family |

Authenticated requests are restricted to the configured origin and do not follow redirects. Signed export downloads use a separate unauthenticated HTTP client. See [configuration and response handling](docs/guides/quick-start.md#configuration-and-errors) for examples and important response distinctions.

## Configuration

`nexthink.NewClientFromEnv` selects the enabled families with `NEXTHINK_API=public`, `web`, or `both` (default: `public`). Set `NEXTHINK_INSTANCE` to the tenant name and `NEXTHINK_REGION` to `us`, `eu`, `pac`, or `meta`.

- Public authentication uses `NEXTHINK_CLIENT_ID` and `NEXTHINK_CLIENT_SECRET`.
- Web authentication uses `NEXTHINK_WEB_AUTH=password` with `NEXTHINK_USERNAME` and `NEXTHINK_PASSWORD` for a local password-only account; `token` with `NEXTHINK_ACCESS_TOKEN`; or `chrome` for an existing signed-in Chrome tab on macOS.
- For application-owned configuration, pass `nexthink.AuthConfig` to `NewClient`; enable one or both credential fields. A disabled family is nil. Each enabled family's required fields are validated before its transport is constructed.

The [quick-start authentication section](docs/guides/quick-start.md#authentication-and-token-lifetime) explains token lifetimes, provider responsibilities, and browser setup. Password mode launches isolated headless Chromium and manages token acquisition without desktop Chrome. Install its pinned driver and browser during runner preparation; no browser is downloaded during API calls. SSO and MFA are not supported by this mode. Chrome mode continues to use an existing user session. Call `defer c.Close()` to release SDK-owned authentication resources.

## Documentation

- [Quick Start Guide](docs/guides/quick-start.md)
- [Web API guide](docs/web-api.md), [coverage inventory](docs/web-api-coverage.md), and [operation catalog](nexthink/web_api/operations.json)
- [NQL guides](docs/guides) and [NQL reference](docs/reference/nql-reference.md)
- [Live validation evidence](docs/lab-validation.md) and [migration notes](docs/migration-lab-validation.md)
- [Go package reference](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-nexthink/nexthink)
- [Nexthink API documentation](https://docs.nexthink.com/api)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) and [GitHub Issues](https://github.com/deploymenttheory/go-sdk-nexthink/issues). New endpoints should include evidenced request/response contracts, resource tests and JSON fixtures, and runnable examples. See the [quick-start verification commands](docs/guides/quick-start.md#developing-the-sdk) for local checks.

## License

Licensed under the [MIT License](LICENSE). This community SDK is not affiliated with or endorsed by Nexthink.
