# ZoomNotifications examples

Use the SDK single entry point with `NEXTHINK_API=web` and browser authentication as described in [the web API guide](../README.md). Run examples from the repository root.

| Example | HTTP operation | Inputs |
| --- | --- | --- |
| [CheckCredentials](CheckCredentials/main.go) | `POST /apigateway/api/v1/integration/zoom/notification-controller/check-credentials` | `NEXTHINK_REQUEST_FILE` (CheckCredentialsRequest) |
| [GetAppInfo](GetAppInfo/main.go) | `GET /apigateway/api/v1/integration/zoom/notification-controller/app-info` | None |

```sh
go run ./examples/nexthink/web_api/zoom_notifications/CheckCredentials
```

GetAppInfo reads the notification URL and credential-presence flags. CheckCredentials takes `{"jwt":"..."}` and uses the legacy JWT check still shipped by the UI; it does not validate the newer OAuth client configuration. It sends form data and exposes status through response metadata. Use LegacyConnectors to manage Zoom configuration and secret entries.

Credential checks are source-backed and unit-tested; no third-party credentials were supplied for a live check. Keep request files private and out of version control.
