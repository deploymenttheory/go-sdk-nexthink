package dex_configuration

import (
	"encoding/json"
	"fmt"
	"strings"
)

func validateManagementStrings(values ...string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("required request field is empty")
		}
	}
	return nil
}

func validateManagementUpdateApplications(r *UpdateApplicationsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Applications == nil {
		return fmt.Errorf("apps is required")
	}
	for _, app := range r.Applications {
		if err := validateManagementStrings(app.UUID); err != nil {
			return err
		}
	}
	return nil
}

func validateManagementGetScoreMetrics(r *GetScoreMetricsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}

func validateManagementUpdateScoreMetrics(r *UpdateScoreMetricsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Metrics == nil {
		return fmt.Errorf("ecScoreMetrics is required")
	}
	for _, metric := range r.Metrics {
		for _, threshold := range []json.RawMessage{metric.Average, metric.Frustrating} {
			if len(threshold) > 0 {
				var number *float64
				if err := json.Unmarshal(threshold, &number); err != nil {
					return fmt.Errorf("threshold must be a number or null: %w", err)
				}
			}
		}
	}
	return nil
}

func validateManagementEnableCampaign(r *EnableCampaignRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.NQLID); err != nil {
		return err
	}
	return nil
}
