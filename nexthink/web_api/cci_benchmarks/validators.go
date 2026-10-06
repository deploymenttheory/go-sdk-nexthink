package cci_benchmarks

import (
	"fmt"
	"strings"
	"time"
)

func validateQuery(r *QueryRequest) error {
	if r == nil || len(r.Queries) == 0 {
		return fmt.Errorf("at least one benchmark query is required")
	}
	for _, q := range r.Queries {
		if strings.TrimSpace(q.Source.Name) == "" || len(q.Metrics) == 0 {
			return fmt.Errorf("benchmark source and metrics are required")
		}
		for _, m := range q.Metrics {
			if strings.TrimSpace(m) == "" {
				return fmt.Errorf("metric cannot be empty")
			}
		}
	}
	if r.TimeZone != "" {
		if _, err := time.LoadLocation(r.TimeZone); err != nil {
			return fmt.Errorf("invalid time zone: %w", err)
		}
	}
	return nil
}
