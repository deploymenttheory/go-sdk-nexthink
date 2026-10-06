# Create software metering

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [resource guide](../README.md). Run from the repository root.

Copy [request.example.json](request.example.json) to a local file and replace `applicationUuids` with applications returned by [GetApplications](../GetApplications/README.md). The fixture UUID is a placeholder, not an enrolled device UUID. Choose a unique `nqlId` and configuration name. The sample meters a WEB application; DESKTOP applications use the corresponding DESKTOP threshold, and HYBRID applications need both threshold types.

Configuration names accept ASCII letters, digits and spaces, up to 255 characters. Underscores and hyphens are not allowed in the name; the separate hash-prefixed `nqlId` can contain underscores.

```sh
NEXTHINK_REQUEST_FILE=/path/to/create.json go run ./examples/nexthink/web_api/software_metering/Create
```

The JSON file contains the configuration directly; do not wrap it in a `configuration` or `variables` object. `Create` returns `{"createConfiguration":true}` rather than a new UUID. Run [List](../List/README.md), match the unique name, and use the returned UUID with Get, Update or Delete. This creates a metering configuration; it does not create an application or enroll a device.
