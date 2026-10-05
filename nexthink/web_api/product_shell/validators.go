package product_shell

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateFlag(flag string) error            { return validation.PathSegment(flag) }
func ValidateMenu(menu string) error            { return validation.PathSegment(menu) }
func ValidateRequest(req json.RawMessage) error { return validation.JSONValue(req) }
