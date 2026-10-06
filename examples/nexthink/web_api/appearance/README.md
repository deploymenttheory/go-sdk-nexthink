# Appearance assets

Configure `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and browser authentication as described in the [web API examples](../README.md). Run commands from the repository root.

Get returns the original image bytes and response media type. The asset name is `menu-logo`, `login-logo` or `login-bg`; the example defaults to menu-logo. Update sends raw image bytes with an image media type. The UI resets an asset by uploading its bundled default image through Update; there is no independent reset/delete endpoint. Menu-logo download bytes matched curl. Tenant branding was not modified; Update has source and unit coverage.

| Method | Guide |
| --- | --- |
| `Get` | [Inputs and example](Get/README.md) |
| `Update` | [Inputs and example](Update/README.md) |

Legacy portal assets use separate explicit per-call `auth.PortalSession` credentials. The configured WebAPI client still needs normal browser credentials or a provider, which is not consulted for these calls. These contracts have source/unit coverage only; no compatible legacy portal session was available.

| Method | Guide |
| --- | --- |
| `GetLegacyAsset` | [Inputs and example](GetLegacyAsset/README.md) |
| `SaveLegacyAsset` | [Inputs, write effects and example](SaveLegacyAsset/README.md) |
