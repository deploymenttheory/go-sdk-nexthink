package campaigns

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ValidateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("content ID is required")
	}
	return nil
}
func ValidateCampaign(r *CampaignInput) error {
	if r == nil {
		return fmt.Errorf("campaign is required")
	}
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLID) == "" {
		return fmt.Errorf("name and nqlId are required")
	}
	if len(r.Questions) == 0 {
		return fmt.Errorf("at least one question is required")
	}
	return nil
}
func ValidateCreateRequest(r *CreateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if !strings.HasPrefix(r.Campaign.NQLID, "#") || len(r.Campaign.NQLID) < 2 {
		return fmt.Errorf("create requires a hash-prefixed nqlId")
	}
	if len(r.Metadata) > 0 && !json.Valid(r.Metadata) {
		return fmt.Errorf("metadata must be valid JSON")
	}
	return ValidateCampaign(&r.Campaign)
}
func ValidateUpdateRequest(r *UpdateRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := ValidateID(r.Campaign.ContentID); err != nil {
		return err
	}
	if strings.TrimSpace(r.Campaign.BCSUID) == "" || strings.TrimSpace(r.Campaign.Status) == "" {
		return fmt.Errorf("bcsUid and status are required")
	}
	return ValidateCampaign(&r.Campaign.CampaignInput)
}
func ValidateListOptions(o *ListOptions) error {
	if o != nil && ((o.PageNumber != nil && *o.PageNumber < 0) || (o.Offset != nil && *o.Offset < 0)) {
		return fmt.Errorf("pagination values cannot be negative")
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

func validateManagementUpdateBranding(r *UpdateBrandingRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.UUID); err != nil {
		return err
	}
	return nil
}

func validateManagementSetStatus(r *SetStatusRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.Status); err != nil {
		return err
	}
	return nil
}

func validateManagementGetByNQLID(r *GetByNQLIDRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.NQLID); err != nil {
		return err
	}
	return nil
}

func validateManagementGetFromLibrary(r *GetFromLibraryRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateManagementStrings(r.LibraryUUID); err != nil {
		return err
	}
	return nil
}

func validateManagementGetWithV6(r *GetWithV6Request) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
