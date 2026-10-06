# GetMetricBreakdown

Run with `NEXTHINK_API=web`, browser authentication and `NEXTHINK_REQUEST_FILE` pointing to a copy of `request.example.json` with your lab identifiers.

```sh
go run ./examples/nexthink/web_api/application_experience/GetMetricBreakdown
```

The request uses the exact variables of the UI operation. Partial GraphQL data is printed before errors.

Set `desktopAppExperienceId` to the ID of a **desktop** application returned by the Applications service. A web application is not a desktop analytics fixture. Replace this exact field in the request; it is separate from fields named `id` or `applicationId` in other operations.
