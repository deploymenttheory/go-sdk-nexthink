# GetDynamicMenu

Configure the Web API client as described in the [web examples guide](../../README.md). Set `NEXTHINK_MENU` to an `apiUrl` value from `ProductShell.GetMenu().Result.Menus`, such as `investigations-menu`, then run:

```sh
go run ./examples/nexthink/web_api/product_shell/GetDynamicMenu
```

The argument is a dynamic menu key. A visible title or a static section name such as `administration` is not interchangeable and may return HTTP 404. Discover the key from the current tenant instead of treating an arbitrary example key as an API failure.
