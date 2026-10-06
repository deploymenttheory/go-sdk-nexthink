# Import

Export takes MANUAL or COMPUTED plus a content UUID. Import takes contentFile containing serialized definition JSON (not base64); metadata is null in the UI request. Import returns a generated docUid. Use distinct names/NQL IDs when importing a copy. GetValidationPatterns returns the server's naming rules.

Use the authentication and inputs in the [resource guide](../README.md).

Copy [request.example.json](request.example.json), replace fixture identifiers and fields for your target, then set `NEXTHINK_REQUEST_FILE` to the edited file.

From the repository root:

```sh
go run ./examples/nexthink/web_api/custom_fields/Import
```
