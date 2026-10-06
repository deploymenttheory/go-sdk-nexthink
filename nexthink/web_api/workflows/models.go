package workflows

import "encoding/json"

// WorkflowSummary is returned by List; Get supplies the complete management view.
type WorkflowSummary struct {
	UUID           string         `json:"uuid"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Status         string         `json:"status"`
	Copy           bool           `json:"copy"`
	Builtin        bool           `json:"builtin"`
	LastUpdateTime string         `json:"lastUpdateTime"`
	TriggerMethods TriggerMethods `json:"triggerMethods"`
	Versions       []Version      `json:"versions"`
}
type Workflow struct {
	UUID           string          `json:"uuid"`
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Status         string          `json:"status"`
	Description    *string         `json:"description"`
	LastUpdateTime string          `json:"lastUpdateTime"`
	TriggerMethods TriggerMethods  `json:"triggerMethods"`
	Versions       []Version       `json:"versions"`
	ScheduledTasks []ScheduledTask `json:"scheduledTasks"`
	EventTriggers  []EventTrigger  `json:"eventTriggers"`
	Targets        *Targets        `json:"targets"`
}
type TriggerMethods struct {
	ManualEnabled                bool `json:"manualEnabled"`
	ManualAllowMultipleInstances bool `json:"manualAllowMultipleInstances"`
	APIEnabled                   bool `json:"apiEnabled"`
	SchedulingEnabled            bool `json:"schedulingEnabled"`
	EventEnabled                 bool `json:"eventEnabled"`
}
type Version struct {
	Description *string     `json:"description"`
	Status      string      `json:"status"`
	Version     int         `json:"version"`
	Valid       bool        `json:"valid"`
	Definition  string      `json:"definition"`
	Parameters  []Parameter `json:"parameters"`
}
type Parameter struct {
	ID               string   `json:"id"`
	AllowCustomValue bool     `json:"allowCustomValue"`
	Description      *string  `json:"description"`
	Options          []string `json:"options"`
}

// ScheduledTask and EventTrigger preserve unvalidated nested value schemas.
type (
	ScheduledTask map[string]json.RawMessage
	EventTrigger  map[string]json.RawMessage
	Targets       struct {
		Device     *Target           `json:"device"`
		User       *Target           `json:"user"`
		VDISession *VDISessionTarget `json:"vdiSession"`
	}
)

type Target struct {
	Enabled bool `json:"enabled"`
}
type VDISessionTarget struct {
	Enabled           bool   `json:"enabled"`
	AllowUserOverride bool   `json:"allowUserOverride"`
	Type              string `json:"type"`
}
type ListResponse struct {
	Workflows []WorkflowSummary `json:"workflows"`
}
type GetResponse struct {
	Workflow *Workflow `json:"workflow"`
}

// ExportResponse contains the opaque export string returned by the service.
// Its encoding is intentionally preserved for compatible import tooling.
type ExportResponse struct {
	Content *string `json:"exportWorkflowById"`
}

// WorkflowInput is the UI management write contract. ID is the hash-prefixed
// NQL ID; UUID identifies an existing workflow for Update. Definition is XML.
// Server-only fields such as lastUpdateTime and version.valid are excluded.
type WorkflowInput struct {
	UUID           string          `json:"uuid,omitempty"`
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Description    *string         `json:"description"`
	Status         string          `json:"status"`
	TriggerMethods TriggerMethods  `json:"triggerMethods"`
	Versions       []VersionInput  `json:"versions"`
	ScheduledTasks []ScheduledTask `json:"scheduledTasks"`
	EventTriggers  []EventTrigger  `json:"eventTriggers"`
	Targets        *Targets        `json:"targets,omitempty"`
}
type VersionInput struct {
	Description *string     `json:"description"`
	Status      string      `json:"status"`
	Version     int         `json:"version"`
	Definition  string      `json:"definition"`
	Parameters  []Parameter `json:"parameters"`
}
type CreateRequest struct {
	Workflow           WorkflowInput `json:"WorkflowInput"`
	BuiltinLibraryUUID *string       `json:"builtinLibraryUuid,omitempty"`
}
type CreateResponse struct {
	Workflow *Workflow `json:"createWorkflow"`
}
type UpdateResponse struct {
	Workflow *Workflow `json:"updateWorkflow"`
}
type DeleteResponse struct {
	Deleted bool `json:"deleteWorkflow"`
}

// ConnectorDefinition describes a built-in workflow connector and its actions.
type ConnectorDefinition struct {
	ConnectorID string                       `json:"connectorId"`
	Enabled     bool                         `json:"enabled"`
	Labels      map[string]map[string]string `json:"labels"`
	IconID      string                       `json:"iconId"`
	Actions     []ConnectorAction            `json:"actions"`
}
type ConnectorAction struct {
	ActionID   string                       `json:"actionId"`
	Enabled    bool                         `json:"enabled"`
	Labels     map[string]map[string]string `json:"labels"`
	Parameters []ConnectorParameter         `json:"parameters"`
	Outputs    *[]ConnectorOutput           `json:"outputs,omitempty"`
}
type ConnectorParameter struct {
	ParameterID string                       `json:"parameterId"`
	Labels      map[string]map[string]string `json:"labels"`
}
type ConnectorOutput struct {
	OutputID string                       `json:"outputId"`
	Labels   map[string]map[string]string `json:"labels"`
}

// ConnectorCredential is the workflow-specific view of generic third-party credentials.
// Manage their lifecycle through WebAPI.ConnectorCredentials.
type ConnectorCredential struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	HasAuthenticationInBody bool   `json:"hasAuthenticationInBody"`
}

type GetFromLibraryRequest struct {
	ContentID string `json:"contentId"`
}
type GetFromLibraryResponse struct {
	Workflow *LibraryWorkflow `json:"copyWorkflow"`
}
type SetActiveRequest struct {
	UUID     string `json:"uuid"`
	Activate bool   `json:"activate"`
}
type SetActiveResponse struct {
	Workflow *Workflow `json:"activateWorkflow"`
}

// ImportRequest.Content is the exported workflow JSON text. The service requires
// workflow.versions even for an empty list; an empty workflow export can omit it.
type ImportRequest struct {
	Content string `json:"content"`
}
type ImportedWorkflow struct {
	ID string `json:"id"`
}
type ImportResponse struct {
	Workflow *ImportedWorkflow `json:"importWorkflow"`
}

// LibraryWorkflow retains nullable persistence metadata for a template not yet installed.
type LibraryWorkflow struct {
	Workflow
	UUID           *string `json:"uuid"`
	LastUpdateTime *string `json:"lastUpdateTime"`
}
