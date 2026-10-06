# Collaboration Tools

The call-quality and MS Teams Rooms dashboards use the shared dashboard and data-exploration services. `CollaborationTools.GetCallInsights` provides the additional AI insight endpoint used by Device View. Integration configuration remains under `LegacyConnectors`, `TeamsCredentials`, `AzureADCredentials` and `ZoomNotifications` as applicable; no duplicate collaboration-specific configuration CRUD is needed.

## Discover dashboard IDs

Configure browser authentication using the root quick start, then query the same product areas as the UI:

```sh
export NEXTHINK_API=web
export NEXTHINK_PRODUCT_AREAS=collaboration,collaboration-tools
go run ./examples/nexthink/web_api/dashboards/GetProductShellMenu
```

The lab returned three built-in dashboards: Call quality, MS Teams rooms, and Call quality – Selected device. Menu entries include localized titles and URLs such as `/dash/collaboration-tools/system/<id>`. Use that product area and ID to fetch the full definition:

```sh
export NEXTHINK_CONTENT_ID="<id from the menu>"
export NEXTHINK_PRODUCT_AREA=collaboration-tools
go run ./examples/nexthink/web_api/dashboards/Get
```

The dashboard contains tabs, widgets, filters and permission/allowed-operation flags. Honor those flags: the presence of a built-in dashboard does not authorize editing or deleting it. Ordinary editable dashboard LCRUD remains under `Dashboards`.

## Query dashboard data

Widget `config.nqlQuery` values execute through `DataExploration.Query`. For example, this call-quality summary uses the same query as the built-in dashboard:

```json
{
  "query": "collaboration.sessions\n| summarize Total_streams = count()",
  "limit": 5
}
```

Save that JSON to a local file and run:

```sh
export NEXTHINK_REQUEST_FILE="<widget request JSON path>"
go run ./examples/nexthink/web_api/data_exploration/Query
```

Use the widget's time range and filters when reproducing its UI values. The selected-device dashboard also needs its device filter. Teams Rooms equipment widgets depend on remote-action output tables supplied by the relevant Library content; those are data prerequisites rather than separate API routes.

The discovery pass validated all three dashboard definitions (13 tabs), shared widget-query execution and Teams/Zoom insight calls with curl before SDK replay. Call telemetry and installed Teams Rooms remote-action content remain necessary to validate populated charts and equipment results.

## Individual call view

The separate Collaboration Experience call-view UI uses the same `DataExploration` operations. It starts with this query and a `fieldFilter` on `collaboration/session/call/id` for an existing call ID:

```nql
collaboration.sessions
| list session.call.id, application.type, session.call.start_time, session.call.end_time, call.type
| limit 1
```

Its participant tables, audio/video/screen-share metrics, user and device context and network charts also use NQL queries. The current call-view frontend contains 12 shared GraphQL operation documents already implemented by `DataExploration`; it does not require a second set of collaboration CRUD wrappers. Populated call-view acceptance requires real call records.
