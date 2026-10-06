# GetValidationPatterns

Export takes MANUAL or COMPUTED plus a content UUID. Import takes contentFile containing serialized definition JSON (not base64); metadata is null in the UI request. Import returns a generated docUid. Use distinct names/NQL IDs when importing a copy. GetValidationPatterns returns the server's naming rules.

Use the authentication and inputs in the [resource guide](../README.md).

From the repository root:

```sh
go run ./examples/nexthink/web_api/custom_fields/GetValidationPatterns
```
