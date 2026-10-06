# Dashboards examples

Use the shared [authentication setup](../README.md). From the repository root, run `go run ./examples/nexthink/web_api/dashboards/List`. Write examples require `NEXTHINK_REQUEST_FILE` containing an explicit request; replace the synthetic values with your intended lab target.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Delete/request.example.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID`, `NEXTHINK_PRODUCT_AREA` | — |
| [List](List/main.go) | None | — |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |

Create/Get/Update/Delete use the UI GraphQL documents, with fragments resolved and Apollo `@api` and `@client` directives removed. List uses the shared `dashboards` content key. Update/Delete need the latest revision and should use `meta.productArea` and `meta.type` from Get for their context. Get accepts `NEXTHINK_PRODUCT_AREA` as an optional example input.

GraphQL examples emit partial data before returning an error so a newly created ID remains available for cleanup. Polymorphic widget, filter and layout values are preserved as JSON. Nested widget/filter/tab mutations, clone and import/export remain tracked follow-up operations; gateway access alone is not full dashboard-operation coverage.

All five examples and a separate curl lifecycle passed on an empty private dashboard, with typed Get compared to curl, update read-back and cleanup.

## Additional browser operations

Use `NEXTHINK_API=web`, `NEXTHINK_WEB_AUTH=chrome` (or a browser token), and the lab instance/region. Supply `NEXTHINK_REQUEST_FILE` for examples with a request file, `NEXTHINK_CONTENT_ID` for ID arguments, and `NEXTHINK_EXECUTION_ID` for execution polling. Requests are synthetic templates; replace identifiers with your intended targets.

- [CreateWidget](CreateWidget/main.go) — [request](CreateWidget/request.example.json)
- [UpdateWidget](UpdateWidget/main.go) — [request](UpdateWidget/request.example.json)
- [DeleteWidget](DeleteWidget/main.go) — [request](DeleteWidget/request.example.json)
- [CreateFilter](CreateFilter/main.go) — [request](CreateFilter/request.example.json)
- [UpdateFilter](UpdateFilter/main.go) — [request](UpdateFilter/request.example.json)
- [DeleteFilter](DeleteFilter/main.go) — [request](DeleteFilter/request.example.json)
- [CreateTab](CreateTab/main.go) — [request](CreateTab/request.example.json)
- [UpdateTab](UpdateTab/main.go) — [request](UpdateTab/request.example.json)
- [UpdateTabs](UpdateTabs/main.go) — [request](UpdateTabs/request.example.json)
- [DeleteTab](DeleteTab/main.go) — [request](DeleteTab/request.example.json)
- [UpdateLayout](UpdateLayout/main.go) — [request](UpdateLayout/request.example.json)
- [Export](Export/main.go) — [request](Export/request.example.json)
- [Duplicate](Duplicate/main.go) — [request](Duplicate/request.example.json)
- [Import](Import/main.go) — [request](Import/request.example.json)

Create/Update/Import/Delete and upload methods write data. Test/StartTest contacts the configured destination or starts a server test; review the target first. Re-fetch revisions between dashboard mutations.

Nested operations require the current dashboard revision and product/type context. Read the dashboard again after each mutation; a tab context also needs its tab ID. Widget/filter configuration is a discriminated JSON union. When deleting the last widget, omit `layout` rather than sending an empty nodes array. Import takes `{ "content": <dashboardExport result> }`.
