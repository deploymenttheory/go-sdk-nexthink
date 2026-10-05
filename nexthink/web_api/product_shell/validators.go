package product_shell

import (
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateFlag(flag string) error            { return validation.PathSegment(flag) }
func ValidateMenu(menu string) error            { return validation.PathSegment(menu) }
func ValidateRequest(req json.RawMessage) error { return validation.JSONValue(req) }

func ValidateClaimsRequest(request *ClaimsRequest) error {
	if request == nil {
		return fmt.Errorf("claims request is required")
	}
	switch request.Method {
	case "hasAllClaims", "hasAnyClaim":
		if request.Claims == nil {
			return fmt.Errorf("claims array is required")
		}
	case "hasPatternClaim":
		if request.PatternClaim == nil || *request.PatternClaim == "" {
			return fmt.Errorf("patternClaim is required")
		}
	case "hasClaimValue":
		if request.Claim == nil || *request.Claim == "" || !json.Valid(request.Value) {
			return fmt.Errorf("claim and JSON value are required")
		}
	default:
		return fmt.Errorf("unsupported claims validation method")
	}
	return nil
}
