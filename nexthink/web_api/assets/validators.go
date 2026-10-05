package assets

import (
	"fmt"
	"mime"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateUpload(r *UploadRequest) error {
	if r == nil || strings.TrimSpace(r.Name) == "" || len(r.Data) == 0 {
		return fmt.Errorf("filename and nonempty data are required")
	}
	if strings.ContainsAny(r.Name, "\r\n") {
		return fmt.Errorf("filename cannot contain header newlines")
	}
	media, params, err := mime.ParseMediaType(r.MediaType)
	if err != nil || len(params) != 0 || media != r.MediaType || !strings.Contains(media, "/") {
		return fmt.Errorf("media type must be a valid MIME type without parameters")
	}
	return nil
}
