package recommendations

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// List retrieves the recommendation collection shown by Forge. The UI performs
// search, category filtering, selection and sorting locally on this collection.
func (s *Service) List(ctx context.Context) (*[]Recommendation, *interfaces.Response, error) {
	var result []Recommendation
	response, err := s.client.Get(ctx, Endpoint, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateStatus changes a recommendation's lifecycle status and optional note.
// It returns the server's updated recommendation, without performing remediation.
func (s *Service) UpdateStatus(ctx context.Context, id string, request *UpdateStatusRequest) (*Recommendation, *interfaces.Response, error) {
	if err := validateUpdate(id, request); err != nil {
		return nil, nil, err
	}
	var result Recommendation
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id)+"/status", request, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
