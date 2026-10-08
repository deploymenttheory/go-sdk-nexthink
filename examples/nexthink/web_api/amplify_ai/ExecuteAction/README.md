# Amplify AI ExecuteAction

Uses `client.WebAPI.AmplifyAI` with browser-token or password authentication. Configure `NEXTHINK_API=web`, instance, region and web authentication.

This contract comes from Nexthink Amplify extension 1.34.0. The lab does not expose this AI backend: curl returned HTTP 404 for both read routes on the extension API hostname and the portal gateway. Successful live behavior is not yet verified.

Copy and customize `request.example.json`, then set `NEXTHINK_REQUEST_FILE`. The identifiers are placeholders.

This operation executes an action associated with a resolution step. Use only an approved lab fixture.

Populate `actionExecutionParams` from the selected resolution step. Alternatively use the extension remote-action widget payload with `remoteActionId`, `deviceId`, `params`, `targets` and the resolution identifiers. Inputs are action-specific.

```sh
go run ./examples/nexthink/web_api/amplify_ai/ExecuteAction
```
