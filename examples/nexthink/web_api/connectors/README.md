# Connectors examples

Use the shared [authentication setup](../README.md). From the repository root run `go run ./examples/nexthink/web_api/connectors/List`.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [List](List/main.go) | None | — |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [ListTemplates](ListTemplates/main.go) | None | — |
| [GetTemplate](GetTemplate/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [ListManualCustomFields](ListManualCustomFields/main.go) | `NEXTHINK_DATA_MODEL_OBJECT` | — |

Create requires a caller-generated UUID in `content_id`; replace `fixture-template` and credential/mapping values with values from `ListTemplates`, `GetTemplate` and your intended lab configuration. Schedules use Quartz cron, not five-field Unix cron. The sample uses 2099. **The lab server forced `enabled:true` on Create even when false was supplied.** Update honored false; do not rely on a disabled Create preventing execution.

Update sends the complete definition and uses its `content_id` in the path. Delete accepts a universal connector ID and returns HTTP metadata (204 observed). List also contains legacy connectors, which these v1 methods cannot manage. GetTemplate uses `NEXTHINK_CONTENT_ID` for the template ID; field lookup uses `NEXTHINK_DATA_MODEL_OBJECT`, for example `device/mobile_device`.

All eight examples passed with a disposable no-auth `.invalid` credential and a 2099 schedule. Template/detail/list responses matched curl. Templates are server-provided; no template create/update/delete contract was observed. Async test methods are listed below; shared legacy configurations use the separate `LegacyConnectors` resource. Tests contact the referenced destination and require a suitable test fixture.

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [StartTest](StartTest/main.go) — [request](StartTest/request.example.json)
- [GetTest](GetTest/main.go)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

`StartTest` starts an asynchronous test against the referenced credential. Poll `GetTest` with `testExecutionId`. `COMPLETED` describes execution completion; inspect each partition's error/status code before treating the connector as working.
