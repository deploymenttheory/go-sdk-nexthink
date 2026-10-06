package dashboards

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateRevision(revision int) error {
	if revision < 1 {
		return fmt.Errorf("revision must be positive")
	}
	return nil
}
func ValidateInput(r *DashboardInput) error {
	if r == nil || strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("dashboard title is required")
	}
	if r.DefaultTimeRange.Value < 1 || strings.TrimSpace(r.DefaultTimeRange.Unit) == "" {
		return fmt.Errorf("positive time range and unit are required")
	}
	if b := r.DefaultTimeRange.ByDuration; b != nil && (b.Value < 1 || strings.TrimSpace(b.Unit) == "") {
		return fmt.Errorf("positive by-duration value and unit are required")
	}
	return nil
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := ValidateID(r.ID); err != nil {
		return err
	}
	if err := ValidateRevision(r.Revision); err != nil {
		return err
	}
	return ValidateInput(&r.Dashboard)
}
func ValidateDeleteRequest(r *DeleteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := ValidateID(r.DashboardID); err != nil {
		return err
	}
	return ValidateRevision(r.Revision)
}

func validateMutationContext(id string, revision int) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	return ValidateRevision(revision)
}
func validateMutationRequest(request any) error {
	switch r := request.(type) {
	case *WidgetRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.DashboardID, r.Revision); err != nil {
			return err
		}
		if err := ValidateID(r.Widget.ID); err != nil {
			return err
		}
		if r.Widget.Type == "" || !json.Valid(r.Widget.Config) || len(r.Layout.Nodes) == 0 {
			return fmt.Errorf("widget type, config and nonempty layout are required")
		}
	case *DeleteWidgetRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.DashboardID, r.Revision); err != nil {
			return err
		}
		return ValidateID(r.ID)
	case *FilterRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.DashboardID, r.Revision); err != nil {
			return err
		}
		if err := ValidateID(r.Filter.ID); err != nil {
			return err
		}
		if r.Filter.Type == "" || !json.Valid(r.Filter.Config) {
			return fmt.Errorf("filter type and config are required")
		}
	case *DeleteFilterRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.Dashboard.ID, r.Dashboard.Revision); err != nil {
			return err
		}
		return ValidateID(r.FilterID)
	case *MutationContext:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		return validateMutationContext(r.DashboardID, r.Revision)
	case *UpdateTabRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.DashboardID, r.Revision); err != nil {
			return err
		}
		return ValidateID(r.TabID)
	case *UpdateTabsRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.DashboardID, r.Revision); err != nil {
			return err
		}
		if r.Tabs == nil {
			return fmt.Errorf("tabs are required")
		}
		for _, tab := range r.Tabs {
			if err := ValidateID(tab.ID); err != nil {
				return err
			}
		}
	case *UpdateLayoutRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		if err := validateMutationContext(r.ID, r.Revision); err != nil {
			return err
		}
		if r.ProductArea == "" || r.DashboardType == "" || len(r.Layout.Nodes) == 0 {
			return fmt.Errorf("productArea, dashboardType and layout are required")
		}
	case *DuplicateRequest:
		if r == nil {
			return fmt.Errorf("request is required")
		}
		return validateMutationContext(r.Dashboard.ID, r.Dashboard.Revision)
	case *ImportRequest:
		if r == nil || !json.Valid(r.Content.Dashboard) {
			return fmt.Errorf("dashboard export content is required")
		}
	}
	return nil
}
