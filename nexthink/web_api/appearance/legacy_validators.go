package appearance

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func validateLegacyAsset(r *SaveLegacyAssetRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateName(r.Name); err != nil {
		return err
	}
	for _, v := range []json.RawMessage{r.ID, r.Version} {
		var scalar any
		if json.Unmarshal(v, &scalar) != nil {
			return fmt.Errorf("id and version must be JSON string or number values from GetLegacyAsset")
		}
		switch value := scalar.(type) {
		case string:
			if value == "" {
				return fmt.Errorf("id and version cannot be empty")
			}
		case float64:
		default:
			return fmt.Errorf("id and version must be JSON string or number values from GetLegacyAsset")
		}
	}
	if strings.TrimSpace(r.Filename) == "" {
		return fmt.Errorf("filename is required")
	}
	if r.Blob == "" {
		return fmt.Errorf("base64 blob is required")
	}
	if _, err := base64.StdEncoding.DecodeString(r.Blob); err != nil {
		return fmt.Errorf("blob must contain base64 bytes without a data URL prefix")
	}
	return nil
}

// The frontend treats result.error.code as a JavaScript truthy value.
func legacyCodeTruthy(raw json.RawMessage) bool {
	var v any
	if json.Unmarshal(raw, &v) != nil {
		return false
	}
	switch value := v.(type) {
	case nil:
		return false
	case bool:
		return value
	case string:
		return value != ""
	case float64:
		return value != 0
	default:
		return true
	}
}
