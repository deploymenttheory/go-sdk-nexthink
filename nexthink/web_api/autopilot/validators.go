package autopilot

import (
	"fmt"
	"strings"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("id is required")
	}
	return nil
}
func validateUpdateWebSearch(r *BooleanValue) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateUpdateFilesystemTool(r *BooleanValue) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	return nil
}
func validateUpdateAgentName(r *StringValue) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.Value) == "" {
		return fmt.Errorf("agent name is required")
	}
	return nil
}
func validateSaveSettings(r *Settings) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Revision < 0 {
		return fmt.Errorf("revision must not be negative")
	}
	return nil
}
func validateReplaceWebSearchDomains(r *WebSearchDomainsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Revision != nil && *r.Revision < 0 {
		return fmt.Errorf("revision must not be negative")
	}
	if r.Domains == nil {
		return fmt.Errorf("domains is required (empty array clears domains)")
	}
	for _, d := range r.Domains {
		if strings.TrimSpace(d.Host) == "" {
			return fmt.Errorf("domain host is required")
		}
	}
	return nil
}
func validateCreateApproval(r *Approval) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateID(r.ID); err != nil {
		return err
	}
	if strings.TrimSpace(r.ActionType) == "" {
		return fmt.Errorf("actionType is required")
	}
	return nil
}
func validateUpdateApproval(r *ApprovalInput) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ActionType) == "" {
		return fmt.Errorf("actionType is required")
	}
	return nil
}
func validateCreateTicket(r *TicketRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.ConversationID == "" || r.UserSID == "" || r.DeviceLicenseID == "" {
		return fmt.Errorf("conversationId, userSid and deviceLicenseId are required")
	}
	if r.CategorizationRowIndex != nil && *r.CategorizationRowIndex < 0 {
		return fmt.Errorf("categorizationRowIndex must not be negative")
	}
	return nil
}
func validateUpdateAgentActionInputs(r *AgentActionInputsRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if r.Inputs == nil {
		return fmt.Errorf("inputs is required (empty array clears inputs)")
	}
	for _, i := range r.Inputs {
		if i.ID == "" {
			return fmt.Errorf("input id is required")
		}
		if i.Options == nil {
			return fmt.Errorf("input options must be an array")
		}
	}
	return nil
}
