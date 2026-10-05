package content_administration

import "encoding/json"

// ConfigurationResponse retains the configurable listing and permission schemas.
type ConfigurationResponse struct {
	Labels             json.RawMessage   `json:"labels"`
	Storage            json.RawMessage   `json:"storage"`
	Storages           []json.RawMessage `json:"storages,omitempty"`
	ContentPermissions json.RawMessage   `json:"contentPermissions"`
	Features           json.RawMessage   `json:"features"`
}
type ListResponse struct {
	User ContentUser `json:"user"`
	Rows []Content   `json:"rows"`
}
type ContentUser struct {
	ID string `json:"id"`
}
type Content struct {
	ContentType       string            `json:"contentType"`
	ContentID         string            `json:"contentId"`
	ContentOwner      string            `json:"contentOwner"`
	Title             string            `json:"title"`
	Active            bool              `json:"active"`
	ResourceName      string            `json:"resourceName"`
	Tags              []json.RawMessage `json:"tags"`
	IsCopyFromLibrary bool              `json:"isCopyFromLibrary"`
	Revision          int               `json:"revision"`
	CreatedBy         string            `json:"createdBy"`
	UpdatedBy         string            `json:"updatedBy"`
}
