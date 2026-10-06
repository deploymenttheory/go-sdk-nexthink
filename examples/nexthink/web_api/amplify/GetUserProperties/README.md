# Amplify GetUserProperties

Uses the root SDK web client and browser-token or password authentication. Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and the selected web authentication variables.

Set `NEXTHINK_CONTENT_ID` to `users[].userUid.value` (not the SID) returned by Amplify Search.

```sh
go run ./examples/nexthink/web_api/amplify/GetUserProperties
```
