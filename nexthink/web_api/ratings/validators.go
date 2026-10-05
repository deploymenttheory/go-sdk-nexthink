package ratings

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
func ValidateInput(r *RatingInput) error {
	if r == nil || strings.TrimSpace(r.TargetField) == "" || strings.TrimSpace(r.TargetFieldLabel) == "" {
		return fmt.Errorf("target field and label are required")
	}
	if r.DefaultValueInEnumeration < 0 || r.DefaultValueInEnumeration > 3 {
		return fmt.Errorf("default enumeration must be between 0 and 3")
	}
	if len(r.Conditions) == 0 {
		return fmt.Errorf("at least one rating condition is required")
	}
	for k, q := range r.Conditions {
		if k != "1" && k != "2" && k != "3" {
			return fmt.Errorf("rating conditions must use 1, 2 or 3")
		}
		if strings.TrimSpace(q) == "" {
			return fmt.Errorf("condition NQL cannot be empty")
		}
	}
	return nil
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := ValidateID(r.RatingID); err != nil {
		return err
	}
	if err := ValidateRevision(r.Revision); err != nil {
		return err
	}
	return ValidateInput(&r.RatingInput)
}
