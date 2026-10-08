package workspace

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"io"
	"strings"
)

// Chat creates or continues a Workspace conversation. It waits for the HTTP body
// to complete and returns decoded SSE events, including partial events on error.
// The shared transport response-body limit applies to the buffered stream.
// Cancel ctx to stop waiting; CancelConversation also cancels server-side work.
// Prompts may invoke tools according to the account's configured capabilities.
func (s *Service) Chat(ctx context.Context, request *ChatRequest) (*ChatResponse, *interfaces.Response, error) {
	if err := validateChat(request); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, ChatEndpoint, request, map[string]string{"Content-Type": "application/json", "Accept": "text/event-stream"}, nil)
	if err != nil {
		if response != nil && response.StatusCode >= 200 && response.StatusCode < 300 && len(response.Body) > 0 {
			result, decodeErr := decodeEvents(response.Body)
			return result, response, errors.Join(err, decodeErr)
		}
		return nil, response, err
	}
	if response == nil {
		return nil, nil, fmt.Errorf("workspace: missing chat response")
	}
	result, err := decodeEvents(response.Body)
	return result, response, err
}
func decodeEvents(body []byte) (*ChatResponse, error) {
	result := &ChatResponse{Events: []Event{}}
	body = bytes.TrimPrefix(body, []byte{0xef, 0xbb, 0xbf})
	body = bytes.ReplaceAll(bytes.ReplaceAll(body, []byte("\r\n"), []byte("\n")), []byte("\r"), []byte("\n"))
	reader := bufio.NewReader(bytes.NewReader(body))
	var data []string
	kind, id := "", ""
	var problems []error
	dispatch := func() {
		if len(data) == 0 {
			kind = ""
			return
		}
		raw := strings.Join(data, "\n")
		data = nil
		if raw == "[DONE]" {
			kind = ""
			return
		}
		if !json.Valid([]byte(raw)) {
			problems = append(problems, fmt.Errorf("workspace: invalid JSON in SSE event"))
			kind = ""
			return
		}
		if kind == "" {
			kind = "message"
		}
		result.Events = append(result.Events, Event{Type: kind, ID: id, Data: json.RawMessage(raw)})
		var payload struct {
			Type      string `json:"type"`
			ErrorText string `json:"errorText"`
			Message   string `json:"message"`
		}
		_ = json.Unmarshal([]byte(raw), &payload)
		if kind == "error" || payload.Type == "error" {
			problems = append(problems, fmt.Errorf("workspace chat stream error: %s %s", payload.ErrorText, payload.Message))
		}
		kind = ""
	}
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSuffix(line, "\n")
			if line == "" {
				dispatch()
			} else if !strings.HasPrefix(line, ":") {
				field, value, _ := strings.Cut(line, ":")
				value = strings.TrimPrefix(value, " ")
				switch field {
				case "data":
					data = append(data, value)
				case "event":
					kind = value
				case "id":
					if !strings.ContainsRune(value, 0) {
						id = value
					}
				}
			}
		}
		if err != nil {
			if !errors.Is(err, io.EOF) {
				problems = append(problems, err)
			}
			break
		}
	}
	if len(data) > 0 {
		problems = append(problems, fmt.Errorf("workspace: truncated SSE event"))
	}
	if len(result.Events) == 0 && len(problems) == 0 {
		problems = append(problems, fmt.Errorf("workspace: response contained no SSE JSON events"))
	}
	return result, errors.Join(problems...)
}
