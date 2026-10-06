package knowledge_bases

import (
	"encoding/base64"
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_administration"
)

// ListResponse includes uploaded knowledge bases and connector-backed sources.
type ListResponse struct {
	User content_administration.ContentUser `json:"user"`
	Rows []Summary                          `json:"rows"`
}
type Summary struct {
	content_administration.Content
	TargetGroup  json.RawMessage `json:"targetGroup"`
	Articles     int             `json:"articles"`
	Type         string          `json:"type"`
	LastUpdate   int64           `json:"lastUpdate"`
	Status       string          `json:"status"`
	Errors       json.RawMessage `json:"errors"`
	NQLID        string          `json:"nqlId"`
	TemplateName string          `json:"templateName"`
}
type CreateRequest struct {
	ContentID    string  `json:"contentId"`
	FileName     string  `json:"fileName"`
	TargetGroup  string  `json:"targetGroup"`
	ITSMHost     string  `json:"itsmHost"`
	UserCriteria *string `json:"userCriteria,omitempty"`
}

// UploadRequest.Data contains original CSV bytes; the service base64-encodes once.
type UploadRequest struct {
	OriginalFileName string `json:"originalFileName"`
	Data             []byte `json:"data"`
}
type UploadResponse struct {
	ContentID string `json:"contentId"`
	FileName  string `json:"fileName"`
}
type StartMultipartRequest struct {
	OriginalFileName string `json:"originalFileName"`
}
type MultipartUpload struct {
	UploadResponse
	UploadID string `json:"uploadId"`
	Bucket   string `json:"bucket"`
}
type MultipartContext struct {
	ContentID string `json:"contentId"`
	FileName  string `json:"fileName"`
	UploadID  string `json:"uploadId"`
}

// UploadPartRequest contains a slice of the base64-encoded WHOLE file, not an
// independently encoded arbitrary binary chunk. The UI uses 6,990,508-character
// chunks. Encode once, then split; do not add padding to intermediate chunks.
type UploadPartRequest struct {
	MultipartContext
	PartNumber   int    `json:"partNumber"`
	EncodedChunk string `json:"encodedChunk"`
}
type UploadedPart struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"etag"`
	Checksum   string `json:"checksum"`
}
type CompleteMultipartRequest struct {
	MultipartContext
	Parts []UploadedPart `json:"parts"`
}
type CompleteMultipartResponse struct {
	ContentID string `json:"contentId"`
	ETag      string `json:"etag"`
}

// DownloadURLResponse preserves the wire's base64 URL. DecodeURL returns the signed
// URL; download it separately without forwarding the Nexthink bearer token.
type DownloadURLResponse struct {
	URL string `json:"url"`
}

func (r DownloadURLResponse) DecodeURL() (string, error) {
	data, err := base64.StdEncoding.DecodeString(r.URL)
	return string(data), err
}

// Content is an indexed knowledge content summary. Ingestion is asynchronous.
type Content struct {
	ContentID  string `json:"contentId"`
	Source     string `json:"source"`
	Type       string `json:"type"`
	ItemsCount int    `json:"itemsCount"`
	LastUpdate string `json:"lastUpdate"`
}
