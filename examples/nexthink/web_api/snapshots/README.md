# Snapshots and custom trends

Configure `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and browser authentication as described in the [web API examples](../README.md). Run commands from the repository root.

Snapshots use the UI custom-trend definition API and the shared `trends` content listing. NQL IDs begin with `#` and contain lowercase letters, digits and underscores. The server requires a `list` statement and rejects `summarize` and identifying fields, including fields used only in filters. The provided hardware-manufacturer filter matches no real device. Retention approval is explicit; the SDK never enables it implicitly. Create/Get/Update/List/Export/Import/Delete passed on disposable zero-match definitions and were cleaned up. Export projects Get into portable definition JSON; Import uses Create, so choose a new name and NQL ID. There are no separate HTTP import/export endpoints.

| Method | Guide |
| --- | --- |
| `List` | [Inputs and example](List/README.md) |
| `Get` | [Inputs and example](Get/README.md) |
| `Export` | [Inputs and example](Export/README.md) |
| `Delete` | [Inputs and example](Delete/README.md) |
| `Create` | [Inputs and example](Create/README.md) |
| `Update` | [Inputs and example](Update/README.md) |
| `Import` | [Inputs and example](Import/README.md) |
