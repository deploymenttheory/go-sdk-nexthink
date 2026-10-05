package content_administration

import (
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateConfiguration(
	configuration string,
) error {
	return validation.PathSegment(configuration)
}
