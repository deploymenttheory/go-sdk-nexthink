# ListRemoteActionsForQuery

`POST /apigateway/act/api/v2/remote-action/large`

This reads the UI API.

Copy `request.example.json`, replace the NQL with the UI-generated action eligibility query and its time context, and set `NEXTHINK_REQUEST_FILE` to that file. This endpoint expects the result schema generated for the UI action menu; an ordinary device list query can return `NQL_QUERY_WRONG_RESULT`.

From the repository root:

```sh
NEXTHINK_API=web NEXTHINK_WEB_AUTH=chrome go run ./examples/nexthink/web_api/action_executions/ListRemoteActionsForQuery
```
