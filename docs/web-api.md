# Web APIs

`client.WebAPI` targets `https://{instance}.{region}.nexthink.cloud`. It is constructed through `nexthink.NewClient`, alongside `client.PublicAPI` when both credential sets are supplied. Use `web_api.Operations()` to enumerate the [evidenced catalog](../nexthink/web_api/operations.json).

## Authentication

- `chrome.New(origin)` explicitly enables macOS Chrome automation. It finds an open tab at that exact origin and reads only the current access token and expiry from the signed-in session. Enable Chrome's **View → Developer → Allow JavaScript from Apple Events** and macOS Automation permission when prompted. No browser profile files, refresh tokens, or login credentials are read.
- `auth.StaticToken(value, expiry)` accepts a caller-provided token. A zero expiry means the caller manages its lifetime.
- `auth.TokenProviderFunc` supports application-owned authentication. Providers must observe context cancellation and be safe for concurrent calls.
- `client.PublicAPI.GetTokenManager()` reuses client credentials where permitted. This is not a substitute for a user token: live checks returned different statuses depending on endpoint and identity.

Expired Chrome sessions return `auth.ErrSessionExpired`; sign in again in Chrome. The SDK does not automate interactive authentication. Tokens are attached only to the configured origin, and authenticated redirects are rejected.

## Operations

Typed models cover the observed Collector links, shell result envelopes, user/module data, license status, content listings, device profiles and saved queries. Resource packages cover Collector management, product shell, license status, content administration, device configuration, NQL editor, saved NQL query CRUD, and the eleven observed GraphQL endpoints. Saved-query writes use `nql`; read responses use `nqlQuery`. Create and update return the saved query, including its content ID.

Use `client.WebAPI.Do(ctx, operationID, web_api.Request{...})` for catalog operations with evolving schemas. Responses remain `json.RawMessage`, preserving unknown fields. `PathParams` substitutes escaped route placeholders. `Query` adds URL parameters. `Body` supplies the request payload; `ContentType` optionally overrides the JSON default. No arbitrary URL or Authorization-header override is accepted by this method.

Use `client.WebAPI.GraphQL.Execute(ctx, "graphql.workflows", graphql.GraphQLRequest{Query: ...})` for GraphQL endpoints. The helper reports response errors even on HTTP 200, while retaining partial data. Confirmed `__typename` probes establish endpoint reachability only. Production GraphQL introspection was disabled in the tested management APIs; the catalog does not claim full query/mutation schema coverage.

The catalog distinguishes replayed operations from bundle evidence. A bundle reference establishes an endpoint and method only when both appear in the first-party client code. API gateway 403/404 responses on guessed URLs do not establish that an endpoint exists. Static assets and third-party telemetry services are not SDK operations.

## Discovery limits

The lab exposed 64 first-party entry bundles and 77 manifest assets. These contain additional API base paths and dynamic routes whose full method, payload, or authorization contract is not established. They are not presented as validated SDK methods. Collector configuration writes and telemetry submissions have not been replayed because they affect shared tenant settings or create telemetry records. AmplifyAI and Hypervisor public contracts remain unverified; product names or credential permissions alone are insufficient to invent those endpoints.

Undocumented APIs can change without notice. See [validation evidence](lab-validation.md) for tested behavior and remaining lab work.

## Layout and configuration

`nexthink/nexthink.go` wires `PublicAPIClient` and `WebAPIClient` into one `Client`. Each resource under `nexthink/web_api/` has `constants.go`, `crud.go`, `models.go`, `validators.go`, resource-local tests, and JSON fixtures in `mocks/`. Both families depend on the same `interfaces.HTTPClient` and `client.Transport` implementation.

For environment configuration, set `NEXTHINK_API` to `public` (default), `web`, or `both`. Public authentication reads `NEXTHINK_CLIENT_ID` and `NEXTHINK_CLIENT_SECRET`. Web authentication uses `NEXTHINK_WEB_AUTH=token` with `NEXTHINK_ACCESS_TOKEN`, or `NEXTHINK_WEB_AUTH=chrome` for an existing Chrome session. Both require `NEXTHINK_INSTANCE` and `NEXTHINK_REGION`.

A nil API-family field means that family was not configured. Provider tokens are validated at request time, including expiry and cancellation. Static tokens with a supplied expiry are also checked during construction. A public token is never silently substituted for a web session.

NQL editor calls accept `nql_editor.Document`, `PositionRequest`, and `ValidationRequest`. Completion items retain opaque fields for resolve round trips. Invalid NQL returns diagnostics under HTTP 200; an empty hover or resolve object is also a legitimate successful response.
