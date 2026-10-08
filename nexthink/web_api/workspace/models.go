package workspace

import "encoding/json"

// Document preserves all fields of polymorphic UI responses, including future fields.
type Document map[string]json.RawMessage

// Conversation exposes stable fields and retains unknown fields when round-tripping JSON.
type Conversation struct {
	ID                   string                     `json:"id,omitempty"`
	Title                *string                    `json:"title,omitempty"`
	Messages             []json.RawMessage          `json:"messages,omitempty"`
	Status               string                     `json:"status,omitempty"`
	Skill                json.RawMessage            `json:"skill,omitempty"`
	Automation           json.RawMessage            `json:"automation,omitempty"`
	IsUnread             *bool                      `json:"is_unread,omitempty"`
	IsStarred            *bool                      `json:"is_starred,omitempty"`
	IsDiscontinued       *bool                      `json:"is_discontinued,omitempty"`
	IsInterrupted        *bool                      `json:"is_interrupted,omitempty"`
	FirstUnreadMessageID *string                    `json:"first_unread_message_id,omitempty"`
	Shared               []json.RawMessage          `json:"shared,omitempty"`
	OriginHash           *string                    `json:"origin_hash,omitempty"`
	Tags                 []string                   `json:"tags,omitempty"`
	Files                []json.RawMessage          `json:"files,omitempty"`
	CreatedAt            string                     `json:"created_at,omitempty"`
	UpdatedAt            string                     `json:"updated_at,omitempty"`
	HasMoreBefore        *bool                      `json:"has_more_before,omitempty"`
	HasMoreAfter         *bool                      `json:"has_more_after,omitempty"`
	BeforeCursor         *string                    `json:"before_cursor,omitempty"`
	AfterCursor          *string                    `json:"after_cursor,omitempty"`
	AdditionalFields     map[string]json.RawMessage `json:"-"`
	present              map[string]json.RawMessage
}
type PageOptions struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
	Around string `json:"around,omitempty"`
	From   string `json:"from,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
type ConversationUpdate struct {
	Title     *string   `json:"title,omitempty"`
	IsStarred *bool     `json:"is_starred,omitempty"`
	Tags      *[]string `json:"tags,omitempty"`
}
type ConversationFileRequest struct {
	File     string `json:"file"`
	MIMEType string `json:"mime_type"`
	Filename string `json:"filename"`
}
type FileResponse struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}
type MCPRequest struct {
	Server         string          `json:"server"`
	Method         string          `json:"method"`
	Params         json.RawMessage `json:"params"`
	ConversationID string          `json:"conversation_id,omitempty"`
	ShareHash      string          `json:"share_hash,omitempty"`
}
type ChatContent struct {
	Type string  `json:"type"`
	Text *string `json:"text,omitempty"`
	ID   string  `json:"id,omitempty"`
}
type ChatMessage struct {
	Author    string             `json:"author"`
	Content   []ChatContent      `json:"content"`
	ClientID  string             `json:"client_id,omitempty"`
	Steps     *[]json.RawMessage `json:"steps,omitempty"`
	Artifacts []json.RawMessage  `json:"artifacts,omitempty"`
	Feedback  json.RawMessage    `json:"feedback,omitempty"`
}
type ChatRequest struct {
	ID               string          `json:"id"`
	Messages         []ChatMessage   `json:"messages"`
	ConversationMode string          `json:"conversation_mode"`
	Source           string          `json:"source"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
	OriginHash       string          `json:"origin_hash,omitempty"`
	SkillID          string          `json:"skill_id,omitempty"`
	ResumeValue      json.RawMessage `json:"resume_value,omitempty"`
	Context          json.RawMessage `json:"context,omitempty"`
}
type Event struct {
	Type string          `json:"type"`
	ID   string          `json:"id,omitempty"`
	Data json.RawMessage `json:"data"`
}
type ChatResponse struct {
	Events []Event `json:"events"`
}
