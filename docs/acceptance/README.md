# SDK acceptance testing

The [2026-10-06 method matrix](2026-10-06.json) accounts for **586 exported resource methods**: **358 passed**, **228 blocked**, and **0 unresolved SDK failures** after corrections. This is not full live acceptance. Counts include query/export convenience methods and routing helpers; they are not counts of distinct HTTP endpoints.

| API family | Passed | Blocked | Total |
| --- | ---: | ---: | ---: |
| Public API | 21 | 7 | 28 |
| Web API | 337 | 221 | 558 |
| Total | 358 | 228 | 586 |

The run began from merged PR56, commit `611dfaccf2d525ffeb304b748b41eeaa978243fe`, and retested the corrections on `test/systematic-sdk-acceptance`. After PR57 merged as `3bc2f0394614e6c2fefbfd9b9c7eca6fbbf30d5e`, a curl-led follow-up on `fix/curl-acceptance-followups` retested 53 methods and established 24 additional positive passes. Results describe the lab on this date, not a vendor compatibility guarantee.

## What the statuses mean

- **pass**: a positive live operation, with curl preflight and SDK/example evidence, or an explicitly identified convenience/routing helper exercised successfully. Export helpers also checked completion and downloaded output. Lists may legitimately be empty.
- **blocked**: insufficient positive evidence because of missing telemetry/fixtures, permission or feature restrictions, backend failures, incomplete curl comparison, or an operation deliberately not exercised. A successful transport response containing null metrics, GraphQL errors, or a failed business outcome is not positive acceptance.
- **failed**: an unresolved SDK/example defect or unexplained curl/SDK disagreement. Confirmed defects were corrected and retested before this report.

Each row states its reason. An empty evidence array means that method was not exercised live; unit coverage does not replace acceptance. Some blocked rows have successful transport/schema comparisons but lack populated data. Historical evidence from earlier PRs was used to prepare requests, not counted as a fresh pass.

## Curl-led follow-up

Several prior errors came from the acceptance requests rather than SDK implementation: the monitor field query used the wrong GraphQL gateway, requests retained placeholder IDs, dynamic menus used an arbitrary section name, and role permissions used an IAM profile identifier instead of the built-in role identifier requested by the UI. These were corrected and compared against fresh curl results. Three analytics requests now succeed but remain classified as telemetry gaps because their metrics are empty.

Corrected committed examples and JSON request fixtures now show the required binary query parameter, standalone diagnostic definition, leaf DEX metric ID and desktop application configuration. Campaign documentation/tests distinguish real publication/retirement transitions from invalid same-state transitions and the independent legacy V6 representation. The enrichment example now shows the full custom-field URI.

New positive checks include connector LCRUD, knowledge multipart upload, sharing reads, campaign status transitions, monitor fields, public enrichment and web custom-field updates/CSV imports. Both curl and SDK writes to the disposable manual field were verified by distinct values in subsequent NQL reads; asynchronous HTTP 200/202 acknowledgements alone were not counted.

The remaining 228 rows now carry explicit `blocker` categories:

| Category | Count |
| --- | ---: |
| Not tested | 90 |
| Requires a suitable fixture | 50 |
| Requires populated telemetry | 64 |
| Permission or route/feature availability | 16 |
| Requires retention consent | 6 |
| Reproduced server error | 2 |

The two server-error rows are software-metering usage distribution (redacted subgraph error) and support Ethernet drilldown (HTTP 500). Other exact UI routes still return 401/403/404 through both curl and SDK, including Autopilot/Forge and support disk/drive drilldowns. Those results do not establish a client-side defect or justify inventing a replacement endpoint.

All six shared follow-up content fixtures were removed after dependent reads/writes completed; integration fixtures were separately removed or cleared. No campaign deliveries, remote-action executions or connector test executions occurred.

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
- IAM changes, sharing grants, tenant/device settings, branding and diagnostic ingestion were not exercised against shared configuration.
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
