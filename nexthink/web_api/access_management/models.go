package access_management

import "encoding/json"

// Response is the UI envelope. A false Status.Success is returned as a StatusError,
// even when the server returns HTTP 200; the original result and metadata survive.
type Response[T any] struct {
	Status *Status `json:"status,omitempty"`
	Result *T      `json:"result,omitempty"`
}
type Status struct {
	Success     bool           `json:"success"`
	Code        string         `json:"code,omitempty"`
	Description string         `json:"description,omitempty"`
	Errors      []StatusDetail `json:"errors"`
}
type StatusDetail struct {
	Code        string            `json:"code"`
	Params      []json.RawMessage `json:"params"`
	Description string            `json:"description"`
}
type StatusError struct{ Status Status }

func (e *StatusError) Error() string {
	return "access management: " + e.Status.Code + ": " + e.Status.Description
}

type EntityIdentity struct {
	ID            *int64 `json:"id,omitempty"`
	PortalID      *int64 `json:"portalId,omitempty"`
	EntityVersion *int64 `json:"entityVersion,omitempty"`
}
type IDRequest struct {
	ID int64 `json:"id"`
}
type UserReference struct {
	ID       int64 `json:"id"`
	PortalID int64 `json:"portalId"`
}
type DeleteEntityRequest struct {
	ID            int64  `json:"id"`
	EntityVersion int64  `json:"entityVersion"`
	PortalID      *int64 `json:"portalId,omitempty"`
}
type UserIDRequest struct {
	UserID int64 `json:"userId"`
}
type ProfileIDsRequest struct {
	ProfileIDs []int64 `json:"profileIds"`
}
type OrderByOption struct {
	Field     string `json:"field"`
	Ascending bool   `json:"ascending"`
}
type Cursor struct {
	ID           int64           `json:"id"`
	OrderByValue json.RawMessage `json:"orderByValue"`
}
type ListUsersRequest struct {
	PageSize      int           `json:"pageSize"`
	OrderByOption OrderByOption `json:"orderByOption"`
	SearchOption  string        `json:"searchOption"`
	Cursor        *Cursor       `json:"cursor,omitempty"`
}
type UserRequest struct {
	ID                   *int64            `json:"id,omitempty"`
	PortalID             *int64            `json:"portalId,omitempty"`
	EntityVersion        int64             `json:"entityVersion"`
	Username             string            `json:"username"`
	Fullname             string            `json:"fullname"`
	Email                string            `json:"email"`
	Password             *string           `json:"password,omitempty"`
	PasswordConfirmation *string           `json:"passwordConfirmation,omitempty"`
	UnlimitedSession     bool              `json:"unlimitedSession"`
	AdditionalRoles      []int64           `json:"additionalRoles"`
	ProfileID            string            `json:"profileId"`
	ProfileIDs           []int64           `json:"profileIds"`
	ProfileParams        []json.RawMessage `json:"profileParams"`
}
type CredentialRequest struct {
	ID            *int64  `json:"id,omitempty"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	PermissionIDs []int64 `json:"permissionIds"`
	EntityVersion int64   `json:"entityVersion"`
}

// CredentialSaveResult includes a secret only when returned by creation. Do not log it.
type CredentialSaveResult struct {
	ID            *int64          `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	ClientID      string          `json:"clientId,omitempty"`
	SecretKey     string          `json:"secretKey,omitempty"`
	EntityVersion *int64          `json:"entityVersion,omitempty"`
	Credential    json.RawMessage `json:"credential,omitempty"`
}
type Mapping struct {
	ID            string            `json:"id"`
	SSOGroup      string            `json:"ssoGroup"`
	Idx           int               `json:"idx"`
	ProfileID     string            `json:"profileId"`
	ProfileIDs    []int64           `json:"profileIds"`
	ProfileParams []json.RawMessage `json:"profileParams"`
	ViewDomains   []json.RawMessage `json:"viewDomains"`
	EntityVersion *int64            `json:"entityVersion,omitempty"`
}
type MappingsRequest struct {
	Mappings []Mapping `json:"mappings"`
}
type Certificate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type SSOConfigurationRequest struct {
	SAMLEnabled               bool        `json:"samlEnabled"`
	IssuerURL                 string      `json:"issuerUrl"`
	SSOURL                    string      `json:"ssoUrl"`
	IDPCertificate            Certificate `json:"idpCertificate"`
	GroupsSAMLAttributeName   string      `json:"groupsSamlAttributeName"`
	NameIDFormat              string      `json:"nameIdFormat"`
	FullNameSAMLAttributeName string      `json:"fullNameSamlAttributeName"`
	EmailSAMLAttributeName    string      `json:"emailSamlAttributeName"`
	RequestBinding            string      `json:"requestBinding"`
}
type PermissionValue struct {
	PermissionID int64           `json:"permissionId"`
	Value        json.RawMessage `json:"value"`
}
type SharedContent struct {
	BCSUID        string   `json:"bcsUid"`
	Tag           string   `json:"tag"`
	EntityVersion int64    `json:"entityVersion"`
	Actions       []string `json:"actions"`
	Service       string   `json:"service"`
	Name          string   `json:"name"`
}
type InfinityViewDomain struct {
	Scopes []string `json:"scopes"`
	Level  string   `json:"level"`
}
type MultipleViewDomains struct {
	Classifications []json.RawMessage `json:"classifications"`
}
type RoleRequest struct {
	ID                          *int64               `json:"id,omitempty"`
	PortalID                    *int64               `json:"portalId,omitempty"`
	Name                        string               `json:"name"`
	Description                 string               `json:"description"`
	EntityVersion               int64                `json:"entityVersion"`
	UUID                        string               `json:"uuid,omitempty"`
	Permissions                 []PermissionValue    `json:"permissions"`
	RoleIDs                     []int64              `json:"roleIds"`
	ViewDomains                 []json.RawMessage    `json:"viewDomains"`
	Timezone                    string               `json:"timezone"`
	LandingPage                 string               `json:"landingPage"`
	SharedContents              []SharedContent      `json:"sharedContents"`
	IsSecondary                 bool                 `json:"isSecondary"`
	CombineWithPrimaryLimitedVD bool                 `json:"combineWithPrimaryLimitedVd"`
	LimitedViewDomain           bool                 `json:"limitedViewDomain"`
	InfinityViewDomains         *InfinityViewDomain  `json:"infinityViewDomains,omitempty"`
	MultipleViewDomains         *MultipleViewDomains `json:"multipleViewDomains,omitempty"`
}
type FeaturesRequest struct {
	Features []string `json:"features"`
}
type FeaturesResult struct {
	Features map[string]json.RawMessage `json:"features"`
}
type SharedContentsRequest struct {
	ProfileID  *int64  `json:"profileId,omitempty"`
	ProfileIDs []int64 `json:"profileIds,omitempty"`
	CategoryID string  `json:"categoryId"`
}
type ContentsRequest struct {
	CategoryID                    string   `json:"categoryId"`
	ExcludeIDs                    []string `json:"excludeIds"`
	ExcludeSharedContentProfileID *int64   `json:"excludeSharedContentProfileId,omitempty"`
}
type CSVSettings struct {
	DefaultEncoding string `json:"csvDefaultEncoding"`
	LineTerminator  string `json:"csvLineTerminator"`
	QuoteSymbol     string `json:"csvQuoteSymbol"`
	SeparatorSymbol string `json:"csvSeparatorSymbol"`
}
type DigestSettings struct {
	HierarchyID *int64 `json:"digestHierarchyId,omitempty"`
	ModuleType  string `json:"digestModuleType"`
	SendDigest  bool   `json:"sendDigest"`
}
type PortalAccountInfo struct {
	Password       *string        `json:"password,omitempty"`
	CSVSettings    CSVSettings    `json:"csvSettings"`
	DigestSettings DigestSettings `json:"digestSettings"`
	FinderTimezone string         `json:"finderTimezone"`
}
type AccountRequest struct {
	Username            string             `json:"username"`
	FullName            string             `json:"fullName"`
	Email               string             `json:"email"`
	UITheme             string             `json:"uiTheme,omitempty"`
	Locale              string             `json:"locale,omitempty"`
	PortalMyAccountInfo *PortalAccountInfo `json:"portalMyAccountInfo,omitempty"`
}
type PasswordRequest struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
type ProfileReference struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type SupportAccess struct {
	ID                 json.RawMessage    `json:"id,omitempty"`
	AccessType         string             `json:"accessType"`
	MainProfile        ProfileReference   `json:"mainProfile"`
	AdditionalProfiles []ProfileReference `json:"additionalProfiles"`
	Email              string             `json:"email,omitempty"`
	Expiration         string             `json:"expiration,omitempty"`
	Comment            string             `json:"comment,omitempty"`
	EntityVersion      *int64             `json:"entityVersion,omitempty"`
}
type ContentPermission struct {
	ResourceName string   `json:"resourceName"`
	ContentID    string   `json:"contentId"`
	Actions      []string `json:"actions"`
}
type GrantRoleContentPermissionsRequest struct {
	RoleUUID           string              `json:"roleUuid"`
	ContentPermissions []ContentPermission `json:"contentPermissions"`
}

