# Applications examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/applications/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/applications/mocks/Update_input.json) and `NEXTHINK_CONTENT_ID` |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/applications/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/applications/List
```

List returns a page; use `ListOptions` to paginate. Update and Delete require the current revision. Delete returns that revision as a JSON number.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.
