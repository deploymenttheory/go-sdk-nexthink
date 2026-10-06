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
