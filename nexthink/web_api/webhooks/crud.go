package webhooks

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) ([]Webhook, *interfaces.Response, error) {
	var result []Webhook
	response, err := s.client.Get(ctx, Endpoint+"/all", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Webhook, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Webhook
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *Webhook) (*WriteResult, *interfaces.Response, error) {
	return s.save(ctx, request)
}
func (s *Service) Update(ctx context.Context, request *Webhook) (*WriteResult, *interfaces.Response, error) {
	return s.save(ctx, request)
}
func (s *Service) save(ctx context.Context, request *Webhook) (*WriteResult, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	response, err := s.client.Post(ctx, Endpoint, request, map[string]string{"Content-Type": "application/json"}, nil)
	if err != nil {
		return nil, response, err
	}
	return &WriteResult{Message: string(response.Body)}, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.DeleteWithBody(ctx, Endpoint+"/"+url.PathEscape(id), struct{}{}, map[string]string{"Content-Type": "application/json"}, nil)
}
func (s *Service) GetAvailability(ctx context.Context) (*Availability, *interfaces.Response, error) {
	var result Availability
	response, err := s.client.Get(ctx, Endpoint+"/available", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Test sends a test request immediately. External failures use HTTP 400/503;
// their exact response body remains available in Response alongside the error.
func (s *Service) Test(ctx context.Context, request *TestRequest) (*interfaces.Response, error) {
	if request == nil || request.ConnectorConfig.ConnectorType == "" || request.HTTPMethod == "" {
		return nil, fmt.Errorf("connector configuration and HTTP method are required")
	}
	return s.client.Post(ctx, EndpointTest, request, map[string]string{"Content-Type": "application/json"}, nil)
}

type WebhooksServiceInterface interface {
	List(context.Context) ([]Webhook, *interfaces.Response, error)
	Get(context.Context, string) (*Webhook, *interfaces.Response, error)
	Create(context.Context, *Webhook) (*WriteResult, *interfaces.Response, error)
	Update(context.Context, *Webhook) (*WriteResult, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	GetAvailability(context.Context) (*Availability, *interfaces.Response, error)
	Test(context.Context, *TestRequest) (*interfaces.Response, error)
}

var _ WebhooksServiceInterface = (*Service)(nil)
