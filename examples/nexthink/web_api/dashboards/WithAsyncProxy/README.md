# Dashboards.WithAsyncProxy

Set `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome`, `NEXTHINK_INSTANCE`, and `NEXTHINK_REGION` with an authenticated Chrome portal session.

The independent service copy sends existing dashboard GraphQL operations to the observed long-running proxy and adds `x-nxt-waas-allow-long-running: true`. REST methods keep their original routes. Client redirect restrictions remain in effect.

```sh
go run ./examples/nexthink/web_api/dashboards/WithAsyncProxy
```
