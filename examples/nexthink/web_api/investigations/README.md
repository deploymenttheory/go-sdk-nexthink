# Investigations examples

Use the shared [authentication setup](../README.md). From the repository root, run `go run ./examples/nexthink/web_api/investigations/List`. Write examples require `NEXTHINK_REQUEST_FILE` containing an explicit request; replace the synthetic values with your intended lab target.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Delete/request.example.json) |
| [Export](Export/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [Import](Import/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Import/request.example.json) |
| [List](List/main.go) | None | — |
| [Update](Update/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |

Saved investigations are distinct from saved NQL API queries. Create/Update send `{name,description,nql}`. Reads return `uid` and nested `nql.query`; use `uid` for Get/Update/Delete/Export. Shared list rows call this identifier `contentId`.

The tested server ignores the submitted description and returns an empty string, matching the empty description sent by the UI. Name and query updates were validated. Saving a definition does not execute its query.

Export emits `{name,nqlQuery}` JSON that Import accepts. Change the name before importing alongside the original: duplicate names return error code 102. Import returns HTTP 201 and the saved representation. Delete returns HTTP metadata with an empty body.

All seven examples and separate curl calls passed, including export/import read-back and deletion of both copies.
