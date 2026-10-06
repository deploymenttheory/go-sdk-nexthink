# Web APIs

`client.WebAPI` targets `https://{instance}.{region}.nexthink.cloud`. It is constructed through `nexthink.NewClient`, alongside `client.PublicAPI` when both credential sets are supplied. Use `web_api.Operations()` to enumerate the [evidenced catalog](../nexthink/web_api/operations.json).

## Authentication

- `chrome.New(origin)` explicitly enables macOS Chrome automation. It finds an open tab at that exact origin and reads only the current access token and expiry from the signed-in session. Enable Chrome's **View → Developer → Allow JavaScript from Apple Events** and macOS Automation permission when prompted. No browser profile files, refresh tokens, or login credentials are read.
- `auth.StaticToken(value, expiry)` accepts a caller-provided token. A zero expiry means the caller manages its lifetime.
- `auth.TokenProviderFunc` supports application-owned authentication. Providers must observe context cancellation and be safe for concurrent calls.
- `client.PublicAPI.GetTokenManager()` reuses client credentials where permitted. This is not a substitute for a user token: live checks returned different statuses depending on endpoint and identity.

Expired Chrome sessions return `auth.ErrSessionExpired`; sign in again in Chrome. The SDK does not automate interactive authentication. Tokens are attached only to the configured origin, and authenticated redirects are rejected.

Legacy portal search and branding methods accept an explicit per-call `auth.PortalSession`. These credentials are restricted to three fixed POST routes; those calls omit bearer authentication and do not consult the token provider. Response cookies are not persisted. Root WebAPI construction still requires browser credentials or a token provider. See the [global search](../examples/nexthink/web_api/global_search/README.md) and [appearance](../examples/nexthink/web_api/appearance/README.md) examples for the separate portal-session requirements.

## Operations

Typed resource packages cover content management, integrations, identity, analytics, Collector management, browser administration, and their observed auxiliary operations. The [coverage inventory](web-api-coverage.md) records 557 API operation methods, a dashboard proxy routing helper, and 17 cataloged GraphQL gateways. Saved-query writes use `nql`; read responses use `nqlQuery`. Create and update return the saved query, including its content ID.

Use `client.WebAPI.Do(ctx, operationID, web_api.Request{...})` for catalog operations with evolving schemas. Responses remain `json.RawMessage`, preserving unknown fields. `PathParams` substitutes escaped route placeholders. `Query` adds URL parameters. `Body` supplies the request payload; `ContentType` optionally overrides the JSON default. No arbitrary URL or Authorization-header override is accepted by this method.

Use the typed `Assets.Create` and `Assets.Update` methods for uploads. They supply the required filename header, which the generic `Do` request does not expose, and construct the plain-text data URL from file bytes.

Use `client.WebAPI.GraphQL.Execute(ctx, "graphql.workflows", graphql.GraphQLRequest{Query: ...})` for GraphQL endpoints. The helper reports response errors even on HTTP 200, while retaining partial data. Confirmed `__typename` probes establish endpoint reachability only. Production GraphQL introspection was disabled in the tested management APIs; the catalog does not claim full query/mutation schema coverage.

The catalog distinguishes replayed operations from bundle evidence. A bundle reference establishes an endpoint and method only when both appear in the first-party client code. API gateway 403/404 responses on guessed URLs do not establish that an endpoint exists. Bundled JavaScript and other static application files are not management APIs. `Appearance` covers the observed branding image operations, and `Observability` submits to the fixed Nexthink telemetry proxy; it does not send authenticated requests to arbitrary third-party hosts.

## Typed management resources

`client.WebAPI.Workflows` provides `List`, `Create`, `Get`, `Update`, `Delete` and `Export`. `client.WebAPI.RemoteActions` provides `List`, `Create`, `Get`, `Update`, `Delete`, `GetForView`, `GetContentVolume`, `InspectBashScript`, `InspectPowerShellScript` and `GetPowerShellSignature`. These use the same browser-authenticated transport and return typed GraphQL data envelopes plus HTTP response metadata. Partial data remains available alongside `graphql.GraphQLErrors`.

Script inspection accepts bytes: a macOS tar.gz archive or UTF-8 PowerShell source including its BOM. The SDK performs base64 encoding. Inspection does not execute a script. See the [coverage inventory](web-api-coverage.md) for live evidence, the LCRUD audit and the wider web API surface.

