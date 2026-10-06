# NLPAssistant examples

Use browser-token authentication: set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE_URL`, and `NEXTHINK_BROWSER_TOKEN`. Run examples from the repository root. Each operation below has its own instructions and request JSON where needed.

Chat buffers the server-sent events until completion; it is not a real-time callback API. Message deltas, status events and the final response remain separate events. Data is JSON and can be decoded using `Event.Decode`; polymorphic artifacts and metadata are retained. In-band error events and malformed/truncated streams return partial results plus an error. A chat prompt can invoke Nexthink assistance capabilities, so use a simple informational prompt for connectivity testing. Live validation asks what NQL means and explicitly requests no queries or content creation.

| Operation | Inputs |
|---|---|
| [Chat](./Chat/README.md) | request JSON |
