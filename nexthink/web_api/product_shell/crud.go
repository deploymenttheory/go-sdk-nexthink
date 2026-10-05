package product_shell

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type ProductShellServiceInterface interface {
	GetMenu(ctx context.Context) (*MenuResponse, *interfaces.Response, error)
	GetUser(ctx context.Context) (*UserResponse, *interfaces.Response, error)
	GetModules(ctx context.Context) (*ModulesResponse, *interfaces.Response, error)
	GetConfiguration(ctx context.Context) (*ConfigurationResponse, *interfaces.Response, error)
	GetFlag(ctx context.Context, flag string) (*FlagResponse, *interfaces.Response, error)
	GetDynamicMenu(ctx context.Context, menu string) (json.RawMessage, *interfaces.Response, error)
	ValidateClaims(
		ctx context.Context,
		req json.RawMessage,
	) (json.RawMessage, *interfaces.Response, error)
	PostTelemetry(
		ctx context.Context,
		req json.RawMessage,
	) (json.RawMessage, *interfaces.Response, error)
}

var _ ProductShellServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetMenu calls the observed POST endpoint.
func (s *Service) GetMenu(ctx context.Context) (*MenuResponse, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result MenuResponse
	resp, err := s.client.Post(ctx, EndpointGetMenu, map[string]any{}, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetUser calls the observed GET endpoint.
func (s *Service) GetUser(ctx context.Context) (*UserResponse, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result UserResponse
	resp, err := s.client.Get(ctx, EndpointGetUser, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetModules calls the observed GET endpoint.
func (s *Service) GetModules(ctx context.Context) (*ModulesResponse, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result ModulesResponse
	resp, err := s.client.Get(ctx, EndpointGetModules, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetConfiguration calls the observed GET endpoint.
func (s *Service) GetConfiguration(
	ctx context.Context,
) (*ConfigurationResponse, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result ConfigurationResponse
	resp, err := s.client.Get(ctx, EndpointGetConfiguration, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetFlag calls the observed GET endpoint.
func (s *Service) GetFlag(
	ctx context.Context,
	flag string,
) (*FlagResponse, *interfaces.Response, error) {
	if err := ValidateFlag(flag); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result FlagResponse
	resp, err := s.client.Get(
		ctx,
		fmt.Sprintf(EndpointGetFlag, url.PathEscape(flag)),
		nil,
		headers,
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetDynamicMenu calls the observed GET endpoint.
func (s *Service) GetDynamicMenu(
	ctx context.Context,
	menu string,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateMenu(menu); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Get(
		ctx,
		fmt.Sprintf(EndpointGetDynamicMenu, url.PathEscape(menu)),
		nil,
		headers,
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ValidateClaims calls the observed POST endpoint.
// The payload schema is opaque; this operation has not been replayed in the lab.
func (s *Service) ValidateClaims(
	ctx context.Context,
	req json.RawMessage,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Post(ctx, EndpointValidateClaims, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// PostTelemetry calls the observed POST endpoint.
// The payload schema is opaque; this operation has not been replayed in the lab.
func (s *Service) PostTelemetry(
	ctx context.Context,
	req json.RawMessage,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Post(ctx, EndpointPostTelemetry, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
