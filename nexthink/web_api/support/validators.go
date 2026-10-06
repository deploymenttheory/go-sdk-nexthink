package support

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
func validateSearchRequest(r *SearchRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}
