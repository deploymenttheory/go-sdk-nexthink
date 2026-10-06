# Lab validation — 5 October 2026

This revision was checked against one Nexthink EU lab. It is an alpha compatibility revision, not a claim of complete API coverage or guaranteed correctness. Public contracts were compared with Nexthink's documentation and curl responses before SDK replay. Undocumented routes were taken from first-party browser requests or frontend code; guessed gateway responses were not treated as discovery evidence.

## Public API evidence

| Area | Live result | Remaining limitation |
| --- | --- | --- |
| OAuth client credentials | curl and SDK obtained bearer tokens | One authentication request timed out during the final example run; immediate curl and isolated SDK retry succeeded |
| Remote actions | List: 200, nine records; detail: 200 | All existing actions have API triggering disabled; no execution performed |
| Workflows | List and details: 200; filtered list: 200 with no matches | Existing workflow has API triggering disabled; no execution performed |
| NQL execute | V1 and V2: 200, three device records; parameterized no-match query: 200, zero rows | Bounded lab queries only |
| NQL export | Start: 200 with export ID; status: 200 COMPLETED; CSV download: 200 with header and three rows | Compressed export formats have contract coverage, not live download coverage |
| NQL errors | Missing query: curl and SDK 404, numeric code 100 retained | No exhaustive server-error matrix |
| Data Management | curl and SDK: 202 ACCEPTED, scheduledCount 0, per-device INVALID for an impossible UID | No real device deletion scheduled or completed |
| Enrichment | Published schema checked; request/response contract tests | No positive live enrichment or partial-success fixture |
| Campaigns | Existing service and example compile and pass local tests | No live campaign triggered |
| Spark | Published handoff contract implemented; HTTP 204 contract tested | No Teams handoff sent; dedicated recipient required |

The inventory already contained three Windows devices. Validation did not delete or modify them. Saved queries `#sdk_lab_devices` and `#sdk_lab_parameterized` were created as bounded fixtures. A separate temporary saved-query fixture was created, updated, read back, and deleted through the WebAPI resource.

## Examples and NQL templates

After the unified-client migration, all 19 public read examples passed in one run of the committed [example harness](../scripts/validate-examples.sh), including both `_build_client` examples that Go's `./...` pattern skips. All nine web examples also passed. The earlier OAuth timeout noted above did not recur in this final run. Examples now exit unsuccessfully when execution or export fails. Write examples require explicit JSON fixture input through `NEXTHINK_REQUEST_FILE`.

All 18 NQL templates were submitted to the lab's NQL editor validator. Initial validation exposed five syntax failures; review also found a memory template calculating disk usage. The corrected templates produced empty diagnostics. This validates the generated queries against the lab parser; it does not establish every query's meaning against a representative telemetry dataset.

Query-builder operator order is now preserved, including filters within event relations before aggregation. String helpers escape literals. Templates use correct enum expressions, filter event fields before computation, and calculate memory usage from used and installed memory.

## Undocumented APIs

The [catalog](../nexthink/web_api/operations.json) contains 55 operation entries with endpoint, method, discovery evidence, and observed authentication results. The main client wires both API families. Web resource methods preserve raw JSON where schemas are not established.

- 26 initial read/probe operations returned HTTP 200 through the SDK, covering Collector management, product shell, licensing, content administration, saved-query reads, and eleven GraphQL endpoints.
- Saved-query create/update/read/delete returned 201/200/200/204. Public execution with named parameters also succeeded.
- Follow-up curl and SDK checks returned 200 for device profiles, NQL validation, and NQL highlighting. The alternate device-settings route returned 403 under the same browser identity.
- GraphQL checks used `__typename`; they establish reachability, not support for all queries or mutations. Management introspection was disabled or unavailable. The SDK reports GraphQL errors even when HTTP status is 200 and preserves partial data.
- Browser and OAuth client identities have different access. For example, shell user/configuration reads worked with the browser identity and returned 401 with client credentials. A Collector configuration request returned 503 with client credentials, which is inconclusive.
- Collector/device configuration writes, shell telemetry, claims validation, and dynamic menus remain bundle-evidenced or unreplayed. Their catalog entries explicitly say so. No shared tenant settings were changed to test these routes.

Discovery inspected 64 first-party entry bundles and 77 manifest assets. Additional dynamic routes and payloads remain unresolved. AmplifyAI and Hypervisor public API contracts remain unverified. No API for retrieving a valid Collector Customer Key was established.

After resource restructuring, 33 read/editor/probe checks returned the expected statuses: 32 HTTP 200 responses and the known device-settings HTTP 403. NQL completion, hover and resolve were first replayed with curl, then through resource methods. A fresh temporary saved query again passed create/update/read/delete (201/200/200/204), with read-back assertions and cleanup.

