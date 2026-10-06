# Web API coverage and discovery inventory

The SDK exposes 557 browser API operation methods through `client.WebAPI`, plus an immutable dashboard proxy configuration helper, with a runnable example for each. The latest continuation from merged PR #54 adds 40 operation methods and 11 resource packages. Every addition uses the shared client, browser authentication, resource package layout, unit tests and synthetic JSON fixtures.

These counts describe SDK methods, not distinct URLs. Multiple GraphQL operations share one gateway, and IAM supports both current and legacy route dialects. The HTTP catalog has 391 entries representing 389 distinct method/path pairs, including 17 GraphQL gateways. All 346 curated follow-up HTTP contracts have implementations. Of 207 captured GraphQL document variants (198 distinct operation kind/name pairs), 197 map to implementations and ten are unavailable on the active schema. The [machine-readable inventory](web-api-discovery.json) maps observed contracts to implementations and records obsolete documents separately. Discovery covers inspected first-party UI bundles through 6 October 2026; it cannot establish every API available in every tenant.

## Discovery lead completion after PR #54

| Resource | Added operations | Coverage |
| --- | ---: | --- |
| TeamsCredentials, ZoomNotifications, AzureADCredentials | 4 | Source-observed credential checks and Zoom application metadata |
| UserCommunicationIntegrations | 6 | Communication integration LCRUD and Azure connector listing |
| LegacyConnectors | 2 | Secret-presence checks and partial secret updates |
| QueryBuilder | 3 | Query transformation, drilldown destinations and drilldown transformation |
| CCIBenchmarks | 1 | Batched benchmark queries |
| Dashboards | 1 | Product-shell menu; separate `WithAsyncProxy` helper routes existing GraphQL operations through the observed proxy |
| ApplicationExperience | 1 | Application insights through the observed asynchronous proxy |
| Campaigns | 1 | End-user feature/claim metadata |
| GlobalSearch | 2 | Explicit-session legacy portal token and dashboard search |
| Appearance | 4 | Modern raw-image retrieval/update and legacy portal asset retrieval/save |
| MobileTokens | 5 | Collector mobile enrollment token LCRUD |
| Snapshots | 7 | Custom-trend definition LCRUD and portable export/import helpers |
| UIEvents | 1 | Paged UI event polling with same-origin continuation links |
| Observability | 1 | Raw browser telemetry submission through the fixed proxy |
| CollectorManagement | 1 | Filtered Collector update-status device queries |

All 64 previously outstanding path leads have been reconciled against their inspected first-party consumers: ten specific leads, 52 partially covered prefixes and two proxy routes. Existing operations are mapped where the source uses an already supported contract; new leaf contracts are implemented. Snapshot import/export reuse Create/Get rather than inventing separate routes. Appearance reset uploads the default image through Update. The audit is bounded by the recorded source versions; it does not prove that unshipped, feature-gated or future backend APIs do not exist.

Legacy portal calls use explicitly supplied session credentials on fixed allowlisted routes, with no bearer token or automatic cookie persistence. The modern tenant did not provide a compatible legacy session, so those methods have source and unit validation. The shared transport redacts query values from errors and logs while preserving error inspection. Legacy connector configuration now distinguishes omitted `enabled` from explicit false and accepts the empty mapping arrays used by Teams/Zoom.

The example guard now inspects all resource Go files, not only `crud.go`. This also exposed and filled six existing public NQL helper-example gaps. The root README follows the reference SDK's badge and navigation layout, with a [quick start](guides/quick-start.md) for both authentication families and an explanation of the browser API's role.

## Additions after PR #53

