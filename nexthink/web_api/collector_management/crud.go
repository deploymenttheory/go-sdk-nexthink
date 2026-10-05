package collector_management

import (
	"context"
	"encoding/json"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type CollectorManagementServiceInterface interface {
	GetDownloadLinks(ctx context.Context) (*DownloadLinks, *interfaces.Response, error)
	GetUpdateConfiguration(ctx context.Context) (*UpdateConfiguration, *interfaces.Response, error)
	GetPlatforms(ctx context.Context) ([]QueryValue, *interfaces.Response, error)
	GetVersions(ctx context.Context) ([]QueryValue, *interfaces.Response, error)
	GetGroups(ctx context.Context) ([]QueryValue, *interfaces.Response, error)
	GetTargetVersions(ctx context.Context) ([]QueryValue, *interfaces.Response, error)
	SetUpdateConfiguration(
		ctx context.Context,
		req json.RawMessage,
	) (json.RawMessage, *interfaces.Response, error)
}

var _ CollectorManagementServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func (s *Service) GetDownloadLinks(
	ctx context.Context,
) (*DownloadLinks, *interfaces.Response, error) {
	var result DownloadLinks
	resp, err := s.client.Get(
		ctx,
		EndpointGetDownloadLinks,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) GetUpdateConfiguration(
	ctx context.Context,
) (*UpdateConfiguration, *interfaces.Response, error) {
	var result UpdateConfiguration
	resp, err := s.client.Get(
		ctx,
		EndpointGetUpdateConfiguration,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

func (s *Service) GetPlatforms(ctx context.Context) ([]QueryValue, *interfaces.Response, error) {
	var result []QueryValue
	resp, err := s.client.Get(
		ctx,
		EndpointGetPlatforms,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) GetVersions(ctx context.Context) ([]QueryValue, *interfaces.Response, error) {
	var result []QueryValue
	resp, err := s.client.Get(
		ctx,
		EndpointGetVersions,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) GetGroups(ctx context.Context) ([]QueryValue, *interfaces.Response, error) {
	var result []QueryValue
	resp, err := s.client.Get(
		ctx,
		EndpointGetGroups,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (s *Service) GetTargetVersions(
	ctx context.Context,
) ([]QueryValue, *interfaces.Response, error) {
	var result []QueryValue
	resp, err := s.client.Get(
		ctx,
		EndpointGetTargetVersions,
		nil,
		map[string]string{"Accept": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// SetUpdateConfiguration submits the full configuration, including its revision.
// This shared-tenant write is bundle-evidenced and has not been replayed in the lab.
func (s *Service) SetUpdateConfiguration(
	ctx context.Context,
	req json.RawMessage,
) (json.RawMessage, *interfaces.Response, error) {
	if err := ValidateUpdateConfiguration(req); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	resp, err := s.client.Post(
		ctx,
		EndpointGetUpdateConfiguration,
		req,
		map[string]string{"Accept": "application/json", "Content-Type": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
