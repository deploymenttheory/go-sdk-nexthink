# UpdateInstance

Calls `WebAPI.ProductConfiguration.UpdateInstance`. Use `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK local-account browser provider. See the root README for authentication.

Changes tenant-wide configuration. Review the complete requested change before running. Live mutation acceptance was not performed because these resources are shared settings, rather than disposable objects.

The request key must equal the path key. Nexthink may reject an unchanged value with `Update requests must actually update something`.

Inputs:

- `NEXTHINK_CONFIGURATION_KEY`: exact configuration key, e.g. `platform.assist.web-search`.
- `NEXTHINK_REQUEST_FILE`: reviewed JSON configuration; copy `request.example.json`.

```sh
go run ./examples/nexthink/web_api/product_configuration/UpdateInstance
```

Source: `product-configuration-ui/1.83.4`. Positive unit fixtures validate source-derived wire contracts; they do not imply that shared lab settings were changed.
