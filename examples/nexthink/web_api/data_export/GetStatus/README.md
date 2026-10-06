# GetStatus

Start submits an asynchronous NQL CSV export. GetStatus polls the returned statusId; the result may be a signed download URL. Do not send the browser bearer token to an external signed URL. The lab test exported zero matching device records.

Use the authentication and inputs in the [resource guide](../README.md).

Required context: `NEXTHINK_CONTENT_ID` (see the resource guide for meanings).

From the repository root:

```sh
go run ./examples/nexthink/web_api/data_export/GetStatus
```
