package global_search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"io"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Search buffers and decodes all category events. It retains partial results
// alongside malformed-stream or in-band provider errors. HTTP metadata remains
// available with transport errors.
func (s *Service) Search(ctx context.Context, r *SearchRequest) (*SearchResponse, *interfaces.Response, error) {
	if err := validateRequest(r); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, r, map[string]string{"Accept": "application/x-json-stream", "Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	if response == nil {
		return nil, nil, fmt.Errorf("global search: missing response")
	}
	result, err := decode(response.Body)
	return result, response, err
}
func decode(body []byte) (*SearchResponse, error) {
	result := &SearchResponse{Events: []SearchEvent{}}
	decoder := json.NewDecoder(bytes.NewReader(body))
	var problems []error
	for {
		var event *SearchEvent
		err := decoder.Decode(&event)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			problems = append(problems, fmt.Errorf("global search stream: %w", err))
			break
		}
		if event == nil {
			problems = append(problems, fmt.Errorf("global search stream: null event"))
			break
		}
		result.Events = append(result.Events, *event)
		if event.ErrorResponse != nil {
			problems = append(problems, event.ErrorResponse)
		}
	}
	return result, errors.Join(problems...)
}
