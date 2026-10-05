package investigations

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *InvestigationInput) error {
	if r == nil || strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQL) == "" {
		return fmt.Errorf("investigation name and NQL are required")
	}
	return nil
}
func ValidateImport(r *ExportDocument) error {
	if r == nil || strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLQuery) == "" {
		return fmt.Errorf("import name and nqlQuery are required")
	}
	return nil
}
