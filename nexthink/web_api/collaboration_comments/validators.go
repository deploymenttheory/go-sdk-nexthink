package collaboration_comments

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func validateReference(name, value string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." {
		return fmt.Errorf("%s is required", name)
	}
	return nil
}
func validateMessage(message string) error {
	if strings.TrimSpace(message) == "" || utf8.RuneCountInString(message) > 1000 {
		return fmt.Errorf("message must contain 1 to 1000 characters")
	}
	return nil
}
func validateRequest(request any) error {
	switch r := request.(type) {
	case *IdentifierRequest:
		if r == nil || len(r.Identifiers) == 0 || strings.TrimSpace(r.URL) == "" {
			return fmt.Errorf("identifiers and url are required")
		}
		for _, id := range r.Identifiers {
			if strings.TrimSpace(id) == "" {
				return fmt.Errorf("identifiers cannot contain blank values")
			}
		}
	case *CreateMessageRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateReference("id", r.ID); err != nil {
			return err
		}
		return validateMessage(r.Message)
	case *EditMessageRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		return validateMessage(r.Message)
	default:
		return fmt.Errorf("request is required")
	}
	return nil
}
