# CCIBenchmarks.Query

Reads the browser API using the SDK root client. Set `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_REQUEST_FILE` to a local JSON request file.

```json
{
  "queries": [
    {
      "source": {
        "name": "dex_login_duration"
      },
      "metrics": [
        "avg_login_duration_per_device"
      ]
    }
  ]
}
```

```sh
go run ./examples/nexthink/web_api/cci_benchmarks/Query
```

Requires an authenticated Chrome portal session. Query rewrites and benchmark queries do not change tenant configuration.
