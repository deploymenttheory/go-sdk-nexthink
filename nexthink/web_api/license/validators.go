package license

import (
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateFeature(feature string) error { return validation.PathSegment(feature) }
