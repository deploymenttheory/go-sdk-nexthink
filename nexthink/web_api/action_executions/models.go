package action_executions

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type ListRemoteActionsResponseItemTargeting struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ManualEnabled    *bool                      `json:"manualEnabled,omitempty"`
}
type ListRemoteActionsResponseItemScriptInfo struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Inputs           *[]RemoteActionInput       `json:"inputs,omitempty"`
	Outputs          *[]RemoteActionOutput      `json:"outputs,omitempty"`
	HasScriptMacOS   *bool                      `json:"hasScriptMacOs,omitempty"`
	HasScriptWindows *bool                      `json:"hasScriptWindows,omitempty"`
	RunAs            *string                    `json:"runAs,omitempty"`
}
type ListRemoteActionsResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage               `json:"-"`
	ID               *string                                  `json:"id,omitempty"`
	Name             *string                                  `json:"name,omitempty"`
	Targeting        *ListRemoteActionsResponseItemTargeting  `json:"targeting,omitempty"`
	ScriptInfo       *ListRemoteActionsResponseItemScriptInfo `json:"scriptInfo,omitempty"`
	Description      *string                                  `json:"description,omitempty"`
	Purpose          *[]string                                `json:"purpose,omitempty"`
	TargetingEntity  *RemoteActionTargetingEntity             `json:"targetingEntity,omitempty"`
}
type ListRemoteActionsResponse []ListRemoteActionsResponseItem
type GetRemoteActionResponseTargeting struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ManualEnabled    *bool                      `json:"manualEnabled,omitempty"`
}
type GetRemoteActionResponseScriptInfo struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Inputs           *[]RemoteActionInput       `json:"inputs,omitempty"`
	Outputs          *[]RemoteActionOutput      `json:"outputs,omitempty"`
	HasScriptMacOS   *bool                      `json:"hasScriptMacOs,omitempty"`
	HasScriptWindows *bool                      `json:"hasScriptWindows,omitempty"`
	RunAs            *string                    `json:"runAs,omitempty"`
}
type GetRemoteActionResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage         `json:"-"`
	ID               *string                            `json:"id,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Targeting        *GetRemoteActionResponseTargeting  `json:"targeting,omitempty"`
	ScriptInfo       *GetRemoteActionResponseScriptInfo `json:"scriptInfo,omitempty"`
	Description      *string                            `json:"description,omitempty"`
	Purpose          *[]string                          `json:"purpose,omitempty"`
	TargetingEntity  *RemoteActionTargetingEntity       `json:"targetingEntity,omitempty"`
}
type ListRemoteActionsForQueryResponseItemTargeting struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ManualEnabled    *bool                      `json:"manualEnabled,omitempty"`
}
type ListRemoteActionsForQueryResponseItemScriptInfo struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Inputs           *[]RemoteActionInput       `json:"inputs,omitempty"`
	Outputs          *[]RemoteActionOutput      `json:"outputs,omitempty"`
	HasScriptMacOS   *bool                      `json:"hasScriptMacOs,omitempty"`
	HasScriptWindows *bool                      `json:"hasScriptWindows,omitempty"`
	RunAs            *string                    `json:"runAs,omitempty"`
}
type ListRemoteActionsForQueryResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                       `json:"-"`
	ID               *string                                          `json:"id,omitempty"`
	Name             *string                                          `json:"name,omitempty"`
	Targeting        *ListRemoteActionsForQueryResponseItemTargeting  `json:"targeting,omitempty"`
	ScriptInfo       *ListRemoteActionsForQueryResponseItemScriptInfo `json:"scriptInfo,omitempty"`
	Description      *string                                          `json:"description,omitempty"`
	Purpose          *[]string                                        `json:"purpose,omitempty"`
	TargetingEntity  *RemoteActionTargetingEntity                     `json:"targetingEntity,omitempty"`
}
type ListRemoteActionsForQueryResponse []ListRemoteActionsForQueryResponseItem
type ExecuteResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	RequestID        *string                    `json:"requestId,omitempty"`
}
type ListActionsResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	Name             *string                    `json:"name,omitempty"`
	ActionType       *string                    `json:"actionType,omitempty"`
}
type ListActionsResponse []ListActionsResponseItem
type GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Type             *string                    `json:"type,omitempty"`
	Value            json.RawMessage            `json:"value,omitempty"`
}
type GetDeviceHistoryResponseItemExecutionDetailsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                                     `json:"-"`
	Status           *string                                                        `json:"status,omitempty"`
	StatusDetails    *GetDeviceHistoryResponseItemExecutionDetailsItemStatusDetails `json:"statusDetails,omitempty"`
	Time             *string                                                        `json:"time,omitempty"`
	DeviceID         *GetDeviceHistoryResponseItemExecutionDetailsItemDeviceID      `json:"deviceId,omitempty"`
}
type GetDeviceHistoryResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                          `json:"-"`
	ID               *string                                             `json:"id,omitempty"`
	ActionType       *string                                             `json:"actionType,omitempty"`
	ExecutionDetails *[]GetDeviceHistoryResponseItemExecutionDetailsItem `json:"executionDetails,omitempty"`
}
type GetDeviceHistoryResponse []GetDeviceHistoryResponseItem

type ListOptions struct {
	IsTargetingManualEnabled                *bool `json:"isTargetingManualEnabled,omitempty"`
	IsTargetingManualAllowMultipleDevices   *bool `json:"isTargetingManualAllowMultipleDevices,omitempty"`
	IsTargetingEntityVDISession             *bool `json:"isTargetingEntityVDISession,omitempty"`
	IsTargetingEntityVDISessionTargetingVDI *bool `json:"isTargetingEntityVDISessionTargetingVDI,omitempty"`
	IsTargetingEntityDevice                 *bool `json:"isTargetingEntityDevice,omitempty"`
	HasScriptWindows                        *bool `json:"hasScriptWindows,omitempty"`
	HasScriptMacOS                          *bool `json:"hasScriptMacOs,omitempty"`
}
type DetailsRequest struct {
	NQLID      string `json:"nql-id"`
	SourceType string `json:"source-type,omitempty"`
}
type NQLRequest struct {
	Platforms   []string `json:"platforms,omitempty"`
	NQLQuery    string   `json:"nqlQuery"`
	NQLTimeZone string   `json:"nqlTimeZone"`
	NQLTimeNow  string   `json:"nqlTimeNow"`
}
type ExecuteRequest struct {
	// Source is the UI nx-source attribution header.
	Source         string            `json:"-"`
	RemoteActionID string            `json:"remoteActionId"`
	Params         map[string]string `json:"params,omitempty"`
	Targets        []Target          `json:"targets,omitempty"`
	TriggerInfo    TriggerInfo       `json:"triggerInfo,omitempty"`
	NQLQuery       string            `json:"nqlQuery,omitempty"`
	NQLTimeZone    string            `json:"nqlTimeZone,omitempty"`
	NQLTimeNow     string            `json:"nqlTimeNow,omitempty"`
}
type HistoryRequest struct {
	Actions []ActionReference `json:"actions"`
}
type Target struct {
	DeviceCollectorUID string `json:"deviceCollectorUid,omitempty"`
	UserSID            string `json:"userSid,omitempty"`
}
type TriggerInfo struct {
	InternalSource   string `json:"internalSource"`
	ResolutionPlanID string `json:"resolutionPlanId,omitempty"`
	ResolutionStepID string `json:"resolutionStepId,omitempty"`
}
type ActionReference struct {
	ID         string `json:"id"`
	ActionType string `json:"actionType"`
}

type RemoteActionInput struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	Name             *string                    `json:"name,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	AllowCustomValue *bool                      `json:"allowCustomValue,omitempty"`
	Options          *[]string                  `json:"options,omitempty"`
	UsedByMacOS      *bool                      `json:"usedByMacOs,omitempty"`
	UsedByWindows    *bool                      `json:"usedByWindows,omitempty"`
}

type RemoteActionOutput struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	Name             *string                    `json:"name,omitempty"`
	Description      *string                    `json:"description,omitempty"`
	Type             *string                    `json:"type,omitempty"`
}

type RemoteActionTargetingEntity struct {
	AdditionalFields                     map[string]json.RawMessage `json:"-"`
	DeviceEnabled                        *bool                      `json:"deviceEnabled,omitempty"`
	VDISessionEnabled                    *bool                      `json:"vdiSessionEnabled,omitempty"`
	VDISessionTargeting                  *string                    `json:"vdiSessionTargeting,omitempty"`
	AllowUserOverrideVDISessionTargeting *bool                      `json:"allowUserOverrideVDISessionTargeting,omitempty"`
}
