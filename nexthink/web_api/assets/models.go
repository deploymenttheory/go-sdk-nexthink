package assets

import "github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"

// UploadRequest carries raw file bytes. The SDK builds the UI's data URL exactly once.
type UploadRequest struct {
	Name      string `json:"name"`
	MediaType string `json:"mediaType"`
	Data      []byte `json:"data"`
}
type CreateResponse struct {
	ContentID string `json:"contentId"`
	Message   string `json:"message"`
}

// SignedURLResponse is a temporary download grant. Treat SignedURL as a secret.
// The SDK does not forward the browser bearer token to this URL.
type SignedURLResponse struct {
	SignedURL string `json:"signedUrl"`
	ExpiresAt int64  `json:"expiresAt"`
	TTL       int64  `json:"ttl"`
}
type Summary struct {
	content_administration.Content
	Size       int64  `json:"size"`
	Extension  string `json:"extension"`
	LastUpdate int64  `json:"lastUpdate"`
}
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
