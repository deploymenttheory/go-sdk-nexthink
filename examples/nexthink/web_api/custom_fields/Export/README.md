# Export

Export takes MANUAL or COMPUTED plus a content UUID. Import takes contentFile containing serialized definition JSON (not base64); metadata is null in the UI request. Import returns a generated docUid. Use distinct names/NQL IDs when importing a copy. GetValidationPatterns returns the server's naming rules.

Use the authentication and inputs in the [resource guide](../README.md).

Required context: `NEXTHINK_CONTENT_ID`, `NEXTHINK_FIELD_TYPE` (see the resource guide for meanings).

From the repository root:

```sh
go run ./examples/nexthink/web_api/custom_fields/Export
```
