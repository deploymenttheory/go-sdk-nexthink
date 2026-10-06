# Query product data

`client.WebAPI.DataExploration.Query` executes ad hoc browser NQL. Configure web authentication as described in the [quick start](../../../../../docs/guides/quick-start.md), then run from the repository root:

```sh
NEXTHINK_REQUEST_FILE=examples/nexthink/web_api/data_exploration/Query/request.platform-audit_logs.json \
  go run ./examples/nexthink/web_api/data_exploration/Query
```

The following aggregate-only requests were validated with curl and the SDK in the lab. They reuse the same query contract; each table does not require a separate service. An empty table can return a zero count, which verifies query access but does not establish populated telemetry coverage.

| Product data | Request |
| --- | --- |
| Agent conversations | [JSON](request.agent-conversations.json) |
| Audit logs | [JSON](request.platform-audit_logs.json) |
| Custom-trend logs | [JSON](request.platform-custom_trends_logs.json) |
| Data-export logs | [JSON](request.platform-data_export_logs.json) |
| Inbound-connector logs | [JSON](request.platform-inbound_connector_logs.json) |
| NQL API logs | [JSON](request.platform-nql_api_logs.json) |
| Tickets | [JSON](request.ticket-tickets.json) |
| Nexthink Usage | [JSON](request.usage-account_actions.json) |
| Collaboration sessions | [JSON](request.collaboration-sessions.json) |
| Mobile devices | [JSON](request.mobile_devices.json) |
| VDI sessions | [JSON](request.vdi_sessions.json) |

Discover available collections with `Dashboards.ListCollections`, fields with `DataExploration.ListFields`, and NQL names with `NQLEditor.Complete`. Role visibility and feature availability determine the accessible schema. The lab did not advertise a dedicated AI Tools table in these schema responses; AI tool configuration and governance use `AITools`.

`NEXTHINK_TIME_CONTEXT_FILE` optionally supplies `timeZone`, `utcOffset`, `isoDateTime` and `appName`. Use identical time context when comparing curl and SDK results. Query execution duration varies between calls. The SDK retains null collection display names, including those observed for NQL API logs.
