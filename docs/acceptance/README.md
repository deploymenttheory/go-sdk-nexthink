# SDK acceptance testing

The [2026-10-06 method matrix](2026-10-06.json) accounts for **692 exported resource methods**: **414 passed**, **278 blocked**, and **0 unresolved SDK failures** after corrections. This is not full live acceptance. Counts include query/export convenience methods and routing helpers; they are not counts of distinct HTTP endpoints.

| API family | Passed | Blocked | Total |
| --- | ---: | ---: | ---: |
| Public API | 21 | 7 | 28 |
| Web API | 393 | 271 | 664 |
| Total | 414 | 278 | 692 |

The run began from merged PR56, commit `611dfaccf2d525ffeb304b748b41eeaa978243fe`, and retested the corrections on `test/systematic-sdk-acceptance`. After PR57 merged as `3bc2f0394614e6c2fefbfd9b9c7eca6fbbf30d5e`, a curl-led follow-up on `fix/curl-acceptance-followups` retested 53 methods and established 24 additional positive passes. Results describe the lab on this date, not a vendor compatibility guarantee.

## What the statuses mean

- **pass**: a positive live operation, with curl preflight and SDK/example evidence, or an explicitly identified convenience/routing helper exercised successfully. Amplify singleton creation is an explicit exception to curl preflight: the SDK creation was independently verified with curl readback; a duplicate creation was not attempted. Export helpers also checked completion and downloaded output. Lists may legitimately be empty.
- **blocked**: insufficient positive evidence because of missing telemetry/fixtures, permission or feature restrictions, backend failures, incomplete curl comparison, or an operation deliberately not exercised. A successful transport response containing null metrics, GraphQL errors, or a failed business outcome is not positive acceptance.
- **failed**: an unresolved SDK/example defect or unexplained curl/SDK disagreement. Confirmed defects were corrected and retested before this report.

Each row states its reason. An empty evidence array means that method was not exercised live; unit coverage does not replace acceptance. Some blocked rows have successful transport/schema comparisons but lack populated data. Historical evidence from earlier PRs was used to prepare requests, not counted as a fresh pass.

## Curl-led follow-up

Several prior errors came from the acceptance requests rather than SDK implementation: the monitor field query used the wrong GraphQL gateway, requests retained placeholder IDs, dynamic menus used an arbitrary section name, and role permissions used an IAM profile identifier instead of the built-in role identifier requested by the UI. These were corrected and compared against fresh curl results. Three analytics requests now succeed but remain classified as telemetry gaps because their metrics are empty.

Corrected committed examples and JSON request fixtures now show the required binary query parameter, standalone diagnostic definition, leaf DEX metric ID and desktop application configuration. Campaign documentation/tests distinguish real publication/retirement transitions from invalid same-state transitions and the independent legacy V6 representation. The enrichment example now shows the full custom-field URI.

New positive checks include connector LCRUD, knowledge multipart upload, sharing reads, campaign status transitions, monitor fields, public enrichment and web custom-field updates/CSV imports. Both curl and SDK writes to the disposable manual field were verified by distinct values in subsequent NQL reads; asynchronous HTTP 200/202 acknowledgements alone were not counted.

After the curl follow-up and product-permission discovery pass, the remaining 278 rows carry explicit `blocker` categories:

| Category | Count |
| --- | ---: |
| Not tested | 104 |
| Requires a suitable fixture | 68 |
| Requires populated telemetry | 65 |
| Permission or route/feature availability | 33 |
| Requires retention consent | 6 |
| Reproduced server error | 2 |

The two server-error rows are software-metering usage distribution (redacted subgraph error) and support Ethernet drilldown (HTTP 500). Other exact UI routes still return 401/403/404 through both curl and SDK, including Autopilot/Forge and support disk/drive drilldowns. Those results do not establish a client-side defect or justify inventing a replacement endpoint.

All six shared follow-up content fixtures were removed after dependent reads/writes completed; integration fixtures were separately removed or cleared. No campaign deliveries, remote-action executions or connector test executions occurred.

## Product-permission discovery after PR59

The next pass started from merged PR59 (`4eb7eab`) on `feat/rbac-product-api-coverage`. The first pass added **84 methods: 45 live passes and 39 outstanding**. All have runnable examples and JSON-backed unit tests; implemented methods are counted even when the lab cannot validate a positive result.

