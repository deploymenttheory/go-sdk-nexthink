# GetLegacyAsset

Configure the [web API client](../../README.md), then set `NEXTHINK_PORTAL_COOKIE` and/or `NEXTHINK_PORTAL_X_AUTH_TOKEN` from a legitimate legacy portal session. `auth.PortalSession` credentials apply only to the exact legacy routes. Browser bearer credentials are not substituted. The configured WebAPI root still requires browser credentials or a token provider; that provider is never consulted for legacy requests. Cookies from responses are not stored automatically.

Set `NEXTHINK_ASSET_NAME` to `menu-logo` (default), `login-logo`, or `login-bg`. Reads the current legacy asset with `defaultImage:false`; prints its id, version, MIME type and base64 image.

```sh
go run ./examples/nexthink/web_api/appearance/GetLegacyAsset
```

Source and synthetic JSON/unit tests validate this legacy contract. No compatible legacy session was available for live testing; no tenant branding was changed. Keep session credentials outside source control and logs.
