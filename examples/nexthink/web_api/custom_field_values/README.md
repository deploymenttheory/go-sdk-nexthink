# Custom Field Values examples

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

Update changes only the requested field on the supplied objectIds. Values are strings; an empty value clears the field, and boolean fields use "1"/"0". HTTP 202 is acceptance, not proof that values are already visible. ValidateCSV makes a dry-run multipart request; ImportCSV queues updates. Use only intended target identifiers.

| Method | Inputs beyond authentication |
| --- | --- |
| [List](List/main.go) | `NEXTHINK_URI` |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` ([sample](Update/request.example.json)) |
| [ValidateCSV](ValidateCSV/main.go) | `NEXTHINK_CSV_FILE`; [sample](ValidateCSV/sample.csv) |
| [ImportCSV](ImportCSV/main.go) | `NEXTHINK_CSV_FILE`; [sample](ImportCSV/sample.csv) |

CSV headers identify both object and field, for example `device.name,device.#fixture_field`. The sample targets a nonexistent device. Replace it and the field identifier deliberately. Files must use UTF-8, comma separators and a supported ID column; see [Nexthink CSV requirements](https://docs.nexthink.com/platform/user-guide/administration/content-management/custom-fields-management).