// RolePermissions preserves permission variants not interpreted by the cockpit UI.
type RolePermissions struct {
	RoleUUID           string              `json:"roleUuid"`
	RoleID             int64               `json:"roleId"`
	ContentPermissions []ContentPermission `json:"contentPermissions"`
}

type UsersResultConfig struct {
	HasSAMLMappings bool   `json:"hasSamlMappings"`
	HasV6Users      bool   `json:"hasV6Users"`
	ShowCrudActions bool   `json:"showCrudActions"`
	CurrentUser     string `json:"currentUser"`
}

type UsersResultSearchInfoUsersItemProfilesItemProfileOriginsItem struct {
	Origin string `json:"origin"`
	Type   string `json:"type"`
}

type UsersResultSearchInfoUsersItemProfilesItem struct {
	ProfileID      int64                                                          `json:"profileId"`
	ProfileOrigins []UsersResultSearchInfoUsersItemProfilesItemProfileOriginsItem `json:"profileOrigins"`
	Name           string                                                         `json:"name"`
	IsSecondary    bool                                                           `json:"isSecondary"`
}

type UsersResultSearchInfoUsersItem struct {
	IsPortalJITEnabled *bool              `json:"isPortalJitEnabled,omitempty"`
	AdditionalRoleIDs  *[]int64           `json:"additionalRoleIds,omitempty"`
	ProfileParams      *[]json.RawMessage `json:"profileParams,omitempty"`

	ID                       int64                                        `json:"id"`
	EntityVersion            int64                                        `json:"entityVersion"`
	PortalID                 int64                                        `json:"portalId"`
	TenantUUID               string                                       `json:"tenantUuid"`
	Username                 string                                       `json:"username"`
	Fullname                 string                                       `json:"fullname"`
	Email                    string                                       `json:"email"`
	Profiles                 []UsersResultSearchInfoUsersItemProfilesItem `json:"profiles"`
	NumberOfSuccessfulLogins int64                                        `json:"numberOfSuccessfulLogins"`
	LastLoginTime            *int64                                       `json:"lastLoginTime,omitempty"`
	UnlimitedSession         bool                                         `json:"unlimitedSession"`
	CanDelete                bool                                         `json:"canDelete"`
	CanEdit                  bool                                         `json:"canEdit"`
	Active                   bool                                         `json:"active"`
	Status                   string                                       `json:"status"`
	CanResetMFA              bool                                         `json:"canResetMfa"`
}

