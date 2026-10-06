# GlobalSearch.GetLegacyAuthToken

See the [resource guide](../README.md) for SDK browser authentication.

This is the legacy `POST /PortalServlet` protocol, validated against UI source and synthetic wire/error tests. A compatible legacy portal session was unavailable in the cloud lab. Supply an authorized portal cookie and/or `x-auth-token` using `NEXTHINK_PORTAL_COOKIE` and `NEXTHINK_PORTAL_X_AUTH_TOKEN`. Keep these credentials out of source files and shell history.

The standard SDK configuration still constructs the web client. This call uses the explicit portal session and suppresses the normal bearer token. The SDK does not store response cookies or rotated tokens; updated credentials are available in response headers for your session owner to handle explicitly. The route and form query are fixed.

The example prints only whether a token was received. In Go, the credential is available as `result.Result.Token`; avoid logging or serializing it.

```sh
go run ./examples/nexthink/web_api/global_search/GetLegacyAuthToken
```
