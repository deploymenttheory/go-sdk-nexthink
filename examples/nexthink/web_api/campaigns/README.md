# Campaigns examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/campaigns/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/campaigns/mocks/Update_input.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/campaigns/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/campaigns/List
```

Create saves a draft. Update preserves the supplied status. Delete returns the deleted content ID. List supports pagination through ListOptions.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.
