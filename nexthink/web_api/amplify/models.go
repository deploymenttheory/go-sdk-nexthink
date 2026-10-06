package amplify

import "encoding/json"

// Property preserves the typed values returned by Amplify. Value can be a string,
// number, boolean, JSON-encoded string, or absent when no value is available.
type Property struct {
	Type        string          `json:"type"`
	Value       json.RawMessage `json:"value,omitempty"`
	Label       string          `json:"label,omitempty"`
	Description string          `json:"description,omitempty"`
}

// Properties retains the server's field names, including userPrincipleName.
// Most fields decode as Property; device disks are a nested array and are kept
// intact instead of forcing every property into the same object shape.
type Properties map[string]json.RawMessage

type SearchRequest struct {
	Keyword string `json:"keyword"`
}

type SearchResponse struct {
	Users   []Properties `json:"users"`
	Devices []Properties `json:"devices"`
}

type DeviceSearchResponse struct {
	Devices []Properties `json:"devices"`
}

// Configuration preserves the full server document, including id, revisionNumber,
// lastUpdated, itsmConfigList and enableUsageDataReporting. Use the latest id and
// revisionNumber when updating, and preserve unrelated application entries.
type Configuration map[string]json.RawMessage

// WebApplication defines the ordered selectors used by the Amplify extension.
// ITSMURL can be a URL pattern accepted by Amplify, so it is not parsed as a
// literal URL. A substitution requires ConfigurationItemRegex.
type WebApplication struct {
	ITSMURL                string `json:"itsmUrl"`
	ConfigurationItem      string `json:"configurationItem"`
	ConfigurationItemRegex string `json:"configurationItemRegex"`
	Substitution           string `json:"substitution"`
}

// ConfigurationRequest replaces the complete ordered application list and usage
// reporting setting. An explicit empty list removes all application entries;
// false disables reporting. Neither value is omitted during serialization.
type ConfigurationRequest struct {
	ITSMConfigList           []WebApplication `json:"itsmConfigList"`
	EnableUsageDataReporting bool             `json:"enableUsageDataReporting"`
}

// InsightsRequest records Amplify extension usage. It does not execute an action.
// Optional pointer fields distinguish false from an omitted browser attribute.
type InsightsRequest struct {
	Action                      string `json:"action"`
	IsSSOEnabled                *bool  `json:"isSsoEnabled,omitempty"`
	ExtensionVersion            string `json:"extensionVersion,omitempty"`
	IsMultiInstanceEnabled      *bool  `json:"isMultiInstanceEnabled,omitempty"`
	DevicePlatform              string `json:"devicePlatform,omitempty"`
	ExtensionOrigin             string `json:"extensionOrigin,omitempty"`
	BrowserName                 string `json:"browserName,omitempty"`
	IsRegistrySetForInstanceURL *bool  `json:"isRegistrySetForInstanceURL,omitempty"`
	ActionTarget                string `json:"actionTarget,omitempty"`
	ActionTargetID              string `json:"actionTargetId,omitempty"`
}