type UsersResultSearchInfoCursor struct {
	ID           int64           `json:"id"`
	OrderByValue json.RawMessage `json:"orderByValue"`
}

type UsersResultSearchInfo struct {
	Users          []UsersResultSearchInfoUsersItem `json:"users"`
	Cursor         UsersResultSearchInfoCursor      `json:"cursor"`
	TotalRows      int64                            `json:"totalRows"`
	ActualPageSize int64                            `json:"actualPageSize"`
	LastPage       bool                             `json:"lastPage"`
}

type UsersResult struct {
	Config     UsersResultConfig     `json:"config"`
	SearchInfo UsersResultSearchInfo `json:"searchInfo"`
}

type UserResultProfilesItemProfileOriginsItem struct {
	Origin string `json:"origin"`
	Type   string `json:"type"`
}

type UserResultProfilesItem struct {
	ProfileID      int64                                      `json:"profileId"`
	ProfileOrigins []UserResultProfilesItemProfileOriginsItem `json:"profileOrigins"`
	Name           string                                     `json:"name"`
	IsSecondary    bool                                       `json:"isSecondary"`
}

type UserResult struct {
	IsPortalJITEnabled *bool              `json:"isPortalJitEnabled,omitempty"`
	AdditionalRoleIDs  *[]int64           `json:"additionalRoleIds,omitempty"`
	ProfileParams      *[]json.RawMessage `json:"profileParams,omitempty"`

	ID                       int64                    `json:"id"`
	EntityVersion            int64                    `json:"entityVersion"`
	PortalID                 int64                    `json:"portalId"`
	TenantUUID               string                   `json:"tenantUuid"`
	Username                 string                   `json:"username"`
	Fullname                 string                   `json:"fullname"`
	Email                    string                   `json:"email"`
	Profiles                 []UserResultProfilesItem `json:"profiles"`
	NumberOfSuccessfulLogins int64                    `json:"numberOfSuccessfulLogins"`
	LastLoginTime            int64                    `json:"lastLoginTime"`
	UnlimitedSession         bool                     `json:"unlimitedSession"`
	CanDelete                bool                     `json:"canDelete"`
	CanEdit                  bool                     `json:"canEdit"`
	Active                   bool                     `json:"active"`
	Status                   string                   `json:"status"`
	CanResetMFA              bool                     `json:"canResetMfa"`
}

