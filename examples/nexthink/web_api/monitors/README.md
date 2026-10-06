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

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [Export](Export/main.go)
- [ExportLibrary](ExportLibrary/main.go)
- [Import](Import/main.go) — [request](Import/request.example.json)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

`Export` returns base64 `content`; call `DecodeContent()` to obtain the JSON text accepted by `ImportRequest.Content`. `ExportLibrary` returns separate content and metadata files. Importing a monitor creates configuration that can evaluate and notify according to its definition; inspect its query, threshold and recipients first.

### Additional management operations

These examples use browser authentication. Request-based examples load `NEXTHINK_REQUEST_FILE`; copy the corresponding `request.example.json` and replace synthetic values. Mutation examples change configuration and should use explicitly selected targets.

- [SetActivity](SetActivity/main.go): `ToggleActivity`.
- [UpdateBuiltIn](UpdateBuiltIn/main.go): `UpdateBuiltInMonitor`.
- [GetLicense](GetLicense/main.go): `License`.
- [ListTags](ListTags/main.go): `Tags`.
- [GetMetadata](GetMetadata/main.go): `MonitorMetaData`.
- [AnalyzeQuery](AnalyzeQuery/main.go): `NqlQueryAnalysis`.
- [GetImpactQuery](GetImpactQuery/main.go): `ImpactQuery`.
- [ListFilterFields](ListFilterFields/main.go): `FilterFields`.

`ListFilterFields` uses the shared visual-editor GraphQL gateway. `GetImpactQuery` needs a meaningful trigger condition; its example supplies a synthetic threshold. `UpdateBuiltIn` is separate from updating custom monitors. The built-in update operation was contract-tested without changing an existing tenant monitor.
