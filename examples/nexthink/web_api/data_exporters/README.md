# DataExporters examples

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [List](List/main.go)
- [Get](Get/main.go)
- [Create](Create/main.go) — [request](Create/request.example.json)
- [Update](Update/main.go) — [request](Update/request.example.json)
- [Delete](Delete/main.go)
- [GetCustomerInfo](GetCustomerInfo/main.go)
- [ListStatuses](ListStatuses/main.go)
- [GetPlaceholders](GetPlaceholders/main.go) — [request](GetPlaceholders/request.example.json)
- [StartTest](StartTest/main.go) — [request](StartTest/request.example.json)
- [GetTest](GetTest/main.go)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Create/Update share a POST upsert. Use a caller-generated UUID and a `#`-prefixed NQL ID. Write enums are numeric: format FILE=0/PAYLOAD=1, file format CSV=0/JSON=1, query type SCHEDULED=0/STREAMING=1. Reads return symbolic names, so do not directly reuse a read response as a write body. `GetPlaceholders` accepts NQL text and encodes the path argument. Poll `GetTest` with the returned execution UUID; the status may initially return 404. Test execution can contact the configured destination even when the saved exporter is disabled.
