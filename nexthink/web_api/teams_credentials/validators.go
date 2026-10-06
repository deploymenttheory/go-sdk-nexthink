package teams_credentials

import (
	"fmt"
	"strings"
)

func ValidateCheckCredentials(r *CheckCredentialsRequest) error {
	if r == nil {
		return fmt.Errorf("credentials request is required")
	}
	if strings.TrimSpace(r.TenantID) == "" {
		return fmt.Errorf("TenantID is required")
	}
	if strings.TrimSpace(r.ClientID) == "" {
		return fmt.Errorf("ClientID is required")
	}
	if strings.TrimSpace(r.ClientSecret) == "" {
		return fmt.Errorf("ClientSecret is required")
	}
	if strings.TrimSpace(r.NationalCloud) == "" {
		return fmt.Errorf("NationalCloud is required")
	}
	return nil
}
