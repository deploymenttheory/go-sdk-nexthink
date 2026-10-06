# AlertHub examples

These read-only browser-token APIs follow first-party UI GraphQL documents. Set `NEXTHINK_API=web`, configure browser authentication, and set `NEXTHINK_REQUEST_FILE` to a JSON file based on the selected example. Replace fixture IDs, dimensions and dates with values available in your tenant. GraphQL partial data is printed before any error is reported.

All operations use `/apigateway/mnt/alert/hub/graphql`. Successful calls depend on licensed features and available telemetry. Synthetic JSON fixtures prove transport and decoding contracts; they are not copies of customer data. Dynamic analytics values, rows and embedded widget documents retain raw JSON.

- [AlertImpactAssessment](AlertImpactAssessment/main.go)
- [EventsDrillDown](EventsDrillDown/main.go)
- [Issues](Issues/main.go)
- [IssuesSelectedPeriod](IssuesSelectedPeriod/main.go)
- [IssuesTimeline](IssuesTimeline/main.go)
- [MonitorConfigView](MonitorConfigView/main.go)
- [AlertTriggerInfo](AlertTriggerInfo/main.go)
- [AlertsImpactedAssociationOverTime](AlertsImpactedAssociationOverTime/main.go)
- [GetAlertOnChangeTimeSeries](GetAlertOnChangeTimeSeries/main.go)
- [GetIssuePeriods](GetIssuePeriods/main.go)
- [GetStaticAlertTimeSeries](GetStaticAlertTimeSeries/main.go)
- [Tags](Tags/main.go)