type MappingsResultProfilesItem struct {
	ID                 int64             `json:"id"`
	Name               string            `json:"name"`
	ProfileHierarchies []json.RawMessage `json:"profileHierarchies"`
	RoleIDs            []json.RawMessage `json:"roleIds"`
	Selectable         bool              `json:"selectable"`
	IsSecondary        bool              `json:"isSecondary"`
	MTPProfileID       int64             `json:"mtpProfileId"`
}

type MappingsResult struct {
	EntityVersion int64                        `json:"entityVersion"`
	Mappings      []json.RawMessage            `json:"mappings"`
	Profiles      []MappingsResultProfilesItem `json:"profiles"`
}

type ProfilesAndRolesResultProfilesItem struct {
	ID                 int64             `json:"id"`
	Name               string            `json:"name"`
	Selectable         bool              `json:"selectable"`
	ProfileHierarchies []json.RawMessage `json:"profileHierarchies"`
	RoleIDs            []json.RawMessage `json:"roleIds"`
	IsSecondary        bool              `json:"isSecondary"`
	MTPProfileID       int64             `json:"mtpProfileId"`
}

type ProfilesAndRolesResult struct {
	Profiles []ProfilesAndRolesResultProfilesItem `json:"profiles"`
	Roles    []json.RawMessage                    `json:"roles"`
}

type ProfilesResultConfig struct {
	HasSAMLMappings bool   `json:"hasSamlMappings"`
	HasV6Users      bool   `json:"hasV6Users"`
	ShowCrudActions bool   `json:"showCrudActions"`
	CurrentUser     string `json:"currentUser"`
}

type ProfilesResultProfilesItemCanEdit struct {
	Status  bool              `json:"status"`
	Reasons []json.RawMessage `json:"reasons"`
}

type ProfilesResultProfilesItemCanDeleteReasonsItem struct {
	Code        string   `json:"code"`
	Params      []string `json:"params"`
	Description string   `json:"description"`
}

type ProfilesResultProfilesItemCanDelete struct {
	Status  bool                                             `json:"status"`
	Reasons []ProfilesResultProfilesItemCanDeleteReasonsItem `json:"reasons"`
}

type ProfilesResultProfilesItem struct {
	ID                    int64                               `json:"id"`
	EntityVersion         int64                               `json:"entityVersion"`
	Name                  string                              `json:"name"`
	Description           string                              `json:"description"`
	PortalID              int64                               `json:"portalId"`
	InformationLevel      string                              `json:"informationLevel"`
	ViewDomains           []json.RawMessage                   `json:"viewDomains"`
	Roles                 []json.RawMessage                   `json:"roles"`
	FinderAccess          bool                                `json:"finderAccess"`
	NumberOfUsers         int64                               `json:"numberOfUsers"`
	NumberOfGroupMappings int64                               `json:"numberOfGroupMappings"`
	CanEdit               ProfilesResultProfilesItemCanEdit   `json:"canEdit"`
	CanDelete             ProfilesResultProfilesItemCanDelete `json:"canDelete"`
	IsSecondary           bool                                `json:"isSecondary"`
	UUID                  string                              `json:"uuid"`
}

type ProfilesResult struct {
	Config   ProfilesResultConfig         `json:"config"`
	Profiles []ProfilesResultProfilesItem `json:"profiles"`
}

type RoleResultProfilePermissionsItemProfileSourcesItem struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleResultProfilePermissionsItem struct {
	PermissionID   int64                                                 `json:"permissionId"`
	Value          string                                                `json:"value"`
	ProfileSources *[]RoleResultProfilePermissionsItemProfileSourcesItem `json:"profileSources,omitempty"`
}

