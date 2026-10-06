package query_builder

import (
	"encoding/json"
	"fmt"
	"strings"
)

func validateQuery(query string) error {
	if strings.TrimSpace(query) == "" {
		return fmt.Errorf("query is required")
	}
	return nil
}
func validateTransform(r *TransformRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return validateQuery(r.Query)
}
func validateListDrilldownDestinations(r *DestinationsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return validateQuery(r.Query)
}
func validateTransformDrilldown(r *DrilldownRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateQuery(r.Query); err != nil {
		return err
	}
	if r.Conditions == nil {
		return nil
	}
	if len(r.Conditions.Identifiers) == 0 || r.Conditions.Values == nil {
		return fmt.Errorf("conditions identifiers and values are required")
	}
	for _, id := range r.Conditions.Identifiers {
		if strings.TrimSpace(id.Identifier) == "" || strings.TrimSpace(id.ObjectURI) == "" {
			return fmt.Errorf("condition identifier and objectURI are required")
		}
	}
	for _, row := range r.Conditions.Values {
		if len(row) != len(r.Conditions.Identifiers) {
			return fmt.Errorf("condition row must match identifiers")
		}
		for _, value := range row {
			if !json.Valid(value) {
				return fmt.Errorf("condition value must be valid JSON")
			}
		}
	}
	return nil
}
