# Content Sharing examples

## Additional browser operations

Set `NEXTHINK_API=web`, `NEXTHINK_INSTANCE`, `NEXTHINK_REGION` and `NEXTHINK_WEB_AUTH=chrome` with Chrome signed into the tenant. Static browser tokens also work through the root client.

SetProfiles changes grants on the supplied content entries; empty actions revoke a grant. The current API uses roleUuid, while legacy sharing uses numeric profileId and a version. Legacy business failures may be present in status under HTTP 200: inspect status.success and status.errors. Lab curl and SDK checks verified actual grants and revocation using disposable content and unassigned roles, followed by cleanup. The legacy application-sharing lifecycle was repeated successfully on 2026-10-08. Legacy owner reads did not succeed with the cloud user identity.

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

For legacy sharing, `GetLegacyProfiles` with `NEXTHINK_SHARED=true` returns profiles already granted access; `false` returns profiles without grants. After `SetLegacyProfiles`, read the matching view to verify the change: a granted profile disappears from the unshared result and reappears there after revocation. Read the response's `result.version` for the update request, and use `GetLegacyActions` to discover the service's supported action values. Application sharing was tested with `service=appex` and the `view` action.