| Resource | Added methods | Coverage |
| --- | ---: | --- |
| AccessManagement | 45 | Users, roles/profiles, SAML, API credentials, account and permission operations; current and legacy routing |
| CollaborationComments | 11 | Comment threads, replies, reactions, archive state, counts and preferences |
| Support, SupportChecklists, SupportTimeline, SupportInsights | 44 | Device search/detail, checklist evaluation, timeline events and drilldowns, generated insights |
| CollaborationTools, VDI | 5 | Call insights and VDI session/context reads |
| WorkflowExecutions, ActionExecutions | 21 | Execution history, metadata, inspection, input and execution operations |
| Autopilot | 21 | Configuration, settings, approvals, calls, knowledge references, conversations, tickets and agent-action inputs |
| GlobalSearch, NLPAssistant | 2 | Category search events and buffered assistant chat events |
| Recommendations | 2 | Knowledge recommendations and status updates |
| DataExport | 2 | Start a browser export and poll its status |
| VisualEditor | 3 | Collections, default columns and filter collections |
| ContentSharing | 8 | Current and legacy sharing permissions, profiles and owner metadata |
| CustomFieldValues | 4 | Field metadata, value updates and CSV validation/import |
| Applications | 2 | Application template list and detail |
| CustomFields, RuleBasedCustomFields | 4 | Validation patterns and manual/computed/rule-based export/import helpers |

The root client exposes `AccessManagement` and `LegacyAccessManagement` using the same transport and credentials, selecting the appropriate route dialect. These do not count as two sets of 45 SDK methods. The new HTTP catalog entries include 34 legacy dialect routes; shared create/update and archive/unarchive URLs are counted once per HTTP method.

Software Metering Create, Update and Delete now include local request samples, and all five original lifecycle examples have operation guides. Create returns a boolean; List supplies the generated configuration UUID. A disposable curl and SDK lifecycle verified the existing Create request contract.

## Additions after PR #52

| Resource | Added methods | Coverage |
| --- | ---: | --- |
| ApplicationExperience | 23 | Application insights, metrics, breakdowns, investigations and suggestions |
| SoftwareMetering | 13 | Configuration details, applications/packages, employees, usage breakdowns/distributions, automatic configuration |
| AlertHub | 12 | Alerts, investigations, filters, monitor details, diagnostics and tags |
| Diagnostics | 13 | Dashboard/widget analysis, executions, devices, trends and investigation queries |
| Benchmark | 7 | Metrics, definitions, comparisons, recommendations and search |
| DexScores | 16 | Scores, trends, breakdowns, drivers, applications and device/user insights |
| CCIInsights | 3 | CCI diagnostic insights and datasets |
| NetworkInsights | 1 | Network graph insights |
| DexConfiguration | 11 | Application, campaign, score and threshold configuration reads/writes |
| DataExploration | 12 | NQL, inspection, collections/field metadata, ratings, filters, breakdowns, menus and item tooltips |
| Library | 20 | Content/packs, metadata, dependencies, locales, installation/update and status operations |
| Workflows | 3 | Library definition, import and activation |
| RemoteActions | 3 | Library definition, import and export |
| Campaigns | 6 | NQL-ID/legacy reads, library definition, branding and status |
| Monitors | 8 | Built-in updates, activity, analysis, impact query, filter fields, tags, metadata and license |
| Dashboards | 3 | Configuration, collections and field metadata |
| Ratings | 4 | Field/action metadata, collected action values and export |
| Checklists | 1 | Built-in library definition |
| Investigations | 3 | NQL execution, query metadata and shareable investigation links |

Each method has a callable example in the [example index](../examples/nexthink/web_api/README.md). A source-level test requires examples for every exported resource method in both API families. Successful compilation alone does not imply successful live execution.

## LCRUD coverage

| Resources | Available lifecycle |
| --- | --- |
| Workflows, RemoteActions, Applications, WritingAssistant, SoftwareMetering, CustomFields, RuleBasedCustomFields, Monitors, Campaigns, NQLQueries | List, create, read, update and delete; auxiliary operations are included where recovered |
| Assets, Checklists, Dashboards, Ratings, Investigations | Full saved-content lifecycle; assets read via signed URL; nested dashboard operations and import/export are included |
| Connectors, LegacyConnectors, Webhooks, DataExporters | Definition lifecycle and recovered test/status operations; subtype validation varies |
| ConnectorCredentials | List/read plus POST upsert and clear/disable; deletion retains the allocated identifier |
| KnowledgeBases | Listing, indexed contents, file/multipart upload, registration, download URL and delete; no independent update contract observed |
| Library | Catalog/pack reads and evidenced installation, update, dependency and status workflows; these are not generic entity CRUD |
| DexConfiguration, DeviceConfiguration, CollectorManagement | Aggregate configuration reads/writes; no independent entity create/delete contract established |
| AccessManagement | User and role/profile lifecycle, credentials and aggregate identity settings; optional legacy route dialect |
| CollaborationComments | Thread/reply lifecycle, reaction and archive operations |
| ContentSharing, CustomFieldValues | Sharing/value reads and writes; CSV validation/import; no independent entity lifecycle invented |
| Autopilot, Recommendations | Observed settings, approval, input and status operations; no unevidenced delete methods |
| Analytics, support, execution insights, metadata, ProductShell, License, NQLEditor | Query/analysis/reference operations; no entity lifecycle invented for read-only or aggregate surfaces |

