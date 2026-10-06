package workspace_agents

import "encoding/json"

// Document preserves all fields of polymorphic UI responses, including future fields.
type Document map[string]json.RawMessage

// Skill exposes stable fields and retains unknown fields when round-tripping JSON.
type Skill struct {
	ID               string                     `json:"id,omitempty"`
	Name             string                     `json:"name,omitempty"`
	Goal             *string                    `json:"goal,omitempty"`
	Instructions     string                     `json:"instructions,omitempty"`
	AutomationIDs    []string                   `json:"automation_ids,omitempty"`
	Source           string                     `json:"source,omitempty"`
	CreatedAt        string                     `json:"created_at,omitempty"`
	UpdatedAt        string                     `json:"updated_at,omitempty"`
	Files            []SkillFile                `json:"files,omitempty"`
	AdditionalFields map[string]json.RawMessage `json:"-"`
	present          map[string]json.RawMessage
}
type SkillRequest struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name"`
	Goal         *string   `json:"goal"`
	Instructions string    `json:"instructions"`
	Files        *[]string `json:"files,omitempty"`
}
type AvailabilityResponse struct {
	Available bool `json:"available"`
}
type FileUploadRequest struct {
	Filename string `json:"filename"`
	MIMEType string `json:"mime_type"`
	Data     []byte `json:"data"`
}
type StartMultipartRequest struct {
	Filename string `json:"filename"`
	MIMEType string `json:"mime_type"`
}
type FileResponse struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}
type MultipartUpload struct {
	UploadID string `json:"uploadId"`
	FileID   string `json:"fileId"`
}
type UploadPartRequest struct {
	MultipartUpload
	PartNumber int    `json:"partNumber"`
	Data       []byte `json:"data"`
}
type UploadedPart struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"etag"`
}
type CompleteMultipartRequest struct {
	MultipartUpload
	Parts []UploadedPart `json:"parts"`
}

// SkillFile describes a file attached to an agent knowledge base.
type SkillFile struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	MIMEType  string `json:"mime_type"`
	Size      int64  `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
