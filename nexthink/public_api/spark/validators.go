package spark

import (
	"fmt"
	"strings"
)

func ValidateHandoffRequest(req *HandoffRequest) error {
	if req == nil || len(req.Message.Parts) == 0 {
		return fmt.Errorf("message.parts must not be empty")
	}
	for i, p := range req.Message.Parts {
		switch p.Type {
		case "TEXT":
			if p.Text == "" || p.MIMEType != "" || p.FileContent != nil {
				return fmt.Errorf("parts[%d]: TEXT requires only text", i)
			}
		case "FILE":
			if p.MIMEType == "" || p.FileContent == nil || p.Text != "" {
				return fmt.Errorf("parts[%d]: FILE requires mimeType and fileContent", i)
			}
		default:
			return fmt.Errorf("parts[%d]: type must be TEXT or FILE", i)
		}
	}
	return nil
}

func ValidateHandoffHeaders(upn, timezone string) error {
	if strings.TrimSpace(upn) == "" || strings.ContainsAny(upn+timezone, "\r\n") {
		return fmt.Errorf("valid user principal name and headers required")
	}
	return nil
}
