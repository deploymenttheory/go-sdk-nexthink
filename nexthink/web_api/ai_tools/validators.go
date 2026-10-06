package ai_tools

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" || strings.ContainsAny(id, "/?#\\") || id == "." || id == ".." {
		return fmt.Errorf("id must be a nonempty path segment")
	}
	return nil
}
func validateRevision(revision int) error {
	if revision < 0 {
		return fmt.Errorf("revision must be nonnegative")
	}
	return nil
}
func validateLanguage(language string) error {
	if language != "" && language != "en" && language != "ja" {
		return fmt.Errorf("language must be empty, en, or ja")
	}
	return nil
}

var nqlIDPattern = regexp.MustCompile(`^#?[a-z0-9_]+$`)

func validateToolIdentity(name, nqlID, description string) error {
	if strings.TrimSpace(name) == "" || utf8.RuneCountInString(name) > 140 {
		return fmt.Errorf("name is required and must not exceed 140 characters")
	}
	if !nqlIDPattern.MatchString(nqlID) || len(nqlID) > 140 {
		return fmt.Errorf("nqlId must contain lowercase letters, digits, underscores, with an optional leading #")
	}
	if utf8.RuneCountInString(description) > 200 {
		return fmt.Errorf("description must not exceed 200 characters")
	}
	return nil
}
func validateToolRequest(request *ToolRequest) error {
	if request == nil {
		return fmt.Errorf("request is required")
	}
	return validateToolIdentity(request.Name, request.NQLID, request.Description)
}
func validateCopilotRequest(request *CopilotRequest) error {
	if request == nil {
		return fmt.Errorf("request is required")
	}
	return validateToolIdentity(request.Name, request.NQLID, request.Description)
}
func validateModule(request *Module) error {
	if request == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateCredentialsRequest(request *CredentialsRequest) error {
	if request == nil || strings.TrimSpace(request.CredentialsReference) == "" {
		return fmt.Errorf("credsRef is required")
	}
	return nil
}
func validateToolInsightsRequest(request *ToolInsightsRequest) error {
	if request == nil || strings.TrimSpace(request.ToolNQLID) == "" {
		return fmt.Errorf("toolNqlId is required")
	}
	return nil
}
func validateGoalRequest(request *GoalRequest) error {
	if request == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(request.Name) == "" || len(request.Metrics) == 0 {
		return fmt.Errorf("name and at least one metric are required")
	}
	if _, err := time.Parse(time.RFC3339Nano, request.Timeline.StartDate); err != nil {
		return fmt.Errorf("timeline.startDate must be an RFC3339 timestamp: %w", err)
	}
	if len(request.Timeline.EndDate) > 0 && string(request.Timeline.EndDate) != "null" {
		var end string
		if err := json.Unmarshal(request.Timeline.EndDate, &end); err != nil {
			return fmt.Errorf("timeline.endDate must be an RFC3339 JSON string or null: %w", err)
		}
		if _, err := time.Parse(time.RFC3339Nano, end); err != nil {
			return fmt.Errorf("timeline.endDate must be an RFC3339 timestamp: %w", err)
		}
	}

	return nil
}
