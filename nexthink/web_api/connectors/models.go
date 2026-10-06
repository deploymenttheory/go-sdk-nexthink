package connectors

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

// ConnectorInput is the UI's complete configuration. ContentID is a caller-generated UUID.
// The lab server forced Enabled=true on Create; Update honored false. Do not rely on
// a disabled create to prevent execution: choose credentials and scheduling accordingly.
type ConnectorInput struct {
	ContentID     string              `json:"content_id"`
	Enabled       bool                `json:"enabled"`
	Template      TemplateReference   `json:"template"`
	General       General             `json:"general"`
	Scheduling    Scheduling          `json:"scheduling"`
	Credentials   CredentialReference `json:"credentials"`
	UserInputs    []UserInput         `json:"user_inputs"`
	Streams       []Stream            `json:"streams"`
	CustomHeaders []Header            `json:"custom_headers"`
}

// Connector is a saved configuration. Empty custom_headers may be absent in responses.
type Connector struct {
	ContentID     string              `json:"content_id"`
	Enabled       bool                `json:"enabled"`
	Template      TemplateReference   `json:"template"`
	General       General             `json:"general"`
	Scheduling    Scheduling          `json:"scheduling"`
	Credentials   CredentialReference `json:"credentials"`
	UserInputs    []UserInput         `json:"user_inputs"`
	Streams       []Stream            `json:"streams"`
	CustomHeaders []Header            `json:"custom_headers,omitempty"`
}
type TemplateReference struct {
	Name string `json:"name"`
}
type General struct {
	NQLID       string `json:"nql_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Scheduling uses Quartz cron syntax, including seconds and an optional year.
type Scheduling struct {
	Cron     string `json:"cron"`
	Timezone string `json:"timezone"`
}
type CredentialReference struct {
	Reference string `json:"cred_ref"`
	Type      string `json:"cred_type"`
}
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type UserInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Stream struct {
	Name           string    `json:"name"`
	Identification []Mapping `json:"identification"`
	Fields         []Mapping `json:"fields"`
}
type Mapping struct {
	SourceField   string `json:"source_field"`
	NexthinkField string `json:"nexthink_field"`
	Enabled       bool   `json:"enabled"`
	BuiltIn       bool   `json:"built_in"`
}

// Summary includes both universal and legacy connectors from the shared UI list.
// A legacy ID cannot be passed to this service's v1 Get/Update/Delete methods.
type Summary struct {
	content_administration.Content
	Title     *string `json:"title"`
	Type      string  `json:"type"`
	Frequency string  `json:"frequency"`
	Enabled   string  `json:"enabled"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}

// Template retains the variable, template-specific schemas as JSON.
// Templates are server-provided; no create/update/delete contract was observed.
type Template struct {
	ID                         string            `json:"id"`
	FriendlyName               string            `json:"friendlyName"`
	AllowCustomHeaders         bool              `json:"allowCustomHeaders"`
	AllowCustomMappings        bool              `json:"allowCustomMappings"`
	AllowedPeriodicity         json.RawMessage   `json:"allowedPeriodicity"`
	CredentialTypes            []string          `json:"credentialTypes"`
	Datasources                []json.RawMessage `json:"datasources"`
	Deprecated                 bool              `json:"deprecated"`
	DeprecationMessage         *string           `json:"deprecationMessage"`
	Destination                json.RawMessage   `json:"destination"`
	ExtractionScript           string            `json:"extractionScript"`
	UsesInternalProxy          bool              `json:"usesInternalProxy"`
	BuiltInMappings            []json.RawMessage `json:"builtInMappings"`
	EnableInFedRamp            bool              `json:"enableInFedRamp"`
	FeatureFlag                *string           `json:"featureFlag"`
	Icon                       *string           `json:"icon"`
	Joins                      []json.RawMessage `json:"joins"`
	InProductDocumentationPath *string           `json:"inProductDocumentationPath"`
	LicenseKey                 *string           `json:"licenseKey"`
	Parameters                 []json.RawMessage `json:"parameters"`
}
type ManualCustomField struct {
	DataModelPath string `json:"dataModelPath"`
	Name          string `json:"name"`
}

// TestRequest executes requests against the referenced third-party destination.
// It does not create a scheduled connector configuration.
type TestRequest struct {
	Credentials   TestCredentials `json:"credentials"`
	CustomHeaders []Header        `json:"custom_headers"`
	TemplateID    string          `json:"templateId"`
	UserInputs    []UserInput     `json:"user_inputs"`
}
type TestCredentials struct {
	Reference string `json:"cred_ref"`
}
type TestExecution struct {
	TestExecutionID string `json:"testExecutionId"`
}

// TestResult reports execution status separately from per-partition failures.
// COMPLETED does not mean that every external request succeeded.
type TestResult struct {
	Status          string       `json:"status"`
	TestExecutionID string       `json:"testExecutionId"`
	Streams         []TestStream `json:"streams,omitempty"`
}
type TestStream struct {
	Name       string          `json:"name"`
	Partitions []TestPartition `json:"partitions"`
}
type TestPartition struct {
	Name                 string          `json:"name"`
	ResponseTimeMillis   int64           `json:"responseTimeMillis"`
	Error                *TestError      `json:"error"`
	ErrorType            *string         `json:"errorType"`
	ResponseBody         json.RawMessage `json:"responseBody"`
	SampleResponseRecord json.RawMessage `json:"sampleResponseRecord"`
	StatusCode           *int            `json:"statusCode"`
	URI                  *string         `json:"uri"`
}
type TestError struct {
	Source string `json:"source"`
	Type   string `json:"type"`
}
