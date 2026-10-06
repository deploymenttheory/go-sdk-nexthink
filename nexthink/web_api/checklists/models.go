package checklists

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

// ChecklistInput is a definition only; saving a checklist does not execute its actions.
type ChecklistInput struct {
	Label              string          `json:"label"`
	Platforms          string          `json:"platforms"`
	Description        string          `json:"description"`
	Categories         []CategoryInput `json:"categories"`
	FieldData          []FieldData     `json:"fieldData"`
	MissingFieldsExist bool            `json:"missingFieldsExist"`
}
type CategoryInput struct {
	ID     string       `json:"id,omitempty"`
	Label  string       `json:"label"`
	Fields []FieldInput `json:"fields"`
}
type FieldInput struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	SubType     string `json:"subType,omitempty"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}
type Category struct {
	ID     string  `json:"id"`
	Label  string  `json:"label"`
	Fields []Field `json:"fields"`
}
type Field struct {
	FieldInput
	Platform string `json:"platform"`
	Rated    bool   `json:"rated"`
}

// Action definitions are polymorphic and retained intact; no action is executed by this service.
type FieldData struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	SubType       string            `json:"subType,omitempty"`
	CustomLabel   string            `json:"customLabel"`
	Documentation string            `json:"documentation"`
	Actions       []json.RawMessage `json:"actions"`
}
type Checklist struct {
	ID                 string          `json:"id"`
	RevisionNumber     int             `json:"revisionNumber"`
	LastUpdated        string          `json:"lastUpdated"`
	Label              string          `json:"label"`
	Platforms          string          `json:"platforms"`
	Description        string          `json:"description"`
	Categories         []Category      `json:"categories"`
	FieldData          []FieldData     `json:"fieldData"`
	Metadata           json.RawMessage `json:"metadata"`
	MissingFieldsExist bool            `json:"missingFieldsExist"`
}
type CreateOptions struct{ LibraryUUID string }
type Summary struct {
	content_administration.Content
	ChecklistLabel       string `json:"checklistLabel"`
	ChecklistDescription string `json:"checklistDescription"`
	RevisionNumber       int    `json:"revisionNumber"`
	LastUpdated          int64  `json:"lastUpdated"`
	Platforms            string `json:"platforms"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}

type ExportDocument struct {
	Label       string          `json:"label"`
	Description string          `json:"description"`
	Platforms   string          `json:"platforms"`
	Categories  []CategoryInput `json:"categories"`
	FieldData   []FieldData     `json:"fieldData"`
	Version     int             `json:"version"`
	Type        string          `json:"type"`
}
type FieldGroup struct {
	Label  string       `json:"label"`
	Fields []FieldInput `json:"fields"`
}

// LibraryDocument omits platforms when the library template does not specify it.
type LibraryDocument struct {
	Label       string             `json:"label"`
	Description string             `json:"description"`
	Platforms   *string            `json:"platforms,omitempty"`
	Categories  []CategoryInput    `json:"categories"`
	FieldData   []LibraryFieldData `json:"fieldData"`
	Version     int                `json:"version"`
	Type        string             `json:"type"`
}

// LibraryFieldData retains an omitted custom label in built-in definitions.
type LibraryFieldData struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	SubType       string            `json:"subType,omitempty"`
	CustomLabel   *string           `json:"customLabel,omitempty"`
	Documentation string            `json:"documentation"`
	Actions       []json.RawMessage `json:"actions"`
}
