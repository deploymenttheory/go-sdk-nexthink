# Visual Editor examples

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

These read-only methods retrieve the inventory collections, default columns and filter metadata used by the monitor visual editor. NEXTHINK_URI is a data-model URI such as device/device.

| Method | Inputs beyond authentication |
| --- | --- |
| [ListBCOs](ListBCOs/main.go) | `NEXTHINK_URI` |
| [GetDefaultColumns](GetDefaultColumns/main.go) | `NEXTHINK_URI` |
| [ListFilterCollections](ListFilterCollections/main.go) | `NEXTHINK_URI` |
