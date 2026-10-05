# Web API coverage and discovery inventory

This inventory separates implemented methods from additional APIs referenced by Nexthink's first-party UI. It reflects the lab UI inspected on 2026-10-05. It is not an exhaustive tenant-independent API specification.

The HTTP catalog has 78 operation entries, including 12 GraphQL gateways. Some concrete listing entries also match the shared parameterized content route; entries are not a count of distinct backend handlers. A reachable GraphQL gateway is not full coverage of its queries and mutations. This expansion adds typed LCRUD resources for workflows, remote actions, applications, Writing Assistant, Software Metering, manual/computed Custom Fields, Monitors, and Campaigns, completes saved-query LCRUD with List, and adds script inspection. GraphQL methods reuse existing gateways and listing methods reuse the content-administration route; method counts are not new HTTP-route counts.

## Newly implemented and live validated

The read/inspection operations below succeeded with curl and with `client.WebAPI`, using the browser session. The nine original read/inspection SDK responses were compared with curl data envelopes. Both management resources also passed independent curl and SDK LCRUD lifecycles on dedicated fixtures, with update read-back and deletion confirmed. Script inspection parses supplied bytes; it does not execute scripts on devices.

| Resource | SDK method | UI GraphQL operation |
| --- | --- | --- |
| Workflows | `Create` / `Update` / `Delete` | `CreateWorkflowMutation` / `UpdateWorkflowMutation` / `DeleteWorkflowMutation` |
| RemoteActions | `Create` / `Update` / `Delete` | `CreateRemoteAction` / `UpdateRemoteAction` / `DeleteRemoteAction` |
| Workflows | `List` | `GetWorkflowsListQuery` |
| Workflows | `Get` | `GetWorkflowQuery` |
| Workflows | `Export` | `ExportWorkflowQuery` |
| RemoteActions | `Get` | `GetRemoteAction` |
| RemoteActions | `GetForView` | `GetRemoteActionByUIdForView` |
| RemoteActions | `GetContentVolume` | `GetContentVolumeDetails` |
| RemoteActions | `InspectBashScript` | `GetInputAndOutputFromBashScript` |
| RemoteActions | `InspectPowerShellScript` | `GetInputAndOutputFromPowershellScript` |
| RemoteActions | `GetPowerShellSignature` | `GetPowershellScriptSignatureInfo` |

Remote Actions `List` uses `/apigateway/content-administration/api/v2/contents/remoteactions`; saved-query `List` uses the same listing service with `fe-nqlapi-nx-content-administration-config`. Both were curl-validated and then called through the SDK. Remote-action list rows preserve management-specific fields, including NQL IDs, targeting and script metadata.

Workflows use `/apigateway/workflows/manage/graphql`; Remote Actions use `/apigateway/act/manage/graphql`. The exact query documents are checked in under each resource's `queries/` directory. They come from the workflow management and Remote Actions UI bundles, including lazily loaded Remote Actions chunks.

`Get` takes the management UUID. Public API execution IDs and management UUIDs are not interchangeable. Export retains the service's opaque string. Script methods accept bytes and base64-encode them once: macOS bytes must be a tar.gz script archive, and PowerShell bytes must include the UTF-8 BOM. The server rejected plain source/base64 source in the wrong formats. Tests cover these wire encodings, null fields, HTTP errors, GraphQL errors under HTTP 200, partial data, malformed responses and transport failures.

The lab returned populated workflow versions, remote-action inputs/outputs and Windows script metadata, and null optional fields. Scheduled-task/event-trigger payloads remain opaque because their nested value schemas were not populated in this lab. Checked-in JSON fixtures and script samples are synthetic, not raw tenant captures. A valid unsigned PowerShell sample returned `NOT_SIGNED`; this does not establish coverage of every signed-script state.

## LCRUD coverage audit

