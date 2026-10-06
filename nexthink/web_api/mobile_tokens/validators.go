package mobile_tokens

import (
	"fmt"
	"strings"
	"time"
)

func validateID(v string) error {
	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("token identifier is required")
	}
	return nil
}
func validateCreate(r *CreateRequest) error {
	if r == nil || strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("token name is required")
	}
	if _, err := time.Parse(time.RFC3339, r.ExpirationDate); err != nil {
		return fmt.Errorf("expirationDate must be RFC3339: %w", err)
	}
	return nil
}
func validateUpdate(r *UpdateRequest) error {
	if r == nil || strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("token name is required")
	}
	if r.Revision < 0 {
		return fmt.Errorf("revision must not be negative")
	}
	return validateID(r.JTI)
}