| Addition | New methods | Live passes | Outstanding |
| --- | ---: | ---: | ---: |
| AITools | 28 | 20 | 8 |
| Amplify | 9 | 9 | 0 |
| AmplifyAI | 11 | 0 | 11 |
| Workspace | 12 | 11 | 1 |
| WorkspaceAgents | 11 | 2 | 9 |
| WorkspaceTasks | 6 | 0 | 6 |
| WorkspaceAssignments | 6 | 2 | 4 |
| ActionExecutions.GetDeviceActions | 1 | 1 | 0 |

AI tool and adoption-goal lifecycles passed curl preflight and SDK replay, with fresh list reads confirming fixture cleanup. Positive reads cover tool configuration, governance counts/trends, metadata and insights. Copilot writes/credential checks need an integration fixture; module singleton writes were not exercised, its read returns 404, and goal insights returns 403. These failures do not prove a missing RBAC grant.

The official Amplify extension supplied the device/user/package/search contracts, which passed complete JSON comparisons, plus usage-event ingestion. Its AI contracts have source and unit validation: two reads return 404 through both the API-host and portal routes, and nine mutation operations were not run. No AI action execution or external ticket resolution was attempted. At that stage the role permitted Amplify viewing/packages but not management; the continuation below enabled management with explicit user authorization and recovered its write contracts.

Workspace chat, conversation updates/deletion/sharing and PNG attachments passed curl and SDK checks. Follow-up readback verified clearing tags with an empty array and starred state with false. All disposable conversations and attachments were removed. Agent creation reports `user_agents_flag_disabled`; task listing reports `automations_feature_flag`. These are explicit feature restrictions despite the Workspace role permission. Assignment lists and unread counts succeed, but specific assignments and connected MCP resources/tools need fixtures. Development-only feature-deployment controls found in UI code are excluded from the customer API catalog and recorded in discovery scope.

The role's ten advertised virtualization dashboard definitions match through the existing `Dashboards.Get` product-area option. Eleven aggregate NQL scenarios also match through `DataExploration.Query`, excluding per-request execution duration: conversations, audit/custom-trend/data-export/inbound-connector/NQL logs, tickets, usage, collaboration sessions, mobile devices and VDI sessions. These checks establish schema access and response correctness, not populated telemetry for every product dashboard. Examples are linked from the [coverage reconciliation](../web-api-coverage.md).

The NQL API-log check found a response-model defect: a nullable collection display name became an empty string. `DMElementInfo.CollectionName` now uses `*string`; callers reading that field must handle nil. A JSON regression test and live replay verify preservation. New Workspace request tests cover explicit empty arrays/null clearing, chat artifacts/feedback, and stream errors; response codecs retain unknown fields and null/absence distinctions.

## Administration and collaboration continuation

The continuation adds **22 methods: five live passes and 17 blocked**. Together, PR60 adds **106 methods across ten new resource packages and existing resources**. Six existing IAM/sharing methods also moved from blocked to passed.

| Addition | New methods | Live passes | Outstanding |
| --- | ---: | ---: | ---: |
| DeviceClassification | 15 | 1 | 14 |
| ProductConfiguration | 3 | 1 | 2 |
| UserClassification | 2 | 1 | 1 |
| Amplify configuration writes | 2 | 2 | 0 |

Classification metadata/download requests explicitly return `NO_RULESET_FOUND` when no CSV has been installed. Shared classification and feature configuration mutations were not exercised merely to populate acceptance evidence. All new methods have resource-local examples and JSON-backed tests, including multipart headers, Unicode descriptions, unknown response fields and explicit empty values.

The user authorized enabling Manage Amplify and leaving it enabled. The newly visible first-party administration frontend established configuration POST and revisioned PUT. SDK creation was independently checked with curl GET; curl and SDK updates verified application add/edit/reorder/removal. Test applications were removed and usage reporting is disabled. An **empty configuration document remains**, because no document-delete operation was observed. Application deletion replaces the ordered application list; it is not a separate DELETE endpoint.

Role create/update/delete and content grant/revoke, plus `ContentSharing.SetProfiles`, passed independent curl and SDK checks with readback. The two unassigned test roles and private dashboard were deleted and their absence verified. No account was assigned either role. JSON tests now reflect direct role-create/update responses and empty successful delete/grant/revoke bodies.

The recovered Collaboration Experience frontend uses 12 GraphQL operations already represented by `DataExploration`. Three built-in dashboard definitions (13 tabs, 223 widgets, 177 NQL widgets) matched curl and SDK, using the new optional product-area menu filter. Call-insights requests now succeed for Teams and Zoom using the Collector UID, but report no call telemetry. That row remains blocked for telemetry; the previous permission/availability failure was not reproduced.

