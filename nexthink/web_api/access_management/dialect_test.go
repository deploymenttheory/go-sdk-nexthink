package access_management

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/access_management/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func TestLegacyDialectContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "DeleteRole" || tt.name == "ChangePassword" || tt.name == "GetSAMLMetadata" {
				return
			}
			path := tt.path
			fixed := strings.HasPrefix(path, Endpoint+"/access") || strings.Contains(path, "/ui/v1/role/") || strings.HasPrefix(path, RolePermissionsEndpoint) || tt.name == "DeleteLegacyProfile"
			if !fixed {
				base := LegacyEndpoint
				if strings.Contains(tt.name, "APICredential") {
					base = LegacyCredentialsEndpoint
				}
				path = strings.Replace(path, Endpoint, base, 1)
			}
			if tt.name == "CreateRole" || tt.name == "UpdateRole" {
				path = LegacyEndpoint + "/ui/profilesave"
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+path, mocks.Responder(200, tt.name+"_success"))
			_, response, err := tt.call(NewService(transport, WithLegacyPortal()))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestLegacyMetadataAndUnsupportedDialectOperations(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport, WithLegacyPortal())
	mock.RegisterResponder("GET", testutil.BaseURL+LegacyEndpoint+"/ui/sso/samlmetadata", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "application/xml", r.Header.Get("Accept"))
		x := httpmock.NewStringResponse(200, "<EntityDescriptor entityID=\"fixture\" />")
		x.Header.Set("Content-Type", "application/xml")
		return x, nil
	})
	result, response, err := s.GetSAMLMetadata(context.Background())
	require.NoError(t, err)
	require.NotNil(t, response)
	assert.True(t, result.Status.Success)
	assert.Equal(t, "<EntityDescriptor entityID=\"fixture\" />", result.Result.SAMLMetadata)
	_, _, err = s.DeleteRole(context.Background(), "fixture")
	require.Error(t, err)
	_, _, err = s.ChangePassword(context.Background(), load[PasswordRequest](t, "ChangePassword_input"))
	require.Error(t, err)
	assert.Equal(t, 1, mock.GetTotalCallCount())
}
func TestBusinessFailurePreservesEnvelopeAndMetadata(t *testing.T) {
	for _, name := range []string{"ListUsers", "CreateRole", "GetFeatureFlags"} {
		t.Run(name, func(t *testing.T) {
			for _, tt := range contractCases(t) {
				if tt.name != name {
					continue
				}
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(200, "error_business"))
				result, response, err := tt.call(NewService(transport))
				var statusErr *StatusError
				require.ErrorAs(t, err, &statusErr)
				assert.Equal(t, "FIXTURE_ERROR", statusErr.Status.Code)
				assert.Len(t, statusErr.Status.Errors, 1)
				require.NotNil(t, result)
				require.NotNil(t, response)
				assert.Equal(t, 200, response.StatusCode)
				assert.JSONEq(t, string(mocks.Fixture("error_business")), string(response.Body))
			}
		})
	}
}
func TestDynamicPermissionAndCategoryKeys(t *testing.T) {
	var result RoleResult
	require.NoError(t, json.Unmarshal([]byte(`{"allPermissions":[{"enabledIf":[{"allOf":{"future_permission":"full"}}]}],"infinityViewDomains":{"classifications":{"locations":{"label":"Locations","scopes":[]}}}}`), &result))
	assert.Equal(t, "full", result.AllPermissions[0].EnabledIf[0].AllOf["future_permission"])
	assert.Equal(t, "Locations", result.InfinityViewDomains.Classifications["locations"].Label)
	var content ContentsResult
	require.NoError(t, json.Unmarshal([]byte(`{"allowedActions":{"custom-action":{"label":"Custom","requiresFullViewDomain":false}}}`), &content))
	assert.Contains(t, content.AllowedActions, "custom-action")
}
func TestRoleUUIDEscaping(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+RolePermissionsEndpoint+"/role%2Fwith%3Freserved", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, RolePermissionsEndpoint+"/role%2Fwith%3Freserved", r.URL.EscapedPath())
		return mocks.Responder(200, "GetRolePermissions_success")(r)
	})
	_, _, err := NewService(transport).GetRolePermissions(context.Background(), "role/with?reserved")
	require.NoError(t, err)
}

func TestSearchRequiresTheUISortFields(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	_, _, err := NewService(transport).ListUsers(context.Background(), &ListUsersRequest{PageSize: 50})
	require.Error(t, err)
	assert.Equal(t, 0, mock.GetTotalCallCount())
}
func TestEmptyMutationResponses(t *testing.T) {
	for _, name := range []string{"CreateRole", "UpdateRole", "DeleteRole", "DeleteSupportAccess", "GrantRoleContentPermissions", "RevokeRoleContentPermissions"} {
		t.Run(name, func(t *testing.T) {
			for _, tt := range contractCases(t) {
				if tt.name != name {
					continue
				}
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(204, ""))
				_, response, err := tt.call(NewService(transport))
				require.NoError(t, err)
				require.NotNil(t, response)
				assert.Equal(t, 204, response.StatusCode)
			}
		})
	}
}
