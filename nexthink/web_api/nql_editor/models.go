package nql_editor

import "encoding/json"

// Document is the editor's current text snapshot. Empty text is valid.
type Document struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Text       string `json:"text"`
	Version    int    `json:"version"`
}

// Position uses zero-based lines and UTF-16 code units, as in the browser editor.
type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}
type ValidationRequest struct {
	Document
	Rules json.RawMessage `json:"rules,omitempty"`
}
type PositionRequest struct {
	Document Document `json:"document"`
	Position Position `json:"position"`
}
type Diagnostic struct {
	Range    Range           `json:"range"`
	Message  string          `json:"message"`
	Severity int             `json:"severity,omitempty"`
	Source   string          `json:"source,omitempty"`
	Code     json.RawMessage `json:"code,omitempty"`
}

// CompletionItem preserves every field for the resolve round trip. The server
// may attach opaque data or additional editor protocol fields.
type CompletionItem map[string]json.RawMessage

// Hover and highlighting retain the editor's evolving protocol schemas.
type (
	Hover        = json.RawMessage
	Highlighting = json.RawMessage
)
