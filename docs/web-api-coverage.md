# Web API coverage and discovery inventory

The SDK exposes 663 browser API operation methods through `client.WebAPI`, plus an immutable dashboard proxy configuration helper (664 web methods), with a runnable example for each. The product-permission discovery pass after merged PR #59 adds 106 methods across ten new resource packages and existing resources. Every addition uses the shared client and browser authentication, resource-local tests and JSON fixtures. Live-tested and source-derived contracts are distinguished below.

These counts describe SDK methods, not distinct URLs. Multiple GraphQL operations share one gateway, and IAM supports both current and legacy route dialects. The HTTP catalog has 497 entries representing 494 distinct method/path pairs, including 17 GraphQL gateways. All 453 curated follow-up HTTP contracts have implementations; the latest continuation adds 22 methods and records an existing menu route's product-area option. Of 207 captured GraphQL document variants (198 distinct operation kind/name pairs), 197 map to implementations and ten are unavailable on the active schema. The [machine-readable inventory](web-api-discovery.json) maps observed contracts to implementations and records obsolete documents separately. Discovery covers inspected first-party UI bundles through 6 October 2026; it cannot establish every API available in every tenant.

## Product coverage reconciliation: role permissions, 6 October 2026

The supplied **LBG - Superuser Sandbox** role editor provides a broader product-area checklist than the previously captured endpoint inventory. Completing the recorded discovery leads does **not** establish complete product coverage. The table below reconciles every supplied area with the current SDK. “Represented” means relevant methods exist, not that every operation has been discovered or passed live acceptance.

Evidence: the current resource implementations, the acceptance matrix, and private first-party captures of the role template, product-shell menu and UI bundles from the lab. The template independently identifies `ai_drive` manage/view/governance claims, `amplify` manage/view/packages claims, and the `assist.nexthink_assist` Workspace claim. These are permission identifiers, not inferred API paths. GetRole confirmed AI Tools manage/view/governance, Workspace and the listed VDI view permissions enabled. Amplify management was initially disabled; it is now enabled alongside view/package visibility following explicit user approval and role readback. That role change remains in place as requested. Feature gates remain independent.

| Role-editor area | Current SDK representation | Remaining coverage work |
| --- | --- | --- |
| Administration | AccessManagement, ProductConfiguration, DeviceClassification, UserClassification | The inspected product-configuration UI adds 20 methods for instance settings, classification CSVs, GeoIP and user organizations. Three reads passed; six need configured rule sets and eleven shared-setting writes remain untested. Product-shell source reuses existing methods. |
| Data privacy and view domain | AccessManagement role permissions, view-domain settings and metadata | Verify field visibility and device-scope round trips and their effect on reads; the presence of generic permission fields is not complete semantic validation. |
| Data model visibility: AI tools, agent conversations, audit logs, mobile, Nexthink Usage, platform logs, tickets | Role permission values; DataExploration/NQL query mechanisms | Eleven aggregate query scenarios now pass curl/SDK comparison: conversations, audit/custom-trend/data-export/inbound-connector/NQL logs, tickets, usage, collaboration sessions, mobile devices and VDI sessions. Dedicated AI Tools table metadata was not advertised in the lab schema. Query examples are linked below. Do not infer entity CRUD from table visibility. |
| Data management | CustomFields, RuleBasedCustomFields, CustomFieldValues, KnowledgeBases, Snapshots, NQLQueries, Ratings, CollectorManagement | All listed capability areas are represented; finish operation-level acceptance, including collector configuration and telemetry-dependent behavior. |
| AI tools | AITools plus shared DataExploration/Dashboards | Added 28 contracts covering application tools, Copilot configuration, governance, insights, module settings and adoption-goal LCRUD. Governance changes use revisioned tool Update. Twenty methods passed live; eight remain unvalidated because of integration fixtures, singleton writes or unavailable routes. |
| Alerts and Diagnostics | Monitors, AlertHub, Diagnostics | Represented; complete analytics acceptance and per-alert permission scenarios. |
| Amplify | Amplify, AmplifyAI and ActionExecutions.GetDeviceActions | Extension 1.34.0 supplied 21 contracts; administration UI 1.34.6 adds configuration Create/Update, closing the inspected management-source gap. Twelve methods passed live. Eleven AI contracts retain source/unit evidence only: two reads returned 404 and nine mutations were not called. Configuration removal uses a full-list update; no document DELETE was observed. |
| Applications | Applications, ApplicationExperience | Represented; verify remaining telemetry-dependent analytics and per-application grants. |
| Campaigns | Campaigns, public campaign execution | Management and trigger operations represented; finish execution/dashboard acceptance and per-campaign grants. |
| Collaboration Tools | CollaborationTools, Dashboards, DataExploration and Teams/Zoom integration services | Three dashboard definitions (13 tabs, 223 widgets, 177 NQL queries) and representative query replays use existing services. Call View UI 2.8.0 maps all twelve GraphQL operations to DataExploration. The inspected source gap is closed; populated widget/call scenarios remain unverified. Call insights now return 200 with no-data explanations, retaining `telemetry_required` status. |
| Desktop Virtualization: Amazon WorkSpaces, Citrix CVAD/DaaS, Microsoft 365, AVD, VMware Horizon and general dashboards | VDI's four methods, shared Dashboards/DataExploration | All ten advertised vendor VM/infrastructure/hypervisor dashboard definitions now match through curl and Dashboards.Get with product area desktop-virtualization. Session views use shared NQL. Populated vendor telemetry and the full set of widget scenarios remain unverified; no duplicate vendor services were added. |
| Device View | Support, SupportTimeline, SupportInsights, SupportChecklists | Represented; remaining telemetry/drilldown acceptance and per-checklist grants. |
| Digital Experience | DexScores, DexConfiguration, Benchmark, CCIInsights, CCIBenchmarks | Score management and dashboard operations represented; complete populated analytics validation and reconcile dashboard navigation against the operation inventory. |
| Investigations | Investigations, GlobalSearch, DataExploration, ContentSharing | Private/shared investigation and search operations represented; finish sharing/access-scope validation. |
| Live dashboards | Dashboards, ContentSharing, CollaborationComments | Saved and nested dashboard operations represented; reconcile product-specific dashboards and outstanding sharing/comment acceptance. |
| Nexthink Library | Library and resource-specific library/import methods | Catalog and installation/update workflows represented; finish outstanding mutation acceptance. |
| Remote actions | RemoteActions, ActionExecutions, public remote-action execution | Create/update and other lifecycle methods exist; finish execution/dashboard acceptance and per-action grants. |
| Software Metering | SoftwareMetering | Lifecycle and analytics represented; remaining telemetry acceptance and a reproduced usage-distribution backend error. |
| Workflows | Workflows, WorkflowExecutions, ConnectorCredentials, public workflow execution | Create/update, execution and credential operations represented; finish execution acceptance and per-workflow grants. |
| Workspace | Workspace, WorkspaceAgents, WorkspaceTasks, WorkspaceAssignments | Added 35 contracts from the loaded current Workspace code. Fifteen passed live, including chat/conversations/files. Custom agents and tasks expose explicit feature gates in the lab; assignment-specific operations need fixtures. Customer UI contracts are distinct from development-only deployment-routing controls. |
| Resource-specific grants | AccessManagement grant/revoke operations and ContentSharing | Six existing role/sharing methods now pass using disposable, unassigned roles and a dashboard; fixtures were cleaned. This validates the tested dashboard scenario, not every alert, application, campaign, checklist, investigation, remote-action or workflow content type. |

