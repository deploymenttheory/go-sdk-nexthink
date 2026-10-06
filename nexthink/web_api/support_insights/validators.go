package support_insights

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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
func validateDates(start, end string) error {
	a, err := time.Parse("2006-01-02T15:04", start)
	if err != nil {
		return fmt.Errorf("start must use YYYY-MM-DDTHH:MM: %w", err)
	}
	b, err := time.Parse("2006-01-02T15:04", end)
	if err != nil {
		return fmt.Errorf("end must use YYYY-MM-DDTHH:MM: %w", err)
	}
	if !b.After(a) {
		return fmt.Errorf("end must be after start")
	}
	return nil
}
func validateInsightsRequest(r *InsightsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateDates(r.StartTime, r.EndTime); err != nil {
		return err
	}
	return nil
}
