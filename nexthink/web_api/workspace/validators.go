package workspace

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func validateID(id string) error {
	if strings.TrimSpace(id) == "" || id == "." || id == ".." {
		return fmt.Errorf("non-empty resource ID is required")
	}
	return nil
}
func pageQuery(o *PageOptions) (string, error) {
	if o == nil {
		return "", nil
	}
	if o.Limit < 0 {
		return "", fmt.Errorf("limit cannot be negative")
	}
	q := url.Values{}
	for k, v := range map[string]string{"before": o.Before, "after": o.After, "around": o.Around, "from": o.From} {
		if v != "" {
			q.Set(k, v)
		}
	}
	if o.Limit > 0 {
		q.Set("limit", strconv.Itoa(o.Limit))
	}
	return querySuffix(q), nil
}
func querySuffix(q url.Values) string {
	if len(q) == 0 {
		return ""
	}
	return "?" + q.Encode()
}
func validateChat(r *ChatRequest) error {
	if r == nil {
		return fmt.Errorf("request is required")
	}
	if err := validateID(r.ID); err != nil {
		return err
	}
	if len(r.Messages) == 0 && len(r.ResumeValue) == 0 {
		return fmt.Errorf("messages or resume_value is required")
	}
	switch r.Source {
	case "investigations", "search", "workspace", "embedded", "assignments", "prism":
	default:
		return fmt.Errorf("unsupported source %q", r.Source)
	}
	if r.ConversationMode != "standard" && r.ConversationMode != "experimental" {
		return fmt.Errorf("conversation_mode must be standard or experimental")
	}
	return nil
}
