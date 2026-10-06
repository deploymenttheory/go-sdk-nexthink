package amplify

import (
	"fmt"
	"strings"
)

func validateSearch(request *SearchRequest) error {
	if request == nil || strings.TrimSpace(request.Keyword) == "" {
		return fmt.Errorf("keyword is required")
	}
	return nil
}

func validateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("identifier is required")
	}
	return nil
}

func validateConfiguration(request *ConfigurationRequest) error {
	if request == nil || request.ITSMConfigList == nil {
		return fmt.Errorf("itsmConfigList is required; use an explicit empty array to clear")
	}
	for _, application := range request.ITSMConfigList {
		if strings.TrimSpace(application.ITSMURL) == "" {
			return fmt.Errorf("itsmUrl is required")
		}
		if application.Substitution != "" && application.ConfigurationItemRegex == "" {
			return fmt.Errorf("configurationItemRegex is required with substitution")
		}
	}
	return nil
}
