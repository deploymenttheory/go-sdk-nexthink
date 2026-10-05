package validation

import (
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrJSONValue   = errors.New("a non-null JSON value is required")
	ErrPathSegment = errors.New("a nonempty single path segment is required")
	ErrJSONObject  = errors.New("a JSON object is required")
)

func PathSegment(value string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." ||
		strings.ContainsAny(value, "/\\") {
		return ErrPathSegment
	}
	return nil
}

func JSONObject(raw json.RawMessage) error {
	var object map[string]json.RawMessage
	if json.Unmarshal(raw, &object) != nil || object == nil {
		return ErrJSONObject
	}
	return nil
}

func JSONValue(raw json.RawMessage) error {
	if !json.Valid(raw) || strings.TrimSpace(string(raw)) == "null" {
		return ErrJSONValue
	}
	return nil
}
