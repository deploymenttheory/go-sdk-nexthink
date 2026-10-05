package software_metering

import (
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("uuid is required")
	}
	return nil
}
func ValidateConfiguration(r *ConfigurationInput) error {
	if r == nil {
		return fmt.Errorf("configuration is required")
	}
	if strings.TrimSpace(r.Name) == "" || !strings.HasPrefix(r.NQLID, "#") || len(r.NQLID) < 2 {
		return fmt.Errorf("name and hash-prefixed nqlId are required")
	}
	if r.LicenseType != "USER" && r.LicenseType != "DEVICE" {
		return fmt.Errorf("licenseType must be USER or DEVICE")
	}
	if len(r.ApplicationUUIDs) == 0 || len(r.Thresholds) == 0 {
		return fmt.Errorf("applications and thresholds are required")
	}
	for _, id := range r.ApplicationUUIDs {
		if err := ValidateID(id); err != nil {
			return err
		}
	}
	for _, t := range r.Thresholds {
		if t.AppType != "WEB" && t.AppType != "DESKTOP" {
			return fmt.Errorf("threshold appType must be WEB or DESKTOP")
		}
		if t.MetricType != "FOCUS" && t.MetricType != "EXECUTION" {
			return fmt.Errorf("threshold metricType must be FOCUS or EXECUTION")
		}
		if t.TimeUnit != "DAY" && t.TimeUnit != "HOUR" && t.TimeUnit != "MINUTE" {
			return fmt.Errorf("threshold timeUnit must be DAY, HOUR or MINUTE")
		}
		if t.Value < 0 {
			return fmt.Errorf("threshold value cannot be negative")
		}
	}
	return nil
}
func ValidateUpdateRequest(id string, r *ConfigurationInput) error {
	if err := ValidateID(id); err != nil {
		return err
	}
	return ValidateConfiguration(r)
}
