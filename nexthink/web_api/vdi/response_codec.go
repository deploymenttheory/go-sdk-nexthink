package vdi

import (
	"encoding/json"
	"reflect"
	"strings"
)

// decodeResponse retains unmodeled feature fields without replacing typed known properties.
func decodeResponse(data []byte, value any) (map[string]json.RawMessage, error) {
	if err := json.Unmarshal(data, value); err != nil {
		return nil, err
	}
	var extra map[string]json.RawMessage
	if err := json.Unmarshal(data, &extra); err != nil {
		return nil, err
	}
	t := reflect.TypeOf(value).Elem()
	for i := 0; i < t.NumField(); i++ {
		key := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		if key != "-" && string(extra[key]) != "null" {
			delete(extra, key)
		}
	}
	return extra, nil
}
func encodeResponse(value any, extra map[string]json.RawMessage) ([]byte, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	if len(extra) == 0 {
		return data, nil
	}
	var result map[string]json.RawMessage
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	t := reflect.TypeOf(value)
	known := map[string]bool{}
	for i := 0; i < t.NumField(); i++ {
		key := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
		known[key] = true
	}
	for key, val := range extra {
		_, present := result[key]
		if !known[key] || (!present && string(val) == "null") {
			result[key] = val
		}
	}
	return json.Marshal(result)
}
func (v *ValidateHostnameResponse) UnmarshalJSON(data []byte) error {
	type wire ValidateHostnameResponse
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ValidateHostnameResponse(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ValidateHostnameResponse) MarshalJSON() ([]byte, error) {
	type wire ValidateHostnameResponse
	return encodeResponse(wire(v), v.AdditionalFields)
}
