# Workflows examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [ListConnectorDefinitions](ListConnectorDefinitions/main.go)
- [ListConnectorCredentials](ListConnectorCredentials/main.go)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Workflow connector definitions and credential references are separate read views. Manage credential configuration through `WebAPI.ConnectorCredentials`; there are no observed independent create/update/delete operations for the built-in workflow connector definitions.

### Additional management operations

These examples use browser authentication. Request-based examples load `NEXTHINK_REQUEST_FILE`; copy the corresponding `request.example.json` and replace synthetic values. Mutation examples change configuration and should use explicitly selected targets.

- [GetFromLibrary](GetFromLibrary/main.go): `GetWorkflowCopyQuery`.
- [SetActive](SetActive/main.go): `ActivateWorkflowMutation`.
- [Import](Import/main.go): `ImportWorkflowMutation`.

`GetFromLibrary` reads an uninstalled library template by its library content ID; it does not create a copy. Library UUID/update-time metadata can be null. `Import` accepts the exported JSON text. For an empty workflow, the lab exporter omitted `workflow.versions`, but the importer required an explicit empty array; the import example includes it. `SetActive` changes activation only and does not execute a workflow.
