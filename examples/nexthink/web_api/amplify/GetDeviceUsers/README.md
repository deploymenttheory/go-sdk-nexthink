# Amplify GetDeviceUsers

Uses the root SDK web client and browser-token or password authentication. Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and the selected web authentication variables.

Set `NEXTHINK_CONTENT_ID` to `devices[].deviceId.value` (Collector UID, not the NQL device UID) returned by Amplify Search.

```sh
go run ./examples/nexthink/web_api/amplify/GetDeviceUsers
```
