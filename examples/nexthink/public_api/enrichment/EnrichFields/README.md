# EnrichFields

Use public client credentials and set `NEXTHINK_REQUEST_FILE` to a reviewed copy of [request.example.json](request.example.json). Replace both the device UID and field URI with a dedicated lab device and an existing writable MANUAL custom field.

```sh
go run ./examples/nexthink/public_api/enrichment/EnrichFields
```

The field name is the full data-model URI, such as `device/device/#sdk_fixture`, not just `#sdk_fixture`. This matches the [Enrichment API contract](https://docs.nexthink.com/api/enrichment/enrich-fields-for-given-objects). A successful request should be followed by a read of the actual target field to verify the intended value; do not treat HTTP success as a substitute for readback.
