# Replace

Calls `WebAPI.UserClassification.Replace`. Use `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK local-account browser provider. See the root README for authentication.

Changes tenant-wide configuration. Review the complete requested change before running. Live mutation acceptance was not performed because these resources are shared settings, rather than disposable objects.

Each `nqlId` includes a leading `#` (for example `#sdk_department`, queried as `user.organization.#sdk_department`). The UI implements add/edit/delete by replacing `customFields` in one PUT. Preserve fields you intend to keep. An explicit empty array clears the complete list.

Inputs:

- `NEXTHINK_REQUEST_FILE`: reviewed JSON configuration; copy `request.example.json`.

```sh
go run ./examples/nexthink/web_api/user_classification/Replace
```

Source: `product-configuration-ui/1.83.4`. Positive unit fixtures validate source-derived wire contracts; they do not imply that shared lab settings were changed.
