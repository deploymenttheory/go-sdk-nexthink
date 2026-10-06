package legacy_connectors

import (
	"fmt"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(id string, r *ConfigurationInput) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("configuration request is required")
	}
	if len(r.ConnectionDetails) == 0 || r.Mapping == nil {
		return fmt.Errorf("connection details and an explicit mapping array are required")
	}
	return nil
}
