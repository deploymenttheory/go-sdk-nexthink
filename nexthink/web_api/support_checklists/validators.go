package support_checklists

import (
	"fmt"
	"strings"
)

func validateID(value string) error {
	if strings.TrimSpace(value) == "" || value == "." || value == ".." || strings.ContainsAny(value, "/\\?#\r\n") {
		return fmt.Errorf("a nonempty path identifier without path separators is required")
	}
	return nil
}
