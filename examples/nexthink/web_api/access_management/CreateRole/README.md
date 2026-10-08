# Create a role

Use `AccessManagement.GetRoleTemplate` to discover current permission IDs and defaults. Permission IDs are tenant metadata; do not send the synthetic ID in the unit fixture to a live tenant.

Supply a `RoleRequest` JSON file with a unique name, description, entityVersion 0, selected permissions and the view-domain/timezone settings from the role editor. Set `NEXTHINK_REQUEST_FILE` and run this example from the repository root after configuring web authentication.

The current IAM API returns `{id, entityVersion, uuid}` directly. `GetRole` accepts the numeric ID; `DeleteRole` accepts the UUID. Creating a role does not assign it to an account. Curl and SDK acceptance used disposable unassigned roles and verified their removal.