## Discovery limits

The initial discovery recorded 64 outstanding leads: specific operations, partially covered API prefixes, and alternate proxy routes. Each now has an implementation or a source audit mapping its observed calls to existing methods. This closes the recorded lead list within the inspected frontend sources; it does not establish coverage of every tenant feature or backend endpoint. The [discovery inventory](web-api-discovery.json) retains per-lead evidence and validation limits.

Live checks cover only the operations and fixtures identified in the validation report. Remaining gaps include compatible legacy portal sessions, populated application-insight fixtures, branding and integration credential writes, Collector configuration changes, and successful telemetry submissions. Source-backed unit fixtures are distinguished from live responses. AmplifyAI and Hypervisor public contracts remain unverified; product names or credential permissions alone are insufficient to invent those endpoints.

Undocumented APIs can change without notice. See [validation evidence](lab-validation.md) for tested behavior and remaining lab work.

## Layout and configuration

`nexthink/nexthink.go` wires `PublicAPIClient` and `WebAPIClient` into one `Client`. Each resource under `nexthink/web_api/` has `constants.go`, `crud.go`, `models.go`, `validators.go`, resource-local tests, and JSON fixtures in `mocks/`. Both families depend on the same `interfaces.HTTPClient` and `client.Transport` implementation.

For environment configuration, set `NEXTHINK_API` to `public` (default), `web`, or `both`. Public authentication reads `NEXTHINK_CLIENT_ID` and `NEXTHINK_CLIENT_SECRET`. Web authentication uses `NEXTHINK_WEB_AUTH=token` with `NEXTHINK_ACCESS_TOKEN`, or `NEXTHINK_WEB_AUTH=chrome` for an existing Chrome session. Both require `NEXTHINK_INSTANCE` and `NEXTHINK_REGION`.

A nil API-family field means that family was not configured. Provider tokens are validated at request time, including expiry and cancellation. Static tokens with a supplied expiry are also checked during construction. A public token is never silently substituted for a web session.

NQL editor calls accept `nql_editor.Document`, `PositionRequest`, and `ValidationRequest`. Completion items retain opaque fields for resolve round trips. Invalid NQL returns diagnostics under HTTP 200; an empty hover or resolve object is also a legitimate successful response.

Saved NQL queries also expose `List`, completing the resource-level LCRUD set. Workflow and remote-action write examples read explicit JSON from `NEXTHINK_REQUEST_FILE`. Delete examples expect `{ "uuid": "..." }` for workflows or `{ "nqlId": "#..." }` for remote actions. Create/update request models preserve the UI field names; custom IDs begin with `#`.

Content summaries preserve nullable ownership/audit fields: `ContentOwner`, `CreatedBy`, and `UpdatedBy` are now `*string`. The UI returns null for system-owned remote actions; converting those values to empty strings lost information.

## Additional management resources and examples

`WebAPI.Applications`, `WritingAssistant`, `SoftwareMetering`, `CustomFields`, `RuleBasedCustomFields`, `Monitors`, and `Campaigns` expose List/Create/Get/Update/Delete alongside the existing Workflows, RemoteActions and NQLQueries resources. Each uses the shared transport and browser token provider. See the [runnable example index](../examples/nexthink/web_api/README.md) for every method and its inputs.

Applications use REST and revision-aware updates/deletion. Writing Assistant and Custom Fields updates also require a revision. Software Metering mutations return booleans; retrieve the created UUID with List. Custom Fields Create returns the definition without its document UID; use List to find it. Custom Fields delete uses `Device`/`User`/`Binary`/`Package`, while create/update use data-model URIs such as `device/device`. Rule-based fields use `WebAPI.RuleBasedCustomFields`; its revision-bearing Delete is a POST with a plain-text acknowledgment.

Monitors expose both a content/document ID and a distinct monitor UUID; Update needs both. Successful update/delete acknowledgments were null. Campaigns create drafts; none of the management lifecycle tests publish or deliver them. Populated question choices were checked against curl, and nullable campaign fields remain nullable.

`WebAPI.Assets`, `Checklists`, `Dashboards`, `Ratings`, and `Investigations` add further management lifecycles. Assets use signed URLs for reads and return empty update/delete responses. Checklists and Ratings require revisions and JSON DELETE bodies. Dashboards preserve polymorphic widget/filter/layout values and use revision plus product/type context. Investigations use `uid`, return nested `nql.query`, and support Export/Import; their tested server ignores description writes. The original 27 examples for these services passed curl-backed lab validation with fixture cleanup. The [discovery inventory](web-api-discovery.json) records the later auxiliary operations and their separate validation evidence.

