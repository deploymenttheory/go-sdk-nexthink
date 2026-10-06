# Content Sharing examples

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

SetProfiles changes grants on the supplied content entries; empty actions revoke a grant. The current API uses roleUuid, while legacy sharing uses numeric profileId and a version. Legacy business failures may be present in status under HTTP 200: inspect status.success and status.errors. The lab only submitted empty profile lists on disposable content; this validates the request and acknowledgment, not a permission change. Legacy owner reads did not succeed with the cloud user identity.

Grant entries require a non-null `actions` array. Use an explicit empty array (`[]`) to revoke that role or profile; modern grants also require a nonblank `roleUuid`.

| Method | Inputs beyond authentication |
| --- | --- |
| [GetActions](GetActions/main.go) | `NEXTHINK_REQUEST_FILE` ([sample](GetActions/request.example.json)) |
| [GetProfiles](GetProfiles/main.go) | `NEXTHINK_REQUEST_FILE` ([sample](GetProfiles/request.example.json)) |
| [GetUser](GetUser/main.go) | `NEXTHINK_CONTENT_ID` |
| [GetLegacyOwner](GetLegacyOwner/main.go) | `NEXTHINK_CONTENT_ID` |
| [SetProfiles](SetProfiles/main.go) | `NEXTHINK_CONTENT_KEY`, `NEXTHINK_REQUEST_FILE` ([sample](SetProfiles/request.example.json)) |
| [GetLegacyActions](GetLegacyActions/main.go) | `NEXTHINK_SERVICE` |
| [GetLegacyProfiles](GetLegacyProfiles/main.go) | `NEXTHINK_SHARED`, `NEXTHINK_REQUEST_FILE` ([sample](GetLegacyProfiles/request.example.json)) |
| [SetLegacyProfiles](SetLegacyProfiles/main.go) | `NEXTHINK_SERVICE`, `NEXTHINK_CONTENT_ID`, `NEXTHINK_CONTENT_NAME`, `NEXTHINK_TAG`, `NEXTHINK_REQUEST_FILE` ([sample](SetLegacyProfiles/request.example.json)) |

Modern sharing applies changes to the specified roles. An empty `profiles` array changes no grants; it does not clear existing sharing. To revoke a particular role, send its `roleUuid` with `actions: []`. Omitted roles keep their existing permissions.
