# Update software metering

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [resource guide](../README.md). Run from the repository root.

Copy [request.example.json](request.example.json), replace the application UUID and other configuration fields, and set `NEXTHINK_CONTENT_ID` to the metering configuration UUID returned by List. The request replaces the supplied configuration fields.

```sh
NEXTHINK_CONTENT_ID=your-configuration-uuid NEXTHINK_REQUEST_FILE=/path/to/update.json go run ./examples/nexthink/web_api/software_metering/Update
```

Preserve the existing configuration NQL ID and selected application UUIDs unless you intend to change them. The response is the server update boolean.
