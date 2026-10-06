package remote_actions

import "encoding/json"

type RemoteAction struct {
	UUID      string `json:"uuid"`
	ContentID string `json:"contentId"`
	RemoteActionConfiguration
}

// RemoteActionConfiguration is returned by Update; it excludes the read-only
// UUID and contentId fields supplied by Get.
type RemoteActionConfiguration struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Origin          *string           `json:"origin"`
	Title           string            `json:"title"`
	ContentType     string            `json:"contentType"`
	ProductArea     string            `json:"productArea"`
	Description     string            `json:"description"`
	Purpose         []string          `json:"purpose"`
	Targeting       Targeting         `json:"targeting"`
	TargetingEntity TargetingEntity   `json:"targetingEntity"`
	ScheduledTasks  []json.RawMessage `json:"scheduledTasks"`
	ScriptInfo      ScriptInfo        `json:"scriptInfo"`
}
type Targeting struct {
	ManualEnabled              bool `json:"manualEnabled"`
	APIEnabled                 bool `json:"apiEnabled"`
	ManualAllowMultipleDevices bool `json:"manualAllowMultipleDevices"`
	SchedulingEnabled          bool `json:"schedulingEnabled"`
	EndPointSchedulingEnabled  bool `json:"endPointSchedulingEnabled"`
	WorkflowEnabled            bool `json:"workflowEnabled"`
	SparkEnabled               bool `json:"sparkEnabled"`
}
type TargetingEntity struct {
	DeviceEnabled                        bool    `json:"deviceEnabled"`
	VDISessionEnabled                    bool    `json:"vdiSessionEnabled"`
	VDISessionTargeting                  *string `json:"vdiSessionTargeting"`
	AllowUserOverrideVDISessionTargeting bool    `json:"allowUserOverrideVDISessionTargeting"`
}
type ScriptInfo struct {
	Inputs                   []Input         `json:"inputs"`
	Outputs                  []Output        `json:"outputs"`
	RunAs                    string          `json:"runAs"`
	ScriptMacOS              *Script         `json:"scriptMacOs"`
	ScriptWindows            *Script         `json:"scriptWindows"`
	TimeoutSeconds           int             `json:"timeoutSeconds"`
	ExecutionServiceDelegate json.RawMessage `json:"executionServiceDelegate"`
}

// Script retains the server's encoded script/archive verbatim.
type Script struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}
type Input struct {
	AllowCustomValue bool     `json:"allowCustomValue"`
	Description      string   `json:"description"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Options          []string `json:"options"`
	UsedByMacOS      bool     `json:"usedByMacOs"`
	UsedByWindows    bool     `json:"usedByWindows"`
}
type Output struct {
	Description   string `json:"description"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	UsedByMacOS   bool   `json:"usedByMacOs"`
	UsedByWindows bool   `json:"usedByWindows"`
}
type RemoteActionView struct {
	UUID        string         `json:"uuid"`
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Title       string         `json:"title"`
	ContentType string         `json:"contentType"`
	ScriptInfo  ViewScriptInfo `json:"scriptInfo"`
}
type ViewScriptInfo struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}
type ContentVolume struct {
	RemoteAction             bool `json:"remoteAction"`
	RemoteActionVolume       int  `json:"remoteActionVolume"`
	CurrentRemoteActionCount int  `json:"currentRemoteActionCount"`
}
type ScriptParameters struct {
	Inputs  []ScriptInput  `json:"inputs"`
	Outputs []ScriptOutput `json:"outputs"`
}
type ScriptInput struct {
	ID string `json:"id"`
}
type ScriptOutput struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type SignatureInfo struct {
	Issuer    *string `json:"issuer"`
	State     string  `json:"state"`
	Signature *string `json:"signature"`
}
type GetResponse struct {
	RemoteAction *RemoteAction `json:"remoteActionByUId"`
}
type ViewResponse struct {
	RemoteAction *RemoteActionView `json:"remoteActionByUIdForView"`
}
type VolumeResponse struct {
	ContentVolume *ContentVolume `json:"contentVolumeDetails"`
}
type BashInspectionResponse struct {
	Parameters *ScriptParameters `json:"bashScriptInputsAndOutputs"`
}
type PowerShellInspectionResponse struct {
	Parameters *ScriptParameters `json:"powershellScriptInputsAndOutputs"`
}
type SignatureResponse struct {
	Signature *SignatureInfo `json:"powershellScriptSignatureInfo"`
}

