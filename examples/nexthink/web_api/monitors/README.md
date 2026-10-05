# Monitors examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/monitors/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/monitors/mocks/Update_input.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/monitors/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/monitors/List
```

Update requires the content/doc UUID, revision, and the separate monitor UUID from Get. The observed Update/Delete acknowledgments are null. Read the monitor again to verify an update.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.
