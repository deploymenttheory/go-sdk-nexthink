package custom_fields

import (
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("docUid is required")
	}
	return nil
}
func ValidateType(t string) error {
	if t != "MANUAL" && t != "COMPUTED" {
		return fmt.Errorf("type must be MANUAL or COMPUTED; rule-based fields use REST")
	}
	return nil
}
func validateCommon(name, id, typ string) error {
	if strings.TrimSpace(name) == "" || !strings.HasPrefix(id, "#") || len(id) < 2 {
		return fmt.Errorf("name and hash-prefixed nqlId are required")
	}
	return ValidateType(typ)
}
func ValidateCreateRequest(r *CreateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateCommon(r.Name, r.NQLID, r.Type); err != nil {
		return err
	}
	if strings.TrimSpace(r.FieldDataType) == "" || strings.TrimSpace(r.DataModelObject) == "" {
		return fmt.Errorf("dataModelObject and fieldDataType are required")
	}
	if r.Type == "COMPUTED" && strings.TrimSpace(r.NQLQuery) == "" {
		return fmt.Errorf("computed fields require nqlQuery")
	}
	return nil
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateCommon(r.Name, r.NQLID, r.Type); err != nil {
		return err
	}
	if r.Revision < 1 || strings.TrimSpace(r.DataModelObject) == "" {
		return fmt.Errorf("positive revision and dataModelObject are required")
	}
	if r.Type == "COMPUTED" && strings.TrimSpace(r.NQLQuery) == "" {
		return fmt.Errorf("computed fields require nqlQuery")
	}
	return ValidateID(r.DocUID)
}
func ValidateDeleteRequest(r *DeleteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateCommon(r.Name, r.NQLID, r.Type); err != nil {
		return err
	}
	if r.Revision < 1 {
		return fmt.Errorf("positive revision is required")
	}
	switch r.DataModelObject {
	case "Device", "User", "Binary", "Package":
	default:
		return fmt.Errorf("delete dataModelObject must be Device, User, Binary or Package")
	}
	return ValidateID(r.DocUID)
}