The final `nexthink.NewClient` configuration was also exercised with both credential sets: `PublicAPI.NQL` returned the enrolled macOS VM, and `WebAPI.ProductShell` returned the signed-in user through the same top-level client. The resource replay and temporary-query CRUD were repeated after the single-entry-point migration and typed-model changes.

The eight web resource packages in the initial validation included request/response contract tests, validation tests and JSON fixtures. They cover all 42 catalog operations with synthetic success and error payloads; mocked success does not promote a bundle-only operation to live-verified status. Positive fixtures reflect observed envelopes and typed fields. Data Management additionally tests a synthetic mixed batch containing `SCHEDULED`, `INVALID` and `FAILED` outcomes; Spark tests its successful empty HTTP 204 response.

## Transport and regression checks

The revision fixes token refresh re-entering its own authentication middleware, coordinates concurrent refreshes, observes cancellation, retains numeric API error codes, and prevents authenticated requests from leaving the configured origin. Token requests use a separate HTTP client while retaining the configured proxy/TLS transport. Debug logging excludes request and response headers/bodies.

Validation completed:

- `go test -race ./... -timeout 180s`
- `go vet ./...`
- Configured `golangci-lint run --fix=false --timeout=5m`: zero issues; existing unmatched-exclusion configuration warnings remain
- Error parser fuzzing: 60 seconds, 388,718 executions
- Ragged V1 result fuzzing: 60 seconds, 414,637 executions
- Credential-value scan of tracked and proposed repository files: zero matches

Regression tests cover concurrent token refresh, cancellation, origin restrictions, numeric/wrapped errors, public service wire contracts, CSV conversion, malformed result rows, query ordering, web API path parameters, and GraphQL partial-data errors. Synthetic fixtures are committed; raw tenant responses, credentials, signed URLs, browser tokens, and Collector keys remain outside the repository.

Hosted CI exposed existing Super-Linter configuration failures with multiple packages and a module root containing no Go files. That wrapper is replaced with explicit formatting and vet checks; the dedicated golangci-lint workflow checks the module and no longer forces a successful exit when findings occur. An earlier dependency-change review was blocked because GitHub reported that Dependency graph was not enabled. Dependency Review passes on the unified-client revision, which introduces no dependency changes.

## Device lab and unfinished validation

The requested 32 GB macOS restore failed at 84%. The approved 64 GB sparse-disk retry also failed during restore. A dedicated copy-on-write clone of the existing stopped macOS 27 golden image was then created with a regenerated MAC address and serial, a 64 GB virtual disk, and 8 GiB RAM. The source image was retained.

Collector 26.8.2.22 installed and its services started. Remote actions were configured to accept trusted or Nexthink-signed scripts. Installing the supplied Customer Key **byte-for-byte, as one line with no trailing newline**, produced a persistent **Connected** status and no decode errors from the current Collector process. Added line breaks, a trailing newline, and CRLF variants produced **“Failed to decode the Customer Key”** despite carrying the same Base64 payload. Preserve the complete key file exactly; PEM-style rewrapping is not valid for this tested Collector. The earlier attribution to a defective supplied key was incorrect.

The dedicated VM subsequently appeared in tenant inventory, increasing the device count from three to four. curl and the SDK both returned `nexthink-sdk-macos` with device and Collector identifiers; the SDK's parameterized lookup returned HTTP 200 and exactly one matching row. Basic enrollment is verified.

Full Disk Access for `nxtsvc` in this VM was explicitly approved. Applying it remains unverified. The VM was restarted headed in a native window and guest exec confirmed the `labrunner` console session plus both Collector processes. Its native display remains black, including after retrying the original 1824×1368 geometry with display refitting disabled. The VM remains running headed; basic enrollment and API reads pass, but Full Disk Access and complete telemetry coverage are not yet verified.

Remaining device-dependent work: verify Full Disk Access and complete telemetry; execute a dedicated signed diagnostic remote action; exercise dedicated workflow and campaign fixtures; test enrichment with read-back; and perform deletion only on an explicitly designated disposable device. Spark requires a dedicated Teams recipient. These are outstanding tests, not passing results.

## References

