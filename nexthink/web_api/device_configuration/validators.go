package device_configuration

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateRequest(req json.RawMessage) error { return validation.JSONObject(req) }

func ValidateSaveProfilesRequest(request *SaveProfilesRequest) error {
	if request == nil || request.Settings == nil {
		return fmt.Errorf("settings are required")
	}
	for _, setting := range request.Settings {
		if strings.TrimSpace(setting.ProfileID) == "" || strings.TrimSpace(setting.Name) == "" || !json.Valid(setting.NewValue) {
			return fmt.Errorf("profile ID, setting name and a valid JSON value are required")
		}
	}
	return nil
}
