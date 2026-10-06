package recommendations

import (
	"fmt"
	"strings"
)

func validateUpdate(id string, request *UpdateStatusRequest) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." {
		return fmt.Errorf("recommendation id is required")
	}
	if request == nil {
		return fmt.Errorf("request is required")
	}
	switch request.Status {
	case StatusNew, StatusInProgress, StatusDone, StatusDismissed:
		return nil
	default:
		return fmt.Errorf("status must be new, in-progress, done or dismissed")
	}
}
