package global_search

import (
	"fmt"
	"strings"
)

func validateRequest(r *SearchRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Search) == "" {
		return fmt.Errorf("search is required")
	}
	if r.MaxResults < 1 {
		return fmt.Errorf("maxResults must be positive")
	}
	return nil
}