type RoleResultProfile struct {
	PortalID            *int64               `json:"portalId,omitempty"`
	InfinityViewDomains *InfinityViewDomain  `json:"infinityViewDomains,omitempty"`
	MultipleViewDomains *MultipleViewDomains `json:"multipleViewDomains,omitempty"`
	LimitedViewDomain   *bool                `json:"limitedViewDomain,omitempty"`
	SharedContents      *[]SharedContent     `json:"sharedContents,omitempty"`

	ID                          *int64                             `json:"id,omitempty"`
	EntityVersion               *int64                             `json:"entityVersion,omitempty"`
	Name                        *string                            `json:"name,omitempty"`
	Description                 *string                            `json:"description,omitempty"`
	Timezone                    string                             `json:"timezone"`
	RoleIDs                     []json.RawMessage                  `json:"roleIds"`
	ViewDomains                 []json.RawMessage                  `json:"viewDomains"`
	Permissions                 []RoleResultProfilePermissionsItem `json:"permissions"`
	LandingPage                 string                             `json:"landingPage"`
	IsSecondary                 *bool                              `json:"isSecondary,omitempty"`
	CombineWithPrimaryLimitedVd *bool                              `json:"combineWithPrimaryLimitedVd,omitempty"`
	UUID                        string                             `json:"uuid"`
}

type RoleResultAllPermissionsItemOptionsItem struct {
	Value    string `json:"value"`
	Ordering int64  `json:"ordering"`
	Text     string `json:"text"`
}

type RoleResultAllPermissionsItemEnabledIfItemAllOf map[string]string

type RoleResultAllPermissionsItemEnabledIfItem struct {
	ValueIfDisabled string                                         `json:"valueIfDisabled"`
	AllOf           RoleResultAllPermissionsItemEnabledIfItemAllOf `json:"allOf"`
	Action          string                                         `json:"action"`
}

type RoleResultAllPermissionsItem struct {
	ID                     int64                                       `json:"id"`
	Name                   string                                      `json:"name"`
	Text                   string                                      `json:"text"`
	PermissionType         string                                      `json:"permissionType"`
	Options                []RoleResultAllPermissionsItemOptionsItem   `json:"options"`
	DefaultValue           string                                      `json:"defaultValue"`
	Scope                  string                                      `json:"scope"`
	Service                string                                      `json:"service"`
	EnabledIf              []RoleResultAllPermissionsItemEnabledIfItem `json:"enabledIf"`
	CategoryID             *string                                     `json:"categoryId,omitempty"`
	ClaimName              string                                      `json:"claimName"`
	RequiresFullViewDomain bool                                        `json:"requiresFullViewDomain"`
	Render                 *string                                     `json:"render,omitempty"`
}

type RoleResultAllCategoriesItemContent struct {
	Service string `json:"service"`
	Tag     string `json:"tag"`
}

type RoleResultAllCategoriesItem struct {
	ID      string                              `json:"id"`
	Text    string                              `json:"text"`
	Content *RoleResultAllCategoriesItemContent `json:"content,omitempty"`
}

type RoleResultInfinityViewDomainsClassificationsEntities struct {
	Label  string            `json:"label"`
	Scopes []json.RawMessage `json:"scopes"`
}

type RoleResultInfinityViewDomainsClassifications map[string]RoleResultInfinityViewDomainsClassificationsEntities

type RoleResultInfinityViewDomains struct {
	Classifications RoleResultInfinityViewDomainsClassifications `json:"classifications"`
	Version         string                                       `json:"version"`
}

type RoleResultMultipleViewDomains struct {
	Classifications []json.RawMessage `json:"classifications"`
}

type RoleResultTimezones struct {
	Default string            `json:"default"`
	Options []json.RawMessage `json:"options"`
}

type RoleResult struct {
	Profile                *RoleResultProfile             `json:"profile,omitempty"`
	Roles                  []json.RawMessage              `json:"roles"`
	HierarchiesForProfiles []json.RawMessage              `json:"hierarchiesForProfiles"`
	AllPermissions         []RoleResultAllPermissionsItem `json:"allPermissions"`
	AllCategories          []RoleResultAllCategoriesItem  `json:"allCategories"`
	InfinityViewDomains    RoleResultInfinityViewDomains  `json:"infinityViewDomains"`
	MultipleViewDomains    RoleResultMultipleViewDomains  `json:"multipleViewDomains"`
	Timezones              RoleResultTimezones            `json:"timezones"`
}

