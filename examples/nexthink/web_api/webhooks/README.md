# Webhooks examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [List](List/main.go)
- [Get](Get/main.go)
- [Create](Create/main.go) — [request](Create/request.example.json)
- [Update](Update/main.go) — [request](Update/request.example.json)
- [Delete](Delete/main.go)
- [GetAvailability](GetAvailability/main.go)
- [Test](Test/main.go) — [request](Test/request.example.json)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Create/Update share a POST upsert and require a caller-generated UUID. `Communication.Payload` holds original bytes and JSON encodes them to base64. `TestRequest.Payload` is plain text. `Test` immediately contacts the destination; an HTTP 400/503 returns an error plus response metadata/body. Saving a disabled configuration does not itself send the test.
