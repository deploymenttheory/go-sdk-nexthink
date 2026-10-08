package device_classification

import "encoding/json"

// Ruleset is the metadata returned for an uploaded device classification CSV.
type Ruleset struct {
	Name             string                     `json:"name,omitempty"`
	Description      string                     `json:"description,omitempty"`
	Filename         string                     `json:"filename,omitempty"`
	LastModified     int64                      `json:"lastModified,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
	present          map[string]json.RawMessage
}

// RulesetUpload sends name and base64 UTF-8 description as headers and CSV as a multipart file.
// Set CSV to nil for a metadata-only update; create requires a nonempty file.
// In example JSON, CSV bytes use standard JSON base64 encoding.
type RulesetUpload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename,omitempty"`
	CSV         []byte `json:"csv,omitempty"`
}
type GeoIPConfiguration struct {
	PublicIP           bool `json:"platform.geoipoptin.public-ip-opt-in"`
	Country            bool `json:"platform.geoipoptin.country-opt-in"`
	CountrySubdivision bool `json:"platform.geoipoptin.country-subdivision-opt-in"`
	City               bool `json:"platform.geoipoptin.city-opt-in"`
}
