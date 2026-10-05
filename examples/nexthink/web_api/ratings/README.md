# Ratings examples

Use the shared [authentication setup](../README.md). From the repository root, run `go run ./examples/nexthink/web_api/ratings/List`. Write examples require `NEXTHINK_REQUEST_FILE` containing an explicit request; replace the synthetic values with your intended lab target.

| Method | Inputs beyond authentication | Sample |
| --- | --- | --- |
| [Create](Create/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Create/request.example.json) |
| [Delete](Delete/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Delete/request.example.json) |
| [Get](Get/main.go) | `NEXTHINK_CONTENT_ID` | — |
| [List](List/main.go) | None | — |
| [Update](Update/main.go) | `NEXTHINK_REQUEST_FILE` | [JSON](Update/request.example.json) |

Ratings apply to an inventory field. Create accepts the target field URI and label, a default enumeration, and conditions keyed by `1` (poor), `2` (average), or `3` (good). Each condition is a complete NQL query, not just a where expression.

Update requires `ratingId` and `revision` in its JSON; the revision is also sent as a query parameter. Delete takes the current revision, sends an empty JSON object, and returns a JSON boolean. Read timestamps are strings while content-list modification times are integers.

Curl and all five examples passed on a previously unrated field with queries matching no devices. Update read-back passed and the fixture was deleted. Field/remote-action metadata, last-value lookup and export are captured as pending auxiliary APIs.
