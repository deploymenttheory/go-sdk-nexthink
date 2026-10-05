// Package spark hands off conversations to Nexthink Spark in Microsoft Teams.
// API reference: https://docs.nexthink.com/api/spark/handoff-api
package spark

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

const EndpointHandoff = "/api/v1/spark/handoff"

// Part is a TEXT or FILE message part. FileContent is the API's encoded content string.
type Part struct {
	Type        string  `json:"type"`
	Text        string  `json:"text,omitempty"`
	MIMEType    string  `json:"mimeType,omitempty"`
	FileContent *string `json:"fileContent,omitempty"`
}
type Message struct {
	Parts []Part `json:"parts"`
}
type HandoffRequest struct {
	Message  Message           `json:"message"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

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

// Handoff sends a conversation to the specified user's Teams account.
// Successful requests return HTTP 204 with no JSON response body.
func (s *Service) Handoff(ctx context.Context, upn, timezone string, req *HandoffRequest) (*interfaces.Response, error) {
	if strings.TrimSpace(upn) == "" || strings.ContainsAny(upn+timezone, "\r\n") {
		return nil, fmt.Errorf("valid user principal name and headers required")
	}
	if err := ValidateHandoffRequest(req); err != nil {
		return nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json", "User-Principal-Name": upn}
	if timezone != "" {
		headers["Timezone"] = timezone
	}
	return s.client.Post(ctx, EndpointHandoff, req, headers, nil)
}
