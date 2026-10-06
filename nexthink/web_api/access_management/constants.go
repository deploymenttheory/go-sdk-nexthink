// Package access_management implements the browser UI access-management contracts.
package access_management

const Endpoint = "/apigateway/iam"
const LegacyEndpoint = "/nxarmproxy"
const LegacyCredentialsEndpoint = "/apigateway/nxarmmt"
const RolePermissionsEndpoint = "/apigateway/nxarmrole/api/v1/role"

// Option selects the UI dialect without changing the authentication mechanism.
type Option func(*Service)

// WithLegacyPortal selects the legacy Portal endpoints. API credentials use nxarmmt;
// other Portal operations use nxarmproxy. Support access and role grants remain IAM.
func WithLegacyPortal() Option { return func(s *Service) { s.legacy = true } }
