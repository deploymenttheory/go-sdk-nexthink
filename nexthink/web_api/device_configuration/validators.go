package device_configuration

import (
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateRequest(req json.RawMessage) error { return validation.JSONObject(req) }
