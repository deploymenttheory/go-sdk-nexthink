# Start

Start submits an asynchronous NQL CSV export. GetStatus polls the returned statusId; the result may be a signed download URL. Do not send the browser bearer token to an external signed URL. The lab test exported zero matching device records.

Use the authentication and inputs in the [resource guide](../README.md).

Copy [request.example.json](request.example.json), replace fixture identifiers and fields for your target, then set `NEXTHINK_REQUEST_FILE` to the edited file.

From the repository root:

```sh
go run ./examples/nexthink/web_api/data_export/Start
```