| Existing resource | List | Create | Read | Update | Delete | Evidence or limitation |
| --- | --- | --- | --- | --- | --- | --- |
| Workflows | `List` | `Create` | `Get` | `Update` | `Delete` | Complete lifecycle passed with curl and SDK; inactive workflow, no triggers |
| RemoteActions | `List` | `Create` | `Get` | `Update` | `Delete` | Complete lifecycle passed with curl and SDK; all triggers disabled; script saved but never executed |
| Applications | `List` (paged) | `Create` | `Get` | `Update` | `Delete` | Curl and all five SDK examples passed; revision-aware update/delete; unused `.invalid` URL and collection enhancements disabled |
| WritingAssistant | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and all five SDK examples passed; uses a dedicated disposable application |
| SoftwareMetering | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and all five SDK examples passed; real threshold payload; mutations return booleans, Create does not return a UUID |
| CustomFields | `List` | `Create` | `Get` | `Update` | `Delete` | MANUAL and COMPUTED lifecycles passed with curl and SDK examples; RULE_BASED uses the separate `RuleBasedCustomFields` service; all five examples also passed |
| RuleBasedCustomFields | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and all five SDK examples passed; List filters the shared listing; Delete is POST with revision and object type; rules matched no devices |
| Monitors | `List` | `Create` | `Get` | `Update` | `Delete` | Custom metric monitor lifecycle passed with curl and SDK examples; no notifications; UUID differs from content ID; built-in update and other monitor operations still pending |
| Campaigns (web) | `List` (paged) | `Create` | `Get` | `Update` | `Delete` | Curl and SDK examples passed on drafts, including populated answer choices; no publication or delivery; status changes, branding and other operations remain pending |
| NQLQueries | `List` | `Create` | `Get` | `Update` | `Delete` | CRUD validated in PR #48; listing added and live validated here |
| Assets | `List` | `Create` | `GetSignedURL` | `Update` | `Delete` | Curl and SDK lifecycles passed; raw data URL upload; update/delete 204; replacement bytes compared; filename stays unchanged |
| Checklists | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and SDK lifecycles passed; populated field metadata; revision query and nonempty DELETE body |
| Dashboards | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and SDK lifecycles passed; revision/context-aware; nested widget/filter/tab mutations and import/export still pending |
| Ratings | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and SDK lifecycles passed on previously unrated field; complete NQL conditions; DELETE boolean response |
| Investigations | `List` | `Create` | `Get` | `Update` | `Delete` | Curl and all seven SDK examples passed, including Export/Import; uses uid; server ignores description |
| CollectorManagement | Version/platform/group queries | No independent entity-create route observed | Links/configuration | `SetUpdateConfiguration` | No independent delete route observed | Aggregate configuration API; write is implemented but not replayed against shared tenant settings |
| DeviceConfiguration | `GetProfiles` | No profile-create route observed | Profiles/settings | `SetProfiles`, `SetSettings` | No profile-delete route observed | `SetProfiles` updates named settings on existing profiles using `{settings:[{profileId,name,newValue}]}`; curl and SDK saves preserved all seven values and completed onboarding. Legacy settings GET returned 403; its write remains unverified. |
| ProductShell | Menu/modules | Not established as an entity resource | User/config/flags | No entity update route observed | No entity delete route observed | Shell and claims helpers are not an entity-management LCRUD implementation |
| License | No list route established | No create route observed | Feature status | No update route observed | No delete route observed | Entitlement status reads; do not invent license mutations |
| ContentAdministration | `List` | Delegated to product management APIs | Configuration/list | Delegated to product management APIs | Delegated to product management APIs | Shared listing infrastructure; remote-action/workflow/query writes belong to their own resources |
| NQLEditor | Not an entity collection | Not an entity operation | Highlighting/hover | Validation/completion requests are analysis operations | Not an entity operation | Saved-query lifecycle belongs to NQLQueries |
| GraphQL gateways | Incomplete | Incomplete | Generic Execute | Incomplete | Incomplete | Reachability and a generic GraphQL helper do not count as typed product-resource coverage |

Create/update use separate write models. Workflow writes omit `lastUpdateTime` and version `valid`; definitions are XML strings. Custom create IDs must start with `#`. Workflow Update/Delete use the workflow UUID, while remote-action Update/Delete use the NQL ID; RemoteAction Get uses content UUID. Mutation responses retain any partial data alongside GraphQL errors, including newly created identifiers needed for cleanup.

Workflow copy/activation/import and remote-action library/export operations were also recovered. They remain outside the implemented LCRUD set and are not claimed as validated methods. Broader product gateways below remain incomplete until each available LCRUD contract is implemented and tested. Absence of a route in the inspected bundles is not proof that an operation cannot exist.

## Additional API areas found in first-party code

Paths below are references or base paths unless explicitly stated otherwise. Their presence does not establish every method, input schema, permission or successful response. They must not be counted as implemented SDK resources.

