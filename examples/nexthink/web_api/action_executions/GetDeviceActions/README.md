# GetDeviceActions

Lists the actions available for a device in the Amplify extension. This does not execute actions.

Configure the root client for web authentication (`NEXTHINK_API=web`). Set `NEXTHINK_DEVICE_ID` to `devices[].deviceId.value` returned by `WebAPI.Amplify.Search`: this is the Collector UID, not the NQL device UID.

```sh
go run ./examples/nexthink/web_api/action_executions/GetDeviceActions
```

The lab returned an empty JSON array with HTTP 200 in both curl and the SDK. A populated device-specific action list still requires a suitable action fixture.
