# CreateCopilot

Calls `POST /apigateway/aidex/config/v2/aitools/ms-copilot`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/ai_tools/CreateCopilot/request.example.json"
go run ./examples/nexthink/web_api/ai_tools/CreateCopilot
```

Review the target and request before running. This operation changes configuration or validates integration credentials. Use a disposable tool or goal for acceptance tests.