| Area | Observed path or base path | Source UI bundle | Remaining work |
| --- | --- | --- | --- |
| Investigations | `/apigateway/inv/`, `/apigateway/store/investigation`, `/apigateway/query-builder` | `investigationsApp.js` | Saved content LCRUD and export/import implemented and live validated; query execution/metadata/link helpers captured but pending |
| Custom fields | `/apigateway/nedm/customfields/graphql`, `/apigateway/nedm/customfields/api/v1/rbcf`, `/apigateway/tlm/customfields/api/v1` | custom-field-manager default export | Manual, computed and rule-based definition LCRUD implemented and live validated; value imports and other auxiliary operations remain pending |
| Alert hub and monitors | `/apigateway/mnt/alert/hub`, `/apigateway/mnt/alert/config/graphql` | `nxAlertHub.js`, `nxmonitorconfig.js` | Custom metric monitor LCRUD implemented; built-in update, activity, metadata, import/export and other types still require follow-up |
| Dashboards | `/apigateway/dash`, `/apigateway/dash/graphql`, `/apigateway/proxy/request/dash-graphql-gateway` | dash-web, network-view | Dashboard LCRUD implemented and live validated; widget/filter/tab mutations, clone and import/export remain pending |
| Campaign management | `/apigateway/euf-gateway/graphql`, `/apigateway/api/v1/euf/features` | `euf-manager.js` | Core web LCRUD implemented and live validated; publication, branding, library and multilingual scenarios still require follow-up |
| Connectors and integrations | `/apigateway/connector`, `/apigateway/connector/v1`, `/apigateway/user-communication-integrations/api` | Teams/Zoom enrichers, integrations manager | Universal connector LCRUD and test contracts recovered; list/templates reads passed; connector writes remain pending |
| Workflow connectors and execution insights | `/apigateway/workflows/manage/api/externals/third-party-connectors/v1/connectors`, `/apigateway/workflow-executions-insights/api` | `eaContentManagerUI.js` | Recover REST method/parameter contracts |
| Knowledge bases and files | `/apigateway/knowledge-manager/api/v1/knowledgebase`, `/apigateway/knowledge-manager/api/v1/file` | `knowledgeBasesUi.js` | List/detail plus upload/import contracts |
| Assets | `/apigateway/asset-manager/api/v1/assets`, `/apigateway/api/asset/signed-url` | `assetManagerUi.js` | LCRUD implemented with signed-URL read; file replacement bytes verified; signed URLs kept private |
| Content sharing and library | `/apigateway/coad/v1`, `/apigateway/library-service/v1` | `contentBuiltinUi.js`, application experience | Sharing/library APIs extend beyond current content-administration lists |
| Access management | `/apigateway/nxarmmt/`, `/apigateway/iam/`, `/apigateway/nxarmrole/api/`, `/apigateway/iam/ui/` | `nxarm.js`, autopilot cockpit | Recover read-only role/permission APIs before write contracts |
| Ratings | `/apigateway/nedm/ratings/api` | nedm-ratings default export | Revision-aware LCRUD implemented and live validated; export and metadata contracts remain pending |
| Collaboration comments | `/apigateway/rtc/api/v1/documents`, `/apigateway/rtc/api/v1/user-mentions`, `/apigateway/rtc/api/v1/identifier` | rtc-comments chunk 268 | Document/comment read contracts and pagination |
| Support, timelines and checklists | `/apigateway/atl/support-be/api`, `/apigateway/atl/support-device-timeline-be/api`, `/apigateway/atl/support-checklist-config-be/api`, `/apigateway/atl/support-checklist-values-be/api`, `/apigateway/atl/support-views-insights-be/api` | `supportFe.js`, `supportChecklistConfigFe.js` | Checklist LCRUD implemented; device reads, action execution and checklist auxiliary APIs remain separate follow-up work |
| VDI and action execution | `/apigateway/vdi/vdi-service/api`, `/apigateway/act/api/v2` | `vdiUi.js`, workflows, support | Recover VDI reads and compare execution APIs with public API support |
| Application experience | `/apigateway/bus/appexgw/rest/v1/templates`, `/apigateway/proxy/request/appex-insights-application` | `bsapp.js` | Template and application insights contracts |
| Data export | `/apigateway/dataexport/api/v1` | application experience, metering | Job creation/status/download contract; separate from public NQL export |
| NLP and assistance | `/apigateway/nlp/nlp-gateway/api` | assist-ui chunk 706 | Recover application operations and feature availability |
| Autopilot | `/apigateway/autopilot-cockpit-backend/`, `/apigateway/autopilot-cockpit-itsm/`, `/apigateway/recommendations-be/` | autopilot cockpit | Recover recommendations and configuration reads |
| Global search | `/apigateway/global-search` | `globalsearchFe.js` | Recover query, filtering and result pagination |
| DEX and benchmark | Existing GraphQL gateways in `operations.json` | respective product bundles | DEX score/application/campaign configuration mutations recovered but not yet typed or live validated; benchmark analytical operations still require an operation-level audit |
| Visual editor and value provider | Existing GraphQL gateways, `/apigateway/visual-editor/api/collections/bco`, `/apigateway/visual-editor/api/columns` | `nxmonitorconfig.js` | Recover schema/filter/value operations |

