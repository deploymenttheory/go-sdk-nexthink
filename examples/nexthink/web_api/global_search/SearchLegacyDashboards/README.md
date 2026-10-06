# GlobalSearch.SearchLegacyDashboards

See the [resource guide](../README.md) for SDK browser authentication.

This is the legacy `POST /PortalServlet` protocol, validated against UI source and synthetic wire/error tests. A compatible legacy portal session was unavailable in the cloud lab. Supply an authorized portal cookie and/or `x-auth-token` using `NEXTHINK_PORTAL_COOKIE` and `NEXTHINK_PORTAL_X_AUTH_TOKEN`. Keep these credentials out of source files and shell history.

The standard SDK configuration still constructs the web client. This call uses the explicit portal session and suppresses the normal bearer token. The SDK does not store response cookies or rotated tokens; updated credentials are available in response headers for your session owner to handle explicitly. The route and form query are fixed.

Limits apply separately to personal, published and role-based results. The UI requests one extra item per category for its local “has more” check; the SDK sends your limits exactly. Nonzero `resultStatus.code` returns both the envelope and an error, including when HTTP status is 200.

```sh
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/global_search/SearchLegacyDashboards/request.example.json"
go run ./examples/nexthink/web_api/global_search/SearchLegacyDashboards
```
