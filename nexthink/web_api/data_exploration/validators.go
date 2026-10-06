package data_exploration

import (
	"fmt"
	"strings"
)

func validateRequest(request any) error {
	switch r := request.(type) {
	case *QueryRequest:
		if r == nil || strings.TrimSpace(r.Query) == "" {
			return fmt.Errorf("query is required")
		}
	case *QueryInput:
		if r == nil || strings.TrimSpace(r.Query) == "" {
			return fmt.Errorf("query is required")
		}
	case *FieldsRequest:
		if r == nil || (r.CollectionURI == "" && len(r.FieldURIs) == 0) {
			return fmt.Errorf("collection or fields are required")
		}
	case *SystemRatingsRequest:
		if r == nil || len(r.FieldURIs) == 0 {
			return fmt.Errorf("fields are required")
		}
	case *FilterValuesRequest:
		if r == nil || r.Input.CollectionURI == "" || r.Input.FieldURI == "" {
			return fmt.Errorf("collection and field are required")
		}
	case *BreakdownFieldsRequest:
		if r == nil || r.NQLVariables.Query == "" {
			return fmt.Errorf("query is required")
		}
	case *BreakdownInsightsRequest:
		if r == nil || r.NQLVariables.Query == "" {
			return fmt.Errorf("query is required")
		}
	case *ByDurationsRequest:
		if r == nil || r.TimeRange.Type == "" {
			return fmt.Errorf("time range is required")
		}
	case *OrganisationRequest:
		if r == nil || r.CollectionURI == "" {
			return fmt.Errorf("collection is required")
		}
	case *ItemMetaRequest:
		if r == nil || r.FieldURI.DMURI == "" || r.Item == "" {
			return fmt.Errorf("field and item are required")
		}
	case *MenuRequest:
		if r == nil || len(r.Inputs) == 0 {
			return fmt.Errorf("menu inputs are required")
		}
	}
	return nil
}
