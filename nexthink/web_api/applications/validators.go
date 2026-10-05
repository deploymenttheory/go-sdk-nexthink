package applications

import (
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." || strings.ContainsAny(id, "/\\") {
		return fmt.Errorf("application ID must be a nonempty path segment")
	}
	return nil
}
func ValidateInput(r *ApplicationInput, update bool) error {
	if r == nil {
		return fmt.Errorf("application is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.Category) == "" {
		return fmt.Errorf("name and category are required")
	}
	if r.Revision < 0 || (update && r.Revision == 0) {
		return fmt.Errorf("update requires a positive revision; create revision cannot be negative")
	}
	if len(r.URLs) == 0 && r.Desktop == nil && r.Network == nil {
		return fmt.Errorf("web URLs, desktop configuration or network configuration is required")
	}
	return nil
}
func ValidateListOptions(o *ListOptions) error {
	if o == nil {
		return nil
	}
	if o.PageNumber < 0 || o.PageSize < 0 {
		return fmt.Errorf("pagination values cannot be negative")
	}
	if o.SortName != "" && o.SortName != "asc" && o.SortName != "desc" {
		return fmt.Errorf("sort must be asc or desc")
	}
	return nil
}
func ValidateRevision(revision int) error {
	if revision < 1 {
		return fmt.Errorf("revision must be positive")
	}
	return nil
}
