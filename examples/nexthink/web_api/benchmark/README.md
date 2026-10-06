# Benchmark examples

These read-only browser-token APIs follow first-party UI GraphQL documents. Set `NEXTHINK_API=web`, configure browser authentication, and set `NEXTHINK_REQUEST_FILE` to a JSON file based on the selected example. Replace fixture IDs, dimensions and dates with values available in your tenant. GraphQL partial data is printed before any error is reported.

All operations use `/apigateway/benchmark/graphql`. Successful calls depend on licensed features and available telemetry. Synthetic JSON fixtures prove transport and decoding contracts; they are not copies of customer data. Dynamic analytics values, rows and embedded widget documents retain raw JSON.

- [GetBinaryProductMaps](GetBinaryProductMaps/main.go)
- [GetProfile](GetProfile/main.go)
- [GetProfileSearchItems](GetProfileSearchItems/main.go)
- [GetProfileVersions](GetProfileVersions/main.go)
- [LookupBenchmark](LookupBenchmark/main.go)
- [ProductProperties](ProductProperties/main.go)
- [ProfileProperties](ProfileProperties/main.go)
