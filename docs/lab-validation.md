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

The inventory already contained three Windows devices. Validation did not delete or modify them. Saved queries `#sdk_lab_devices` and `#sdk_lab_parameterized` were created as bounded fixtures. A separate temporary saved-query fixture was created, updated, read back, and deleted through the experimental SDK.

## Examples and NQL templates

All 19 read examples passed on the final code, including both `_build_client` examples that Go's `./...` pattern skips. The committed [example harness](../scripts/validate-examples.sh) ran 18 successfully in one pass; `StartNQLExport` encountered the OAuth timeout noted above and passed on isolated retry. Examples now exit unsuccessfully when execution or export fails. Write examples require explicit JSON fixture input through `NEXTHINK_REQUEST_FILE`.

All 18 NQL templates were submitted to the lab's NQL editor validator. Initial validation exposed five syntax failures; review also found a memory template calculating disk usage. The corrected templates produced empty diagnostics. This validates the generated queries against the lab parser; it does not establish every query's meaning against a representative telemetry dataset.

Query-builder operator order is now preserved, including filters within event relations before aggregation. String helpers escape literals. Templates use correct enum expressions, filter event fields before computation, and calculate memory usage from used and installed memory.

## Undocumented APIs

The [catalog](../nexthink/experimental/operations.json) contains 42 operations with endpoint, method, discovery evidence, and observed authentication results. The opt-in experimental client preserves raw JSON where schemas are not established.

- 26 initial read/probe operations returned HTTP 200 through the SDK, covering Collector management, product shell, licensing, content administration, saved-query reads, and eleven GraphQL endpoints.
- Saved-query create/update/read/delete returned 201/200/200/204. Public execution with named parameters also succeeded.
- Follow-up curl and SDK checks returned 200 for device profiles, NQL validation, and NQL highlighting. The alternate device-settings route returned 403 under the same browser identity.
- GraphQL checks used `__typename`; they establish reachability, not support for all queries or mutations. Management introspection was disabled or unavailable. The SDK reports GraphQL errors even when HTTP status is 200 and preserves partial data.
- Browser and OAuth client identities have different access. For example, shell user/configuration reads worked with the browser identity and returned 401 with client credentials. A Collector configuration request returned 503 with client credentials, which is inconclusive.
- Collector/device configuration writes, shell telemetry, claims validation, dynamic menus, and several editor operations remain bundle-evidenced or unreplayed. Their catalog entries explicitly say so. No shared tenant settings were changed to test these routes.

Discovery inspected 64 first-party entry bundles and 77 manifest assets. Additional dynamic routes and payloads remain unresolved. AmplifyAI and Hypervisor public API contracts remain unverified. No API for retrieving a valid Collector Customer Key was established.

## Transport and regression checks

The revision fixes token refresh re-entering its own authentication middleware, coordinates concurrent refreshes, observes cancellation, retains numeric API error codes, and prevents authenticated requests from leaving the configured origin. Token requests use a separate HTTP client while retaining the configured proxy/TLS transport. Debug logging excludes request and response headers/bodies.

Validation completed:

- `go test -race ./... -timeout 180s`
- `go vet ./...`
- Configured `golangci-lint run --fix=false --timeout=5m`: zero issues; existing unmatched-exclusion configuration warnings remain
- Error parser fuzzing: 60 seconds, 388,718 executions
- Ragged V1 result fuzzing: 60 seconds, 414,637 executions
- Credential-value scan of tracked and proposed repository files: zero matches

Regression tests cover concurrent token refresh, cancellation, origin restrictions, numeric/wrapped errors, public service wire contracts, CSV conversion, malformed result rows, query ordering, experimental path parameters, and GraphQL partial-data errors. Synthetic fixtures are committed; raw tenant responses, credentials, signed URLs, browser tokens, and Collector keys remain outside the repository.

Hosted CI exposed existing Super-Linter configuration failures with multiple packages and a module root containing no Go files. That wrapper is replaced with explicit formatting and vet checks; the dedicated golangci-lint workflow checks the module and no longer forces a successful exit when findings occur. Dependency Review remains blocked because GitHub reports that Dependency graph is not enabled for this repository.

## Device lab and unfinished validation

The requested 32 GB macOS restore failed at 84%. The approved 64 GB sparse-disk retry also failed during restore. A dedicated copy-on-write clone of the existing stopped macOS 27 golden image was then created with a regenerated MAC address and serial, a 64 GB virtual disk, and 8 GiB RAM. The source image was retained.

Collector 26.8.2.22 installed and its services started. Remote actions were configured to accept trusted or Nexthink-signed scripts. Enrollment remains blocked by the Collector's local error **“Failed to decode the Customer Key”**, before it can set the HTTP authentication header. The supplied payload is structurally valid Base64; preserving the original file and normalizing only line wrapping both produced the same error. The original complete lab-issued key or Nexthink support diagnosis is still required.

Full Disk Access for `nxtsvc` in this VM was explicitly approved. Applying it remains unverified because the VM runner's experimental VNC connection crashed or displayed a black framebuffer after login. The VM is not yet a verified enrolled fixture.

Remaining device-dependent work: verify enrollment and telemetry; execute a dedicated signed diagnostic remote action; exercise dedicated workflow and campaign fixtures; test enrichment with read-back; and perform deletion only on an explicitly designated disposable device. Spark requires a dedicated Teams recipient. These are outstanding tests, not passing results.

## References

- [Nexthink API documentation](https://docs.nexthink.com/api)
- [Data Management device deletions](https://docs.nexthink.com/api/data-management/schedule-device-deletions)
- [Spark handoff](https://docs.nexthink.com/api/spark/handoff-api)
- [NQL investigation examples](https://docs.nexthink.com/platform/user-guide/investigations/investigations-nql-examples)
- [Migration notes](migration-lab-validation.md)
