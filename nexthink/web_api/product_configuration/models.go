package product_configuration

import "encoding/json"

// InstanceConfiguration maps a configuration key to its value. Observed values
// include boolean Assist settings and the strings "enabled"/"disabled" for AI tools.
type InstanceConfiguration map[string]json.RawMessage
type InstanceResponse struct {
	Data InstanceConfiguration `json:"data"`
}
