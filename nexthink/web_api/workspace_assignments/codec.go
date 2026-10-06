package workspace_assignments

import (
	"encoding/json"
	"reflect"
	"strings"
)

// The codec preserves unknown fields and the distinction between omitted and null
// values while allowing callers to edit typed fields before serializing a response.
func (v *Assignment) UnmarshalJSON(data []byte) error {
	type wire Assignment
	var decoded wire
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	decoded.present = map[string]json.RawMessage{}
	t := reflect.TypeOf(decoded)
	for i := 0; i < t.NumField(); i++ {
		key := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if key == "" || key == "-" {
			continue
		}
		if _, ok := fields[key]; ok {
			decoded.present[key] = fields[key]
			delete(fields, key)
		}
	}
	decoded.AdditionalFields = fields
	*v = Assignment(decoded)
	return nil
}
func (v Assignment) MarshalJSON() ([]byte, error) {
	fields := map[string]json.RawMessage{}
	for k, data := range v.AdditionalFields {
		fields[k] = data
	}
	value := reflect.ValueOf(v)
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		key := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
		if key == "" || key == "-" {
			continue
		}
		field := value.Field(i)
		if _, exists := v.present[key]; !exists && field.IsZero() {
			continue
		}
		if original, exists := v.present[key]; exists && string(original) == "null" && field.IsZero() {
			fields[key] = json.RawMessage("null")
			continue
		}
		data, err := json.Marshal(field.Interface())
		if err != nil {
			return nil, err
		}
		fields[key] = data
	}
	return json.Marshal(fields)
}
