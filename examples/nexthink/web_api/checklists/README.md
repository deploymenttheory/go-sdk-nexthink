# Checklists examples

Use the shared [authentication setup](../README.md). From the repository root, run `go run ./examples/nexthink/web_api/checklists/List`. Write examples require `NEXTHINK_REQUEST_FILE` containing an explicit request; replace the synthetic values with your intended lab target.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Delete/request.example.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [List](List/main.go) | None | — |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |

Create and Update accept checklist definitions. Update/Delete require the latest `revisionNumber` from Get; the revision is sent as a query parameter. Delete preserves the UI's `{"a":"fix"}` JSON body and accepts an empty successful response. Create optionally accepts a `LibraryUUID` through `CreateOptions` in Go.

Field-data entries require `id` and `type`, plus `subType` when supplied by the UI. The examples include a property with a custom label and documentation. Action definitions are retained as JSON because they vary by action type; saving a checklist does not execute them. Use Get to preserve category IDs when editing an existing checklist.

Curl and all five examples passed, with populated field data, update read-back and deletion. The auxiliary methods below expose grouped-field metadata, library reads, and checklist import/export.

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [Export](Export/main.go)
- [Import](Import/main.go) — [request](Import/request.example.json)
- [ListGroupedFields](ListGroupedFields/main.go)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

## Additional metadata and query operations

Set `NEXTHINK_API=web` and `NEXTHINK_WEB_AUTH=chrome`, or supply a browser access token as described in the main examples README. These operations do not modify saved configuration.

| Example | Input |
| --- | --- |
| [GetFromLibrary](GetFromLibrary/main.go) | `NEXTHINK_CONTENT_ID` (library UUID) |

Run from the repository root: `go run ./examples/nexthink/web_api/checklists/GetFromLibrary`. Query operations read tenant data; use a restricted query and limit. Library reads retrieve definitions without installing content.
