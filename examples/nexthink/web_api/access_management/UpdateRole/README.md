# Update a role

Read the role with `AccessManagement.GetRole`, preserve its permission values and scope, and supply the current numeric `id`, `uuid` and `entityVersion` in `NEXTHINK_REQUEST_FILE`. Change only the intended fields. The same POST route handles creation and updates; the response contains `{id, entityVersion, uuid}` directly.

Run `go run ./examples/nexthink/web_api/access_management/UpdateRole` after configuring web authentication. A fresh GetRole readback should confirm the intended changes. The acceptance run used unassigned roles rather than changing an existing account's access.
