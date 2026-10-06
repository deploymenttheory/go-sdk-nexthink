package workflow_executions

import (
	"encoding/json"
	"fmt"
	"strings"
)

func validateID(value string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." || strings.ContainsAny(value, "/\\?#\r\n") {
		return fmt.Errorf("a nonempty path identifier without path separators is required")
	}
	return nil
}
func queryParameters(value any) (map[string]string, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	var values map[string]json.RawMessage
	if err = json.Unmarshal(encoded, &values); err != nil {
		return nil, err
	}
	result := map[string]string{}
	for k, v := range values {
		if string(v) == "null" {
			continue
		}
		var str string
		if json.Unmarshal(v, &str) == nil {
			if str != "" {
				result[k] = str
			}
		} else {
			result[k] = string(v)
		}
	}
	return result, nil
}
func validateTimelineOptions(r *TimelineOptions) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ExecutionStatus) == "" {
		return fmt.Errorf("executionStatus is required")
	}
	return nil
}
func validateListOptions(r *ListOptions) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateExecuteRequest(r *ExecuteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.WorkflowUUID) == "" {
		return fmt.Errorf("workflowUuid is required")
	}
	if len(r.Targets) == 0 {
		return fmt.Errorf("at least one target is required")
	}
	for _, t := range r.Targets {
		if t.DeviceCollectorUID == "" && t.UserSID == "" {
			return fmt.Errorf("a target must identify a device or user")
		}
	}
	return nil
}
func validateExecuteNQLRequest(r *ExecuteNQLRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.WorkflowUUID) == "" {
		return fmt.Errorf("workflowUuid is required")
	}
	if strings.TrimSpace(r.NQLQuery) == "" {
		return fmt.Errorf("nqlQuery is required")
	}
	if strings.TrimSpace(r.NQLTimeZone) == "" {
		return fmt.Errorf("nqlTimeZone is required")
	}
	if strings.TrimSpace(r.NQLTimeNow) == "" {
		return fmt.Errorf("nqlTimeNow is required")
	}
	return nil
}
