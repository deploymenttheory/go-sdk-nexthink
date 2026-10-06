package nlp_assistant

import (
	"fmt"
	"strings"
)

func validateRequest(r *ChatRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id is required")
	}
	if r.Source != "search" && r.Source != "investigations" {
		return fmt.Errorf("source must be search or investigations")
	}
	if r.ConversationMode != "standard" && r.ConversationMode != "experimental" {
		return fmt.Errorf("conversation_mode must be standard or experimental")
	}
	if len(r.Messages) == 0 {
		return fmt.Errorf("messages are required")
	}
	for _, m := range r.Messages {
		if m.Author != "agent" && m.Author != "user" && m.Author != "system" {
			return fmt.Errorf("unknown message author %q", m.Author)
		}
		if strings.TrimSpace(m.Content) == "" {
			return fmt.Errorf("message content is required")
		}
	}
	return nil
}
