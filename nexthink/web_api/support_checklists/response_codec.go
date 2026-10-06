package support_checklists

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
func (v *ListPropertiesResponseItemValuesItem) UnmarshalJSON(data []byte) error {
	type wire ListPropertiesResponseItemValuesItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListPropertiesResponseItemValuesItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListPropertiesResponseItemValuesItem) MarshalJSON() ([]byte, error) {
	type wire ListPropertiesResponseItemValuesItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListPropertiesResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListPropertiesResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListPropertiesResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListPropertiesResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListPropertiesResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
func (v *ListResponseItem) UnmarshalJSON(data []byte) error {
	type wire ListResponseItem
	var decoded wire
	extra, err := decodeResponse(data, &decoded)
	if err != nil {
		return err
	}
	*v = ListResponseItem(decoded)
	v.AdditionalFields = extra
	return nil
}
func (v ListResponseItem) MarshalJSON() ([]byte, error) {
	type wire ListResponseItem
	return encodeResponse(wire(v), v.AdditionalFields)
}
