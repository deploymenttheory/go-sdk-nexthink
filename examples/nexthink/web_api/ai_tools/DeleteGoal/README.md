# DeleteGoal

Calls `DELETE /apigateway/aidex/config/v1/goal/{id}`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_CONTENT_ID="<existing resource ID>"
export NEXTHINK_REVISION="<latest _rev from Get>"
go run ./examples/nexthink/web_api/ai_tools/DeleteGoal
```

Review the target and request before running. This operation changes configuration or validates integration credentials. Use a disposable tool or goal for acceptance tests.

Use the current revision to avoid overwriting concurrent edits. Update sends a replacement document; retain monitoring and campaign settings when changing governance.
