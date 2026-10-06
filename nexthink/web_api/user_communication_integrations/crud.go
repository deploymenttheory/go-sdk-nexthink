package user_communication_integrations

import (
	"context"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) ([]Integration, *interfaces.Response, error) {
	var result []Integration
	response, err := s.client.Get(ctx, Endpoint, nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}
func (s *Service) Get(ctx context.Context, id string) (*Integration, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	var result Integration
	response, err := s.client.Get(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *IntegrationInput) (*Integration, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Integration
	response, err := s.client.Post(ctx, Endpoint, request, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Update(ctx context.Context, id string, request *IntegrationInput) (*Integration, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	var result Integration
	response, err := s.client.Put(ctx, Endpoint+"/"+url.PathEscape(id), request, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Delete(ctx context.Context, id string) (*interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, err
	}
	return s.client.Delete(ctx, Endpoint+"/"+url.PathEscape(id), nil, nil, nil)
}
func (s *Service) ListAzureConnectors(ctx context.Context) ([]AzureConnector, *interfaces.Response, error) {
	var result []AzureConnector
	response, err := s.client.Get(ctx, Endpoint+"/externals/azureconnectors", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

type UserCommunicationIntegrationsServiceInterface interface {
	List(context.Context) ([]Integration, *interfaces.Response, error)
	Get(context.Context, string) (*Integration, *interfaces.Response, error)
	Create(context.Context, *IntegrationInput) (*Integration, *interfaces.Response, error)
	Update(context.Context, string, *IntegrationInput) (*Integration, *interfaces.Response, error)
	Delete(context.Context, string) (*interfaces.Response, error)
	ListAzureConnectors(context.Context) ([]AzureConnector, *interfaces.Response, error)
}

var _ UserCommunicationIntegrationsServiceInterface = (*Service)(nil)
