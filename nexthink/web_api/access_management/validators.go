package access_management

import (
	"fmt"
	"reflect"
	"strings"
)

func validateReference(name, value string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." {
		return fmt.Errorf("%s is required and must be a path reference", name)
	}
	return nil
}
func validateRequest(operation string, request any) error {
	if request == nil || (reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()) {
		return fmt.Errorf("request is required")
	}
	switch r := request.(type) {
	case *IDRequest:
		if r.ID <= 0 {
			return fmt.Errorf("id must be positive")
		}
	case *UserReference:
		if r.ID <= 0 {
			return fmt.Errorf("id must be positive")
		}
	case *DeleteEntityRequest:
		if r.ID <= 0 {
			return fmt.Errorf("id must be positive")
		}
	case *UserIDRequest:
		if r.UserID <= 0 {
			return fmt.Errorf("userId must be positive")
		}
	case *ListUsersRequest:
		if r.PageSize <= 0 || strings.TrimSpace(r.OrderByOption.Field) == "" {
			return fmt.Errorf("pageSize and orderByOption.field are required")
		}
	case *ProfileIDsRequest:
		if len(r.ProfileIDs) == 0 {
			return fmt.Errorf("profileIds are required")
		}
	case *UserRequest:
		if strings.TrimSpace(r.Username) == "" || strings.TrimSpace(r.Fullname) == "" || strings.TrimSpace(r.Email) == "" || r.ProfileID == "" {
			return fmt.Errorf("username, fullname, email and profileId are required")
		}
		if operation == "UpdateUser" && (r.ID == nil || *r.ID <= 0) {
			return fmt.Errorf("id is required to update a user")
		}
	case *CredentialRequest:
		if strings.TrimSpace(r.Name) == "" || len(r.PermissionIDs) == 0 {
			return fmt.Errorf("name and permissionIds are required")
		}
		if operation == "UpdateAPICredential" && (r.ID == nil || *r.ID <= 0) {
			return fmt.Errorf("id is required to update a credential")
		}
	case *RoleRequest:
		if strings.TrimSpace(r.Name) == "" {
			return fmt.Errorf("name is required")
		}
		if operation == "UpdateRole" && (r.ID == nil || *r.ID <= 0) {
			return fmt.Errorf("id is required to update a role")
		}
	case *ContentsRequest:
		if r.CategoryID == "" {
			return fmt.Errorf("categoryId is required")
		}
	case *SharedContentsRequest:
		if r.CategoryID == "" || (r.ProfileID == nil && len(r.ProfileIDs) == 0) {
			return fmt.Errorf("categoryId and profileId or profileIds are required")
		}
	case *PasswordRequest:
		if r.OldPassword == "" || r.NewPassword == "" || r.NewPassword != r.ConfirmPassword {
			return fmt.Errorf("oldPassword and matching newPassword/confirmPassword are required")
		}
	case *GrantRoleContentPermissionsRequest:
		if r.RoleUUID == "" || len(r.ContentPermissions) == 0 {
			return fmt.Errorf("roleUuid and contentPermissions are required")
		}
	case *SupportAccess:
		if (r.AccessType != "INDIVIDUAL" && r.AccessType != "GROUP") || r.MainProfile.ID <= 0 {
			return fmt.Errorf("accessType and mainProfile are required")
		}
	}
	return nil
}
