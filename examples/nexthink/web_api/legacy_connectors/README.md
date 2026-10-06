# LegacyConnectors examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [List](List/main.go) — [request](List/request.example.json)
- [Get](Get/main.go)
- [Create](Create/main.go) — [request](Create/request.example.json)
- [Update](Update/main.go) — [request](Update/request.example.json)
- [Delete](Delete/main.go)
- [SaveSecrets](SaveSecrets/main.go) — [request](SaveSecrets/request.example.json)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Create/Update use the same POST upsert; the identifier selects the legacy connector subtype (for example `AZURE_AD-1`). List options can filter by type/enabled. Secret values are write-only; omit secret submission to leave them unchanged. The lab validated disabled Azure AD configuration with a synthetic nonfunctional secret, not authentication against an external directory.

## Collaboration connector extensions

[HasSecrets](HasSecrets/main.go) performs `GET /secret/{id}` and returns whether the HTTP status is 200. A 404 remains an error with response metadata; no secret values are decoded. [UpdateSecrets](UpdateSecrets/main.go) performs the partial `PATCH /secret/{id}` using the same `SecretRequest` JSON as SaveSecrets. Both require `NEXTHINK_CONTENT_ID`; UpdateSecrets also requires `NEXTHINK_REQUEST_FILE`. The UI uses connector IDs `ms_teams-1` and `zoom-1`. Secret writes change the configured integration credentials.

Create/Update also accept the collaboration configuration shape: `{"connectionDetails":[{"connectionKey":"client_id","connectionValue":"..."}],"mapping":[],"runTime":"23:30"}`. An explicit empty mapping array is valid. ConnectorName and TimeZone are optional. Enabled is now `*bool`: nil omits the field for collaboration configurations, while a pointer to false explicitly disables configurations that support it.

```sh
go run ./examples/nexthink/web_api/legacy_connectors/HasSecrets
go run ./examples/nexthink/web_api/legacy_connectors/UpdateSecrets
```
