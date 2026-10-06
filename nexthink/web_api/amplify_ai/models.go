package amplify_ai

import "encoding/json"

// AnalysisRequest matches the Amplify extension's generation request. Generating
// an analysis may consume AI capacity; it is not a read-only lookup.
type AnalysisRequest struct {
	TicketID          string         `json:"ticketId,omitempty"`
	DeviceID          string         `json:"deviceId"`
	ForceRegeneration bool           `json:"forceRegeneration"`
	OnTheFly          *TicketContext `json:"onTheFly,omitempty"`
}
type TicketContext struct {
	Title          string `json:"title,omitempty"`
	Description    string `json:"description,omitempty"`
	TicketNumber   string `json:"ticketNumber,omitempty"`
	ServiceNowHost string `json:"serviceNowHost,omitempty"`
}
type Feedback struct {
	Rating  int    `json:"rating"`
	Message string `json:"message,omitempty"`
}
type FeedbackRequest struct {
	Feedback         Feedback `json:"feedback"`
	ResolutionPlanID string   `json:"resolutionPlanId"`
}
type Metric struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MetricRequest struct {
	Metric Metric `json:"metric"`
}
type TicketRetrievalDurationRequest struct {
	DurationSeconds float64 `json:"durationSeconds"`
	Status          string  `json:"status"`
}

// ExecuteActionRequest preserves the two action payload variants in the extension:
// actionExecutionParams with resolution identifiers, or remote-action widget
// parameters (remoteActionId, deviceId, params, targets and optional NQL targeting).
// Both use resolutionPlanId and resolutionStepId. Action inputs vary by content.
type ExecuteActionRequest map[string]json.RawMessage

type ExecuteUserActionRequest struct {
	ResolutionPlanID string `json:"resolutionPlanId"`
	ResolutionStepID string `json:"resolutionStepId"`
}
type RefreshResolutionStepRequest struct {
	DeviceID string `json:"device_id"`
	ID       string `json:"id"`
}
type UpdateResolutionStepRequest struct {
	Status       string  `json:"status"`
	ID           string  `json:"id"`
	ClosingNotes *string `json:"closing_notes,omitempty"`
}
type ResolveTicketRequest struct {
	SysID string `json:"sysId"`
}
