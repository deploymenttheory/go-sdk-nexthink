package nql_editor

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ValidateDocument(doc Document) error {
	if strings.TrimSpace(doc.URI) == "" || strings.TrimSpace(doc.LanguageID) == "" {
		return fmt.Errorf("document uri and languageId are required")
	}
	return nil
}

func ValidateValidationRequest(req *ValidationRequest) error {
	if req == nil {
		return fmt.Errorf("validation request is required")
	}
	if len(req.Rules) > 0 && !json.Valid(req.Rules) {
		return fmt.Errorf("rules must be valid JSON")
	}
	return ValidateDocument(req.Document)
}

func ValidatePositionRequest(req *PositionRequest) error {
	if req == nil {
		return fmt.Errorf("position request is required")
	}
	if req.Position.Line < 0 || req.Position.Character < 0 {
		return fmt.Errorf("position must not be negative")
	}
	return ValidateDocument(req.Document)
}

func ValidateCompletionItem(item CompletionItem) error {
	var label string
	if json.Unmarshal(item["label"], &label) != nil || label == "" {
		return fmt.Errorf("completion label is required")
	}
	for _, value := range item {
		if !json.Valid(value) {
			return fmt.Errorf("completion fields must be valid JSON")
		}
	}
	return nil
}
