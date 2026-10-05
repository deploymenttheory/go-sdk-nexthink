package device_configuration

import (
	"context"
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type DeviceConfigurationServiceInterface interface {
	GetProfiles(ctx context.Context) (*ProfilesResponse, *interfaces.Response, error)
	SetProfiles(
		ctx context.Context,
		req json.RawMessage,
	) (json.RawMessage, *interfaces.Response, error)
	GetSettings(ctx context.Context) (json.RawMessage, *interfaces.Response, error)
	SetSettings(
		ctx context.Context,
		req json.RawMessage,
	) (json.RawMessage, *interfaces.Response, error)
}

var _ DeviceConfigurationServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetProfiles calls the observed GET endpoint.
func (s *Service) GetProfiles(
	ctx context.Context,
) (*ProfilesResponse, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result ProfilesResponse
	resp, err := s.client.Get(ctx, EndpointGetProfiles, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// SetProfiles calls the observed PUT endpoint.
// The payload schema is opaque; this operation has not been replayed in the lab.
func (s *Service) SetProfiles(
	ctx context.Context,
	req json.RawMessage,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Put(ctx, EndpointSetProfiles, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetSettings calls the observed GET endpoint.
func (s *Service) GetSettings(ctx context.Context) (json.RawMessage, *interfaces.Response, error) {
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Get(ctx, EndpointGetSettings, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// SetSettings calls the observed PUT endpoint.
// The payload schema is opaque; this operation has not been replayed in the lab.
func (s *Service) SetSettings(
	ctx context.Context,
	req json.RawMessage,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	resp, err := s.client.Put(ctx, EndpointSetSettings, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