## Corrections found

- Public remote actions now preserve `targeting.sparkEnabled` and `targetingEntity.vdiSessionTargeting`.
- Content-administration listings now retain workflow NQL IDs, timestamps, trigger methods, status and versions.
- Library dependencies preserve an omitted optional filename instead of emitting an empty filename.
- Software-metering Create/Update validate the UI's name contract: ASCII letters, digits and spaces, up to 255 characters. Invalid names previously reached a backend that returned an opaque subgraph error.
- Support examples explain that device routes need `device.collector.uid`; using `device.uid` can silently return empty data.
- The public example harness supplies explicit query/download prerequisites and isolates each example's outputs. The headless-password example is run separately with its own prerequisites.

JSON regression fixtures are synthetic; tenant responses and credentials are not committed.

## Remaining acceptance gaps

The matrix is the complete checklist. Principal gaps are:

- Analytics requiring populated application, VDI, call, alert, execution or historical telemetry. A device being enrolled does not populate every feature.
- Public action/workflow execution, campaign delivery, Spark handoff and device deletion. These require purpose-built enabled execution targets, writable fields or dedicated recipients/deletion targets.
- Account/identity changes, broader resource-specific grants, tenant/device settings, branding and diagnostic ingestion remain unverified. Disposable unassigned-role lifecycles and dashboard sharing grants/revocations passed in the continuation.
- Snapshot creation requires explicit extended-retention consent. No consent flag was enabled by this run.
- Legacy portal operations require a separate cookie/x-auth-token session; modern bearer authentication does not validate them.
- Connector test executions and real third-party delivery still require suitable isolated destinations/credentials. Connector CRUD was verified using a nonexecuting future schedule and immediate disabling; no outbound test was started.
- The initial knowledge deletion race was resolved in the follow-up by observing completed ingestion before deletion. Both fresh curl and SDK lifecycle deletions then succeeded without retries.

Disposable content was cleaned up and verified through fresh reads. Four test credential slots remain allocated because the API clears/disables slots rather than deleting them; their test credentials were cleared and disabled. No pending test content cleanup remained at completion.

## Verification

Full `go test -race ./... -timeout 180s`, `go vet ./...`, repository lint and the source-inventory validator passed. Fresh live headless login matched curl; controlled headless browser fixtures passed with the race detector. The public harness exercised both client constructors, the NQL tutorials and resource examples. Six late calls hit a transient DNS failure; all six targeted reruns passed and are recorded separately from the earlier successful service comparisons.

## Repeat and extend the pass

Inventory and report validation require no credentials:

```sh
go run ./scripts/acceptance > /tmp/nexthink-methods.json
go run ./scripts/acceptance -report docs/acceptance/2026-10-06.json
go test -race ./... -timeout 180s
go vet ./...
```

CI checks that the report contains exactly one valid result for every current exported `Service` method. Adding a method requires adding an honest result or categorized blocker; CI does not run live tenant writes or upgrade blocked rows to passed.

For public read/example acceptance, set the existing public client environment variables and these bounded fixtures, then run `bash scripts/validate-examples.sh`:

- `NEXTHINK_QUERY_ID`: a saved, bounded query.
- `NEXTHINK_REQUEST_FILE`: readable JSON containing that query's `queryId` and any required parameters.
- `NEXTHINK_EXPORT_ID` and `NEXTHINK_DOWNLOAD_URL`: a completed export and its current download URL.
- `NEXTHINK_REMOTE_ACTION_ID` and `NEXTHINK_WORKFLOW_ID`: existing read-only detail targets.

The harness reports its private output directory. It skips password authentication because that requires local-account credentials and the pinned browser runtime. Run its separate example or the opt-in test:

```sh
# Requires NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_USERNAME,
# NEXTHINK_PASSWORD, curl and the installed pinned browser runtime.
NEXTHINK_LIVE_PASSWORD_TEST=1 go test ./nexthink/auth/password \
  -run TestLivePasswordAuthentication -count=1 -timeout 3m
```

For web operations, use each resource's committed example and request JSON. Establish the same request with curl first; compare populated fields or bytes and verify write readback. Keep fixture creation and cleanup in a journal. Do not replay prior captured mutation IDs.

Raw curl responses, SDK output and fixture journals are private under `/private/tmp/nexthink-sdk-acceptance`. The local `current/evidence-index.json` maps each matrix `private-run:` reference to its captures. These files contain tenant data and are intentionally excluded from the repository. This report is a sanitized audit, not a self-contained replay of private fixtures.
