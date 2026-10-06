package support_insights

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetCrashesResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Insights         *string                    `json:"insights,omitempty"`
}
type GetCPUUsageResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Insights         *string                    `json:"insights,omitempty"`
}
type GetMemoryUsageResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Insights         *string                    `json:"insights,omitempty"`
}

type InsightsRequest struct {
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	StaticEpochDate *int64 `json:"staticEpochDate,omitempty"`
}
