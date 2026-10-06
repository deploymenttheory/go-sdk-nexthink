package remote_actions

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

func ValidateScript(script []byte) error {
	if len(script) == 0 {
		return fmt.Errorf("script bytes are required")
	}
	return nil
}

func ValidateCreateRequest(request *RemoteActionInput) error {
	if err := ValidateUpdateRequest(request); err != nil {
		return err
	}
	if !strings.HasPrefix(request.ID, "#") {
		return fmt.Errorf("custom remote action NQL ID must start with #")
	}
	return nil
}

func ValidateUpdateRequest(request *RemoteActionInput) error {
	if request == nil || strings.TrimSpace(request.ID) == "" ||
		strings.TrimSpace(request.Name) == "" {
		return fmt.Errorf("remote action NQL ID and name are required")
	}
	return nil
}

// ValidateNQLID checks the management identifier used by Update and Delete.
func ValidateNQLID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("remote action NQL ID is required")
	}
	return nil
}

func validateManagementStrings(values ...string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("required request field is empty")
		}
	}
	return nil
}

func validateManagementExport(r *ExportRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.ID); err != nil {
		return err
	}
	return nil
}

func validateManagementImport(r *ImportRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.RemoteAction.ID); err != nil {
		return err
	}
	return ValidateCreateRequest(&r.RemoteAction)
}

func validateManagementGetFromLibrary(r *GetFromLibraryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.ID); err != nil {
		return err
	}
	return nil
}
