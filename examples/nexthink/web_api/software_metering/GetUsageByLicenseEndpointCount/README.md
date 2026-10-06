# GetUsageByLicenseEndpointCount

Run with `NEXTHINK_API=web`, browser authentication and `NEXTHINK_REQUEST_FILE` pointing to a copy of `request.example.json` with your lab identifiers.

```sh
go run ./examples/nexthink/web_api/software_metering/GetUsageByLicenseEndpointCount
```

The request uses the exact variables of the UI operation. Partial GraphQL data is printed before errors.

Set `configurationUuid` to the actual metering configuration UUID returned by the SoftwareMetering service. This operation uses `configurationUuid`; other metering examples may instead name that variable `uuid`. An application UUID identifies a different resource.
