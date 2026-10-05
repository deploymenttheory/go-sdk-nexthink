package license

import "encoding/json"

type FeatureStatusResponse struct {
	Data FeatureStatus `json:"data"`
}
type FeatureStatus struct {
	Enabled bool            `json:"enabled"`
	Quota   json.RawMessage `json:"quota"`
}
