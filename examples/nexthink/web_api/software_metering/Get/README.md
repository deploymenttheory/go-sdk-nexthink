# Get software metering

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [resource guide](../README.md). Run from the repository root.

Set `NEXTHINK_CONTENT_ID` to a metering configuration UUID returned by List. This reads the configuration, thresholds and selected applications.

```sh
NEXTHINK_CONTENT_ID=your-configuration-uuid go run ./examples/nexthink/web_api/software_metering/Get
```
