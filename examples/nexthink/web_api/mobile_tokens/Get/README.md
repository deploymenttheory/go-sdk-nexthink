# Get

See the [resource guide](../README.md) for authentication, behavior and validation limits.

Set `NEXTHINK_CONTENT_ID` to the ID returned by Create or List.

```sh
go run ./examples/nexthink/web_api/mobile_tokens/Get
```

Stdout contains token metadata with `jwtToken` omitted. To save the full response, including the enrollment credential, explicitly set `NEXTHINK_OUTPUT_FILE` to a new local path. The example creates it with owner-only permissions (`0600`) and refuses to overwrite an existing file. Stdout stays redacted when a file is saved.
