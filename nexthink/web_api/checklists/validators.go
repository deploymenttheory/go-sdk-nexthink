package checklists

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateRevision(revision int) error {
	if revision < 1 {
		return fmt.Errorf("revision must be positive")
	}
	return nil
}
func ValidateInput(r *ChecklistInput) error {
	if r == nil || strings.TrimSpace(r.Label) == "" {
		return fmt.Errorf("checklist label is required")
	}
	switch r.Platforms {
	case "All", "Windows", "macOS", "Linux":
	default:
		return fmt.Errorf("platforms must be All, Windows, macOS or Linux")
	}
	count := 0
	for _, c := range r.Categories {
		if strings.TrimSpace(c.Label) == "" {
			return fmt.Errorf("category label is required")
		}
		for _, f := range c.Fields {
			if strings.TrimSpace(f.ID) == "" || strings.TrimSpace(f.Type) == "" {
				return fmt.Errorf("field ID and type are required")
			}
			count++
		}
	}
	if count == 0 {
		return fmt.Errorf("at least one field is required")
	}
	for _, data := range r.FieldData {
		if strings.TrimSpace(data.ID) == "" || strings.TrimSpace(data.Type) == "" {
			return fmt.Errorf("field data requires an ID and type")
		}
	}
	return nil
}
