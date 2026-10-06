package workspace_agents

import (
	"fmt"
	"strings"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." {
		return fmt.Errorf("non-empty resource ID is required")
	}
	return nil
}
func validateSkill(r *SkillRequest) error {
	if r == nil || strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.Instructions) == "" {
		return fmt.Errorf("name and instructions are required")
	}
	return nil
}
