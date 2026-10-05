package device_configuration

import "encoding/json"

type ProfilesResponse struct {
	OnboardingCompleted bool           `json:"onboardingCompleted"`
	Profiles            []ProfileEntry `json:"profiles"`
}

// Profile and category contents depend on enabled products and tenant settings.
type ProfileEntry struct {
	Profile    json.RawMessage   `json:"profile"`
	Categories []json.RawMessage `json:"categories"`
}
