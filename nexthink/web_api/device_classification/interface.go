package device_classification

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type DeviceClassificationServiceInterface interface {
	GetOrganization(context.Context) ([]Ruleset, *interfaces.Response, error)
	CreateOrganization(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	UpdateOrganization(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	DownloadOrganization(context.Context) ([]byte, *interfaces.Response, error)
	GetLocationType(context.Context) ([]Ruleset, *interfaces.Response, error)
	CreateLocationType(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	UpdateLocationType(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	DownloadLocationType(context.Context) ([]byte, *interfaces.Response, error)
	GetVPNEgress(context.Context) ([]Ruleset, *interfaces.Response, error)
	CreateVPNEgress(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	UpdateVPNEgress(ctx context.Context, request *RulesetUpload) (*interfaces.Response, error)
	DownloadVPNEgress(context.Context) ([]byte, *interfaces.Response, error)
	DeleteVPNEgress(context.Context) (*interfaces.Response, error)
	GetGeoIP(context.Context) (*GeoIPConfiguration, *interfaces.Response, error)
	UpdateGeoIP(ctx context.Context, request *GeoIPConfiguration) (*GeoIPConfiguration, *interfaces.Response, error)
}

var _ DeviceClassificationServiceInterface = (*Service)(nil)