- [Nexthink API documentation](https://docs.nexthink.com/api)
- [Data Management device deletions](https://docs.nexthink.com/api/data-management/schedule-device-deletions)
- [Spark handoff](https://docs.nexthink.com/api/spark/handoff-api)
- [NQL investigation examples](https://docs.nexthink.com/platform/user-guide/investigations/investigations-nql-examples)
- [Migration notes](migration-lab-validation.md)

## Web management expansion after PR #48

Nine typed management operations were replayed with curl and then through the single SDK entry point: workflow list/get/export; remote-action get/view/quota; Bash input/output inspection; PowerShell input/output inspection and signature inspection. All returned HTTP 200 with successful GraphQL data. Every typed SDK data envelope matched its curl response. The two new resource examples also completed successfully. Pre-existing tenant workflows and remote actions were read only. Dedicated temporary fixtures were subsequently created for lifecycle tests.

Script inspection established that macOS input is a base64-encoded tar.gz archive, while PowerShell input is base64-encoded UTF-8 source including its BOM. Incorrect formats returned GraphQL errors under HTTP 200. The SDK accepts the underlying bytes and encodes them once. The synthetic PowerShell sample returned `NOT_SIGNED`. No script was executed by these inspection calls.

The new resources follow the existing package layout, add nine successful JSON response fixtures plus request/error fixtures, and share a typed GraphQL data decoder that preserves partial data and errors. Full `go test -race ./...`, `go vet ./...`, and configured lint checks passed locally. The [coverage inventory](web-api-coverage.md) lists additional UI API areas and distinguishes recovered mutation documents from implemented, live-validated methods.

Following the LCRUD coverage audit, both Workflows and Remote Actions passed create/list/read/update/read/delete through the SDK. Separate curl create/update/read/delete runs also passed. Workflow fixtures remained inactive with every trigger disabled; remote-action fixtures had every trigger disabled and were never executed. All fixture deletions returned true. Saved NQL query List was curl- and SDK-validated, completing its existing CRUD surface. Six mutation success fixtures and mutation partial-error fixtures were added alongside listing fixtures.

All ten new Workflows/Remote Actions LCRUD examples ran successfully against dedicated fixtures (including a second Get after Update), followed by the saved-query List example. Both example-created objects were deleted. Live listing also exposed nullable content ownership/audit fields; those now retain null rather than decoding to empty strings.

## Full web service and example pass

Application management, Writing Assistant, Software Metering, manual Custom Fields, rule-based Custom Fields, custom metric Monitors, and Campaigns each passed independent curl and SDK-example LCRUD lifecycles. All 35 new management examples ran successfully, with a second Get after Update. All created objects were deleted. Applications used an unused `.invalid` URL with enhanced collection disabled; monitors had no notification recipients, matched no devices and had an impossible trigger threshold; campaigns remained drafts. No device action, campaign delivery or message was triggered.

The application response and a campaign response containing a single-answer question, two answer choices and a final message matched curl exactly. Additional live calls validated the typed claims helper, existing-profile setting save, the current collector reads, shell configuration/modules/menu/flags, dynamic `bus-menu`, NQL completion/hover/resolve/highlighting, the generic GraphQL example, and remote-action inspection examples. The profile save reread the tenant state before writing all seven unchanged setting values, then verified both values and completed onboarding.

The alternate device-settings GET again returned 403. It is recorded as a validation gap, not a successful response. Collector update configuration and telemetry write examples compile but have not been live replayed. Public helper examples were added to close the source inventory; their addition alone does not establish fresh live execution evidence.

The UI audit additionally recovered Monitor auxiliary mutations, Campaign status/branding operations, and DEX configuration mutations. These auxiliary operations are follow-up contracts, not claimed implemented coverage. Rule-based Custom Fields were implemented after their separate REST lifecycle passed. The coverage inventory records the remaining broader endpoint discovery work.

Rule-based Custom Fields also passed curl and all five SDK examples, with successful update read-back and deletion. The creation and update payloads preserve rule enum identifiers; response-only `deleted` flags are not sent back. List filters the shared Custom Fields listing to the RULE_BASED subtype.

Computed Custom Fields also passed the existing five examples with update read-back and deletion. The test used one include clause, one compute clause and a final list, following [Nexthink's query constraints](https://docs.nexthink.com/platform/user-guide/administration/content-management/custom-fields-management). The calculated query filtered events to a synthetic non-matching device name.

The individual saved-query Create/List/Get/Update/Delete examples were rerun on a disposable query, with curl read-back after Create and Update and a listing check after Delete. The first read used the client-supplied creation ID and returned 404: the server had generated a different ID. Using the returned ID completed the lifecycle and cleanup. The example index now calls out this behavior. Remote Actions GetForView and Workflows Export also ran successfully; Export now emits the actual opaque content on stdout for saving to a file.

Continuing discovery recovered dashboard management, widget/filter/tab mutations and ratings CRUD from 22 additional lazy-loaded UI files. The [discovery inventory](web-api-discovery.json) records 33 follow-up HTTP contracts and 205 GraphQL document variants, with nine successful read replays across assets, knowledge sources, checklist metadata, ratings and dashboards. The dashboard administration content key is `dashboards`; `custom_dashboards` is a product-area value and returned 404 when incorrectly used as a list key. Follow-up writes in that inventory remain unimplemented and unvalidated.

## Content management continuation after PR #50

Assets, Checklists, Dashboards, Ratings and Investigations passed separate curl and SDK/example lifecycles. All 27 added examples ran successfully; update read-back and deletion were checked. The Investigation export was imported under a distinct name, read back and deleted. Shared content listings confirmed cleanup. No connector configuration, device action or notification was triggered.

Populated checklist field metadata required `type` as well as `id`; the server rejected its omission. Ratings conditions required complete NQL queries. Investigation descriptions were ignored even with HTTP 200, so update validation checked name and NQL instead. Asset updates returned 204, changed revision and bytes, and retained the original filename; downloaded replacement bytes matched exactly. Investigation read representations and the Checklists, Ratings and Dashboards representations matched curl.

The newly exercised DELETE bodies exposed and fixed a transport bug: Resty required `SetMethodDeleteAllowPayload(true)`. Previously `DeleteWithBody` discarded its body, and its old success test did not inspect it. The regression now verifies the server receives the expected JSON.

Universal connector listing, its 18 templates and the nine workflow connector types returned HTTP 200. Connector CRUD/test contracts were captured from the first-party repository code, but no existing connector was modified or tested against a third-party system. The discovery inventory marks these as pending implementation.

Local validation for this continuation: `go test -race ./...`, `go vet ./...` and `golangci-lint run --fix=false` all passed. The changed-file credential scan found no lab secrets, browser tokens or signed URLs. Relative documentation links and JSON fixtures were checked.

## Connector and credential validation — 2026-10-06

From main after PR #51, `codex/nexthink-web-connectors-expansion` adds two shared-client resources. Browser curl first validated both lifecycles, then all 14 new Go examples passed. Create/list/get/update/get/delete behavior was checked with disposable fixtures. Typed template collections (18, unordered), individual templates, enabled credentials, credential details, connector details and shared connector lists matched curl. The custom-field lookup for `device/mobile_device` returned 200 with an empty array.

Create required a caller-generated UUID and valid Quartz scheduling. The server returned `enabled:true` despite an explicit false; that fixture was immediately removed. Subsequent tests used a reserved `.invalid` destination, no authentication secrets, and a 2099 schedule. Connector Update persisted false, custom headers and changed description. Delete returned 204 and removed the object from the shared listing. No test-execution endpoint was invoked.

Credential Create returned 201; Update returned 200. Both use the same POST upsert. UI-style Delete returned 201, cleared connection details/secrets and persisted `enabled:false`; the ID remains allocated and Get still succeeds. Enabled listings excluded the cleared credentials. All fixture connectors were deleted and fixture credentials cleared/disabled; existing objects were unchanged.

Nullable template icon/documentation and credential-list names initially failed typed-versus-curl comparison. Models and synthetic fixtures now preserve those nulls. No tenant payloads or credentials were copied into test fixtures. Populated manual-field results, nonempty connector mapping execution, secret writes and third-party test execution remain outside live validation.

A separate workflow connector-credential route was recovered from the first-party workflow UI and returned an empty array (200). It is recorded as pending typed implementation rather than inferred to share the generic credential schema.

## PR #52 continuation — 6 October 2026

All 56 additional examples were exercised with the lab browser session after curl established the contracts. Fifty-five completed successfully. `Webhooks.Test` correctly returned an SDK error for the HTTP 503 produced by the deliberately unreachable `.invalid` destination; curl confirmed the same failure. This validates error propagation, not successful webhook delivery.

| Area | Live evidence | Limits |
| --- | --- | --- |
| Connector tests | Async start and result polling; completed response included partition-level connection failure | No successful third-party records; populated successful records have synthetic unit fixtures |
| Workflow references | Connector definitions and populated credential view matched curl | Built-in definitions are reference data; credentials use the shared credential service for writes |
| Legacy connectors | Disabled Azure AD configuration LCRUD, list/detail comparisons, synthetic secret save | Other legacy subtypes and real credential authentication not tested |
| Webhooks | Disabled configuration LCRUD, availability, base64 saved payload and plain-text test payload | Test destination intentionally unreachable; no real recipient contacted |
| Data Exporter | Disabled generic HTTP configuration LCRUD, customer info, status list, placeholders, async test | Test succeeded with zero matching records; initial status 404 observed; non-generic protocols untested |
| Knowledge bases | Single-file and multipart upload, register, populated contents/list, signed URL, delete | One-part multipart lifecycle tested live; split encoded chunks and query escaping unit-tested. No update operation observed |
| Dashboards | Widget/filter/tab create/update/delete; tab ordering, layout, duplicate, export/import | Heading widget and terms filter tested; other widget/filter variants preserve their configuration JSON but are not individually replayed |
| Checklists | Versioned export/import round trip and grouped fields | Import returns no ID; located imported fixture through List |
| Monitors | Export, library export and import round trip | Custom metric monitor only, with no notifications and an unreachable threshold |

Single and multipart knowledge downloads matched the original CSV bytes. Signed downloads used curl without a Nexthink bearer token. Processing is asynchronous: one immediate delete returned a backend persistence error; retry succeeded. Collection comparisons waited for ingestion to settle. Every disposable configuration/content object was removed; credential clear-and-disable leaves its allocated ID, as designed by the UI. Legacy secret tests used a synthetic, nonfunctional value.

Wire tests caught the shared transport rejecting successful plain-text acknowledgments. Calls that request raw response metadata now retain non-JSON bodies; typed JSON calls continue to reject non-JSON success responses. Further fixes preserve legacy list-only nullable descriptions and absent workflow action outputs. Unit fixtures cover populated success data, HTTP failures, malformed responses, transport failures, validation, and GraphQL partial data. No tokens, raw lab captures, real credentials or signed URLs are committed.

Final checks: `go test -race ./...`, `go vet ./...`, and `golangci-lint run --fix=false --timeout 10m` passed (zero lint issues). The source-level example guard, modified Markdown links, JSON fixture parsing, whitespace checks and tenant/secret artifact scan also passed.

## Browser operation completion after PR #52 — 6 October 2026

This pass reconciles the earlier pending-operation statements above. Those sections describe their historical validation stage; the current implementation inventory is [web-api-discovery.json](web-api-discovery.json), summarized in [coverage](web-api-coverage.md).

| Batch | Added methods | Live outcome |
| --- | ---: | --- |
| Application experience and metering | 36 | 34 read examples matched curl; usage distribution returned the same backend GraphQL error; automatic metering returned CREATED with curl and ALREADY_CONFIGURED on SDK repetition |
| Alert Hub, Diagnostics, Benchmark, DEX scores, CCI and network insights | 52 | 42 successful examples; 10 missing-fixture/backend errors matched curl, including partial data |
| Workflow, remote-action, campaign, monitor and DEX configuration helpers | 31 | 22 successful curl/SDK comparisons; two expected legacy/business errors; seven persistent configuration/status writes not applied |
| Data exploration | 12 | All examples succeeded and matched curl; populated inspection, breakdown metadata and optional NQL selections also checked |
| Dashboard metadata, ratings, checklist library and investigations | 11 | All examples succeeded and matched curl; query results populated; disposable rating exported then deleted |
| Library | 20 | Seven successful read examples; 13 installation/custom-pack operations have source and unit validation without live mutation or a suitable custom-pack fixture |

Benchmark search results were compared as unordered records because server ordering changes. The library catalog changed order and some content versions between curl and SDK calls: six other library reads matched curl exactly, while all catalog records were additionally compared against the SDK's own raw HTTP response to verify decoding without attributing server changes to the model. Query execution comparisons exclude only server-measured duration and, for investigation execution, the per-request server timestamp. Deterministic response fields matched.

Management validation used disposable inactive workflows, disabled remote actions, draft campaigns and inactive monitors. All task-created definitions and metering/application fixtures were removed. Library definitions were read without installation. Existing branding, DEX aggregate settings and active campaign/monitor configuration were not changed. Workflow import required an explicit empty `versions` array when importing an otherwise versionless exported definition. Campaign legacy reads require a v6-backed campaign; a newly created draft does not supply that record.

Curl showed the visual-editor gateway returning base64-encoded gzip text while claiming `Content-Encoding: gzip` for `Accept: application/json`. `Accept: */*` returned correctly framed gzip and allowed the SDK response to match curl. The GraphQL client now uses that negotiation, with a transport-level regression. Dashboard collection/field metadata requires `x-nxt-waas-iso-date-time`, timezone and UTC-offset headers; these are now supplied. Library models preserve nullable identifiers/titles and absent checklist labels found during full-response comparisons.

Ten obsolete GraphQL documents failed active-schema validation and are documented as unavailable, rather than exposed as usable methods. The unused library content-type configuration helper returned 404. These exclusions are separate from implemented operations whose live responses lack telemetry or return business errors.

Validation includes full `go test -race ./...`, `go vet ./...` and `golangci-lint run --fix=false --timeout 10m`, alongside targeted curl/example comparisons. Every new method has a runnable example, synthetic positive/request JSON fixtures and error-path tests. No claim is made that live configuration writes, library installation, every telemetry subtype or every uninspected UI API has been validated.

The catalog replay also exposed response-body loss after Resty streamed typed JSON decoding. Per-request buffering now retains `interfaces.Response.Body`; byte access preserves leading/trailing whitespace in raw downloads. Real HTTP server regressions cover small and 2 MiB JSON, gzip, malformed JSON, binary bytes and response-size limit enforcement. The fix applies to both API families through their shared transport.

After the buffering fix, live `Library.ListContents` retained 2,164,916 response bytes for 810 records; its typed result matched that same HTTP body exactly. The raw response is retained alongside parsed models, with configured response-size limits unchanged.

## Browser identity, support and helper APIs after PR #53 — 6 October 2026

This continuation adds 174 methods with runnable examples, request/positive-response JSON fixtures and error-path tests. Curl established available lab contracts before SDK replay. The table separates successful calls from source-derived methods which could not be fully exercised in this tenant.

| Batch | Added methods | Live outcome and limits |
| --- | ---: | --- |
| Access management, collaboration comments, recommendations | 58 | 21 safe IAM/current and legacy credential reads exactly matched curl. Feature flags returned 404, comment reads 403 and recommendations 401. Account, identity, role, credential and notification mutations were not applied. |
| Support, checklist/timeline/insight views, VDI and execution services | 70 | 39 successful SDK reads, 36 matching separate curl responses exactly. Three generated insight narratives vary across requests. No workflow or action execution was triggered. Some device drilldowns fail or lack suitable telemetry; VDI/checklist execution fixtures are unavailable. |
| Autopilot, global search, NLP assistant | 23 | Search and informational assistant chat succeeded. Search category ordering/ranking and assistant text vary; a controlled unmatched search matched curl after sorting category events. Autopilot returned 401; agent-action GraphQL returned the expected missing-permission error. No Autopilot settings, approvals or tickets were changed. |
| Export jobs, visual editor, sharing, custom-field values, templates and field import/export | 23 | 22 successful SDK methods. Legacy sharing owner lookup returned 400 for the cloud user identity. Manual/rule-based export and manual import matched the observed contracts; computed export was not replayed. |

The last batch used a zero-record export job, disposable application/custom-field definitions and the dedicated SDK test VM. Modern and legacy sharing calls submitted empty profile lists on task-created content: they validate transport/acknowledgment, not grant or revoke behavior. Modern sharing applies deltas to named roles; an omitted role retains its permissions. Template catalog comparisons included populated selector-system metadata, requiring preservation of its JSON object shape. Export status preserves the absence of a download URL while a job is pending.

Custom-field value writes and clearing returned 202; acceptance is not proof of eventual device-value persistence. CSV dry-run and import used a nonexistent device name and returned 202 with an empty array. No existing device values were imported. All four task-created custom-field definitions and the sharing application were deleted; a final field listing confirmed their absence. Signed export URLs and raw tenant responses remain outside the repository.

The Software Metering Create example had no local request sample, and its guide did not explain the boolean result or subsequent UUID lookup. Create/Update/Delete now have local JSON samples; all five lifecycle examples have operation guides. The existing Create wire request passed a fresh disposable curl and SDK create/read/delete lifecycle, and the metering configuration and application were removed. The UUID in the sample represents an application, not a device.

Global Search and NLP Assistant buffer their event responses using the shared transport. They do not expose realtime callbacks or SSE automatic reconnection. Legacy PortalServlet token and dashboard-search calls need separate cookie/x-auth-token authentication; the bearer-only lab request returned 403, and those contracts are documented without claiming working SDK support.

Final integration checks passed: `go test -race ./...`, `go vet ./...`, and `golangci-lint run --fix=false --timeout 10m` (zero issues). The example coverage guard, 474 JSON documents, eight intentionally malformed response fixtures, 147 Markdown files, staged whitespace and tenant/secret artifact scan were checked. Review also added regressions for VDI plain-text health responses, legacy sharing error codes and explicit empty-action arrays used to revoke individual grants.

## Remaining discovery leads after PR #54 — 6 October 2026

The source pass resolves the 64 recorded outstanding leads into implemented contracts, existing-method mappings or source initialization/static/non-HTTP behavior. It adds 40 API operation methods and one immutable dashboard proxy helper, with 11 new resource packages. The source audit records the inspected bundle versions; unknown or future backend functionality remains outside that bounded claim.

| Area | Live evidence | Limits |
| --- | --- | --- |
| Integration credentials and communication settings | Three safe SDK reads matched curl; two secret-presence reads returned expected 404s | No real credentials, integration changes or outbound checks were sent |
| QueryBuilder, CCI benchmarks, dashboard menu/proxy | Six curl/SDK JSON comparisons passed, including all query-builder methods and a GraphQL request through the proxy | Application insights had no configured application fixture and returned 404; null drilldown conditions returned 500, while the source-derived selected-row condition succeeded |
| Mobile enrollment tokens | Separate curl and SDK create/list/get/update/delete lifecycles; full response comparisons and update read-back | Tokens were unused and short-lived; no mobile device enrolled. List omits JWT values and the model preserves that absence |
| Snapshots/custom trends | Curl and SDK LCRUD plus export/import, definition/list comparisons and update read-back | Zero-match hardware-manufacturer filter; no identifying fields. All temporary definitions were deleted |
| Collector device query, appearance and events | Populated Collector query matched curl; menu-logo bytes matched exactly; UI polling returned an empty page | Existing branding was not changed; populated events/pagination and image updates have source/unit validation |
| Observability proxy | Empty curl submission returned the expected 403 | No valid telemetry batch was submitted; successful byte-preserving transport is unit-tested |
| EUF metadata and legacy portal adapters | Feature metadata returned 200 and matched the SDK | Legacy search/branding uses separate explicit session credentials; no compatible session was available in the modern tenant |

Snapshot validation required a `list` statement, rejected `summarize`, and rejected identifying fields even when they appeared only in a filter. The working fixture filtered on a nonexistent hardware manufacturer and projected only that field. Its NQL ID required the leading `#`. Import creates a new definition; export projects the existing Get response into the portable JSON representation. Every temporary mobile token and snapshot created in this pass was deleted.

Legacy connector input validation was corrected to match actual Teams/Zoom saves: name, enabled and timezone may be omitted, and mapping may be an explicit empty array. `ConfigurationInput.Enabled` is now `*bool`, retaining the difference between omitted and explicit false. Tests cover both forms.

Legacy portal requests are restricted to the three observed POST routes. They suppress bearer authentication, never consult the bearer provider, and do not persist cookies to subsequent calls. Real-server regressions exercise cookie/header isolation, subsequent bearer requests and query-value redaction in errors and logs. Legacy endpoints are implemented from shipped source and synthetic fixtures; this does not establish successful portal authentication in the cloud lab.

The expanded example guard covers exported service methods in extension files as well as `crud.go`. Six existing public NQL export-helper examples were added and compiled without live exports. The README and quick-start guide use the SDK's actual environment variables and signatures; all three complete documentation programs compiled.

Final checks passed: full `go test -race ./...`, `go vet ./...`, and `golangci-lint run --fix=false --timeout 10m` (zero issues). The final inventory/source-audit guard and scoped-authentication suites also passed with race detection. Across this pass, 24 newly added operation methods and the dashboard proxy path completed successful curl/SDK validation. Final listings confirmed fixture cleanup. Mobile token examples now redact JWTs on stdout; optional full output uses a newly created private file without overwriting existing files.

## Headless local-password authentication — 6 October 2026

The SDK authenticated with a password-only local lab account through its own headless Chromium process, with no existing Chrome session, visible window or exported browser state. The root `NewClient` username/password configuration returned HTTP 200 from the Applications listing. A separate non-root Linux ARM64 container, with credentials injected only at runtime, completed the same read successfully.

The opt-in live provider test obtained a server-issued access token with a five-minute lifetime and a refresh token. It compared the typed SDK Applications read with curl, waited until after the actual access-token expiry, and repeated the comparison successfully. The complete run took 307 seconds and recorded **one browser login and one OAuth refresh request**. No token expiry was fabricated, and no second browser login was needed. Authentication state stayed in memory; browser contexts were closed after acquisition.

The first live attempt revealed that Okta loads a discovery iframe from a different origin during the local login flow. The adapter now blocks that child-frame navigation without misclassifying it as corporate SSO; foreign top-level navigation still fails, and credential origins remain restricted. The container smoke test also found a root-installed Playwright Node executable that was not executable by the runtime user. The Dockerfile corrects that single executable's permissions during image construction.

Controlled browser fixtures passed under race detection on macOS: combined and staged forms, invalid credentials, MFA, locked accounts, password changes, foreign redirects/iframes, changed DOM, cancellation and concurrent login acquisition. Provider and transport regressions cover rotation, rejected refresh credentials, password-login fallback, timeout sharing, late 401s, cancellation after a successful rotation, redaction and cleanup. Invalid-password, account-lockout and revocation scenarios were tested with fixtures rather than changing the lab account. Linux/macOS/Windows fixture jobs are defined separately from the opt-in live test and use no tenant credentials.

Full SDK race tests, `go vet ./...`, final focused authentication race tests and lint passed. The container image was built and exercised as a non-root user. This establishes unattended login and renewal for this lab's local-account policy; it does not establish support for corporate SSO, MFA or every Nexthink tenant. Those challenges return explicit errors. See the [headless authentication example](../examples/nexthink/_build_client/headless_password/README.md) for deployment and opt-in live-test commands.

## Systematic acceptance after PR56 — 2026-10-06

At this historical stage, the run accounted for all 586 exported resource methods: 334 positive passes and 252 explicit blockers, with no unresolved SDK failures after the recorded corrections. The [acceptance report](acceptance/README.md) and [per-method matrix](acceptance/2026-10-06.json) now include later follow-ups. They separate populated live success from empty telemetry, permission/feature restrictions and unexercised mutations and record cleanup details, discovered defects and repeat commands.

## Curl-led follow-up after PR57 — 2026-10-06

Retested 53 methods with corrected source-derived requests and retained disposable fixtures. This historical stage recorded 358 positive passes and 228 categorized validation gaps. Twenty-four previously blocked methods passed; three further request/schema failures were resolved but remained telemetry gaps. Enrichment and custom-field CSV/value writes were checked by actual NQL value readback. All follow-up content fixtures were cleaned after dependent tests finished. The report corrects the earlier erroneous monitor-schema diagnosis: ListFilterFields already used the correct visual-editor gateway; the original curl harness did not.


## Initial product-permission discovery after PR59 — 6 October 2026

The LBG - Superuser Sandbox role was read through the IAM API. At this initial stage, AI Tools/governance, Workspace and the listed VDI view permissions were enabled; Amplify view and packages were enabled, but Amplify management was disabled. The continuation below records the subsequently approved role change. Feature availability remains separate from role permissions.

The initial discovery added 84 methods across AITools, Amplify, AmplifyAI, Workspace, WorkspaceAgents, WorkspaceTasks, WorkspaceAssignments and the existing ActionExecutions resource. Curl preflight followed by SDK replay established 45 positive methods; 39 lacked positive acceptance. The historical total at this stage was 403 passed and 267 outstanding across 670 exported methods, including helpers.

Disposable AI tools/adoption goals and Workspace conversations/attachments were cleaned up. Fresh list/readback checks verified cleanup and explicit empty-tag/false update behavior. Amplify usage ingestion wrote labeled acceptance events. No Amplify AI actions or external ticket resolution were executed; AI mutation fixtures remain source/unit validated only. The lab explicitly gates custom agents and tasks.

Ten product-menu virtualization dashboard definitions matched the existing SDK without new vendor-specific endpoints. Eleven aggregate table scenarios matched the browser NQL API, excluding per-request duration. This uncovered and fixed nullable collection-name loss for NQL API logs; a JSON regression fixture and corrected live replay cover the change.

All added methods had runnable examples and JSON-backed tests. Full race tests, vet, lint and inventory checks passed at this stage. Amplify management and collaboration breadth remained discovery work until the continuation below.

## Product-permission continuation — 6 October 2026

The current [acceptance report](acceptance/README.md) records **414 passed and 278 outstanding methods across 692 exported methods**, including helpers. Of these, 664 are web methods: 663 operations and the dashboard proxy helper. The combined PR adds 106 methods across ten new packages and existing resources. This continuation contributes 22 new methods; the existing dashboard menu method also gains the UI's product-area filter. The HTTP catalog contains 497 entries representing 494 method/path pairs, and the curated follow-up inventory contains 453 implemented contracts. These are method and observed-contract counts, not an exhaustive count of Nexthink's backend endpoints.

Amplify administration UI 1.34.6 established configuration Create/Update, closing the inspected administration-source gap. Following explicit user approval, Manage Amplify was enabled and independently read back; it remains enabled as requested. The initially absent singleton configuration was created once through the SDK and verified independently with curl GET, without issuing a duplicate curl POST. Curl and SDK updates exercised adding, editing, reordering and removing test applications. Cleanup removed every test application and verified `itsmConfigList: []` with `enableUsageDataReporting: false`. The empty configuration document remains because no document DELETE was observed in the UI; application removal uses full-list replacement. Both new methods passed.

Product-configuration UI 1.83.4 supplied 20 methods in `ProductConfiguration`, `DeviceClassification` and `UserClassification`. Instance configuration, GeoIP and user-organization reads passed curl/SDK JSON comparison. Six classification reads returned matching `NO_RULESET_FOUND` responses and need configured CSV rule sets. Eleven writes remain untested because they change shared tenant configuration rather than isolated disposable objects. User-organization edits/removals replace the full list; only VPN egress has an observed DELETE. The inspected product-shell consumer reused existing services and added no further endpoint.

Collaboration discovery recovered three advertised dashboards and Call View UI 2.8.0. The dashboard definitions contain 13 tabs, 223 widgets and 177 NQL queries; full definitions matched curl and SDK, and representative queries matched through DataExploration. All twelve GraphQL operations in the inspected Call View bundle map to existing DataExploration methods. Teams/Zoom connector consumers also reuse existing integration services. This closes the inspected source gap without creating duplicate wrappers. Fresh Teams and Zoom call-insights requests returned HTTP 200 through curl and the SDK using a Collector UID, but reported no call data. The method therefore retains `telemetry_required` status; populated calls and the full widget scenarios remain unvalidated.

Six existing IAM/sharing methods gained positive acceptance through disposable, unassigned roles and a dashboard: role Create/Update/Delete, role content Grant/Revoke and ContentSharing.SetProfiles. Curl and SDK readbacks verified the changes. Both roles and the dashboard were removed, with cleanup verified. This tests the dashboard permission scenario, not every product's content-specific grants.

The new methods include runnable examples and JSON-backed tests. Full race tests, vet, lint and the inventory/example guard passed for the continuation. The [coverage reconciliation](web-api-coverage.md) distinguishes completed source audits from remaining fixture, telemetry, feature-availability and shared-setting validation. Neither a successful empty response nor a source-audit mapping establishes full product acceptance.
