# Create

See the [resource guide](../README.md) for authentication, behavior and validation limits.

Copy [request.example.json](request.example.json), replace the synthetic values, and set `NEXTHINK_REQUEST_FILE` to its path.

```sh
go run ./examples/nexthink/web_api/mobile_tokens/Create
```

Stdout contains token metadata with `jwtToken` omitted. To save the full response, including the enrollment credential, explicitly set `NEXTHINK_OUTPUT_FILE` to a new local path. The example creates it with owner-only permissions (`0600`) and refuses to overwrite an existing file. Stdout stays redacted when a file is saved.
