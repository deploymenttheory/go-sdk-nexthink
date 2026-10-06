# Delete a role

Set `NEXTHINK_ID` to the role UUID returned by CreateRole or ListRoles. It is not the numeric profile ID used by GetRole. After configuring web authentication, run this example from the repository root.

Current IAM returns HTTP 200 with an empty body. The SDK result is nil; an encoded JSON `null` is a successful empty result. Verify removal with ListRoles. The acceptance fixtures were unassigned to users and SSO groups before deletion.
