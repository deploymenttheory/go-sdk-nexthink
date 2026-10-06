Browser session authentication is required. Set `NEXTHINK_API=web` and your instance/browser authentication settings, then run `go run ./examples/nexthink/web_api/recommendations/List`. The list is a JSON array; search, category filtering and record selection are performed locally by the shipped UI.

`UpdateStatus` requires `NEXTHINK_ID` and `NEXTHINK_REQUEST_FILE`; the directory contains a synthetic request example. Status values are `new`, `in-progress`, `done` and `dismissed`. Omit `note` to leave it unspecified, or provide a string. The response contains the updated recommendation.

| Method | HTTP contract |
| --- | --- |
| `List` | `GET /apigateway/recommendations-be/v1/forge/knowledge-recommendations` |
| `UpdateStatus` | `PUT /apigateway/recommendations-be/v1/forge/knowledge-recommendations/{id}/status` |

Both methods follow the Forge frontend 0.10.0 source and have positive/error fixture tests. The lab list returns401; no existing recommendation status was changed. Related conversation identifiers are exposed by `WebAPI.Autopilot.GetRecommendationConversationIDs`. The inspected UI contains no create, single-record retrieval or delete HTTP contract.
