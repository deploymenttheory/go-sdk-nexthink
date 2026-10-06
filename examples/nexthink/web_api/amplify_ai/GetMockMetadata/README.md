# Amplify AI GetMockMetadata

Uses `client.WebAPI.AmplifyAI` with browser-token or password authentication. Configure `NEXTHINK_API=web`, instance, region and web authentication.

This contract comes from Nexthink Amplify extension 1.34.0. The lab does not expose this AI backend: curl returned HTTP 404 for both read routes on the extension API hostname and the portal gateway. Successful live behavior is not yet verified.

Set `NEXTHINK_CONTENT_ID` to an existing ticket sys ID.

```sh
go run ./examples/nexthink/web_api/amplify_ai/GetMockMetadata
```
