# CreateModule

Calls `POST /apigateway/aidex/config/v1/module`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/ai_tools/CreateModule/request.example.json"
go run ./examples/nexthink/web_api/ai_tools/CreateModule
```

Review the target and request before running. This operation changes configuration or validates integration credentials. Use a disposable tool or goal for acceptance tests.

The module is a tenant-wide singleton. An unconfigured tenant returns 404. Creating or updating it can change campaign exclusions and opt-in settings; the example has opt-in disabled.
