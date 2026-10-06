# List software metering

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [resource guide](../README.md). Run from the repository root.

List all metering configuration summaries, including their UUIDs. No request file is required. Use a configuration UUID for Get, Update and Delete; application UUIDs come from GetApplications instead.

```sh
go run ./examples/nexthink/web_api/software_metering/List
```
