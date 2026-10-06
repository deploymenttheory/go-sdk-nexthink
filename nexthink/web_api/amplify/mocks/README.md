# Fixture provenance

Successful response fixtures are sanitized, size-limited captures from Amplify extension API curl calls against the lab on 2026-10-06. All eight read responses were compared with the SDK output before sanitizing. Identifying names, user SIDs, UUIDs, serial numbers and IP addresses are replaced. Missing property values, numeric memory, nested disk arrays and the server's `userPrincipleName` spelling are retained.

`GetConfiguration_success.json` is the observed empty configuration collection. It does not establish the shape of every configuration item. The extension consumes an `itsmConfigList` on each item; additional configuration fields remain lossless JSON.

PostInsights has an empty HTTP 200 response. Its request fixture identifies the SDK acceptance origin. Error fixtures are controlled unit-test inputs rather than claims about a live response.
