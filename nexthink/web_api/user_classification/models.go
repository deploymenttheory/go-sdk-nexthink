package user_classification

import "encoding/json"

// Field.NQLID includes the leading #, as in user.organization.#department.
type Field struct {
	NQLID       string `json:"nqlId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type ListResponse struct {
	CustomFields []Field         `json:"customFields"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
}

// ReplaceRequest replaces the complete user-organization field list. An empty
// array removes every field; nil is rejected to avoid accidentally sending null.
type ReplaceRequest struct {
	CustomFields []Field `json:"customFields"`
}