Workflows and remote actions have genuine create/update operations; their library preview methods retrieve templates without installing them. Workflow management UUIDs, NQL IDs, public execution IDs and library UUIDs have different uses. Operation examples document their inputs. Dashboard mutations carry revision and product/tab context. Rating/checklist deletion carries revision and the observed request body.

## Validation and corrections

Curl established the contracts before SDK replay. Successful SDK representations were compared with decoded curl responses, including populated collections, aggregate query results, library templates and nullable fields. Analytics without suitable telemetry can return empty data or backend errors; those outcomes are recorded rather than described as successful analytics. Persistent aggregate configuration and library installation writes have source-derived contracts and unit tests but were not applied to existing tenant settings or installed content. See [lab validation](lab-validation.md) for the exact limits.

The shared GraphQL client now negotiates `Accept: */*`, matching the browser. One visual-editor endpoint returned base64 text mislabeled as gzip when sent `Accept: application/json`; a curl negotiation matrix reproduced the issue. Wildcard negotiation returned genuine gzip. A regression test exercises compressed responses through the real transport. GraphQL requests can also carry per-request UI time-context headers without serializing those headers into the JSON body. Dashboard and data-exploration resolvers require that context.

The shared transport also buffers typed responses so `interfaces.Response.Body` remains available after JSON decoding. It uses byte access rather than the whitespace-trimming string accessor, preserving binary downloads and exact response bytes. Tests cover small, large, compressed and malformed responses plus configured size limits.

Models retain the observed nullability of library workflow identifiers/timestamps and remote-action titles. Checklist library models preserve absent custom labels. Partial GraphQL data remains available alongside GraphQL errors under HTTP 200. Tests inspect method/path, headers and JSON payloads, full positive projections, HTTP errors, partial GraphQL results, malformed JSON, transport failures and validation before requests.

## Obsolete documents and remaining discovery

Ten bundled GraphQL documents are incompatible with the active tenant schema (confirmed through schema checks and representative curl failures): eight legacy application overview operations (`AveragePageViewsPerEmployee`, `ErrorCount`, `NumberOfEmployees`, `OverviewTooltips`, `PageLoadTime`, `TransactionTime`, `UsageTime`, `WaitingTime`) and two old Alert Hub views (`DeviceView`, `DiagnosticView`). They are recorded as unavailable, with modern alternatives where established. They are not exported as working SDK methods. The library's unused `getContentTypeConfig` helper returned 404 and is also recorded separately.

The concrete operation backlog and path leads are reconciled independently. The previously recorded API prefixes, specific paths and alternate proxy leads now have source-audit outcomes in the machine-readable inventory. A recovered consumer either maps to an implementation, shares an existing method, or is explained as initialization/static/non-HTTP behavior without an independent endpoint. Legacy portal authentication now has explicit scoped support, with no live-success claim for the cloud lab.

A generic GraphQL gateway still does not establish typed coverage of every possible operation. The ten obsolete documents above remain unsupported by the active tenant schema; source-audit closure is not a claim of universal Nexthink coverage or full live validation.

Fixtures are synthetic. Tokens, credentials, raw tenant captures and signed download URLs remain outside the repository. Unit tests and examples demonstrate the implemented contracts; they do not establish exhaustive Nexthink coverage or successful execution of every product subtype.
