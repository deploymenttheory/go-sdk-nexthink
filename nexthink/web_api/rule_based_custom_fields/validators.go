package rule_based_custom_fields

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *FieldInput, update bool) error {
	if r == nil {
		return fmt.Errorf("field is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLID) == "" || strings.TrimSpace(r.InventoryObject) == "" || r.Type != "RULE_BASED" {
		return fmt.Errorf("name, nqlId, inventoryObject and RULE_BASED type are required")
	}
	if update && (r.Revision == nil || *r.Revision < 1) {
		return fmt.Errorf("update requires a positive revision")
	}
	if len(r.TagConditions) == 0 {
		return fmt.Errorf("at least one rule is required")
	}
	for _, rule := range r.TagConditions {
		if strings.TrimSpace(rule.Label) == "" || strings.TrimSpace(rule.Statement) == "" {
			return fmt.Errorf("rule label and statement are required")
		}
	}
	return nil
}
func ValidateDeleteRequest(r *DeleteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLID) == "" || r.Type != "RULE_BASED" || r.Revision < 1 {
		return fmt.Errorf("name, nqlId, RULE_BASED type and positive revision are required")
	}
	switch r.InventoryObject {
	case "Device", "User", "Binary", "Package":
		return nil
	default:
		return fmt.Errorf("delete inventoryObject must be Device, User, Binary or Package")
	}
}
