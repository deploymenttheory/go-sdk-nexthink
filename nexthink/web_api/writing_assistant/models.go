package writing_assistant

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"

type CreateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Instructions  string `json:"instructions"`
	ApplicationID string `json:"applicationId"`
	Tool          string `json:"tool"`
}
type UpdateRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Revision     int    `json:"revision"`
}
type WritingAssistant struct {
	ID           string `json:"id"`
	Tool         string `json:"tool"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Revision     int    `json:"revision"`
}
type Identifier struct {
	ID string `json:"id"`
}
type DeleteResult struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}
type CreateResponse struct {
	WritingAssistant *Identifier `json:"createWritingAssistant"`
}
type UpdateResponse struct {
	WritingAssistant *Identifier `json:"updateWritingAssistant"`
}
type DeleteResponse struct {
	WritingAssistant *DeleteResult `json:"deleteWritingAssistant"`
}
type GetResponse struct {
	WritingAssistant *WritingAssistant `json:"writingAssistant"`
}
type Summary struct {
	content_administration.Content
	Description string `json:"description"`
	Tool        string `json:"tool"`
	LastUpdate  int64  `json:"lastUpdate"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
