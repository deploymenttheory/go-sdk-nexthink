# Update

See the [resource guide](../README.md) for authentication, behavior and validation limits.

Copy [request.example.json](request.example.json), replace the synthetic values, and set `NEXTHINK_REQUEST_FILE` to its path. Set `NEXTHINK_CONTENT_ID` to the ID returned by Create or List.

```sh
go run ./examples/nexthink/web_api/snapshots/Update
```
