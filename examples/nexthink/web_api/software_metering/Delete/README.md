# Delete software metering

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [resource guide](../README.md). Run from the repository root.

Copy [request.example.json](request.example.json) and replace `uuid` with the metering configuration UUID returned by List. This removes that configuration.

```sh
NEXTHINK_REQUEST_FILE=/path/to/delete.json go run ./examples/nexthink/web_api/software_metering/Delete
```