`DeviceConfiguration.SetProfiles` accepts `*SaveProfilesRequest` and returns `*ProfilesResponse`. Supply named setting changes in `settings`; it is not a collection replacement. `ProductShell.ValidateClaims` accepts `*ClaimsRequest` and returns `*ClaimsResponse`, with the boolean at `Result.Result`. These replace the previous raw JSON signatures.

`WebAPI.Connectors` provides universal connector LCRUD, `ListTemplates`, `GetTemplate` and `ListManualCustomFields`. Supply a new UUID in `content_id` and use Quartz scheduling syntax. In the tested lab, Create forced `enabled:true` even when false was requested; Update honored false. Treat creation as potentially scheduling an integration. List includes legacy entries, whose IDs do not belong to the v1 management routes. Templates have no observed independent write operations.

`WebAPI.ConnectorCredentials` provides LCRUD plus `ListIDs`. Create and Update both invoke the UI's POST upsert and do not enforce absence/existence. Delete follows the UI's clear-and-disable POST, leaving the ID allocated. Secret values are optional write-only entries; omit `secret` to retain existing secrets during update. Get/List do not retrieve secrets. The lab lifecycle used no-auth fixtures; secret submission has synthetic wire tests only. See [connector examples](../examples/nexthink/web_api/connectors/README.md) and [credential examples](../examples/nexthink/web_api/connector_credentials/README.md).

## Browser integration and content operations

`WebAPI.Connectors`: Asynchronous tests return an execution ID; COMPLETED can contain per-partition failures. Lab used .invalid destination; successful external records are synthetic unit fixtures. See [connectors examples](../examples/nexthink/web_api/connectors/README.md).

`WebAPI.Workflows`: Workflow-specific reference views. Credential writes use WebAPI.ConnectorCredentials; connector definitions have no observed independent LCRUD. See [workflows examples](../examples/nexthink/web_api/workflows/README.md).

`WebAPI.KnowledgeBases`: Single upload sends base64 of the whole CSV as application/octet-stream. Multipart slices that encoded string. Create returns 202; contents/list are eventually consistent. Download URL is base64 and signed; download without the bearer token. No independent update operation observed. See [knowledge_bases examples](../examples/nexthink/web_api/knowledge_bases/README.md).

`WebAPI.LegacyConnectors`: Legacy configuration POST upsert; connector type is the URL ID. Save returns plain text. List includes nullable names absent from Get. Disabled Azure AD configuration tested with synthetic secret; other legacy subtypes not validated. See [legacy_connectors examples](../examples/nexthink/web_api/legacy_connectors/README.md).

`WebAPI.Webhooks`: Caller-generated UUID; Create/Update share POST upsert and plain-text acknowledgment. Saved payload is base64, test payload is plain text. Test sends immediately; .invalid destination returned HTTP 503. See [webhooks examples](../examples/nexthink/web_api/webhooks/README.md).

`WebAPI.DataExporters`: Create/Update share POST upsert with plain-text acknowledgment. Write enums are numeric, read enums are names. List uses $deleted=false. Placeholders path argument is base64 NQL. Send-test is asynchronous; status may initially return 404. Lab test succeeded with zero matching records. See [data_exporters examples](../examples/nexthink/web_api/data_exporters/README.md).

`WebAPI.Dashboards`: Revision and product/type context required. Omit layout when deleting the last widget. Resolved UI documents omit @api/@client. Import accepts the dashboardExport document. Polymorphic config fields retained. See [dashboards examples](../examples/nexthink/web_api/dashboards/README.md).

`WebAPI.Checklists`: Export is a versioned JSON document; Import posts it and returns an empty acknowledgment. List locates the imported ID. See [checklists examples](../examples/nexthink/web_api/checklists/README.md).

`WebAPI.Monitors`: Export content is base64; Import expects decoded JSON text. ExportLibrary returns content and metadata files. Disposable no-notification monitor round trip tested. See [monitors examples](../examples/nexthink/web_api/monitors/README.md).

The transport accepts non-JSON acknowledgments when no decoded result is requested. Typed JSON calls still reject successful non-JSON responses. API failures retain their HTTP response body and metadata.
