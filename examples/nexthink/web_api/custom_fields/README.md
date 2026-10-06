# CustomFields examples

Configure `NEXTHINK_API=web`, instance, region, and browser authentication as described in the [example index](../README.md). Run examples from the repository root.

| Example | Required inputs |
| --- | --- |
| [List](List/main.go) | None (first page where paginated) |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` matching [Create_input.json](../../../../nexthink/web_api/custom_fields/mocks/Create_input.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` and `NEXTHINK_CUSTOM_FIELD_TYPE` (`MANUAL` or `COMPUTED`) |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` matching [Update_input.json](../../../../nexthink/web_api/custom_fields/mocks/Update_input.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` matching [Delete_input.json](../../../../nexthink/web_api/custom_fields/mocks/Delete_input.json) |

```sh
go run ./examples/nexthink/web_api/custom_fields/List
```

These methods manage MANUAL and COMPUTED fields. Rule-based fields use the adjacent RuleBasedCustomFields resource. Create omits docUid in its response; find it through List. Delete uses Device/User/Binary/Package, whereas create/update use object URIs.

Replace synthetic IDs and revisions with the object you intend to manage. Create and Update write the supplied object; Delete removes it. GraphQL examples print partial data before reporting errors so returned identifiers remain available.

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

Export takes MANUAL or COMPUTED plus a content UUID. Import takes contentFile containing serialized definition JSON (not base64); metadata is null in the UI request. Import returns a generated docUid. Use distinct names/NQL IDs when importing a copy. GetValidationPatterns returns the server's naming rules.

| Method | Inputs beyond authentication |
| --- | --- |
| [GetValidationPatterns](GetValidationPatterns/main.go) | None |
| [Export](Export/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_FIELD_TYPE` |
| [Import](Import/main.go) | `NEXTHINK_REQUEST_FILE` ([sample](Import/request.example.json)) |
