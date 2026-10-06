# GetLegacyActions

SetProfiles changes grants on the supplied content entries; empty actions revoke a grant. The current API uses roleUuid, while legacy sharing uses numeric profileId and a version. Legacy business failures may be present in status under HTTP 200: inspect status.success and status.errors. The lab only submitted empty profile lists on disposable content; this validates the request and acknowledgment, not a permission change. Legacy owner reads did not succeed with the cloud user identity.

Use the authentication and inputs in the [resource guide](../README.md).

Required context: `NEXTHINK_SERVICE` (see the resource guide for meanings).

From the repository root:

```sh
go run ./examples/nexthink/web_api/content_sharing/GetLegacyActions
```
