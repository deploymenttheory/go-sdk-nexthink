package custom_field_values

import (
	"fmt"
	"io"
	"strings"
)

func validateUpdate(r *UpdateRequest) error {
	if r == nil || strings.TrimSpace(r.URI) == "" || len(r.ObjectIDs) == 0 {
		return fmt.Errorf("field URI and target object IDs are required")
	}
	for _, id := range r.ObjectIDs {
		if strings.TrimSpace(id) == "" {
			return fmt.Errorf("object ID must not be empty")
		}
	}
	return nil
}
func validateUpload(name string, reader io.Reader, size int64) error {
	if strings.TrimSpace(name) == "" || reader == nil || size <= 0 {
		return fmt.Errorf("CSV filename, reader and positive size are required")
	}
	return nil
}
