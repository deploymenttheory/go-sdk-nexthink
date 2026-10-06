package access_management

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/access_management/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var result T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &result))
	return &result
}

type contractCase struct {
	name, method, path string
	hasBody, empty     bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "ListUsers", method: "POST", path: "/apigateway/iam/ui/usersearch", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.ListUsers(ctx, load[ListUsersRequest](t, "ListUsers_input"))
	}},
		{name: "CreateUser", method: "POST", path: "/apigateway/iam/ui/usersave", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateUser(ctx, load[UserRequest](t, "CreateUser_input"))
		}},
		{name: "UpdateUser", method: "POST", path: "/apigateway/iam/ui/usersave", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateUser(ctx, load[UserRequest](t, "UpdateUser_input"))
		}},
		{name: "GetUser", method: "POST", path: "/apigateway/iam/ui/userview", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetUser(ctx, load[UserReference](t, "GetUser_input"))
		}},
		{name: "DeleteUser", method: "POST", path: "/apigateway/iam/ui/userdelete", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteUser(ctx, load[DeleteEntityRequest](t, "DeleteUser_input"))
		}},
		{name: "ListMappings", method: "POST", path: "/apigateway/iam/ui/mappings", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListMappings(ctx) }},
		{name: "UpdateMappings", method: "POST", path: "/apigateway/iam/ui/mappingssave", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateMappings(ctx, load[MappingsRequest](t, "UpdateMappings_input"))
		}},
		{name: "GetViewDomains", method: "POST", path: "/apigateway/iam/ui/viewdomains", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetViewDomains(ctx, load[MappingsRequest](t, "GetViewDomains_input"))
		}},
		{name: "GetSSOConfiguration", method: "GET", path: "/apigateway/iam/ui/sso/config", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetSSOConfiguration(ctx) }},
		{name: "UpdateSSOConfiguration", method: "POST", path: "/apigateway/iam/ui/sso/config", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateSSOConfiguration(ctx, load[SSOConfigurationRequest](t, "UpdateSSOConfiguration_input"))
		}},
		{name: "ListAPICredentials", method: "POST", path: "/apigateway/iam/ui/apicredentials", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListAPICredentials(ctx) }},
		{name: "GetAPICredential", method: "POST", path: "/apigateway/iam/ui/apicredentialview", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetAPICredential(ctx, load[IDRequest](t, "GetAPICredential_input"))
		}},
		{name: "CreateAPICredential", method: "POST", path: "/apigateway/iam/ui/apicredentialsave", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateAPICredential(ctx, load[CredentialRequest](t, "CreateAPICredential_input"))
		}},
		{name: "UpdateAPICredential", method: "POST", path: "/apigateway/iam/ui/apicredentialsave", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateAPICredential(ctx, load[CredentialRequest](t, "UpdateAPICredential_input"))
		}},
		{name: "DeleteAPICredential", method: "POST", path: "/apigateway/iam/ui/apicredentialdelete", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteAPICredential(ctx, load[IDRequest](t, "DeleteAPICredential_input"))
		}},
		{name: "ListAPICredentialPermissions", method: "POST", path: "/apigateway/iam/ui/apicredentialpermissions", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListAPICredentialPermissions(ctx) }},
		{name: "ListProfilesAndRoles", method: "POST", path: "/apigateway/iam/ui/profilesroles", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListProfilesAndRoles(ctx) }},
		{name: "ListRoles", method: "POST", path: "/apigateway/iam/ui/profiles", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListRoles(ctx) }},
		{name: "CreateRole", method: "POST", path: "/apigateway/iam/ui/v2/role", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateRole(ctx, load[RoleRequest](t, "CreateRole_input"))
		}},
		{name: "UpdateRole", method: "POST", path: "/apigateway/iam/ui/v2/role", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateRole(ctx, load[RoleRequest](t, "UpdateRole_input"))
		}},
		{name: "GetRole", method: "POST", path: "/apigateway/iam/ui/profileview", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetRole(ctx, load[IDRequest](t, "GetRole_input"))
		}},
		{name: "GetRoleTemplate", method: "POST", path: "/apigateway/iam/ui/profileview", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetRoleTemplate(ctx) }},
		{name: "GetRoleSummary", method: "POST", path: "/apigateway/iam/ui/profilesummary", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetRoleSummary(ctx, load[ProfileIDsRequest](t, "GetRoleSummary_input"))
		}},
		{name: "DeleteLegacyProfile", method: "POST", path: "/nxarmproxy/ui/profiledelete", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteLegacyProfile(ctx, load[DeleteEntityRequest](t, "DeleteLegacyProfile_input"))
		}},
		{name: "DeleteRole", method: "DELETE", path: "/apigateway/iam/ui/v2/role/delete/fixture-id", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.DeleteRole(ctx, "fixture-id") }},
		{name: "GetSAMLMetadata", method: "GET", path: "/apigateway/iam/ui/sso/samlmetadata", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetSAMLMetadata(ctx) }},
		{name: "GetFeatureFlags", method: "POST", path: "/apigateway/iam/flags", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetFeatureFlags(ctx, load[FeaturesRequest](t, "GetFeatureFlags_input"))
		}},
		{name: "GetSharedContents", method: "POST", path: "/apigateway/iam/ui/content", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetSharedContents(ctx, load[SharedContentsRequest](t, "GetSharedContents_input"))
		}},
		{name: "ListContents", method: "POST", path: "/apigateway/iam/ui/contentall", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListContents(ctx, load[ContentsRequest](t, "ListContents_input"))
		}},
		{name: "GetAccount", method: "GET", path: "/apigateway/iam/ui/myaccount", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetAccount(ctx) }},
		{name: "UpdateAccount", method: "PUT", path: "/apigateway/iam/ui/myaccount", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateAccount(ctx, load[AccountRequest](t, "UpdateAccount_input"))
		}},
		{name: "ResetMFA", method: "POST", path: "/apigateway/iam/ui/myaccount/resetmfa", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ResetMFA(ctx) }},
		{name: "ResetUserMFA", method: "POST", path: "/apigateway/iam/ui/user/resetmfa", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ResetUserMFA(ctx, load[UserIDRequest](t, "ResetUserMFA_input"))
		}},
		{name: "ResetUserPassword", method: "POST", path: "/apigateway/iam/ui/user/resetPassword", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ResetUserPassword(ctx, load[UserIDRequest](t, "ResetUserPassword_input"))
		}},
		{name: "ResendActivationEmail", method: "POST", path: "/apigateway/iam/ui/user/resendactivation", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ResendActivationEmail(ctx, load[UserIDRequest](t, "ResendActivationEmail_input"))
		}},
		{name: "UnlockUser", method: "POST", path: "/apigateway/iam/ui/user/unlock", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UnlockUser(ctx, load[UserIDRequest](t, "UnlockUser_input"))
		}},
		{name: "ChangePassword", method: "POST", path: "/apigateway/iam/ui/myaccount/password/change", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ChangePassword(ctx, load[PasswordRequest](t, "ChangePassword_input"))
		}},
		{name: "ListSupportAccess", method: "GET", path: "/apigateway/iam/access", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListSupportAccess(ctx) }},
		{name: "GetSupportAccess", method: "GET", path: "/apigateway/iam/access/fixture-id", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetSupportAccess(ctx, "fixture-id") }},
		{name: "CreateSupportAccess", method: "POST", path: "/apigateway/iam/access", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateSupportAccess(ctx, load[SupportAccess](t, "CreateSupportAccess_input"))
		}},
		{name: "UpdateSupportAccess", method: "PUT", path: "/apigateway/iam/access/fixture-id", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateSupportAccess(ctx, "fixture-id", load[SupportAccess](t, "UpdateSupportAccess_input"))
		}},
		{name: "DeleteSupportAccess", method: "DELETE", path: "/apigateway/iam/access/fixture-id", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.DeleteSupportAccess(ctx, "fixture-id") }},
		{name: "GetRolePermissions", method: "GET", path: "/apigateway/nxarmrole/api/v1/role/fixture-id", hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetRolePermissions(ctx, "fixture-id") }},
		{name: "GrantRoleContentPermissions", method: "POST", path: "/apigateway/iam/ui/v1/role/grant", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GrantRoleContentPermissions(ctx, load[GrantRoleContentPermissionsRequest](t, "GrantRoleContentPermissions_input"))
		}},
		{name: "RevokeRoleContentPermissions", method: "POST", path: "/apigateway/iam/ui/v1/role/revoke", hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.RevokeRoleContentPermissions(ctx, load[GrantRoleContentPermissionsRequest](t, "RevokeRoleContentPermissions_input"))
		}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					b, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(b))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 1, mock.GetTotalCallCount())
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			if !tt.empty {
				actual, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			} else {
				assert.Nil(t, result)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(response.Body))
			}
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, failure := range []struct {
				name      string
				responder httpmock.Responder
			}{{"HTTP400", mocks.Responder(400, "error_validation")}, {"HTTP401", mocks.Responder(401, "error_unauthorized")}, {"HTTP403", mocks.Responder(403, "error_forbidden")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(r *http.Request) (*http.Response, error) {
				x := httpmock.NewStringResponse(200, "{broken")
				x.Header.Set("Content-Type", "application/json")
				return x, nil
			}}} {
				t.Run(failure.name, func(t *testing.T) {
					if tt.empty && failure.name == "Malformed" {
						return
					}
					transport, mock := testutil.NewTransport(t)
					mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, failure.responder)
					result, response, err := tt.call(NewService(transport))
					require.Error(t, err)
					assert.Nil(t, result)
					if failure.name != "Transport" {
						require.NotNil(t, response)
					}
				})
			}
		})
	}
}
