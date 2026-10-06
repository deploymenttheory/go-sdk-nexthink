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

func validatePortalSession(session *PortalSession) error { return session.Validate() }
func validateLegacySearch(r *LegacyDashboardSearchRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Search) == "" {
		return fmt.Errorf("search is required")
	}
	if r.MaxPersonalResults < 1 || r.MaxPublishedResults < 1 || r.MaxRoleBasedResults < 1 {
		return fmt.Errorf("all result limits must be positive")
	}
	return nil
}
