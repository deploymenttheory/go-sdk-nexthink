package autopilot

import "encoding/json"

type Configuration struct {
	Learning       bool   `json:"learning"`
	WebSearch      bool   `json:"webSearch"`
	Escalation     bool   `json:"escalation"`
	FilesystemTool bool   `json:"filesystemTool"`
	AgentName      string `json:"agentName"`
}
type BooleanValue struct {
	Value bool `json:"value"`
}
type StringValue struct {
	Value string `json:"value"`
}
type Settings struct {
	Revision       int64                  `json:"revision"`
	Calls          *SavedCalls            `json:"calls,omitempty"`
	Configuration  *UploadedConfiguration `json:"configuration,omitempty"`
	Categorization *Categorization        `json:"categorization,omitempty"`
}
type SavedCalls struct {
	Saved    []SavedCall `json:"saved"`
	Selected *string     `json:"selected,omitempty"`
}
type SavedCall struct {
	CallID            string                      `json:"callId"`
	Parameters        map[string]ParameterMapping `json:"parameters,omitempty"`
	FallbackURL       string                      `json:"fallbackUrl"`
	TicketURL         *string                     `json:"ticketUrl,omitempty"`
	TicketDisplayname *string                     `json:"ticketDisplayname,omitempty"`
}

// ParameterMapping is the source-discriminated union used in escalation settings.
type ParameterMapping struct {
	Source      string  `json:"source"`
	CustomValue *string `json:"customValue,omitempty"`
	Field       *string `json:"field,omitempty"`
	Column      *string `json:"column,omitempty"`
}
type UploadedConfiguration struct {
	Filename   string   `json:"filename"`
	CSVContent *string  `json:"csvContent,omitempty"`
	UploadedAt *float64 `json:"uploadedAt,omitempty"`
}
type Categorization struct {
	Columns []CategorizationColumn `json:"columns"`
	Rows    [][]string             `json:"rows"`
}
type CategorizationColumn struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level *int   `json:"level,omitempty"`
}
type CategorizationDownload struct {
	Filename string `json:"filename"`
	Data     []byte `json:"data"`
}
type WebSearchDomains struct {
	Revision        *int64            `json:"revision,omitempty"`
	BaselineVersion *string           `json:"baselineVersion,omitempty"`
	Domains         []WebSearchDomain `json:"domains"`
}
type WebSearchDomain struct {
	ID      string `json:"id"`
	Host    string `json:"host"`
	Origin  string `json:"origin"`
	Enabled bool   `json:"enabled"`
}
type WebSearchDomainsRequest struct {
	Revision *int64                 `json:"revision,omitempty"`
	Domains  []WebSearchDomainInput `json:"domains"`
}
type WebSearchDomainInput struct {
	Host    string `json:"host"`
	Enabled bool   `json:"enabled"`
}
type Approval struct {
	ID string `json:"_id"`
	ApprovalInput
}
type ApprovalInput struct {
	NeedsApproval bool   `json:"needsApproval"`
	BeforeMessage string `json:"beforeMessage"`
	DuringMessage string `json:"duringMessage"`
	AfterMessage  string `json:"afterMessage"`
	ActionType    string `json:"actionType"`
	Revision      *int64 `json:"revision,omitempty"`
}
type CallsResponse struct {
	Calls []Call `json:"calls"`
}
type Call struct {
	ID         string          `json:"id"`
	NQLID      string          `json:"nqlId"`
	Name       string          `json:"name"`
	Parameters []CallParameter `json:"parameters,omitempty"`
	Outputs    []CallParameter `json:"outputs,omitempty"`
}
type CallParameter struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
type ArticleCount struct {
	ArticleCount int64 `json:"articleCount"`
}
type KnowledgeConnectors struct {
	Connectors []KnowledgeConnector `json:"connectors"`
}
type KnowledgeConnector struct {
	ID            *string `json:"id"`
	Name          *string `json:"name"`
	Connector     *string `json:"connector"`
	Status        *string `json:"status"`
	LastUpdate    *string `json:"lastUpdate"`
	ProcessedRows *int64  `json:"processedRows"`
	ReceivedRows  *int64  `json:"receivedRows"`
}
type TicketRequest struct {
	ConversationID         string  `json:"conversationId"`
	UserSID                string  `json:"userSid"`
	DeviceLicenseID        string  `json:"deviceLicenseId"`
	ShortDescription       string  `json:"shortDescription"`
	ConversationSummary    string  `json:"conversationSummary"`
	WorkNotes              string  `json:"workNotes"`
	ActionsTaken           string  `json:"actionsTaken"`
	CategorizationRowIndex *int    `json:"categorizationRowIndex,omitempty"`
	Category               *string `json:"category,omitempty"`
	Subcategory            *string `json:"subcategory,omitempty"`
}

