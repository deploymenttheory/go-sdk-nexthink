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
