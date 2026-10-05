package workflows

import (
	"fmt"
	"strings"
)

// ValidateUUID rejects empty identifiers; GraphQL transmits IDs as variables,
// so route escaping and path-segment restrictions do not apply.
func ValidateUUID(uuid string) error {
	if strings.TrimSpace(uuid) == "" {
		return fmt.Errorf("UUID is required")
	}
	return nil
}

func ValidateCreateRequest(request *CreateRequest) error {
	if request == nil {
		return fmt.Errorf("create request is required")
	}
	if request.Workflow.UUID != "" {
		return fmt.Errorf("create does not accept an existing workflow UUID")
	}
	if !strings.HasPrefix(request.Workflow.ID, "#") {
		return fmt.Errorf("custom workflow NQL ID must start with #")
	}
	return validateInput(&request.Workflow)
}

func ValidateUpdateRequest(request *WorkflowInput) error {
	if err := validateInput(request); err != nil {
		return err
	}
	return ValidateUUID(request.UUID)
}

func validateInput(request *WorkflowInput) error {
	if request == nil || strings.TrimSpace(request.ID) == "" ||
		strings.TrimSpace(request.Name) == "" ||
		strings.TrimSpace(request.Status) == "" {
		return fmt.Errorf("workflow ID, name and status are required")
	}
	return nil
}
