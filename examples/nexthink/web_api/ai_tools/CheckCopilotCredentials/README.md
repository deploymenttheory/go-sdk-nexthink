# CheckCopilotCredentials

Calls `POST /apigateway/aidex/check-credentials/ms-copilot`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/ai_tools/CheckCopilotCredentials/request.example.json"
go run ./examples/nexthink/web_api/ai_tools/CheckCopilotCredentials
```

Review the target and request before running. This operation changes configuration or validates integration credentials. Use a disposable tool or goal for acceptance tests.
