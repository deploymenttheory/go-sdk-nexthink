package workflow_executions

import "encoding/json"

// Polymorphic NQL values, event payloads and plugin-defined details retain their exact JSON.
type GetTimelineResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields         map[string]json.RawMessage `json:"-"`
	ID                       *string                    `json:"id,omitempty"`
	Name                     *string                    `json:"name,omitempty"`
	Type                     *string                    `json:"type,omitempty"`
	Status                   *string                    `json:"status,omitempty"`
	StartTime                *string                    `json:"startTime,omitempty"`
	EndTime                  *string                    `json:"endTime,omitempty"`
	Details                  json.RawMessage            `json:"details,omitempty"`
	ExtraInformation         json.RawMessage            `json:"extraInformation,omitempty"`
	RepeatingActionClickable *bool                      `json:"repeatingActionClickable,omitempty"`
}
type GetTimelineResponse []GetTimelineResponseItem
type GetTimelineV2ResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields         map[string]json.RawMessage `json:"-"`
	ID                       *string                    `json:"id,omitempty"`
	Name                     *string                    `json:"name,omitempty"`
	Type                     *string                    `json:"type,omitempty"`
	Status                   *string                    `json:"status,omitempty"`
	StartTime                *string                    `json:"startTime,omitempty"`
	EndTime                  *string                    `json:"endTime,omitempty"`
	Details                  json.RawMessage            `json:"details,omitempty"`
	ExtraInformation         json.RawMessage            `json:"extraInformation,omitempty"`
	RepeatingActionClickable *bool                      `json:"repeatingActionClickable,omitempty"`
}
type GetTimelineV2Response []GetTimelineV2ResponseItem
type ListActivitiesResponseItem struct {
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ActivityName     *string                    `json:"activityName,omitempty"`
	ActivityType     *string                    `json:"activityType,omitempty"`
	StartTime        *string                    `json:"startTime,omitempty"`
	TransactionOrder *int                       `json:"transactionOrder,omitempty"`
}
type ListActivitiesResponse []ListActivitiesResponseItem
type GetRemoteActionDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type GetCustomFieldsDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type GetCampaignDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type GetFunctionDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type GetMessageDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type GetSAPIDetailsResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	Status           *string                    `json:"status,omitempty"`
	Inputs           json.RawMessage            `json:"inputs,omitempty"`
	Outputs          json.RawMessage            `json:"outputs,omitempty"`
}
type ListWorkflowsResponseItemTriggerMethods struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ManualEnabled    *bool                      `json:"manualEnabled,omitempty"`
}
type ListWorkflowsResponseItemVersionsItemParametersItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	AllowCustomValue *bool                      `json:"allowCustomValue,omitempty"`
	Options          *[]string                  `json:"options,omitempty"`
}
type ListWorkflowsResponseItemVersionsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                             `json:"-"`
	Version          *int64                                                 `json:"version,omitempty"`
	Parameters       *[]ListWorkflowsResponseItemVersionsItemParametersItem `json:"parameters,omitempty"`
}
type ListWorkflowsResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage               `json:"-"`
	ID               *string                                  `json:"id,omitempty"`
	UUID             *string                                  `json:"uuid,omitempty"`
	Name             *string                                  `json:"name,omitempty"`
	Status           *string                                  `json:"status,omitempty"`
	TriggerMethods   *ListWorkflowsResponseItemTriggerMethods `json:"triggerMethods,omitempty"`
	Versions         *[]ListWorkflowsResponseItemVersionsItem `json:"versions,omitempty"`
}
type ListWorkflowsResponse []ListWorkflowsResponseItem
type GetWorkflowResponseTriggerMethods struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ManualEnabled    *bool                      `json:"manualEnabled,omitempty"`
}
type GetWorkflowResponseVersionsItemParametersItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ID               *string                    `json:"id,omitempty"`
	AllowCustomValue *bool                      `json:"allowCustomValue,omitempty"`
	Options          *[]string                  `json:"options,omitempty"`
}
type GetWorkflowResponseVersionsItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage                       `json:"-"`
	Version          *int64                                           `json:"version,omitempty"`
	Parameters       *[]GetWorkflowResponseVersionsItemParametersItem `json:"parameters,omitempty"`
}
type GetWorkflowResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage         `json:"-"`
	ID               *string                            `json:"id,omitempty"`
	UUID             *string                            `json:"uuid,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Status           *string                            `json:"status,omitempty"`
	TriggerMethods   *GetWorkflowResponseTriggerMethods `json:"triggerMethods,omitempty"`
	Versions         *[]GetWorkflowResponseVersionsItem `json:"versions,omitempty"`
}
type GetResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ExecutionUUID    *string                    `json:"executionUuid,omitempty"`
	WorkflowUUID     *string                    `json:"workflowUuid,omitempty"`
	Status           *string                    `json:"status,omitempty"`
}
type GetHistoryResponseItem struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	ThinkletID       *string                    `json:"thinkletId,omitempty"`
	ThinkletName     *string                    `json:"thinkletName,omitempty"`
	Type             *string                    `json:"type,omitempty"`
	Status           *string                    `json:"status,omitempty"`
	StartTime        *string                    `json:"startTime,omitempty"`
	EndTime          *string                    `json:"endTime,omitempty"`
	Details          json.RawMessage            `json:"details,omitempty"`
}
type GetHistoryResponse []GetHistoryResponseItem
type ExecuteResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	RequestUUID      *string                    `json:"requestUuid,omitempty"`
	ExecutionsUuids  *[]string                  `json:"executionsUuids,omitempty"`
}
type ExecuteNQLResponse struct {
	// AdditionalFields retains feature-dependent fields unknown to this SDK version.
	AdditionalFields map[string]json.RawMessage `json:"-"`
	RequestUUID      *string                    `json:"requestUuid,omitempty"`
	ExecutionsUuids  *[]string                  `json:"executionsUuids,omitempty"`
}

type TimelineOptions struct {
	ExecutionStatus string `json:"executionStatus"`
	TerminationDate string `json:"terminationDate,omitempty"`
}
type ListOptions struct {
	TriggerMethod            string `json:"triggerMethod,omitempty"`
	FetchOnlyActiveWorkflows *bool  `json:"fetchOnlyActiveWorkflows,omitempty"`
	FetchWorkflowContent     *bool  `json:"fetchWorkflowContent,omitempty"`
	Target                   string `json:"target,omitempty"`
}
type ExecuteRequest struct {
	// Source is the UI nx-source attribution header.
	Source       string            `json:"-"`
	WorkflowUUID string            `json:"workflowUuid"`
	Params       map[string]string `json:"params,omitempty"`
	Targets      []Target          `json:"targets,omitempty"`
	TriggerInfo  *TriggerInfo      `json:"triggerInfo,omitempty"`
}
type ExecuteNQLRequest struct {
	Source       string            `json:"-"`
	WorkflowUUID string            `json:"workflowUuid"`
	Params       map[string]string `json:"params,omitempty"`
	NQLQuery     string            `json:"nqlQuery"`
	NQLTimeZone  string            `json:"nqlTimeZone"`
	NQLTimeNow   string            `json:"nqlTimeNow"`
	TriggerInfo  *TriggerInfo      `json:"triggerInfo,omitempty"`
}
type Target struct {
	DeviceCollectorUID string `json:"deviceCollectorUid,omitempty"`
	UserSID            string `json:"userSid,omitempty"`
}

// TriggerInfo attributes workflow requests to the UI product and VDI targeting choice.
type TriggerInfo struct {
	InternalSource string `json:"internalSource"`
	Extra          string `json:"extra,omitempty"`
}
