package workspace_assignments

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." {
		return fmt.Errorf("non-empty resource ID is required")
	}
	return nil
}
func querySuffix(q url.Values) string {
	if len(q) == 0 {
		return ""
	}
	return "?" + q.Encode()
}
func assignmentQuery(o *AssignmentOptions) string {
	q := url.Values{}
	if o != nil {
		for _, v := range o.Sources {
			q.Add("sources", v)
		}
		if o.Sort != "" {
			q.Set("sort", o.Sort)
		}
	}
	return querySuffix(q)
}
func validateAssignment(r *AssignmentUpdate) error {
	if r == nil || r.Revision < 0 || (r.State == "" && len(r.AssigneeID) == 0) {
		return fmt.Errorf("revision and a state or assigneeId update are required")
	}
	if r.State != "" && r.State != "open" && r.State != "done" && r.State != "dismissed" {
		return fmt.Errorf("state must be open, done or dismissed")
	}
	if len(r.AssigneeID) > 0 {
		var v *string
		if err := json.Unmarshal(r.AssigneeID, &v); err != nil {
			return fmt.Errorf("assigneeId must be a JSON string or null: %w", err)
		}
	}
	return nil
}
