# Manage content-specific role permissions

Supply `NEXTHINK_REQUEST_FILE` with a role UUID and `contentPermissions` entries containing the actual content ID, resource name and actions. Obtain resource names from ContentAdministration.List and allowed actions from ContentSharing.GetActions; the unit fixture's values are illustrative.

After configuring web authentication, run this operation's example from the repository root. Current IAM returns HTTP 200 with an empty body, represented by a nil SDK result. Verify the grant or revocation through ContentSharing.GetProfiles or the role's shared-content view.

Curl and SDK acceptance granted and revoked dashboard view permission only on an unassigned disposable role and a private test dashboard. Neither HTTP acknowledgement alone nor a permission label proves the resulting access scope.
