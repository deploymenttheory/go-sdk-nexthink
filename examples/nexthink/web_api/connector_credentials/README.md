# ConnectorCredentials examples

Use the shared [authentication setup](../README.md). From the repository root run `go run ./examples/nexthink/web_api/connector_credentials/List`.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [List](List/main.go) | None | — |
| [ListIDs](ListIDs/main.go) | None | — |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [Create](Create/main.go) | `NEXTHINK_CONTENT_ID` and `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Update](Update/main.go) | `NEXTHINK_CONTENT_ID` and `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_CONTENT_ID` | — |

Use IDs shaped `conn_cr-<number>`. `ListIDs` includes allocated disabled IDs; the UI chooses one greater than the maximum suffix. Allocation is not atomic, and Create/Update are the same POST upsert: neither method guarantees create-only or update-only behavior. Coordinate IDs when multiple writers use the lab.

Create/Update require an explicit request file. The no-auth sample points at `.invalid` and contains no real secrets. Set `config.runTime` explicitly (the UI uses `23:30`). Secret entries are optional, write-only values; omit `secret` to preserve existing secrets during update. Do not publish request files containing secrets.

Delete follows the observed UI clear-and-disable behavior through POST; it does not physically remove the record. The ID stays in `ListIDs`, Get returns a disabled configuration, and the enabled-only List excludes it. All six examples passed on isolated fixtures, which were cleared/disabled afterward. Secret writes have synthetic request tests and were not exercised against a third-party system.