// TicketResult includes application-level FAILED results returned over HTTP 200.
type TicketResult struct {
	Status            string  `json:"status"`
	TicketID          *string `json:"ticketId,omitempty"`
	TicketNumber      *string `json:"ticketNumber,omitempty"`
	TicketURL         *string `json:"ticketUrl,omitempty"`
	TicketDisplayname *string `json:"ticketDisplayname,omitempty"`
	Error             *string `json:"error,omitempty"`
	FailureReason     *string `json:"failureReason,omitempty"`
}
type ConversationIDs struct {
	ConversationIDs []string `json:"conversationIds"`
}
type Conversation struct {
	Messages          []ConversationMessage `json:"messages"`
	LastTicketCreated *CreatedTicket        `json:"lastTicketCreated,omitempty"`
}
type ConversationMessage struct {
	MessageID   string              `json:"messageId"`
	Timestamp   float64             `json:"timestamp"`
	MessageStep string              `json:"messageStep"`
	Content     *string             `json:"content,omitempty"`
	Reasoning   *string             `json:"reasoning,omitempty"`
	ToolCall    *ToolCall           `json:"toolCall,omitempty"`
	Action      *ConversationAction `json:"action,omitempty"`
}
type ToolCall struct {
	ToolCallID    string          `json:"toolCallId"`
	ToolName      string          `json:"toolName"`
	Args          json.RawMessage `json:"args,omitempty"`
	OutcomeStatus string          `json:"outcomeStatus"`
	Result        json.RawMessage `json:"result,omitempty"`
}
type ConversationAction struct {
	Purpose       *string `json:"purpose,omitempty"`
	ActionOutcome *string `json:"actionOutcome,omitempty"`
	ActionName    *string `json:"actionName,omitempty"`
}
type CreatedTicket struct {
	SysID            *string         `json:"sysid,omitempty"`
	TicketURL        *string         `json:"ticketUrl,omitempty"`
	ShortDescription *string         `json:"shortDescription,omitempty"`
	LongDescription  *string         `json:"longDescription,omitempty"`
	Category         *string         `json:"category,omitempty"`
	Subcategory      *string         `json:"subcategory,omitempty"`
	Segments         []TicketSegment `json:"segments,omitempty"`
	WorkNotes        *string         `json:"workNotes,omitempty"`
	TicketNumber     *string         `json:"ticketNumber,omitempty"`
	DisplayName      *string         `json:"displayName,omitempty"`
}
type TicketSegment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type AgentAction struct {
	ContentID   *string          `json:"contentId,omitempty"`
	Title       *string          `json:"title,omitempty"`
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Purpose     []string         `json:"purpose,omitempty"`
	ScriptInfo  *AgentScriptInfo `json:"scriptInfo"`
}
type AgentScriptInfo struct {
	RunAs          *string            `json:"runAs,omitempty"`
	TimeoutSeconds *int               `json:"timeoutSeconds,omitempty"`
	Inputs         []AgentActionInput `json:"inputs"`
	ScriptWindows  *AgentScript       `json:"scriptWindows,omitempty"`
	ScriptMacOS    *AgentScript       `json:"scriptMacOs,omitempty"`
	Impact         *AgentImpact       `json:"impact,omitempty"`
}
type AgentActionInput struct {
	AllowCustomValue bool     `json:"allowCustomValue"`
	Description      string   `json:"description"`
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Options          []string `json:"options"`
	UsedByMacOS      bool     `json:"usedByMacOs"`
	UsedByWindows    bool     `json:"usedByWindows"`
	Type             string   `json:"type"`
}
type AgentScript struct {
	Name string `json:"name"`
}
type AgentImpact struct {
	Windows *PlatformImpact `json:"windows"`
	MacOS   *PlatformImpact `json:"macos"`
}
type PlatformImpact struct {
	Behavior       *string `json:"behavior"`
	ExpectedImpact *string `json:"expectedImpact"`
}
type AgentActionInputsRequest struct {
	Inputs []AgentActionInput `json:"inputs"`
}
