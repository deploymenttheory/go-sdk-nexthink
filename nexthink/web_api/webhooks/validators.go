package webhooks

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

var uuidPattern = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *Webhook) error {
	if r == nil || !uuidPattern.MatchString(r.UUID) {
		return fmt.Errorf("caller-generated UUID is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.CredentialID) == "" || strings.TrimSpace(r.NQLCondition) == "" {
		return fmt.Errorf("name, credential ID and NQL condition are required")
	}
	if r.Type != "Event" || len(r.Communications) == 0 {
		return fmt.Errorf("Event type and communications are required")
	}
	for _, c := range r.Communications {
		if c.HTTPMethod != "POST" && c.HTTPMethod != "PUT" && c.HTTPMethod != "PATCH" {
			return fmt.Errorf("communication method must be POST, PUT or PATCH")
		}
	}
	return nil
}
