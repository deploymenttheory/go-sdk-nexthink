package nql_queries

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateContentID(contentID string) error { return validation.PathSegment(contentID) }
func ValidateSaveQueryRequest(req *SaveQueryRequest, update bool) error {
	if req == nil || strings.TrimSpace(req.NQLAPIID) == "" || strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.NQL) == "" {
		return fmt.Errorf("query ID, name and NQL are required")
	}
	if update {
		return ValidateContentID(req.ContentID)
	}
	return nil
}
