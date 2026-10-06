package access_management

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct {
	client interfaces.HTTPClient
	legacy bool
}

func NewService(c interfaces.HTTPClient, options ...Option) *Service {
	s := &Service{client: c}
	for _, option := range options {
		option(s)
	}
	return s
}
func (s *Service) endpoint(credentials bool) string {
	if s.legacy {
		if credentials {
			return LegacyCredentialsEndpoint
		}
		return LegacyEndpoint
	}
	return Endpoint
}
func execute[T any](ctx context.Context, s *Service, verb, path string, body any) (*T, *interfaces.Response, error) {
	var result T
	var response *interfaces.Response
	var err error
	headers := map[string]string{"Accept": "application/json", "Content-Language": "en"}
	switch verb {
	case "GET":
		response, err = s.client.Get(ctx, path, nil, headers, &result)
	case "POST":
		response, err = s.client.Post(ctx, path, body, headers, &result)
	case "PUT":
		response, err = s.client.Put(ctx, path, body, headers, &result)
	case "DELETE":
		response, err = s.client.Delete(ctx, path, nil, headers, &result)
	default:
		return nil, nil, fmt.Errorf("unsupported method %s", verb)
	}
	if err != nil {
		return nil, response, err
	}
	if response != nil && len(response.Body) > 0 {
		var envelope struct {
			Status *Status `json:"status"`
		}
		if json.Unmarshal(response.Body, &envelope) == nil && envelope.Status != nil && !envelope.Status.Success {
			return &result, response, &StatusError{Status: *envelope.Status}
		}
	}
	return &result, response, nil
}

// ListUsers implements the observed ARM UI POST /ui/usersearch operation.
func (s *Service) ListUsers(ctx context.Context, request *ListUsersRequest) (*Response[UsersResult], *interfaces.Response, error) {
	if err := validateRequest("ListUsers", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[UsersResult]](ctx, s, "POST", s.endpoint(false)+"/ui/usersearch", request)
}

// CreateUser implements the observed ARM UI POST /ui/usersave operation.
func (s *Service) CreateUser(ctx context.Context, request *UserRequest) (*Response[EntityIdentity], *interfaces.Response, error) {
	if err := validateRequest("CreateUser", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[EntityIdentity]](ctx, s, "POST", s.endpoint(false)+"/ui/usersave", request)
}

// UpdateUser implements the observed ARM UI POST /ui/usersave operation.
func (s *Service) UpdateUser(ctx context.Context, request *UserRequest) (*Response[EntityIdentity], *interfaces.Response, error) {
	if err := validateRequest("UpdateUser", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[EntityIdentity]](ctx, s, "POST", s.endpoint(false)+"/ui/usersave", request)
}

// GetUser implements the observed ARM UI POST /ui/userview operation.
func (s *Service) GetUser(ctx context.Context, request *UserReference) (*Response[UserResult], *interfaces.Response, error) {
	if err := validateRequest("GetUser", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[UserResult]](ctx, s, "POST", s.endpoint(false)+"/ui/userview", request)
}

