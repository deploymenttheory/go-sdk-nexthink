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

All eight examples passed with a disposable no-auth `.invalid` credential and a 2099 schedule. Template/detail/list responses matched curl. Templates are server-provided; no template create/update/delete contract was observed. Async tests and legacy connectors remain pending.
