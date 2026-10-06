package monitors

import (
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("content/doc UUID is required")
	}
	return nil
}
func ValidateInput(r *MonitorInput) error {
	if r == nil {
		return fmt.Errorf("monitor is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLID) == "" || strings.TrimSpace(r.MonitorType) == "" || strings.TrimSpace(r.MetricType) == "" {
		return fmt.Errorf("name, nqlId, monitorType and metricType are required")
	}
	return nil
}
func ValidateCreateRequest(r *MonitorInput) error {
	if err := ValidateInput(r); err != nil {
		return err
	}
	if r.UUID != nil {
		return fmt.Errorf("create must omit uuid")
	}
	return nil
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := ValidateID(r.DocUUID); err != nil {
		return err
	}
	if r.Revision < 1 {
		return fmt.Errorf("revision must be positive")
	}
	if r.Monitor.UUID == nil || strings.TrimSpace(*r.Monitor.UUID) == "" {
		return fmt.Errorf("update requires the monitor uuid returned by Get")
	}
	return ValidateInput(&r.Monitor)
}
func ValidateDeleteRequest(r *DeleteRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.MonitorType != "METRIC" && r.MonitorType != "REAL_TIME" {
		return fmt.Errorf("monitorType must be METRIC or REAL_TIME")
	}
	return ValidateID(r.DocUUID)
}

func validateManagementStrings(values ...string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("required request field is empty")
		}
	}
	return nil
}

func validateManagementSetActivity(r *SetActivityRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.DocUUID, r.Activity); err != nil {
		return err
	}
	return nil
}

func validateManagementUpdateBuiltIn(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.DocUUID); err != nil {
		return err
	}
	return ValidateUpdateRequest(r)
}

func validateManagementGetMetadata(r *GetMetadataRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.NQLQuery); err != nil {
		return err
	}
	return nil
}

func validateManagementAnalyzeQuery(r *AnalyzeQueryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.NQLQuery); err != nil {
		return err
	}
	return nil
}

func validateManagementGetImpactQuery(r *GetImpactQueryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.Input.Query); err != nil {
		return err
	}
	return nil
}

func validateManagementListFilterFields(r *ListFilterFieldsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.Collection.CollectionURI, r.Collection.Type); err != nil {
		return err
	}
	return nil
}
