package dashboards

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

type Duration struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}
type TimeRangeInput struct {
	Duration
	ByDuration *Duration `json:"byDuration,omitempty"`
}
type TimeRange struct {
	Duration
	ByDuration *Duration `json:"byDuration"`
}
type DashboardInput struct {
	Title            string         `json:"title"`
	DefaultTimeRange TimeRangeInput `json:"defaultTimeRange"`
}
type UpdateRequest struct {
	ID            string         `json:"id"`
	Revision      int            `json:"revision"`
	ProductArea   string         `json:"productArea,omitempty"`
	DashboardType string         `json:"dashboardType,omitempty"`
	Dashboard     DashboardInput `json:"dashboard"`
}
type DeleteRequest struct {
	DashboardID   string `json:"dashboardId"`
	Revision      int    `json:"revision"`
	ProductArea   string `json:"productArea,omitempty"`
	DashboardType string `json:"dashboardType,omitempty"`
}
type GetOptions struct{ ProductArea string }
type Meta struct {
	ProductArea string `json:"productArea"`
	Type        string `json:"type"`
}
type ReadMeta struct {
	Meta
	BuiltinContentMetadata json.RawMessage `json:"builtinContentMetadata"`
}
type CreatedDashboard struct {
	ID               string    `json:"id"`
	Revision         int       `json:"revision"`
	Title            string    `json:"title"`
	DefaultTimeRange TimeRange `json:"defaultTimeRange"`
	Meta             Meta      `json:"meta"`
}
type UpdatedDashboard struct {
	CreatedDashboard
	Description *string `json:"description"`
}
type UserPermissions struct {
	Write bool `json:"write"`
}
type AllowedOperations struct {
	Clone  bool `json:"clone"`
	Delete bool `json:"delete"`
	Edit   bool `json:"edit"`
	Export bool `json:"export"`
	Share  bool `json:"share"`
}

// Polymorphic widget, layout and filter documents retain all fields selected by the UI.
type Dashboard struct {
	ID                string            `json:"id"`
	Title             string            `json:"title"`
	Revision          int               `json:"revision"`
	DefaultTimeRange  TimeRange         `json:"defaultTimeRange"`
	Description       *string           `json:"description"`
	FixedTimeRange    bool              `json:"fixedTimeRange"`
	Widgets           []json.RawMessage `json:"widgets"`
	Layout            json.RawMessage   `json:"layout"`
	Meta              ReadMeta          `json:"meta"`
	Filters           []json.RawMessage `json:"filters"`
	Tabs              []Tab             `json:"tabs"`
	UserPermissions   UserPermissions   `json:"userPermissions"`
	AllowedOperations AllowedOperations `json:"allowedOperations"`
}
type Tab struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Widgets     []json.RawMessage `json:"widgets"`
	Layout      json.RawMessage   `json:"layout"`
}
type CreateResponse struct {
	Dashboard *CreatedDashboard `json:"createDashboard"`
}
type UpdateResponse struct {
	Dashboard *UpdatedDashboard `json:"updateDashboard"`
}
type GetResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}
type DeleteResponse struct {
	ID *string `json:"deleteDashboard"`
}
type Summary struct {
	content_administration.Content
	ProductArea string `json:"productArea"`
	ModifiedAt  int64  `json:"modifiedAt"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
