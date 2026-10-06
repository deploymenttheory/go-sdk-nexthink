package content_sharing

import "encoding/json"

type ActionsOptions struct {
	ContentKey   string `json:"contentKey"`
	ResourceName string `json:"resourceName"`
}
type ProfilesOptions struct {
	ContentKey   string `json:"contentKey"`
	ResourceName string `json:"resourceName"`
	ContentID    string `json:"contentId"`
	Shared       bool   `json:"shared"`
	BCSName      string `json:"bcsName"`
}
type Action struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type ActionsResponse struct {
	Actions []Action `json:"actions"`
}
type Profile struct {
	ProfileID int      `json:"profileId"`
	Name      string   `json:"name"`
	UUID      string   `json:"uuid"`
	Actions   []string `json:"actions,omitempty"`
}
type ProfilesResponse struct {
	Profiles []Profile `json:"profiles"`
	Version  int       `json:"version"`
}
type User struct {
	Fullname string `json:"fullname"`
}
type UserResponse struct {
	User User `json:"user"`
}
type ProfileGrant struct {
	RoleUUID string `json:"roleUuid"`
	// Actions must be non-nil; an explicit empty slice revokes this role's grant.
	Actions []string `json:"actions"`
}

// Empty Actions revoke a role's grant; an empty Profiles array changes no grants.
type ShareContent struct {
	ContentID    string         `json:"contentId"`
	ContentName  string         `json:"contentName"`
	ResourceName string         `json:"resourceName"`
	Profiles     []ProfileGrant `json:"profiles"`
}
type LegacyOptions struct {
	Service   string `json:"service"`
	ContentID string `json:"contentId"`
	BCSName   string `json:"bcsName"`
	Tag       string `json:"tag"`
}
type LegacyActionsResponse struct {
	Version int      `json:"version"`
	Actions []Action `json:"actions"`
}
type LegacyProfile struct {
	ProfileID int      `json:"profileId"`
	Name      string   `json:"name"`
	Actions   []string `json:"actions"`
}
type LegacyProfiles struct {
	Version  int             `json:"version"`
	Profiles []LegacyProfile `json:"profiles"`
}
type LegacyStatus struct {
	Success            bool              `json:"success"`
	Code               string            `json:"code,omitempty"`
	Errors             []string          `json:"errors"`
	RestrictedProfiles []json.RawMessage `json:"restrictedProfiles"`
}

// Legacy APIs also report business failures inside Status under HTTP 200.
type LegacyProfilesResponse struct {
	Result *LegacyProfiles `json:"result"`
	Status LegacyStatus    `json:"status"`
}
type LegacyProfileGrant struct {
	ProfileID int `json:"profileId"`
	// Actions must be non-nil; an explicit empty slice revokes this profile's grant.
	Actions []string `json:"actions"`
}
type LegacyUpdateRequest struct {
	Version  int                  `json:"version"`
	Profiles []LegacyProfileGrant `json:"profiles"`
}
