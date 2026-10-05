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
