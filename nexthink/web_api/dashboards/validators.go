package dashboards

import (
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
