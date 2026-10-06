# Data Export examples

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

Start submits an asynchronous NQL CSV export. GetStatus polls the returned statusId; the result may be a signed download URL. Do not send the browser bearer token to an external signed URL. The lab test exported zero matching device records.

| Method | Inputs beyond authentication |
| --- | --- |
| [Start](Start/main.go) | `NEXTHINK_REQUEST_FILE` ([sample](Start/request.example.json)) |
| [GetStatus](GetStatus/main.go) | `NEXTHINK_CONTENT_ID` |
