package vdi

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetGlobalHealthResponse string
type ValidateHostnameResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Valid            *bool                      `json:"valid,omitempty"`
}

type TimelineRequest struct {
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	BucketSizeInSeconds int    `json:"bucketSizeInSeconds"`
}
type HealthRequest struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type HostnameRequest struct {
	Hostname string `json:"hostname"`
}

// TimelineSection is keyed by metric/group ID. Result is either a time-series object or an application-series array in the UI contract.
type TimelineSection struct {
	Result json.RawMessage            `json:"result"`
	Meta   map[string]json.RawMessage `json:"meta"`
}
type GetSessionTimelineResponse map[string]TimelineSection
type GetHypervisorTimelineResponse map[string]TimelineSection

// TimelineSeries models the common object alternative of TimelineSection.Result.
type TimelineSeries struct {
	TimeSeries []TimelineBucket `json:"timeSeries"`
}
type TimelineBucket struct {
	StartTime string                     `json:"startTime"`
	EndTime   string                     `json:"endTime"`
	Data      map[string]json.RawMessage `json:"data"`
}