// DeleteUser implements the observed ARM UI POST /ui/userdelete operation.
func (s *Service) DeleteUser(ctx context.Context, request *DeleteEntityRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("DeleteUser", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/userdelete", request)
}

// ListMappings implements the observed ARM UI POST /ui/mappings operation.
func (s *Service) ListMappings(ctx context.Context) (*Response[MappingsResult], *interfaces.Response, error) {
	return execute[Response[MappingsResult]](ctx, s, "POST", s.endpoint(false)+"/ui/mappings", struct{}{})
}

// UpdateMappings implements the observed ARM UI POST /ui/mappingssave operation.
func (s *Service) UpdateMappings(ctx context.Context, request *MappingsRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("UpdateMappings", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/mappingssave", request)
}

// GetViewDomains implements the observed ARM UI POST /ui/viewdomains operation.
func (s *Service) GetViewDomains(ctx context.Context, request *MappingsRequest) (*Response[ViewDomainsResult], *interfaces.Response, error) {
	if err := validateRequest("GetViewDomains", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[ViewDomainsResult]](ctx, s, "POST", s.endpoint(false)+"/ui/viewdomains", request)
}

// GetSSOConfiguration implements the observed ARM UI GET /ui/sso/config operation.
func (s *Service) GetSSOConfiguration(ctx context.Context) (*Response[SSOConfiguration], *interfaces.Response, error) {
	return execute[Response[SSOConfiguration]](ctx, s, "GET", s.endpoint(false)+"/ui/sso/config", nil)
}

// UpdateSSOConfiguration implements the observed ARM UI POST /ui/sso/config operation.
func (s *Service) UpdateSSOConfiguration(ctx context.Context, request *SSOConfigurationRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("UpdateSSOConfiguration", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/sso/config", request)
}

// ListAPICredentials implements the observed ARM UI POST /ui/apicredentials operation.
func (s *Service) ListAPICredentials(ctx context.Context) (*Response[CredentialsResult], *interfaces.Response, error) {
	return execute[Response[CredentialsResult]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentials", struct{}{})
}

// GetAPICredential implements the observed ARM UI POST /ui/apicredentialview operation.
func (s *Service) GetAPICredential(ctx context.Context, request *IDRequest) (*Response[CredentialResult], *interfaces.Response, error) {
	if err := validateRequest("GetAPICredential", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[CredentialResult]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentialview", request)
}

// CreateAPICredential implements the observed ARM UI POST /ui/apicredentialsave operation.
func (s *Service) CreateAPICredential(ctx context.Context, request *CredentialRequest) (*Response[CredentialSaveResult], *interfaces.Response, error) {
	if err := validateRequest("CreateAPICredential", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[CredentialSaveResult]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentialsave", request)
}

// UpdateAPICredential implements the observed ARM UI POST /ui/apicredentialsave operation.
func (s *Service) UpdateAPICredential(ctx context.Context, request *CredentialRequest) (*Response[CredentialSaveResult], *interfaces.Response, error) {
	if err := validateRequest("UpdateAPICredential", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[CredentialSaveResult]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentialsave", request)
}

// DeleteAPICredential implements the observed ARM UI POST /ui/apicredentialdelete operation.
func (s *Service) DeleteAPICredential(ctx context.Context, request *IDRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("DeleteAPICredential", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentialdelete", request)
}

// ListAPICredentialPermissions implements the observed ARM UI POST /ui/apicredentialpermissions operation.
func (s *Service) ListAPICredentialPermissions(ctx context.Context) (*Response[CredentialPermissionsResult], *interfaces.Response, error) {
	return execute[Response[CredentialPermissionsResult]](ctx, s, "POST", s.endpoint(true)+"/ui/apicredentialpermissions", struct{}{})
}

// ListProfilesAndRoles implements the observed ARM UI POST /ui/profilesroles operation.
func (s *Service) ListProfilesAndRoles(ctx context.Context) (*Response[ProfilesAndRolesResult], *interfaces.Response, error) {
	return execute[Response[ProfilesAndRolesResult]](ctx, s, "POST", s.endpoint(false)+"/ui/profilesroles", struct{}{})
}

// ListRoles implements the observed ARM UI POST /ui/profiles operation.
func (s *Service) ListRoles(ctx context.Context) (*Response[ProfilesResult], *interfaces.Response, error) {
	return execute[Response[ProfilesResult]](ctx, s, "POST", s.endpoint(false)+"/ui/profiles", struct{}{})
}

// CreateRole implements the observed ARM UI POST /ui/v2/role operation.
func (s *Service) CreateRole(ctx context.Context, request *RoleRequest) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateRequest("CreateRole", request); err != nil {
		return nil, nil, err
	}
	path := s.endpoint(false) + "/ui/v2/role"
	if s.legacy {
		path = s.endpoint(false) + "/ui/profilesave"
	}
	return execute[json.RawMessage](ctx, s, "POST", path, request)
}

// UpdateRole implements the observed ARM UI POST /ui/v2/role operation.
func (s *Service) UpdateRole(ctx context.Context, request *RoleRequest) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateRequest("UpdateRole", request); err != nil {
		return nil, nil, err
	}
	path := s.endpoint(false) + "/ui/v2/role"
	if s.legacy {
		path = s.endpoint(false) + "/ui/profilesave"
	}
	return execute[json.RawMessage](ctx, s, "POST", path, request)
}

// GetRole implements the observed ARM UI POST /ui/profileview operation.
func (s *Service) GetRole(ctx context.Context, request *IDRequest) (*Response[RoleResult], *interfaces.Response, error) {
	if err := validateRequest("GetRole", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[RoleResult]](ctx, s, "POST", s.endpoint(false)+"/ui/profileview", request)
}

// GetRoleTemplate implements the observed ARM UI POST /ui/profileview operation.
func (s *Service) GetRoleTemplate(ctx context.Context) (*Response[RoleResult], *interfaces.Response, error) {
	return execute[Response[RoleResult]](ctx, s, "POST", s.endpoint(false)+"/ui/profileview", struct{}{})
}

// GetRoleSummary implements the observed ARM UI POST /ui/profilesummary operation.
func (s *Service) GetRoleSummary(ctx context.Context, request *ProfileIDsRequest) (*Response[RoleResult], *interfaces.Response, error) {
	if err := validateRequest("GetRoleSummary", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[RoleResult]](ctx, s, "POST", s.endpoint(false)+"/ui/profilesummary", request)
}

// DeleteLegacyProfile implements the observed ARM UI POST /ui/profiledelete operation.
func (s *Service) DeleteLegacyProfile(ctx context.Context, request *DeleteEntityRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("DeleteLegacyProfile", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", LegacyEndpoint+"/ui/profiledelete", request)
}

// DeleteRole implements the observed ARM UI DELETE /ui/v2/role/delete/{id} operation.
func (s *Service) DeleteRole(ctx context.Context, id string) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateReference("id", id); err != nil {
		return nil, nil, err
	}
	if s.legacy {
		return nil, nil, fmt.Errorf("DeleteRole uses the IAM UUID contract; use DeleteLegacyProfile for Portal roles")
	}
	return execute[json.RawMessage](ctx, s, "DELETE", s.endpoint(false)+"/ui/v2/role/delete/"+url.PathEscape(id)+"", nil)
}

// GetSAMLMetadata implements the observed ARM UI GET /ui/sso/samlmetadata operation.
func (s *Service) GetSAMLMetadata(ctx context.Context) (*Response[SAMLMetadata], *interfaces.Response, error) {
	if s.legacy {
		response, err := s.client.Get(ctx, LegacyEndpoint+"/ui/sso/samlmetadata", nil, map[string]string{"Accept": "application/xml", "Content-Language": "en"}, nil)
		if err != nil {
			return nil, response, err
		}
		result := &Response[SAMLMetadata]{Status: &Status{Success: true, Errors: []StatusDetail{}}, Result: &SAMLMetadata{SAMLMetadata: string(response.Body)}}
		return result, response, nil
	}
	return execute[Response[SAMLMetadata]](ctx, s, "GET", s.endpoint(false)+"/ui/sso/samlmetadata", nil)
}

// GetFeatureFlags implements the observed ARM UI POST /flags operation.
func (s *Service) GetFeatureFlags(ctx context.Context, request *FeaturesRequest) (*Response[FeaturesResult], *interfaces.Response, error) {
	if err := validateRequest("GetFeatureFlags", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[FeaturesResult]](ctx, s, "POST", s.endpoint(false)+"/flags", request)
}

// GetSharedContents implements the observed ARM UI POST /ui/content operation.
func (s *Service) GetSharedContents(ctx context.Context, request *SharedContentsRequest) (*Response[SharedContentsResult], *interfaces.Response, error) {
	if err := validateRequest("GetSharedContents", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[SharedContentsResult]](ctx, s, "POST", s.endpoint(false)+"/ui/content", request)
}

// ListContents implements the observed ARM UI POST /ui/contentall operation.
func (s *Service) ListContents(ctx context.Context, request *ContentsRequest) (*Response[ContentsResult], *interfaces.Response, error) {
	if err := validateRequest("ListContents", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[ContentsResult]](ctx, s, "POST", s.endpoint(false)+"/ui/contentall", request)
}

// GetAccount implements the observed ARM UI GET /ui/myaccount operation.
func (s *Service) GetAccount(ctx context.Context) (*Response[Account], *interfaces.Response, error) {
	return execute[Response[Account]](ctx, s, "GET", s.endpoint(false)+"/ui/myaccount", nil)
}

// UpdateAccount implements the observed ARM UI PUT /ui/myaccount operation.
func (s *Service) UpdateAccount(ctx context.Context, request *AccountRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("UpdateAccount", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "PUT", s.endpoint(false)+"/ui/myaccount", request)
}

// ResetMFA implements the observed ARM UI POST /ui/myaccount/resetmfa operation.
func (s *Service) ResetMFA(ctx context.Context) (*Response[json.RawMessage], *interfaces.Response, error) {
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/myaccount/resetmfa", struct{}{})
}

// ResetUserMFA implements the observed ARM UI POST /ui/user/resetmfa operation.
func (s *Service) ResetUserMFA(ctx context.Context, request *UserIDRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("ResetUserMFA", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/user/resetmfa", request)
}

// ResetUserPassword implements the observed ARM UI POST /ui/user/resetPassword operation.
func (s *Service) ResetUserPassword(ctx context.Context, request *UserIDRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("ResetUserPassword", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/user/resetPassword", request)
}

// ResendActivationEmail implements the observed ARM UI POST /ui/user/resendactivation operation.
func (s *Service) ResendActivationEmail(ctx context.Context, request *UserIDRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("ResendActivationEmail", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/user/resendactivation", request)
}

// UnlockUser implements the observed ARM UI POST /ui/user/unlock operation.
func (s *Service) UnlockUser(ctx context.Context, request *UserIDRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("UnlockUser", request); err != nil {
		return nil, nil, err
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/user/unlock", request)
}

// ChangePassword implements the observed ARM UI POST /ui/myaccount/password/change operation.
func (s *Service) ChangePassword(ctx context.Context, request *PasswordRequest) (*Response[json.RawMessage], *interfaces.Response, error) {
	if err := validateRequest("ChangePassword", request); err != nil {
		return nil, nil, err
	}
	if s.legacy {
		return nil, nil, fmt.Errorf("legacy password changes use UpdateAccount with PortalMyAccountInfo.Password")
	}
	return execute[Response[json.RawMessage]](ctx, s, "POST", s.endpoint(false)+"/ui/myaccount/password/change", request)
}

// ListSupportAccess implements the observed ARM UI GET /access operation.
func (s *Service) ListSupportAccess(ctx context.Context) (*[]SupportAccess, *interfaces.Response, error) {
	return execute[[]SupportAccess](ctx, s, "GET", Endpoint+"/access", nil)
}

// GetSupportAccess implements the observed ARM UI GET /access/{id} operation.
func (s *Service) GetSupportAccess(ctx context.Context, id string) (*SupportAccess, *interfaces.Response, error) {
	if err := validateReference("id", id); err != nil {
		return nil, nil, err
	}
	return execute[SupportAccess](ctx, s, "GET", Endpoint+"/access/"+url.PathEscape(id)+"", nil)
}

// CreateSupportAccess implements the observed ARM UI POST /access operation.
func (s *Service) CreateSupportAccess(ctx context.Context, request *SupportAccess) (*SupportAccess, *interfaces.Response, error) {
	if err := validateRequest("CreateSupportAccess", request); err != nil {
		return nil, nil, err
	}
	return execute[SupportAccess](ctx, s, "POST", Endpoint+"/access", request)
}

// UpdateSupportAccess implements the observed ARM UI PUT /access/{id} operation.
func (s *Service) UpdateSupportAccess(ctx context.Context, id string, request *SupportAccess) (*SupportAccess, *interfaces.Response, error) {
	if err := validateReference("id", id); err != nil {
		return nil, nil, err
	}
	if err := validateRequest("UpdateSupportAccess", request); err != nil {
		return nil, nil, err
	}
	return execute[SupportAccess](ctx, s, "PUT", Endpoint+"/access/"+url.PathEscape(id)+"", request)
}

// DeleteSupportAccess implements the observed ARM UI DELETE /access/{id} operation.
func (s *Service) DeleteSupportAccess(ctx context.Context, id string) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateReference("id", id); err != nil {
		return nil, nil, err
	}
	return execute[json.RawMessage](ctx, s, "DELETE", Endpoint+"/access/"+url.PathEscape(id)+"", nil)
}

// GetRolePermissions implements the observed ARM UI GET /apigateway/nxarmrole/api/v1/role/{id} operation.
func (s *Service) GetRolePermissions(ctx context.Context, id string) (*RolePermissions, *interfaces.Response, error) {
	if err := validateReference("id", id); err != nil {
		return nil, nil, err
	}
	return execute[RolePermissions](ctx, s, "GET", "/apigateway/nxarmrole/api/v1/role/"+url.PathEscape(id)+"", nil)
}

// GrantRoleContentPermissions implements the observed ARM UI POST /ui/v1/role/grant operation.
func (s *Service) GrantRoleContentPermissions(ctx context.Context, request *GrantRoleContentPermissionsRequest) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateRequest("GrantRoleContentPermissions", request); err != nil {
		return nil, nil, err
	}
	return execute[json.RawMessage](ctx, s, "POST", Endpoint+"/ui/v1/role/grant", request)
}

// RevokeRoleContentPermissions implements the observed ARM UI POST /ui/v1/role/revoke operation.
func (s *Service) RevokeRoleContentPermissions(ctx context.Context, request *GrantRoleContentPermissionsRequest) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateRequest("RevokeRoleContentPermissions", request); err != nil {
		return nil, nil, err
	}
	return execute[json.RawMessage](ctx, s, "POST", Endpoint+"/ui/v1/role/revoke", request)
}