### Effect on the acceptance gaps

The [current acceptance report](acceptance/README.md) records **692 exported methods: 418 passed and 274 without completed positive acceptance** across both API families. Counts include helpers and only implemented methods, not unknown product endpoints. Before this discovery pass, the baseline was **586 methods: 358 passed and 228 outstanding**. The table below is that historical baseline, not the current category distribution.

| Existing acceptance category | Methods | Interpretation after reconciliation |
| --- | ---: | --- |
| Not tested | 90 | Outstanding test work; not evidence of an external permission blocker. |
| Telemetry required | 64 | Role permissions do not create device events, sessions, calls or usage data. |
| Fixture required | 50 | Appropriate objects or identifiers still need to be prepared or discovered. |
| Permission or availability | 16 | Unresolved HTTP failures or route/fixture availability, not 16 proven RBAC denials. Includes Autopilot/recommendations 401s, comments/call insights 403s, a gateway authorization-format error and several 404s. Compare effective claims, licensing, feature routing and successful UI requests before assigning a cause. |
| Consent required | 6 | Previously recorded action-authorization limits; assigning a product permission is separate from authorizing a particular test action. |
| Server error | 2 | Reproduced backend failures; the supplied role does not resolve their cause. |

The inspected Amplify administration and collaboration dashboard/Call View sources now have implementation or shared-method mappings. The latter adds no new endpoint: its twelve GraphQL operations reuse DataExploration, while the product-area menu filter extends an existing method. Vendor dashboard definitions and eleven data-table scenarios also reuse existing APIs; runnable examples are documented under [Dashboards.Get](../examples/nexthink/web_api/dashboards/Get/README.md) and [DataExploration.Query](../examples/nexthink/web_api/data_exploration/Query/README.md). Remaining work includes populated telemetry, shared-setting writes, gated features and content-specific permission scenarios. This source audit does not establish exhaustive product coverage. LCRUD applies where an entity lifecycle exists; read-only dashboards, aggregate settings and data-table visibility do not imply five separate operations.

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
| CollaborationComments | 11 | Comment threads, replies, archive state, user mentions and document identifier resolution |
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
| ProductConfiguration, DeviceClassification, UserClassification | Instance configuration read/create/update; classification CSV read/create/update/download, VPN egress deletion and GeoIP read/update; user organizations read and full-list replacement. No additional list/delete lifecycle inferred |
| AITools | Application-tool and Copilot configuration LCRUD, governance through Update, adoption-goal LCRUD and aggregate module settings; no module delete contract observed |
| Workspace, WorkspaceAgents, WorkspaceTasks, WorkspaceAssignments | Conversations created/continued through Chat; conversation edits/deletion/sharing/files, custom-agent LCRUD and knowledge uploads, task LCRUD/reconciliation, assignment reads/updates |
| Amplify, AmplifyAI | Extension search/device/user/package reads and usage events; configuration read/create/update with full application-list replacement, no observed document DELETE; AI analysis, feedback, resolution-plan and action/ticket operations |
| CollaborationComments | Thread/reply lifecycle, archive operations, mentions and document identifier resolution |
| ContentSharing, CustomFieldValues | Sharing/value reads and writes; CSV validation/import; no independent entity lifecycle invented |
| Autopilot, Recommendations | Observed settings, approval, input and status operations; no unevidenced delete methods |
| Analytics, support, execution insights, metadata, ProductShell, License, NQLEditor | Query/analysis/reference operations; no entity lifecycle invented for read-only or aggregate surfaces |

