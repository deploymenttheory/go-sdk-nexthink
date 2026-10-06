# Amplify PostInsights

Uses the root SDK web client and browser-token or password authentication. Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and the selected web authentication variables.

Copy `request.example.json`, edit the values, and set `NEXTHINK_REQUEST_FILE` to that file.

This call records usage telemetry and changes the Amplify usage history. The sample labels the origin as `sdk-acceptance`. It does not perform the action named in the event.

```sh
go run ./examples/nexthink/web_api/amplify/PostInsights
```
