package action_executions

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
func validateListOptions(r *ListOptions) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateDetailsRequest(r *DetailsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.NQLID) == "" {
		return fmt.Errorf("nql-id is required")
	}
	return nil
}
func validateNQLRequest(r *NQLRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
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
func validateExecuteRequest(r *ExecuteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.RemoteActionID) == "" {
		return fmt.Errorf("remoteActionId is required")
	}
	if (len(r.Targets) > 0) == (r.NQLQuery != "") {
		return fmt.Errorf("exactly one of targets or nqlQuery is required")
	}
	if r.NQLQuery != "" && (r.NQLTimeZone == "" || r.NQLTimeNow == "") {
		return fmt.Errorf("NQL targeting requires nqlTimeZone and nqlTimeNow")
	}
	for _, t := range r.Targets {
		if t.DeviceCollectorUID == "" {
			return fmt.Errorf("a remote-action target must identify a device")
		}
	}
	return nil
}
func validateHistoryRequest(r *HistoryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if len(r.Actions) == 0 {
		return fmt.Errorf("at least one action is required")
	}
	for _, a := range r.Actions {
		if strings.TrimSpace(a.ID) == "" || strings.TrimSpace(a.ActionType) == "" {
			return fmt.Errorf("action id and actionType are required")
		}
	}
	return nil
}
