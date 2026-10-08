package device_classification

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

// GetOrganization returns uploaded CSV rule-set metadata; an unconfigured tenant returns NO_RULESET_FOUND404.
func (s *Service) GetOrganization(ctx context.Context) ([]Ruleset, *interfaces.Response, error) {
	var result []Ruleset
	response, err := s.client.Get(ctx, Endpoint+"/entities", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// CreateOrganization installs the tenant-wide classification rule set.
func (s *Service) CreateOrganization(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "POST", Endpoint+"/entities", request, true)
}

// UpdateOrganization replaces the tenant-wide classification rule set.
func (s *Service) UpdateOrganization(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "PUT", Endpoint+"/entities", request, false)
}

// DownloadOrganization returns the uploaded CSV bytes without re-encoding.
func (s *Service) DownloadOrganization(ctx context.Context) ([]byte, *interfaces.Response, error) {
	response, data, err := s.client.GetBytes(ctx, Endpoint+"/entities/download", nil, nil)
	return data, response, err
}

// GetLocationType returns uploaded CSV rule-set metadata; an unconfigured tenant returns NO_RULESET_FOUND404.
func (s *Service) GetLocationType(ctx context.Context) ([]Ruleset, *interfaces.Response, error) {
	var result []Ruleset
	response, err := s.client.Get(ctx, Endpoint+"/location-type", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// CreateLocationType installs the tenant-wide classification rule set.
func (s *Service) CreateLocationType(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "POST", Endpoint+"/location-type", request, true)
}

// UpdateLocationType replaces the tenant-wide classification rule set.
func (s *Service) UpdateLocationType(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "PUT", Endpoint+"/location-type", request, false)
}

// DownloadLocationType returns the uploaded CSV bytes without re-encoding.
func (s *Service) DownloadLocationType(ctx context.Context) ([]byte, *interfaces.Response, error) {
	response, data, err := s.client.GetBytes(ctx, Endpoint+"/location-type/download", nil, nil)
	return data, response, err
}

// GetVPNEgress returns uploaded CSV rule-set metadata; an unconfigured tenant returns NO_RULESET_FOUND404.
func (s *Service) GetVPNEgress(ctx context.Context) ([]Ruleset, *interfaces.Response, error) {
	var result []Ruleset
	response, err := s.client.Get(ctx, Endpoint+"/vpn-egress", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// CreateVPNEgress installs the tenant-wide classification rule set.
func (s *Service) CreateVPNEgress(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "POST", Endpoint+"/vpn-egress", request, true)
}

// UpdateVPNEgress replaces the tenant-wide classification rule set.
func (s *Service) UpdateVPNEgress(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error) {
	return s.upload(ctx, "PUT", Endpoint+"/vpn-egress", request, false)
}

// DownloadVPNEgress returns the uploaded CSV bytes without re-encoding.
func (s *Service) DownloadVPNEgress(ctx context.Context) ([]byte, *interfaces.Response, error) {
	response, data, err := s.client.GetBytes(ctx, Endpoint+"/vpn-egress/download", nil, nil)
	return data, response, err
}

// DeleteVPNEgress removes the tenant-wide VPN egress classification rule set.
func (s *Service) DeleteVPNEgress(ctx context.Context) (*interfaces.Response, error) {
	return s.client.Delete(ctx, Endpoint+"/vpn-egress", nil, nil, nil)
}

// GetGeoIP reads tenant geolocation collection granularity.
func (s *Service) GetGeoIP(ctx context.Context) (*GeoIPConfiguration, *interfaces.Response, error) {
	var result GeoIPConfiguration
	response, err := s.client.Get(ctx, Endpoint+"/configuration/geoip", nil, nil, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateGeoIP replaces tenant geolocation collection settings. False values are sent explicitly.
func (s *Service) UpdateGeoIP(ctx context.Context, request *GeoIPConfiguration) (*GeoIPConfiguration, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, errRequestRequired
	}
	var result GeoIPConfiguration
	response, err := s.client.Put(ctx, Endpoint+"/configuration/geoip", request, map[string]string{"Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
