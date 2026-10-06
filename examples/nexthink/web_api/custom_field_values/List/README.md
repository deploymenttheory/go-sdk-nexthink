# List

Update changes only the requested field on the supplied objectIds. Values are strings; an empty value clears the field, and boolean fields use "1"/"0". HTTP 202 is acceptance, not proof that values are already visible. ValidateCSV makes a dry-run multipart request; ImportCSV queues updates. Use only intended target identifiers.

Use the authentication and inputs in the [resource guide](../README.md).

Required context: `NEXTHINK_URI` (see the resource guide for meanings).

From the repository root:

```sh
go run ./examples/nexthink/web_api/custom_field_values/List
```
