# DeleteVPNEgress

Calls `WebAPI.DeviceClassification.DeleteVPNEgress`. Use `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK local-account browser provider. See the root README for authentication.

Changes tenant-wide configuration. Review the complete requested change before running. Live mutation acceptance was not performed because these resources are shared settings, rather than disposable objects.

```sh
go run ./examples/nexthink/web_api/device_classification/DeleteVPNEgress
```

Source: `product-configuration-ui/1.83.4`. Positive unit fixtures validate source-derived wire contracts; they do not imply that shared lab settings were changed.
