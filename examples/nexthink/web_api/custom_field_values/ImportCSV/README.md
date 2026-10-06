# ImportCSV

Update changes only the requested field on the supplied objectIds. Values are strings; an empty value clears the field, and boolean fields use "1"/"0". HTTP 202 is acceptance, not proof that values are already visible. ValidateCSV makes a dry-run multipart request; ImportCSV queues updates. Use only intended target identifiers.

Set `NEXTHINK_CSV_FILE` to your CSV path; [sample.csv](sample.csv) shows the format. Configure authentication as described in the [resource guide](../README.md), then run:

```sh
go run ./examples/nexthink/web_api/custom_field_values/ImportCSV
```
