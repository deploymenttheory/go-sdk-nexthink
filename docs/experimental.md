# Experimental browser APIs

`nexthink/experimental` targets `https://{instance}.{region}.nexthink.cloud`. It is independent of the public API host and requires an explicit `auth.TokenProvider`. Use `experimental.Operations()` to enumerate the [evidenced catalog](../nexthink/experimental/operations.json).

## Authentication

- `chrome.New(origin)` explicitly enables macOS Chrome automation. It finds an open tab at that exact origin and reads only the current access token and expiry from the signed-in session. Enable Chrome's **View → Developer → Allow JavaScript from Apple Events** and macOS Automation permission when prompted. No browser profile files, refresh tokens, or login credentials are read.
- `auth.StaticToken(value, expiry)` accepts a caller-provided token. A zero expiry means the caller manages its lifetime.
- `auth.TokenProviderFunc` supports application-owned authentication. Providers must observe context cancellation and be safe for concurrent calls.
- `publicClient.GetTokenManager()` reuses client credentials where permitted. This is not a substitute for a user token: live checks returned different statuses depending on endpoint and identity.

Expired Chrome sessions return `auth.ErrSessionExpired`; sign in again in Chrome. The SDK does not automate interactive authentication. Tokens are attached only to the configured origin, and authenticated redirects are rejected.

## Operations

Typed wrappers cover Collector download/configuration reads, product-shell reads, license status, and saved NQL query CRUD. Saved-query writes use `nql`; read responses use `nqlQuery`. Create and update return the saved query, including its content ID.

Use `Do(ctx, operationID, Request{...})` for catalog operations with evolving schemas. Responses remain `json.RawMessage`, preserving unknown fields. `PathParams` substitutes escaped route placeholders. `Query` adds URL parameters. `Body` supplies the request payload; `ContentType` optionally overrides the JSON default. No arbitrary URL or Authorization-header override is accepted by this method.

Use `GraphQL(ctx, "graphql.workflows", GraphQLRequest{Query: ...})` for GraphQL endpoints. The helper reports response errors even on HTTP 200, while retaining partial data. Confirmed `__typename` probes establish endpoint reachability only. Production GraphQL introspection was disabled in the tested management APIs; the catalog does not claim full query/mutation schema coverage.

The catalog distinguishes replayed operations from bundle evidence. A bundle reference establishes an endpoint and method only when both appear in the first-party client code. API gateway 403/404 responses on guessed URLs do not establish that an endpoint exists. Static assets and third-party telemetry services are not SDK operations.

## Discovery limits

The lab exposed 64 first-party entry bundles and 77 manifest assets. These contain additional API base paths and dynamic routes whose full method, payload, or authorization contract is not established. They are not presented as validated SDK methods. Collector configuration writes and telemetry submissions have not been replayed because they affect shared tenant settings or create telemetry records. AmplifyAI and Hypervisor public contracts remain unverified; product names or credential permissions alone are insufficient to invent those endpoints.

Undocumented APIs can change without notice. See [validation evidence](lab-validation.md) for tested behavior and remaining lab work.
