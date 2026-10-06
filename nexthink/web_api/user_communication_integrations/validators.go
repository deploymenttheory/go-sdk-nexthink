package user_communication_integrations

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *IntegrationInput) error {
	if r == nil {
		return fmt.Errorf("integration request is required")
	}
	if strings.TrimSpace(r.Content.AzureTenantID) == "" && strings.TrimSpace(r.Content.AzureConnectorID) == "" {
		return fmt.Errorf("azure tenant ID or connector ID is required")
	}
	return nil
}