Workflows and remote actions have genuine create/update operations; their library preview methods retrieve templates without installing them. Workflow management UUIDs, NQL IDs, public execution IDs and library UUIDs have different uses. Operation examples document their inputs. Dashboard mutations carry revision and product/tab context. Rating/checklist deletion carries revision and the observed request body.

## Validation and corrections

Curl preflight established observed requests and tenant state before SDK replay. Successful SDK representations were compared with decoded curl responses, including populated collections, aggregate query results, library templates and nullable fields. Amplify's initially absent configuration was created once through the SDK and independently verified with curl GET; no duplicate curl POST was attempted. Curl and SDK updates verified the configuration lifecycle. All test applications were removed, leaving an empty configuration with `itsmConfigList: []` and `enableUsageDataReporting: false`; no document DELETE was observed. The approved Amplify management role change remains enabled. Other outstanding aggregate configuration and library installation writes retain source/unit evidence without live-success claims. Analytics without suitable telemetry can return empty data or backend errors; those outcomes remain validation gaps. See [lab validation](lab-validation.md) for the exact limits.

The shared GraphQL client now negotiates `Accept: */*`, matching the browser. One visual-editor endpoint returned base64 text mislabeled as gzip when sent `Accept: application/json`; a curl negotiation matrix reproduced the issue. Wildcard negotiation returned genuine gzip. A regression test exercises compressed responses through the real transport. GraphQL requests can also carry per-request UI time-context headers without serializing those headers into the JSON body. Dashboard and data-exploration resolvers require that context.

The shared transport also buffers typed responses so `interfaces.Response.Body` remains available after JSON decoding. It uses byte access rather than the whitespace-trimming string accessor, preserving binary downloads and exact response bytes. Tests cover small, large, compressed and malformed responses plus configured size limits.

Models retain the observed nullability of library workflow identifiers/timestamps and remote-action titles. Checklist library models preserve absent custom labels. Partial GraphQL data remains available alongside GraphQL errors under HTTP 200. Tests inspect method/path, headers and JSON payloads, full positive projections, HTTP errors, partial GraphQL results, malformed JSON, transport failures and validation before requests.

## Obsolete documents and remaining discovery

Ten bundled GraphQL documents are incompatible with the active tenant schema (confirmed through schema checks and representative curl failures): eight legacy application overview operations (`AveragePageViewsPerEmployee`, `ErrorCount`, `NumberOfEmployees`, `OverviewTooltips`, `PageLoadTime`, `TransactionTime`, `UsageTime`, `WaitingTime`) and two old Alert Hub views (`DeviceView`, `DiagnosticView`). They are recorded as unavailable, with modern alternatives where established. They are not exported as working SDK methods. The library's unused `getContentTypeConfig` helper returned 404 and is also recorded separately.

The concrete operation backlog and path leads are reconciled independently. The previously recorded API prefixes, specific paths and alternate proxy leads now have source-audit outcomes in the machine-readable inventory. A recovered consumer either maps to an implementation, shares an existing method, or is explained as initialization/static/non-HTTP behavior without an independent endpoint. Legacy portal authentication now has explicit scoped support, with no live-success claim for the cloud lab.

A generic GraphQL gateway still does not establish typed coverage of every possible operation. The ten obsolete documents above remain unsupported by the active tenant schema; source-audit closure is not a claim of universal Nexthink coverage or full live validation.

Fixtures are synthetic. Tokens, credentials, raw tenant captures and signed download URLs remain outside the repository. Unit tests and examples demonstrate the implemented contracts; they do not establish exhaustive Nexthink coverage or successful execution of every product subtype.

## Acceptance continuation — 8 October 2026

The [104-method triage](acceptance/2026-10-08-untested-triage.json) separates actual invocations from prerequisite reviews. Four new passes bring the matrix to 418 passed and 274 outstanding. Partial Library installation responses retain their exact JSON fields; legacy sharing grants and revocations passed independent readback. Three Autopilot reads reproduced backend restrictions. These checks add no SDK method or catalog endpoint. The remaining 97 methods in this triage were not invoked, and older temporary raw evidence is no longer available; see the [acceptance report](acceptance/README.md).
