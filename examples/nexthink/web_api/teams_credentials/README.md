# TeamsCredentials examples

Use the SDK single entry point with `NEXTHINK_API=web` and browser authentication as described in [the web API guide](../README.md). Run examples from the repository root.

| Example | HTTP operation | Inputs |
| --- | --- | --- |
| [CheckCredentials](CheckCredentials/main.go) | `POST /apigateway/api/v1/integration/credentials-teams/subscription-manager/check-credentials` | `NEXTHINK_REQUEST_FILE` (CheckCredentialsRequest) |

```sh
go run ./examples/nexthink/web_api/teams_credentials/CheckCredentials
```

CheckCredentials sends a form containing tenant_id, client_id, client_secret and ms_national_cloud. Supply `GLOBAL` for the cloud value used by the UI default. The private JSON input is converted to form data without trimming secret characters. The call validates credentials against the integration; it is not a Nexthink API client login. HTTP failures preserve response metadata and return an error.

Credential checks are source-backed and unit-tested; no third-party credentials were supplied for a live check. Keep request files private and out of version control.
