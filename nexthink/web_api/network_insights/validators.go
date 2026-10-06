package network_insights

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// requestVariables preserves omitted optional values and validates GraphQL non-null variables before transport.
func requestVariables(request any) (map[string]any, error) {
	value := reflect.ValueOf(request)
	if !value.IsValid() || value.IsNil() {
		return nil, fmt.Errorf("request is required")
	}
	value = value.Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if typ.Field(i).Tag.Get("required") != "true" {
			continue
		}
		switch field.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Map:
			if field.IsNil() {
				return nil, fmt.Errorf("%s is required", typ.Field(i).Name)
			}
		case reflect.String:
			if field.Len() == 0 {
				return nil, fmt.Errorf("%s is required", typ.Field(i).Name)
			}
		}
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("encode request: %w", err)
	}
	var result map[string]any
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
