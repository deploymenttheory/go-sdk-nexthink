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

// MutationContext selects the current revision and optional tab of a dashboard.
type MutationContext struct {
	DashboardID   string `json:"dashboardId"`
	Revision      int    `json:"revision"`
	ProductArea   string `json:"productArea,omitempty"`
	DashboardType string `json:"dashboardType,omitempty"`
	TabID         string `json:"tabId,omitempty"`
}
type DashboardContext struct {
	ID          string `json:"id"`
	Revision    int    `json:"revision"`
	ProductArea string `json:"productArea,omitempty"`
	Type        string `json:"type,omitempty"`
	TabID       string `json:"tabId,omitempty"`
}

// Config is a discriminated union selected by Type; its complete JSON is retained.
type WidgetInput struct {
	ID     string          `json:"id"`
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}
type FilterInput struct {
	ID     string          `json:"id"`
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}
type LayoutInput struct {
	Version int               `json:"version"`
	Nodes   []LayoutNodeInput `json:"nodes"`
}
type LayoutNodeInput struct {
	Leaf   *LayoutLeafInput   `json:"leaf,omitempty"`
	Parent *LayoutParentInput `json:"parent,omitempty"`
}
type LayoutLeafInput struct {
	ID       string  `json:"id"`
	ParentID *string `json:"parentId"`
	WidgetID string  `json:"widgetId"`
}
type LayoutParentInput struct {
	ID            string  `json:"id"`
	ParentID      *string `json:"parentId"`
	FirstChildID  string  `json:"firstChildId"`
	SecondChildID string  `json:"secondChildId"`
	Direction     string  `json:"direction"`
	Split         float64 `json:"split"`
	Grouped       bool    `json:"grouped"`
}
type WidgetRequest struct {
	MutationContext
	Widget WidgetInput `json:"widget"`
	Layout LayoutInput `json:"layout"`
}

// Omit Layout when deleting the last widget. An empty nodes array is rejected by the API.
type DeleteWidgetRequest struct {
	MutationContext
	ID     string       `json:"id"`
	Layout *LayoutInput `json:"layout,omitempty"`
}
type FilterRequest struct {
	MutationContext
	Filter FilterInput `json:"filter"`
}
type DeleteFilterRequest struct {
	Dashboard DashboardContext `json:"dashboard"`
	FilterID  string           `json:"filterId"`
}
type TabInput struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}
type TabUpdateInput struct {
	ID string `json:"id"`
	TabInput
}
type UpdateTabRequest struct {
	MutationContext
	Tab TabInput `json:"tab"`
}
type UpdateTabsRequest struct {
	MutationContext
	Tabs []TabUpdateInput `json:"tabs"`
}
type UpdateLayoutRequest struct {
	ID            string      `json:"id"`
	Revision      int         `json:"revision"`
	ProductArea   string      `json:"productArea"`
	DashboardType string      `json:"dashboardType"`
	TabID         string      `json:"tabId,omitempty"`
	Layout        LayoutInput `json:"layout"`
}
type DuplicateRequest struct {
	Dashboard DashboardContext `json:"dashboard"`
}
type ExportedQuery struct {
	QueryID string `json:"queryId"`
	Query   string `json:"query"`
}
type ExportDocument struct {
	ExportedAt string          `json:"exportedAt"`
	Dashboard  json.RawMessage `json:"dashboard"`
	Queries    []ExportedQuery `json:"queries"`
}
type ImportRequest struct {
	Content ExportDocument `json:"content"`
}

// Operation-specific results retain their exact selection sets.
type RevisionResult struct {
	ID       string `json:"id"`
	Revision int    `json:"revision"`
}
type WidgetResult struct {
	RevisionResult
	Description *string           `json:"description"`
	Widgets     []json.RawMessage `json:"widgets"`
	Layout      json.RawMessage   `json:"layout"`
	Tabs        []Tab             `json:"tabs"`
}
type DeleteWidgetResult struct {
	RevisionResult
	Description *string         `json:"description"`
	Layout      json.RawMessage `json:"layout"`
	Tabs        []Tab           `json:"tabs"`
}
type FilterResult struct {
	RevisionResult
	Filters []json.RawMessage `json:"filters"`
}
type TabResult struct {
	RevisionResult
	Description *string `json:"description"`
	Meta        Meta    `json:"meta"`
	Tabs        []Tab   `json:"tabs"`
}
type DeleteTabResult struct {
	TabResult
	Widgets []json.RawMessage `json:"widgets"`
	Layout  json.RawMessage   `json:"layout"`
}
type DuplicateResult struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
}
type ImportedDashboard struct {
	ID                string            `json:"id"`
	Title             string            `json:"title"`
	ModifiedAt        string            `json:"modifiedAt"`
	AllowedOperations AllowedOperations `json:"allowedOperations"`
	Meta              Meta              `json:"meta"`
}
type CreateWidgetResponse struct {
	Dashboard *WidgetResult `json:"createWidget"`
}
type UpdateWidgetResponse struct {
	Dashboard *WidgetResult `json:"updateWidget"`
}
type DeleteWidgetResponse struct {
	Dashboard *DeleteWidgetResult `json:"deleteWidget"`
}
type CreateFilterResponse struct {
	Dashboard *FilterResult `json:"createFilter"`
}
type UpdateFilterResponse struct {
	Dashboard *FilterResult `json:"updateFilter"`
}
type DeleteFilterResponse struct {
	Dashboard *RevisionResult `json:"deleteFilter"`
}
type CreateTabResponse struct {
	Dashboard *TabResult `json:"createTab"`
}
type UpdateTabResponse struct {
	Dashboard *TabResult `json:"updateTab"`
}
type UpdateTabsResponse struct {
	Dashboard *TabResult `json:"updateTabs"`
}
type DeleteTabResponse struct {
	Dashboard *DeleteTabResult `json:"deleteTab"`
}
type UpdateLayoutResponse struct {
	Dashboard *LayoutResult `json:"updateLayout"`
}
type ExportResponse struct {
	Dashboard *ExportDocument `json:"dashboardExport"`
}
type DuplicateResponse struct {
	Dashboard *DuplicateResult `json:"cloneDashboard"`
}
type ImportResponse struct {
	Dashboard *ImportedDashboard `json:"importDashboard"`
}

type LayoutResult struct {
	ID                string            `json:"id"`
	Title             string            `json:"title"`
	Revision          int               `json:"revision"`
	DefaultTimeRange  TimeRange         `json:"defaultTimeRange"`
	Description       *string           `json:"description"`
	Widgets           []json.RawMessage `json:"widgets"`
	Layout            json.RawMessage   `json:"layout"`
	Meta              ReadMeta          `json:"meta"`
	Filters           []json.RawMessage `json:"filters"`
	Tabs              []Tab             `json:"tabs"`
	UserPermissions   UserPermissions   `json:"userPermissions"`
	AllowedOperations AllowedOperations `json:"allowedOperations"`
}
