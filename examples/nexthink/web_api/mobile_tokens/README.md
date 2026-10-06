# Mobile Collector enrollment tokens

Configure `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and browser authentication as described in the [web API examples](../README.md). Run commands from the repository root.

List, create, retrieve, rename and delete mobile enrollment tokens. Get/Create/Update return redacted metadata on stdout, omitting the JWT credential. To save the complete response, explicitly set `NEXTHINK_OUTPUT_FILE` to a new local file. These examples create the file with owner-only permissions (`0600`) and refuse to overwrite an existing path; stdout remains redacted. Treat the saved file as an enrollment credential. List omits the token value. Revisions come from the latest response. A disposable, unused, short-lived token passed curl and SDK LCRUD and was deleted.

| Method | Guide |
| --- | --- |
| `List` | [Inputs and example](List/README.md) |
| `Get` | [Inputs and example](Get/README.md) |
| `Create` | [Inputs and example](Create/README.md) |
| `Update` | [Inputs and example](Update/README.md) |
| `Delete` | [Inputs and example](Delete/README.md) |
