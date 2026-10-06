package connector_credentials

import (
	"fmt"
	"regexp"
	"strings"
)

var credentialID = regexp.MustCompile(`^conn_cr-[0-9]+$`)

func ValidateID(id string) error {
	if !credentialID.MatchString(id) {
		return fmt.Errorf("credential ID must have the form conn_cr-<number>")
	}
	return nil
}
func ValidateInput(id string, r *CredentialInput) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("credential input is required")
	}
	if r.Config.ConnectorType != "" && r.Config.ConnectorType != id {
		return fmt.Errorf("connectorType must match the credential ID")
	}
	if strings.TrimSpace(r.Config.RunTime) == "" {
		return fmt.Errorf("runTime is required (the UI uses 23:30)")
	}
	if len(r.Config.ConnectionDetails) == 0 {
		return fmt.Errorf("connection details are required; use Delete to clear a credential")
	}
	for _, detail := range r.Config.ConnectionDetails {
		if strings.TrimSpace(detail.Key) == "" {
			return fmt.Errorf("connection key is required")
		}
	}
	return nil
}