// RemoteActionInput uses an NQL ID, not the UUID accepted by Get. Custom action
// IDs start with #. Script.Script is already base64 encoded, matching the UI.
type RemoteActionInput struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Title           string            `json:"title"`
	ContentType     string            `json:"contentType"`
	ProductArea     string            `json:"productArea"`
	Description     string            `json:"description"`
	Purpose         []string          `json:"purpose"`
	Targeting       Targeting         `json:"targeting"`
	TargetingEntity TargetingEntity   `json:"targetingEntity"`
	ScheduledTasks  []json.RawMessage `json:"scheduledTasks"`
	ScriptInfo      ScriptInfo        `json:"scriptInfo"`
	Origin          *string           `json:"origin,omitempty"`
	LibraryUUID     *string           `json:"libraryUuid,omitempty"`
	AIAgent         json.RawMessage   `json:"aiAgent,omitempty"`
}
type CreatedRemoteAction struct {
	ID        string `json:"id"`
	ContentID string `json:"contentId"`
}
type CreateResponse struct {
	RemoteAction *CreatedRemoteAction `json:"createRemoteAction"`
}
type UpdateResponse struct {
	RemoteAction *RemoteActionConfiguration `json:"updateRemoteAction"`
}
type DeleteResponse struct {
	Deleted bool `json:"deleteRemoteAction"`
}
type ListResponse struct {
	User ListUser    `json:"user"`
	Rows []ListEntry `json:"rows"`
}
type ListUser struct {
	ID string `json:"id"`
}
type ListEntry struct {
	ContentType       string            `json:"contentType"`
	ContentID         string            `json:"contentId"`
	ContentOwner      *string           `json:"contentOwner"`
	Title             string            `json:"title"`
	Active            bool              `json:"active"`
	ResourceName      string            `json:"resourceName"`
	Tags              []json.RawMessage `json:"tags"`
	IsCopyFromLibrary bool              `json:"isCopyFromLibrary"`
	Revision          int               `json:"revision"`
	CreatedBy         *string           `json:"createdBy"`
	UpdatedBy         *string           `json:"updatedBy"`
	Targeting         Targeting         `json:"targeting"`
	NQLID             string            `json:"nqlId"`
	ScriptInfo        ListScriptInfo    `json:"scriptInfo"`
	LastUpdated       int64             `json:"lastUpdated"`
}
type ListScriptInfo struct {
	RunAs          string          `json:"runAs"`
	Outputs        []Output        `json:"outputs"`
	Inputs         []ListInput     `json:"inputs"`
	TimeoutSeconds int             `json:"timeoutSeconds"`
	ScriptWindows  *ScriptMetadata `json:"scriptWindows,omitempty"`
	ScriptMacOS    *ScriptMetadata `json:"scriptMacOs,omitempty"`
}
type ListInput struct {
	Input
	Type *string `json:"type,omitempty"`
}
type ScriptMetadata struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type ExportRequest struct {
	ID string `json:"id"`
}
type ExportedRemoteAction struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	Purpose         []string              `json:"purpose"`
	AIAgent         *AIAgentConfiguration `json:"aiAgent"`
	Targeting       Targeting             `json:"targeting"`
	TargetingEntity TargetingEntity       `json:"targetingEntity"`
	ScheduledTasks  []json.RawMessage     `json:"scheduledTasks"`
	ScriptInfo      ScriptInfo            `json:"scriptInfo"`
}
type AIAgentConfiguration struct {
	Category             *string        `json:"category"`
	SafetyClassification *string        `json:"safetyClassification"`
	UserConsentRequired  *bool          `json:"userConsentRequired"`
	Impact               *AIAgentImpact `json:"impact"`
}
type AIAgentImpact struct {
	PreApproval    *string `json:"preApproval"`
	PostApproval   *string `json:"postApproval"`
	AfterExecution *string `json:"afterExecution"`
}
type ExportResponse struct {
	RemoteAction *ExportedRemoteAction `json:"remoteAction"`
}
type ImportRequest struct {
	RemoteAction RemoteActionInput `json:"remoteAction"`
}
type ImportedRemoteAction struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ImportResponse struct {
	RemoteAction *ImportedRemoteAction `json:"createRemoteAction"`
}
type GetFromLibraryRequest struct {
	ID string `json:"id"`
}
type GetFromLibraryResponse struct {
	RemoteAction *LibraryRemoteAction `json:"remoteActionFromLibrary"`
}

// LibraryRemoteAction retains an unset title before a template is installed.
type LibraryRemoteAction struct {
	RemoteActionConfiguration
	Title *string `json:"title"`
}
