# Amplify AI ExecuteUserAction

Uses `client.WebAPI.AmplifyAI` with browser-token or password authentication. Configure `NEXTHINK_API=web`, instance, region and web authentication.

This contract comes from Nexthink Amplify extension 1.34.0. The lab does not expose this AI backend: curl returned HTTP 404 for both read routes on the extension API hostname and the portal gateway. Successful live behavior is not yet verified.

Copy and customize `request.example.json`, then set `NEXTHINK_REQUEST_FILE`. The identifiers are placeholders.

This operation executes the user action associated with a resolution step. Use only an approved lab fixture.

```sh
go run ./examples/nexthink/web_api/amplify_ai/ExecuteUserAction
```
