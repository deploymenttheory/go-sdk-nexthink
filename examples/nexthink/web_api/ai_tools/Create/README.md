# Create

Calls `POST /apigateway/aidex/config/v2/aitools/application`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_REQUEST_FILE="$PWD/examples/nexthink/web_api/ai_tools/Create/request.example.json"
go run ./examples/nexthink/web_api/ai_tools/Create
```

Review the target and request before running. This operation changes configuration or validates integration credentials. Use a disposable tool or goal for acceptance tests.

Custom tool NQL IDs start with `#`; library tool IDs do not. The example uses an `.invalid` domain and disables campaigns.
