package collaboration_tools

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetCallInsightsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Insights         *string                    `json:"insights,omitempty"`
}

type CallInsightsRequest struct {
	ApplicationType string `json:"applicationType"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
}
