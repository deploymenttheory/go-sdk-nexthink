package application_experience

import (
	"fmt"
	"strings"
	"time"
)

func validateApplicationInsights(r *ApplicationInsightsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ApplicationID) == "" || strings.TrimSpace(r.CurrentTimeframe) == "" || strings.TrimSpace(r.PreviousTimeframe) == "" {
		return fmt.Errorf("applicationId and both timeframes are required")
	}
	if _, err := time.LoadLocation(r.Timezone); err != nil || r.Timezone == "" {
		return fmt.Errorf("valid timezone is required")
	}
	if r.Breakdowns == nil {
		return fmt.Errorf("breakdowns are required")
	}
	return nil
}
