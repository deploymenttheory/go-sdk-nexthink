Browser session authentication is required. Set `NEXTHINK_API=web`, the instance URL, and browser-token authentication as described in the main examples README. Run each directory with `go run ./examples/nexthink/web_api/access_management/METHOD`.

Requests use `NEXTHINK_REQUEST_FILE`; each request-based example includes a synthetic `request.example.json`. Replace fixture values with your explicitly intended lab targets before running writes. Methods with path references require the matching `NEXTHINK_ID`, `NEXTHINK_DOCUMENT_ID`, `NEXTHINK_COMMENT_ID`, or `NEXTHINK_REPLY_ID` variable. Mention searches use `NEXTHINK_SEARCH`.

The default is `WebAPI.AccessManagement` (IAM). Set `NEXTHINK_ACCESS_LEGACY=true` to use `WebAPI.LegacyAccessManagement`: Portal operations route to `/nxarmproxy`, API credentials to `/apigateway/nxarmmt`. Support access and role-permission grants use fixed IAM routes. Modern role create/update share `POST /ui/v2/role`; Portal create/update share `/ui/profilesave`. `DeleteLegacyProfile` always uses Portal. Legacy password updates use `UpdateAccount.PortalMyAccountInfo.Password`; IAM uses `ChangePassword`. Legacy SAML XML is normalized to `Response[SAMLMetadata]`.

HTTP 200 can contain `status.success=false`; methods return `StatusError` together with the decoded envelope and HTTP metadata. Opaque acknowledgments remain JSON; role/profile permission variants without a fixed schema remain raw JSON. API credential creation can return `secretKey`; protect its output. Account/security, permission, SSO, support-access and credential mutations are source-evidenced and unit-tested; they are not live-tested against existing lab configuration. `/flags` is present in the shipped client but returned404 in this tenant.

| Method | HTTP contract |
| --- | --- |
| `ListUsers` | `POST /ui/usersearch` |
| `CreateUser` | `POST /ui/usersave` |
| `UpdateUser` | `POST /ui/usersave` |
| `GetUser` | `POST /ui/userview` |
| `DeleteUser` | `POST /ui/userdelete` |
| `ListMappings` | `POST /ui/mappings` |
| `UpdateMappings` | `POST /ui/mappingssave` |
| `GetViewDomains` | `POST /ui/viewdomains` |
| `GetSSOConfiguration` | `GET /ui/sso/config` |
| `UpdateSSOConfiguration` | `POST /ui/sso/config` |
| `ListAPICredentials` | `POST /ui/apicredentials` |
| `GetAPICredential` | `POST /ui/apicredentialview` |
| `CreateAPICredential` | `POST /ui/apicredentialsave` |
| `UpdateAPICredential` | `POST /ui/apicredentialsave` |
| `DeleteAPICredential` | `POST /ui/apicredentialdelete` |
| `ListAPICredentialPermissions` | `POST /ui/apicredentialpermissions` |
| `ListProfilesAndRoles` | `POST /ui/profilesroles` |
| `ListRoles` | `POST /ui/profiles` |
| `CreateRole` | `POST /ui/v2/role` |
| `UpdateRole` | `POST /ui/v2/role` |
| `GetRole` | `POST /ui/profileview` |
| `GetRoleTemplate` | `POST /ui/profileview` |
| `GetRoleSummary` | `POST /ui/profilesummary` |
| `DeleteLegacyProfile` | `POST /ui/profiledelete` |
| `DeleteRole` | `DELETE /ui/v2/role/delete/{id}` |
| `GetSAMLMetadata` | `GET /ui/sso/samlmetadata` |
| `GetFeatureFlags` | `POST /flags` |
| `GetSharedContents` | `POST /ui/content` |
| `ListContents` | `POST /ui/contentall` |
| `GetAccount` | `GET /ui/myaccount` |
| `UpdateAccount` | `PUT /ui/myaccount` |
| `ResetMFA` | `POST /ui/myaccount/resetmfa` |
| `ResetUserMFA` | `POST /ui/user/resetmfa` |
| `ResetUserPassword` | `POST /ui/user/resetPassword` |
| `ResendActivationEmail` | `POST /ui/user/resendactivation` |
| `UnlockUser` | `POST /ui/user/unlock` |
| `ChangePassword` | `POST /ui/myaccount/password/change` |
| `ListSupportAccess` | `GET /access` |
| `GetSupportAccess` | `GET /access/{id}` |
| `CreateSupportAccess` | `POST /access` |
| `UpdateSupportAccess` | `PUT /access/{id}` |
| `DeleteSupportAccess` | `DELETE /access/{id}` |
| `GetRolePermissions` | `GET /apigateway/nxarmrole/api/v1/role/{id}` |
| `GrantRoleContentPermissions` | `POST /ui/v1/role/grant` |
| `RevokeRoleContentPermissions` | `POST /ui/v1/role/revoke` |
