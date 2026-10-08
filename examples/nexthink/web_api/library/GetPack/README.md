# GetPack

Configure `NEXTHINK_API=web` and browser authentication and set `NEXTHINK_REQUEST_FILE` to a copy of `request.example.json` with valid lab identifiers.

```sh
go run ./examples/nexthink/web_api/library/GetPack
```

For a built-in catalog pack that has not been imported as a custom pack, supply
its `fileName` and leave the optional `packUUID` empty. That optional query value
identifies a tenant Library record; using the catalog pack UUID there can return
a Library-not-found error even when the catalog filename is valid.
