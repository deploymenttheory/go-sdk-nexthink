package workspace_assignments

import "encoding/json"

// Document preserves all fields of polymorphic UI responses, including future fields.
type Document map[string]json.RawMessage

// Assignment exposes stable fields and retains unknown fields when round-tripping JSON.
type Assignment struct {
	ID               string                     `json:"id,omitempty"`
	Type             string                     `json:"type,omitempty"`
	Source           string                     `json:"source,omitempty"`
	State            string                     `json:"state,omitempty"`
	Assignee         json.RawMessage            `json:"assignee,omitempty"`
	ConversationID   *string                    `json:"conversationId,omitempty"`
	Title            string                     `json:"title,omitempty"`
	Metadata         map[string]string          `json:"metadata,omitempty"`
	CreatedAt        string                     `json:"createdAt,omitempty"`
	UpdatedAt        string                     `json:"updatedAt,omitempty"`
	Read             bool                       `json:"read,omitempty"`
	ReadAt           *string                    `json:"readAt,omitempty"`
	Revision         int                        `json:"revision,omitempty"`
	PriorityScore    *float64                   `json:"priorityScore,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
	present          map[string]json.RawMessage
}
type AssignmentOptions struct {
	Sources []string `json:"sources,omitempty"`
	Sort    string   `json:"sort,omitempty"`
}
type AssignmentUpdate struct {
	State      string          `json:"state,omitempty"`
	AssigneeID json.RawMessage `json:"assigneeId,omitempty"`
	Revision   int             `json:"revision"`
}
type AssignmentResponse struct {
	Item      *Assignment `json:"item"`
	ErrorCode string      `json:"errorCode,omitempty"`
}
type UnreadCountResponse struct {
	UnreadCount int `json:"unreadCount"`
}
