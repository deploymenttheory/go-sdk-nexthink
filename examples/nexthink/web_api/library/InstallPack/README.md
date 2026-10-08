# InstallPack

Configure `NEXTHINK_API=web` and browser authentication and set `NEXTHINK_REQUEST_FILE` to a copy of `request.example.json` with valid lab identifiers.

```sh
go run ./examples/nexthink/web_api/library/InstallPack
```

This operation installs or changes Library content; select your intended lab content before running.

Use the selected pack from `GetPack` as the request, including its `packUuid`,
`fileName`, and `builtinContent`. Inspect every included item and dependency
before installing. The installation response can contain only the pack name,
installation state, and counts; the SDK preserves that partial response without
inventing empty catalog fields. Read the pack again to inspect its installed items.
