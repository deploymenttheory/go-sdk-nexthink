# GetGoalInsights

Calls `GET /apigateway/aidex/goals/v1/goals/{id}/insights`.

Use the SDK browser-token or local-account password configuration described in the root quick start.

```sh
export NEXTHINK_API=web
export NEXTHINK_CONTENT_ID="<existing resource ID>"
go run ./examples/nexthink/web_api/ai_tools/GetGoalInsights
```

The UI gates this endpoint behind the goal-tracking-insights feature flag. The lab currently returns HTTP 403 even though goal configuration CRUD succeeds.
