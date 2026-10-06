# Web API coverage and discovery inventory

The SDK exposes 343 browser API resource methods through `client.WebAPI`, with a runnable example for each. The continuation from merged PR #52 adds 162 methods and 10 resources. Every addition uses the shared client, browser authentication, resource package layout, unit tests and synthetic JSON fixtures.

These counts describe SDK methods, not distinct URLs. Multiple GraphQL operations share one gateway; document variants may select different fields or target different gateways. The HTTP catalog now has 158 entries, including 16 GraphQL gateways. All 99 follow-up HTTP contracts are implemented. Of 205 captured GraphQL document variants (196 distinct operation kind/name pairs), 195 are mapped to implementations and ten are unavailable on the active schema. The [machine-readable inventory](web-api-discovery.json) maps observed contracts to implementations and records obsolete documents separately. Discovery covers the inspected first-party UI bundles through 6 October 2026; it cannot establish every API available in every tenant.

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
| Analytics, metadata, ProductShell, License, NQLEditor | Query/analysis/reference operations; no entity lifecycle invented for read-only or aggregate surfaces |

Workflows and remote actions have genuine create/update operations; their library preview methods retrieve templates without installing them. Workflow management UUIDs, NQL IDs, public execution IDs and library UUIDs have different uses. Operation examples document their inputs. Dashboard mutations carry revision and product/tab context. Rating/checklist deletion carries revision and the observed request body.

## Validation and corrections

Curl established the contracts before SDK replay. Successful SDK representations were compared with decoded curl responses, including populated collections, aggregate query results, library templates and nullable fields. Analytics without suitable telemetry can return empty data or backend errors; those outcomes are recorded rather than described as successful analytics. Persistent aggregate configuration and library installation writes have source-derived contracts and unit tests but were not applied to existing tenant settings or installed content. See [lab validation](lab-validation.md) for the exact limits.

The shared GraphQL client now negotiates `Accept: */*`, matching the browser. One visual-editor endpoint returned base64 text mislabeled as gzip when sent `Accept: application/json`; a curl negotiation matrix reproduced the issue. Wildcard negotiation returned genuine gzip. A regression test exercises compressed responses through the real transport. GraphQL requests can also carry per-request UI time-context headers without serializing those headers into the JSON body. Dashboard and data-exploration resolvers require that context.

The shared transport also buffers typed responses so `interfaces.Response.Body` remains available after JSON decoding. It uses byte access rather than the whitespace-trimming string accessor, preserving binary downloads and exact response bytes. Tests cover small, large, compressed and malformed responses plus configured size limits.

Models retain the observed nullability of library workflow identifiers/timestamps and remote-action titles. Checklist library models preserve absent custom labels. Partial GraphQL data remains available alongside GraphQL errors under HTTP 200. Tests inspect method/path, headers and JSON payloads, full positive projections, HTTP errors, partial GraphQL results, malformed JSON, transport failures and validation before requests.

## Obsolete documents and remaining discovery

Ten bundled GraphQL documents are incompatible with the active tenant schema (confirmed through schema checks and representative curl failures): eight legacy application overview operations (`AveragePageViewsPerEmployee`, `ErrorCount`, `NumberOfEmployees`, `OverviewTooltips`, `PageLoadTime`, `TransactionTime`, `UsageTime`, `WaitingTime`) and two old Alert Hub views (`DeviceView`, `DiagnosticView`). They are recorded as unavailable, with modern alternatives where established. They are not exported as working SDK methods. The library's unused `getContentTypeConfig` helper returned 404 and is also recorded separately.

The concrete operation backlog is reconciled independently from broader API path leads. Access management, content sharing, collaboration comments, support/device timelines, VDI, workflow execution insights, data export jobs, NLP/assistance, Autopilot and global search still have base-path references requiring additional contract discovery. A path literal is not enough evidence to invent request models or LCRUD. Likewise, a generic GraphQL gateway does not establish typed coverage of every possible operation.

Fixtures are synthetic. Tokens, credentials, raw tenant captures and signed download URLs remain outside the repository. Unit tests and examples demonstrate the implemented contracts; they do not establish exhaustive Nexthink coverage or successful execution of every product subtype.
