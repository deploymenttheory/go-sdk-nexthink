# DexScores examples

These read-only browser-token APIs follow first-party UI GraphQL documents. Set `NEXTHINK_API=web`, configure browser authentication, and set `NEXTHINK_REQUEST_FILE` to a JSON file based on the selected example. Replace fixture IDs, dimensions and dates with values available in your tenant. GraphQL partial data is printed before any error is reported.

All operations use `/apigateway/dex-ec/graphql`. Successful calls depend on licensed features and available telemetry. Synthetic JSON fixtures prove transport and decoding contracts; they are not copies of customer data. Dynamic analytics values, rows and embedded widget documents retain raw JSON.

- [GetCampaign](GetCampaign/main.go)
- [GetDeviceExperience](GetDeviceExperience/main.go)
- [GetDimensionBreakdowns](GetDimensionBreakdowns/main.go)
- [GetDimensionsV2](GetDimensionsV2/main.go)
- [GetInvestigationUrl](GetInvestigationUrl/main.go)
- [GetLeaves](GetLeaves/main.go)
- [GetMetricThreshold](GetMetricThreshold/main.go)
- [GetScores](GetScores/main.go)
- [GetTrend](GetTrend/main.go)
- [GetTrendDevices](GetTrendDevices/main.go)
- [GetTrendEmployeesWithIssues](GetTrendEmployeesWithIssues/main.go)
- [GetTrendImprovement](GetTrendImprovement/main.go)
- [GetTrendScore](GetTrendScore/main.go)
- [GetTrendTimeLost](GetTrendTimeLost/main.go)
- [GetTrendWithRange](GetTrendWithRange/main.go)
- [GetWhatsChanged](GetWhatsChanged/main.go)
