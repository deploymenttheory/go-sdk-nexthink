package legacy_connectors

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(id string, r *ConfigurationInput) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	if r == nil || strings.TrimSpace(r.ConnectorName) == "" {
		return fmt.Errorf("connector name is required")
	}
	if len(r.ConnectionDetails) == 0 || len(r.Mapping) == 0 {
		return fmt.Errorf("connection details and mapping are required")
	}
	return nil
}
