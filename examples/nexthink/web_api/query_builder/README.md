# Query builder

Access these browser APIs through `client.WebAPI.QueryBuilder` with browser authentication.

| Method | Request | Result |
| --- | --- | --- |
| [Transform](Transform/README.md) | NQL query and transform | Rewritten NQL |
| [ListDrilldownDestinations](ListDrilldownDestinations/README.md) | NQL query and optional kind | Destinations and context identifiers |
| [TransformDrilldown](TransformDrilldown/README.md) | Query, selected result rows, destination transform | Rewritten drilldown NQL |

These operations rewrite queries or inspect query metadata; they do not create stored investigations. Pass the transform and context identifiers returned by the service. Condition values preserve JSON scalar types. The UI may send null conditions for queries without context identifiers; the lab currently returns a server error for that variant. All three methods have successful live curl and SDK checks with non-null drilldown conditions.
