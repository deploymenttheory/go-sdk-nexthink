package data_exporters

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateInput(r *ConfigurationInput) error {
	if r == nil {
		return fmt.Errorf("exporter request is required")
	}
	if err := ValidateID(r.UUID); err != nil {
		return err
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.ConnectorID) == "" || !strings.HasPrefix(r.NQLID, "#") {
		return fmt.Errorf("name, connectorId and a # prefixed nqlId are required")
	}
	if len(r.QueryInfos) == 0 {
		return fmt.Errorf("at least one query is required")
	}
	return nil
}
func ValidateTest(r *TestRequest) error {
	if r == nil {
		return fmt.Errorf("test request is required")
	}
	if err := ValidateID(r.UUID); err != nil {
		return err
	}
	if strings.TrimSpace(r.ConnectorID) == "" || strings.TrimSpace(r.QueryInfo.Text) == "" {
		return fmt.Errorf("connectorId and query text are required")
	}
	return nil
}
