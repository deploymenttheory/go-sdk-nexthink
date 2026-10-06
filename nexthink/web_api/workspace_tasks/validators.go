package workspace_tasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." {
		return fmt.Errorf("non-empty resource ID is required")
	}
	return nil
}
func validateTask(r *TaskRequest) error {
	if r == nil || strings.TrimSpace(r.Title) == "" || strings.TrimSpace(r.Prompt) == "" || r.Schedule.Type == "" {
		return fmt.Errorf("title, prompt and schedule type are required")
	}
	if len(r.ExpiresAt) > 0 {
		var value *string
		if err := json.Unmarshal(r.ExpiresAt, &value); err != nil {
			return fmt.Errorf("expires_at must be a JSON timestamp string or null: %w", err)
		}
	}
	return nil
}
