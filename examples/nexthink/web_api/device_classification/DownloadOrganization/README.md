# DownloadOrganization

Calls `WebAPI.DeviceClassification.DownloadOrganization`. Use `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK local-account browser provider. See the root README for authentication.

Read-only. An absent classification ruleset returns `404 NO_RULESET_FOUND`; this is not an empty successful download.

Inputs:

- `NEXTHINK_OUTPUT_FILE`: destination path for exact downloaded CSV bytes (written with mode 0600).

```sh
go run ./examples/nexthink/web_api/device_classification/DownloadOrganization
```

Source: `product-configuration-ui/1.83.4`. Positive unit fixtures validate source-derived wire contracts; they do not imply that shared lab settings were changed.
