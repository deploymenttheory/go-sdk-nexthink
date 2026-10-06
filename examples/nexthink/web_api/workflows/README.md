# Workflows examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [ListConnectorDefinitions](ListConnectorDefinitions/main.go)
- [ListConnectorCredentials](ListConnectorCredentials/main.go)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Workflow connector definitions and credential references are separate read views. Manage credential configuration through `WebAPI.ConnectorCredentials`; there are no observed independent create/update/delete operations for the built-in workflow connector definitions.
