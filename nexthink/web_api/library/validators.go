package library

import (
	"fmt"
	"strings"
)

func validateReference(name, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", name)
	}
	return nil
}
func validateInstallContent(r *ContentInstallationRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateReference("FileName", r.FileName); err != nil {
		return err
	}
	return nil
}
func validateInstallPack(r *Pack) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateReference("PackUUID", r.PackUUID); err != nil {
		return err
	}
	if err := validateReference("FileName", r.FileName); err != nil {
		return err
	}
	return nil
}
func validateUpdateContent(r *ContentUpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateReference("ContentID", r.ContentID); err != nil {
		return err
	}
	if err := validateReference("FileName", r.FileName); err != nil {
		return err
	}
	return nil
}
func validateGetDependenciesStatus(r *DependenciesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Dependencies == nil {
		return fmt.Errorf("dependencies is required")
	}
	return nil
}
func validateInstallDependencies(r *DependenciesRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Dependencies == nil {
		return fmt.Errorf("dependencies is required")
	}
	return nil
}
