package content_administration

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type ContentAdministrationServiceInterface interface {
	GetConfiguration(
		ctx context.Context,
		configuration string,
	) (*ConfigurationResponse, *interfaces.Response, error)
	List(ctx context.Context, configuration string) (*ListResponse, *interfaces.Response, error)
}

var _ ContentAdministrationServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetConfiguration calls the observed GET endpoint.
func (s *Service) GetConfiguration(
	ctx context.Context,
	configuration string,
) (*ConfigurationResponse, *interfaces.Response, error) {
	if err := ValidateConfiguration(configuration); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result ConfigurationResponse
	resp, err := s.client.Get(
		ctx,
		fmt.Sprintf(EndpointGetConfiguration, url.PathEscape(configuration)),
		nil,
		headers,
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// List calls the observed GET endpoint.
func (s *Service) List(
	ctx context.Context,
	configuration string,
) (*ListResponse, *interfaces.Response, error) {
	if err := ValidateConfiguration(configuration); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result ListResponse
	resp, err := s.client.Get(
		ctx,
		fmt.Sprintf(EndpointList, url.PathEscape(configuration)),
		nil,
		headers,
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
