package workspace_tasks

import "encoding/json"

// Document preserves all fields of polymorphic UI responses, including future fields.
type Document map[string]json.RawMessage

// Task exposes stable fields and retains unknown fields when round-tripping JSON.
type Task struct {
	ID               string                     `json:"id,omitempty"`
	Title            string                     `json:"title,omitempty"`
	Prompt           string                     `json:"prompt,omitempty"`
	Enabled          bool                       `json:"enabled,omitempty"`
	Status           string                     `json:"status,omitempty"`
	TimeZone         string                     `json:"time_zone,omitempty"`
	ContentType      string                     `json:"contentType,omitempty"`
	Origin           string                     `json:"origin,omitempty"`
	Schedule         json.RawMessage            `json:"schedule,omitempty"`
	ExpiresAt        *string                    `json:"expires_at,omitempty"`
	CreatedAt        string                     `json:"created_at,omitempty"`
	UpdatedAt        string                     `json:"updated_at,omitempty"`
	SkillID          *string                    `json:"skill_id,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
	present          map[string]json.RawMessage
}
type TaskSchedule struct {
	Type            string   `json:"type"`
	Days            []string `json:"days,omitempty"`
	Time            string   `json:"time,omitempty"`
	IntervalMinutes int      `json:"interval_minutes,omitempty"`
	Datetime        string   `json:"datetime,omitempty"`
}
type TaskRequest struct {
	Title    string       `json:"title"`
	Prompt   string       `json:"prompt"`
	Enabled  bool         `json:"enabled"`
	Schedule TaskSchedule `json:"schedule"`
	// ExpiresAt is a JSON timestamp string, null to clear expiry, or nil to omit.
	ExpiresAt json.RawMessage `json:"expires_at,omitempty"`
	SkillID   *string         `json:"skill_id"`
}
type TaskListResponse struct {
	Items       []Task `json:"items"`
	Page        int    `json:"page"`
	Size        int    `json:"size"`
	ActiveCount *int   `json:"active_count,omitempty"`
	MaxActive   *int   `json:"max_active,omitempty"`
}
type TaskListOptions struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
