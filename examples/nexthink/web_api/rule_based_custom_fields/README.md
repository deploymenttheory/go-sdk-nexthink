# RuleBasedCustomFields examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/rule_based_custom_fields/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/rule_based_custom_fields/mocks/Update_input.json) and `NEXTHINK_CONTENT_ID` |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/rule_based_custom_fields/mocks/Delete_input.json) and `NEXTHINK_CONTENT_ID` |

```sh
go run ./examples/nexthink/web_api/rule_based_custom_fields/List
```

List returns only RULE_BASED rows from the shared listing. Delete uses POST and a revision-bearing body; its acknowledgment is plain text. Delete uses Device/User/Binary/Package, while create/update use object URIs.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

Export returns the rule definition and revision in its importable representation, without changing field values.

| Method | Inputs beyond authentication |
| --- | --- |
| [Export](Export/main.go) | `NEXTHINK_CONTENT_ID` |
