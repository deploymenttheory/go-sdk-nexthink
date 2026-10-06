package zoom_notifications

import (
	"fmt"
	"strings"
)

func ValidateCheckCredentials(r *CheckCredentialsRequest) error {
	if r == nil {
		return fmt.Errorf("credentials request is required")
	}
	if strings.TrimSpace(r.JWT) == "" {
		return fmt.Errorf("JWT is required")
	}
	return nil
}
