# UI event polling

Configure `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and browser authentication as described in the [web API examples](../README.md). Run commands from the repository root.

Poll retrieves one page without starting a background loop. Leave `NEXTHINK_NEXT_HREF` unset for the first call; pass `_links.next.href` for the next page. Absolute links must use the configured tenant origin and the exact messages path. The browser stops polling on HTTP 422; the SDK returns that status as an error for caller control. The live lab returned an empty messages array; populated messages and pagination use synthetic fixtures.

| Method | Guide |
| --- | --- |
| `Poll` | [Inputs and example](Poll/README.md) |
