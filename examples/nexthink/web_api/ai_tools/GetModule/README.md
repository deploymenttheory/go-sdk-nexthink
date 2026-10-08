# GetModule

Calls `GET /apigateway/aidex/config/v1/module`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
go run ./examples/nexthink/web_api/ai_tools/GetModule
```

The module is a tenant-wide singleton. An unconfigured tenant returns 404. Creating or updating it can change campaign exclusions and opt-in settings; the example has opt-in disabled.
