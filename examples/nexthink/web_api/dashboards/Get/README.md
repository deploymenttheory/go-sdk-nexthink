# Read a dashboard

Configure [web authentication](../../../../../docs/guides/quick-start.md), then supply a dashboard ID from its product-menu entry:

```sh
export NEXTHINK_CONTENT_ID=your-dashboard-uuid
export NEXTHINK_PRODUCT_AREA=desktop-virtualization
go run ./examples/nexthink/web_api/dashboards/Get
```

For a URL shaped `/dashboards/desktop-virtualization/<id>`, use its final segment as the content ID and `desktop-virtualization` as the product area. Discover current entries using `ProductShell.GetMenu`; do not reuse IDs captured from another tenant.

The lab's ten VM, infrastructure and hypervisor dashboards for AVD, Windows 365, Citrix CVAD/DaaS, VMware Horizon and Amazon WorkSpaces used this existing method. Curl and SDK dashboard definitions matched. This verifies definition reads, not populated vendor telemetry. Session-overview menu entries use filtered NQL queries, while individual timeline/health contracts are exposed through `VDI`.

Omit `NEXTHINK_PRODUCT_AREA` when reading a standard live dashboard that does not need product context. Built-in product dashboards may not allow the lifecycle operations available for user-created live dashboards; inspect their returned `allowedOperations`.
