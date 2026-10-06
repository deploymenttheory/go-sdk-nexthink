# Browser observability proxy

Configure `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and browser authentication as described in the [web API examples](../README.md). Run commands from the repository root.

Submit forwards an explicit telemetry batch to the fixed Nexthink observability proxy. Its client token is the browser telemetry token, not the Nexthink OAuth client secret. The request JSON uses base64 for `payload` because the Go field contains raw bytes; the HTTP body carries the decoded bytes unchanged. Query parameters match the shipped browser telemetry client, including optional encoding, batch and retry metadata. This is telemetry ingestion, with no LCRUD lifecycle. Empty curl submission returned 403; no valid telemetry was sent to the lab. Successful acknowledgment and byte-preserving submission are unit-tested.

| Method | Guide |
| --- | --- |
| `Submit` | [Inputs and example](Submit/README.md) |
