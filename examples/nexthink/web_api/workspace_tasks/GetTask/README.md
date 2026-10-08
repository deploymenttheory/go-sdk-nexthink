# GetTask

Calls `WebAPI.WorkspaceTasks.GetTask` using the Workspace UI contract.

Read-only.

Use browser-token authentication with `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_ACCESS_TOKEN`, or the SDK headless local-account token provider. See the root README for auth setup.

Inputs:

- `NEXTHINK_RESOURCE_ID`: target resource ID (share hash for GetSharedConversation).

```sh
go run ./examples/nexthink/web_api/workspace_tasks/GetTask
```

Source: `assist-v3-ui/0.268.1`. Tenant feature flags and licensing apply in addition to user permissions. Custom user agents and task automation returned explicit feature-gate denials in the lab. Successful mock fixtures model the observed UI contract; they do not imply live acceptance of a gated operation.
