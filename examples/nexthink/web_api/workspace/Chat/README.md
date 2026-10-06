# Chat

Calls `WebAPI.Workspace.Chat` using the Workspace UI contract.

Creates or continues a conversation. Chat can invoke tools; review the prompt and account capabilities. The sample only asks for a definition. The SDK buffers SSE until completion; the context deadline stops waiting.

Use browser-token authentication with `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK headless local-account token provider. See the root README for auth setup.

Inputs:

- `NEXTHINK_REQUEST_FILE`: path to request JSON; copy and review `request.example.json`.

```sh
go run ./examples/nexthink/web_api/workspace/Chat
```

Source: `assist-v3-ui/0.268.1`. Tenant feature flags and licensing apply in addition to user permissions. Custom user agents and task automation returned explicit feature-gate denials in the lab. Successful mock fixtures model the observed UI contract; they do not imply live acceptance of a gated operation.
