package user_classification

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"strings"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

// List returns fields used for user-organization classification.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, Endpoint, nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Replace replaces the entire tenant list; creation, edits and deletion are all
// represented by editing this collection, as in the UI. There are no per-item routes.
func (s *Service) Replace(ctx context.Context, request *ReplaceRequest) (*interfaces.Response, error) {
	if request == nil || request.CustomFields == nil {
		return nil, fmt.Errorf("customFields array is required (use an explicit empty array to clear)")
	}
	for _, field := range request.CustomFields {
		if strings.TrimSpace(field.NQLID) == "" || strings.TrimSpace(field.Name) == "" {
			return nil, fmt.Errorf("nqlId and name are required for each field")
		}
	}
	return s.client.Put(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json"}, nil)
}
