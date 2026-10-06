# Diagnostics examples

These read-only browser-token APIs follow first-party UI GraphQL documents. Set `NEXTHINK_API=web`, configure browser authentication, and set `NEXTHINK_REQUEST_FILE` to a JSON file based on the selected example. Replace fixture IDs, dimensions and dates with values available in your tenant. GraphQL partial data is printed before any error is reported.

All operations use `/apigateway/diagnostic/graphql`. Successful calls depend on licensed features and available telemetry. Synthetic JSON fixtures prove transport and decoding contracts; they are not copies of customer data. Dynamic analytics values, rows and embedded widget documents retain raw JSON.

- [BinaryInfo](BinaryInfo/main.go)
- [GetDiagnosticContexts](GetDiagnosticContexts/main.go)
- [GetDiagnosticOverview](GetDiagnosticOverview/main.go)
- [GetImpactedAndTotalObjects](GetImpactedAndTotalObjects/main.go)
- [GetStandaloneDashboard](GetStandaloneDashboard/main.go)
- [HierarchyBreakdownDimensions](HierarchyBreakdownDimensions/main.go)
- [IssueTimeseries](IssueTimeseries/main.go)
- [TroubleshootingInsights](TroubleshootingInsights/main.go)
- [GetConfiguration](GetConfiguration/main.go)
- [IssueEventsAndAssociatedObjectsByHierarchyBreakdown](IssueEventsAndAssociatedObjectsByHierarchyBreakdown/main.go)
- [IssueEventsAndAssociatedObjectsByLocationBreakdown](IssueEventsAndAssociatedObjectsByLocationBreakdown/main.go)
- [IssueEventsAndAssociatedObjectsByTechnicalBreakdown](IssueEventsAndAssociatedObjectsByTechnicalBreakdown/main.go)
- [GetDashboard](GetDashboard/main.go)
