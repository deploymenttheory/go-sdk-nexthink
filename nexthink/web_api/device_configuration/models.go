package device_configuration

import "encoding/json"

type ProfilesResponse struct {
	OnboardingCompleted bool           `json:"onboardingCompleted"`
	Profiles            []ProfileEntry `json:"profiles"`
}
type ProfileEntry struct {
	Profile    Profile    `json:"profile"`
	Categories []Category `json:"categories"`
}
type Profile struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	IsDefault bool   `json:"isDefault"`
}
type Category struct {
	Name     string    `json:"name"`
	Order    int       `json:"order"`
	Status   string    `json:"status"`
	Settings []Setting `json:"settings"`
}
type Setting struct {
	Name          string          `json:"name"`
	Order         int             `json:"order"`
	Type          string          `json:"type"`
	Value         json.RawMessage `json:"value"`
	Options       json.RawMessage `json:"options"`
	EffectiveDate json.RawMessage `json:"effectiveDate"`
	IsOnboarded   bool            `json:"isOnboarded"`
}

// SaveProfilesRequest updates settings on existing profiles; it does not replace
// the profile collection. Values are encoded according to each setting's type.
type SaveProfilesRequest struct {
	Settings []SettingChange `json:"settings"`
}
type SettingChange struct {
	ProfileID string          `json:"profileId"`
	Name      string          `json:"name"`
	NewValue  json.RawMessage `json:"newValue"`
}
