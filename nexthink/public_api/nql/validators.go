package nql

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateExecuteRequest validates an NQL execute request
func ValidateExecuteRequest(req *ExecuteRequest) error {
	if req == nil {
		return fmt.Errorf("execute request cannot be nil")
	}

	if err := validateQueryID(req.QueryID); err != nil {
		return err
	}

	return nil
}

// ValidateExportRequest validates an NQL export request
func ValidateExportRequest(req *ExportRequest) error {
	if req == nil {
		return fmt.Errorf("export request cannot be nil")
	}

	if err := validateQueryID(req.QueryID); err != nil {
		return err
	}

	if req.Compression != "" && req.Compression != "NONE" && req.Compression != "GZIP" && req.Compression != "ZSTD" {
		return fmt.Errorf("compression must be NONE, GZIP, or ZSTD")
	}

	if req.Format != "" && req.Format != ExportFormatCSV && req.Format != ExportFormatJSON {
		return fmt.Errorf("format must be either 'csv' or 'json', got: %s", req.Format)
	}

	return nil
}

// ValidateExportID validates an export ID
func ValidateExportID(exportID string) error {
	if exportID == "" {
		return fmt.Errorf("export ID cannot be empty")
	}

	if len(exportID) > MaxExportIDLength {
		return fmt.Errorf("export ID exceeds maximum length of %d characters", MaxExportIDLength)
	}

	return nil
}

// validateQueryID validates a query ID
func validateQueryID(queryID string) error {
	if queryID == "" {
		return fmt.Errorf("query ID is required")
	}

	if !strings.HasPrefix(queryID, "#") {
		return fmt.Errorf("query ID must start with '#', got: %s", queryID)
	}

	if len(queryID) > MaxQueryIDLength {
		return fmt.Errorf("query ID exceeds maximum length of %d characters", MaxQueryIDLength)
	}

	if !regexp.MustCompile(`^#[a-z0-9_]{2,254}$`).MatchString(queryID) {
		return fmt.Errorf("query ID must contain 2 to 254 lowercase letters, digits, or underscores after #")
	}

	return nil
}
