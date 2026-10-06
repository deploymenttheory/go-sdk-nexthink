package nlp_assistant

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

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Chat waits for the HTTP response to complete, then decodes the SSE events.
// Partial events are returned with any decoding or server-reported stream error.
func (s *Service) Chat(ctx context.Context, r *ChatRequest) (*ChatResponse, *interfaces.Response, error) {
	if err := validateRequest(r); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, r, map[string]string{"Accept": "text/event-stream", "Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	if response == nil {
		return nil, nil, fmt.Errorf("NLP assistance: missing response")
	}
	result, err := decode(response.Body)
	return result, response, err
}
func decode(body []byte) (*ChatResponse, error) {
	result := &ChatResponse{Events: []Event{}}
	body = bytes.TrimPrefix(body, []byte{0xef, 0xbb, 0xbf})
	body = bytes.ReplaceAll(bytes.ReplaceAll(body, []byte("\r\n"), []byte("\n")), []byte("\r"), []byte("\n"))
	reader := bufio.NewReader(bytes.NewReader(body))
	var data []string
	eventType, id := "", ""
	var problems []error
	dispatch := func() {
		if len(data) == 0 {
			eventType = ""
			return
		}
		raw := []byte(strings.Join(data, "\n"))
		if !json.Valid(raw) {
			problems = append(problems, fmt.Errorf("NLP assistance: invalid JSON in %q event", eventType))
			data = nil
			eventType = ""
			return
		}
		if eventType == "" {
			eventType = "message"
		}
		event := Event{Type: eventType, ID: id, Data: raw}
		result.Events = append(result.Events, event)
		if eventType == "error" {
			decoded, err := event.Decode()
			if err != nil {
				problems = append(problems, err)
			} else {
				problems = append(problems, &StreamError{Status: decoded.Status, Message: decoded.Message})
			}
		}
		data = nil
		eventType = ""
	}
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			if line == "" {
				dispatch()
			} else if !strings.HasPrefix(line, ":") {
				field, value, found := strings.Cut(line, ":")
				if !found {
					value = ""
				}
				value = strings.TrimPrefix(value, " ")
				switch field {
				case "event":
					eventType = value
				case "data":
					data = append(data, value)
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
	// SSE only dispatches events terminated by a blank line. An incomplete final
	// event is an error rather than silently accepting a truncated response.
	if len(data) > 0 {
		problems = append(problems, fmt.Errorf("NLP assistance: truncated SSE event"))
	}
	if len(result.Events) == 0 && len(problems) == 0 && len(bytes.TrimSpace(body)) > 0 {
		problems = append(problems, fmt.Errorf("NLP assistance: response contained no SSE data events"))
	}
	return result, errors.Join(problems...)
}
