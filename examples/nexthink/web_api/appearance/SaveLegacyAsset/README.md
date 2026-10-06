# SaveLegacyAsset

Configure the [web API client](../../README.md), then set `NEXTHINK_PORTAL_COOKIE` and/or `NEXTHINK_PORTAL_X_AUTH_TOKEN` from a legitimate legacy portal session. `auth.PortalSession` credentials apply only to the exact legacy routes. Browser bearer credentials are not substituted. The configured WebAPI root still requires browser credentials or a token provider; that provider is never consulted for legacy requests. Cookies from responses are not stored automatically.

Set `NEXTHINK_REQUEST_FILE` to your JSON request. Start with `request.example.json`, replace the synthetic blob with real base64 image bytes (no data URL prefix), and use the current `id` and `version` from GetLegacyAsset. **This replaces tenant portal branding.** The SDK sends one save request; call GetLegacyAsset separately for refreshed metadata. Returned in-band errors are retained alongside the error.

```sh
go run ./examples/nexthink/web_api/appearance/SaveLegacyAsset
```

Source and synthetic JSON/unit tests validate this legacy contract. No compatible legacy session was available for live testing; no tenant branding was changed. Keep session credentials outside source control and logs.
