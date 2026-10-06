# Data exploration examples

Configure browser authentication and set `NEXTHINK_API=web`. For methods with input, set `NEXTHINK_REQUEST_FILE` to a copy of that method's `request.example.json`, then run its directory with `go run`. `ListOrganisationFields` requires no input file.

The service supplies the WAAS time headers automatically using UTC and the current time. To reproduce a UI timezone or a fixed clock, set `NEXTHINK_TIME_CONTEXT_FILE` to a file based on [time-context.example.json](time-context.example.json). The UTC offset is minutes west of UTC, matching JavaScript `Date.getTimezoneOffset`: London in summer uses `-60`.

All calls are read-only. Query executes NQL and Inspect returns its metadata. Replace example queries, fields and values with permitted tenant data. Item metadata is supported for `binary/binary/name`; the UI does not request it for arbitrary device fields. Data rows, embedded expressions and dynamic response values retain their JSON representation. GraphQL partial results are printed before an error is reported.

- [Query](Query/main.go)
- [Inspect](Inspect/main.go)
- [GetFilterValues](GetFilterValues/main.go)
- [ListFields](ListFields/main.go)
- [ListSystemRatings](ListSystemRatings/main.go)
- [ListOrganisationFields](ListOrganisationFields/main.go)
- [GetMenu](GetMenu/main.go)
- [ListBreakdownFields](ListBreakdownFields/main.go)
- [GetBreakdownInsights](GetBreakdownInsights/main.go)
- [ListByDurations](ListByDurations/main.go)
- [GetOrganisation](GetOrganisation/main.go)
- [GetItemMeta](GetItemMeta/main.go)
