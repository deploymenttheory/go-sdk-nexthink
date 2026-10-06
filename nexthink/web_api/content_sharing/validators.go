package content_sharing

import (
	"fmt"
	"strings"
)

func required(values ...string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("content key, resource and identifiers must not be empty")
		}
	}
	return nil
}
func validateActions(o *ActionsOptions) error {
	if o == nil {
		return fmt.Errorf("options required")
	}
	return required(o.ContentKey, o.ResourceName)
}
func validateProfiles(o *ProfilesOptions) error {
	if o == nil {
		return fmt.Errorf("options required")
	}
	return required(o.ContentKey, o.ResourceName, o.ContentID)
}
func validateLegacy(o *LegacyOptions) error {
	if o == nil {
		return fmt.Errorf("options required")
	}
	return required(o.Service, o.ContentID)
}
