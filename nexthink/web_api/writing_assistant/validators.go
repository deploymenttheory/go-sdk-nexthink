package writing_assistant

import (
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
func ValidateCreateRequest(r *CreateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.Instructions) == "" || strings.TrimSpace(r.Tool) == "" {
		return fmt.Errorf("name, instructions and tool are required")
	}
	return ValidateID(r.ApplicationID)
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.Instructions) == "" {
		return fmt.Errorf("name and instructions are required")
	}
	if r.Revision < 1 {
		return fmt.Errorf("revision must be positive")
	}
	return ValidateID(r.ID)
}
