# QueryBuilder.TransformDrilldown

Reads the browser API using the SDK root client. Set `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION`, and `NEXTHINK_REQUEST_FILE` to a local JSON request file.

```json
{
  "query": "devices | list device.name",
  "conditions": {
    "identifiers": [
      {
        "objectURI": "device/device/collector/uid",
        "identifier": "device.collector.uid"
      }
    ],
    "values": [
      [
        "00000000-0000-0000-0000-000000000000"
      ]
    ]
  },
  "transform": {
    "objectURI": "user/user",
    "timeSeriesURI": "execution/event"
  }
}
```

```sh
go run ./examples/nexthink/web_api/query_builder/TransformDrilldown
```

Requires an authenticated Chrome portal session. Query rewrites and benchmark queries do not change tenant configuration.