The [machine-readable discovery inventory](web-api-discovery.json) records 56 HTTP contracts (23 now implemented), 205 GraphQL document variants and the original 95 API path literals from the inspected UI files. The original nine discovery reads plus three connector/template listing reads returned 200; new implemented lifecycles have separate curl and SDK validation. These are discovery counts, not counts of implemented endpoints. GraphQL variants are deduplicated by operation kind, name and document hash; different selections can describe the same logical operation. Literal paths can be incomplete base paths. Telemetry and observability submissions, static assets, and third-party services are not automatically SDK resources.

The inventory distinguishes observed methods and payload requirements from successful replays. Asset writes send a text/plain data URL with a filename header. Checklist and rating deletion use revision parameters and explicit request bodies. Dashboard queries contain Apollo client directives and fragment references that must be resolved before wire replay. DEX score mutations update aggregate tenant configuration; no independent entity-create/delete operation was found in that configuration bundle. None of these observations establishes complete product coverage.

Expansion order: Investigations/custom fields, alerts/dashboards/campaigns, then connectors/library/access management. Full available LCRUD is the completion criterion for each management resource. Each increment should include observed request contracts, successful replay evidence, models, JSON fixtures, tests, examples and root-client wiring before being described as covered.

## Complete examples and contract corrections

There are 111 exported web resource methods with runnable examples, plus examples for all 22 public resource methods in `crud.go`. The [web example index](../examples/nexthink/web_api/README.md) lists required inputs and JSON samples. A source-level test requires an example that calls every exported resource method. Compilation is not a claim that every method passed live testing.

The seven additional management resources each have request and response JSON fixtures, HTTP/transport/malformed-response tests, validation tests, and GraphQL partial-data tests where relevant. Fixtures use synthetic IDs and data; browser tokens, raw tenant responses, and signed URLs stay outside the repository. A populated Campaign read and Application read were compared directly with curl. All disposable lifecycle objects were deleted.

`ProductShell.ValidateClaims` now has a typed request for the four UI claim-check methods and a nested boolean response. `DeviceConfiguration.SetProfiles` now uses `SaveProfilesRequest`; it does not replace a profile collection. These typed signatures replace the previous raw JSON signatures. Dynamic menu IDs come from `GetMenu` (`bus-menu` succeeded); arbitrary display names can return 404.

Follow-up work remains visible: the legacy settings route returned 403, Collector configuration writes and telemetry have not been replayed, and the newly discovered API areas above are not all implemented. Full LCRUD for an entity does not imply every auxiliary operation or subtype has been validated.

## Content management continuation from PR #50

The new branch adds Assets, Checklists, Dashboards, Ratings and Investigations through the same `nexthink.NewClient` entry point. All 27 new examples passed in the lab. Curl and SDK lifecycles verified create/list/read/update/read/delete, and Investigation export/import created and removed a second copy. Every disposable object was deleted. Reads for Checklists, Ratings, Dashboards and Investigations matched curl data exactly; asset replacement bytes matched a download through the returned signed URL without forwarding a bearer token.

New wire-contract tests exposed a shared transport defect: `DeleteWithBody` did not enable Resty's DELETE payload option, so bodies were silently dropped. The request now explicitly enables it; the transport regression test and checklist/rating tests assert the transmitted JSON body.

Further discovery recovered universal connector CRUD, asynchronous tests, connector templates and credential references, Investigation query/metadata/link helpers, and more knowledge-upload details. Existing connectors were not modified. These contracts are recorded with their implementation and validation status. Knowledge-file transformation and nested connector schemas still need validation. One referenced Investigation UI chunk returned 404; this and unimplemented nested/auxiliary operations prevent an exhaustive-coverage claim.
