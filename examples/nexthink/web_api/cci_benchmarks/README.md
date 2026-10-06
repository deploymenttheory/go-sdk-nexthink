# CCI benchmarks

Use `client.WebAPI.CCIBenchmarks.Query` to read one or more aggregate industry benchmarks. See [Query](Query/README.md) for a runnable example.

Each query names a benchmark source, requested metrics, and optional filters such as `ext/segment/industry`. `TimeZone` controls the `x-client-timezone` header and defaults to UTC. The result retains partial results alongside service errors. Error details remain JSON because their schema is not defined by the observed browser consumer.

This batch REST contract is separate from `client.WebAPI.CCIInsights` GraphQL operations. It does not expose create, update, or delete operations in the observed UI. The SDK response has been compared with a successful live curl response.
