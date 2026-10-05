package collector_management

import (
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateUpdateConfiguration(req json.RawMessage) error {
	if err := validation.JSONObject(req); err != nil {
		return err
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(req, &object); err != nil {
		return err
	}
	revision, ok := object["configRevision"]
	if !ok || string(revision) == "null" {
		return fmt.Errorf("configRevision is required")
	}
	return nil
}
