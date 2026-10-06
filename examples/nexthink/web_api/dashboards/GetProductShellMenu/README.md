# Dashboards.GetProductShellMenu

Configure browser authentication and `NEXTHINK_API=web` using the root quick start. Local-account password authentication works without an interactive Chrome session.

Use the optional comma-separated product areas to retrieve the same dashboard entries as the UI:

```sh
export NEXTHINK_PRODUCT_AREAS=collaboration,collaboration-tools
go run ./examples/nexthink/web_api/dashboards/GetProductShellMenu
```

The lab returned Call quality, MS Teams rooms and Call quality – Selected device. Their menu URLs identify the product area and dashboard ID to pass to `Dashboards.Get`. The response preserves localized titles and other menu fields as raw JSON.

Leave `NEXTHINK_PRODUCT_AREAS` unset for the original unfiltered request. An empty unfiltered menu does not establish that the tenant has no built-in dashboards.
