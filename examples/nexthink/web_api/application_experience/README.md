# Application experience examples

Configure `NEXTHINK_API=web`, instance, region and browser authentication as described in the [example index](../README.md). Run from the repository root. Copy the request JSON, replace fixture identifiers and set `NEXTHINK_REQUEST_FILE` to that local file.

| Example | Request |
| --- | --- |
| [GetApplicationsOverviewDesktopInvestigations](GetApplicationsOverviewDesktopInvestigations/main.go) | [GetApplicationsOverviewDesktopInvestigationsRequest](GetApplicationsOverviewDesktopInvestigations/request.example.json) |
| [GetAvgNetworkResponseTime](GetAvgNetworkResponseTime/main.go) | [GetAvgNetworkResponseTimeRequest](GetAvgNetworkResponseTime/request.example.json) |
| [GetBinarySuggestion](GetBinarySuggestion/main.go) | [GetBinarySuggestionRequest](GetBinarySuggestion/request.example.json) |
| [GetDeviceCentricMetricBreakdown](GetDeviceCentricMetricBreakdown/main.go) | [GetDeviceCentricMetricBreakdownRequest](GetDeviceCentricMetricBreakdown/request.example.json) |
| [GetFailedConnectionsRatio](GetFailedConnectionsRatio/main.go) | [GetFailedConnectionsRatioRequest](GetFailedConnectionsRatio/request.example.json) |
| [GetInsights](GetInsights/main.go) | [GetInsightsRequest](GetInsights/request.example.json) |
| [GetMetricBreakdown](GetMetricBreakdown/main.go) | [GetMetricBreakdownRequest](GetMetricBreakdown/request.example.json) |
| [GetNumOfCrashesAndDevices](GetNumOfCrashesAndDevices/main.go) | [GetNumOfCrashesAndDevicesRequest](GetNumOfCrashesAndDevices/request.example.json) |
| [GetNumOfCrashesAndEmployees](GetNumOfCrashesAndEmployees/main.go) | [GetNumOfCrashesAndEmployeesRequest](GetNumOfCrashesAndEmployees/request.example.json) |
| [GetNumOfDevices](GetNumOfDevices/main.go) | [GetNumOfDevicesRequest](GetNumOfDevices/request.example.json) |
| [GetNumOfDevicesWithCrashes](GetNumOfDevicesWithCrashes/main.go) | [GetNumOfDevicesWithCrashesRequest](GetNumOfDevicesWithCrashes/request.example.json) |
| [GetNumOfEmployees](GetNumOfEmployees/main.go) | [GetNumOfEmployeesRequest](GetNumOfEmployees/request.example.json) |
| [GetNumOfEmployeesWithCrashes](GetNumOfEmployeesWithCrashes/main.go) | [GetNumOfEmployeesWithCrashesRequest](GetNumOfEmployeesWithCrashes/request.example.json) |
| [OverviewDesktopTooltips](OverviewDesktopTooltips/main.go) | [OverviewDesktopTooltipsRequest](OverviewDesktopTooltips/request.example.json) |
| [TilesDesktopCrashesPerEmployee](TilesDesktopCrashesPerEmployee/main.go) | [TilesDesktopCrashesPerEmployeeRequest](TilesDesktopCrashesPerEmployee/request.example.json) |
| [TilesDesktopNumberOfEmployees](TilesDesktopNumberOfEmployees/main.go) | [TilesDesktopNumberOfEmployeesRequest](TilesDesktopNumberOfEmployees/request.example.json) |
| [TilesErrorCount](TilesErrorCount/main.go) | [TilesErrorCountRequest](TilesErrorCount/request.example.json) |
| [TilesFrustratingPageLoads](TilesFrustratingPageLoads/main.go) | [TilesFrustratingPageLoadsRequest](TilesFrustratingPageLoads/request.example.json) |
| [TilesNumberOfEmployees](TilesNumberOfEmployees/main.go) | [TilesNumberOfEmployeesRequest](TilesNumberOfEmployees/request.example.json) |
| [TilesPageLoadTime](TilesPageLoadTime/main.go) | [TilesPageLoadTimeRequest](TilesPageLoadTime/request.example.json) |
| [TilesTransactionTime](TilesTransactionTime/main.go) | [TilesTransactionTimeRequest](TilesTransactionTime/request.example.json) |
| [TilesUsageTime](TilesUsageTime/main.go) | [TilesUsageTimeRequest](TilesUsageTime/request.example.json) |
| [WebOverviewTooltips](WebOverviewTooltips/main.go) | [WebOverviewTooltipsRequest](WebOverviewTooltips/request.example.json) |

```sh
NEXTHINK_REQUEST_FILE=/path/to/request.json go run ./examples/nexthink/web_api/application_experience/TilesNumberOfEmployees
```

All methods in this service read metrics or metadata. Time filters require `localFrom`, `localTo`, `localNow` and an IANA `timeZone`; choose a range supported by your tenant retention. Desktop and web input types remain distinct. Nullable metric values and polymorphic insight points are preserved. GraphQL examples print partial data before reporting errors.

Validation covered curl and SDK reads; many desktop metrics and web insights were null in the empty lab. Populated fixtures derive from the live GraphQL schema and cover numeric values, time series and union payloads. Legacy bundle queries using `TimeFrameFilterInput` and `overview` are excluded because the active schema rejects them. Dataset queries are exposed through `CCIInsights.GetDataset`.
