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

// Configuration preserves the extension configuration document without guessing
// fields beyond the itsmConfigList consumed by the extension.
type Configuration map[string]json.RawMessage

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