type CredentialResultCredential struct {
	ID            int64   `json:"id"`
	EntityVersion int64   `json:"entityVersion"`
	Name          string  `json:"name"`
	ClientID      string  `json:"clientId"`
	CreatedBy     string  `json:"createdBy"`
	CreatedOn     int64   `json:"createdOn"`
	PermissionIDs []int64 `json:"permissionIds"`
	Description   string  `json:"description"`
}

type CredentialResult struct {
	Credential CredentialResultCredential `json:"credential"`
}

type CredentialsResultCredentialsItem struct {
	ID            int64   `json:"id"`
	EntityVersion int64   `json:"entityVersion"`
	Name          string  `json:"name"`
	ClientID      string  `json:"clientId"`
	CreatedBy     string  `json:"createdBy"`
	CreatedOn     int64   `json:"createdOn"`
	PermissionIDs []int64 `json:"permissionIds"`
	Description   string  `json:"description"`
}

type CredentialsResultPermissionsItem struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Service string `json:"service"`
}

type CredentialsResultConfig struct {
	ShowCrudActions bool `json:"showCrudActions"`
}

type CredentialsResult struct {
	Credentials []CredentialsResultCredentialsItem `json:"credentials"`
	Permissions []CredentialsResultPermissionsItem `json:"permissions"`
	Config      CredentialsResultConfig            `json:"config"`
}

type CredentialPermissionsResultPermissionsItem struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Service string `json:"service"`
}

type CredentialPermissionsResult struct {
	Permissions []CredentialPermissionsResultPermissionsItem `json:"permissions"`
}

type SSOConfigurationIDpCertificate struct {
	Name string `json:"name"`
}

type SSOConfiguration struct {
	SAMLEnabled               bool                           `json:"samlEnabled"`
	PortalSAMLEnabled         bool                           `json:"portalSamlEnabled"`
	IssuerURL                 string                         `json:"issuerUrl"`
	SSOURL                    string                         `json:"ssoUrl"`
	IDpCertificate            SSOConfigurationIDpCertificate `json:"idpCertificate"`
	NameIDFormat              string                         `json:"nameIdFormat"`
	FullNameSAMLAttributeName string                         `json:"fullNameSamlAttributeName"`
	EmailSAMLAttributeName    string                         `json:"emailSamlAttributeName"`
	GroupsSAMLAttributeName   string                         `json:"groupsSamlAttributeName"`
	RequestBinding            string                         `json:"requestBinding"`
}

type Account struct {
	UITheme               *string         `json:"uiTheme,omitempty"`
	DataPrivacy           json.RawMessage `json:"dataPrivacy,omitempty"`
	PortalMyAccountInfo   json.RawMessage `json:"portalMyAccountInfo,omitempty"`
	HasFinderAccess       *bool           `json:"hasFinderAccess,omitempty"`
	ProfileFinderTimezone *string         `json:"profileFinderTimezone,omitempty"`

	Username                string            `json:"username"`
	FullName                string            `json:"fullName"`
	Email                   string            `json:"email"`
	MTPRoles                []string          `json:"mtpRoles"`
	DataPrivacyRestrictions []json.RawMessage `json:"dataPrivacyRestrictions"`
	ViewDomain              string            `json:"viewDomain"`
	UserType                string            `json:"userType"`
	LandingPage             string            `json:"landingPage"`
	CanResetMFA             bool              `json:"canResetMfa"`
	Locale                  string            `json:"locale"`
}

type SAMLMetadata struct {
	SAMLMetadata string `json:"samlMetadata"`
}

type ViewDomainsResult struct {
	Mappings []json.RawMessage `json:"mappings"`
}

type ContentsResultAllowedActions map[string]AllowedAction

type ContentsResultContentsItem struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	Tags   []json.RawMessage `json:"tags"`
	Origin string            `json:"origin"`
}

type ContentsResult struct {
	AllowedActions ContentsResultAllowedActions `json:"allowedActions"`
	Contents       []ContentsResultContentsItem `json:"contents"`
}

type SharedContentsResultAllowedActions map[string]AllowedAction

type SharedContentsResult struct {
	AllowedActions SharedContentsResultAllowedActions `json:"allowedActions"`
	SharedContents []json.RawMessage                  `json:"sharedContents"`
}

type AllowedAction struct {
	Label                  string `json:"label"`
	RequiresFullViewDomain bool   `json:"requiresFullViewDomain"`
}
