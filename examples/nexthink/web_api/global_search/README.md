# GlobalSearch examples

Use browser-token authentication: set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE_URL`, and `NEXTHINK_BROWSER_TOKEN`. Run examples from the repository root. Each operation below has its own instructions and request JSON where needed.

Search buffers the server’s concatenated JSON objects until completion. The response retains each category event, provider metadata and unknown result fields. An in-band provider failure returns both partial results and an error. `maxResults` is per category; `loadMoreUrl` is a UI navigation link, not a REST pagination token. Live curl and SDK validation use read-only searches.

| Operation | Inputs |
|---|---|
| [Search](./Search/README.md) | request JSON |
