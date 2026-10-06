package appearance

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Get returns the original image bytes; Content-Type is available in response metadata.
func (s *Service) Get(ctx context.Context, name AssetName) ([]byte, *interfaces.Response, error) {
	if err := validateName(name); err != nil {
		return nil, nil, err
	}
	resp, err := s.client.Get(ctx, Endpoint+"/"+string(name), nil, map[string]string{"Accept": "image/*"}, nil)
	if err != nil {
		return nil, resp, err
	}
	if resp == nil {
		return nil, nil, fmt.Errorf("appearance returned no response")
	}
	return resp.Body, resp, nil
}

// Update uploads raw image bytes. The UI resets an asset by uploading the default image through this same operation.
func (s *Service) Update(ctx context.Context, name AssetName, contentType string, data []byte) (*UpdateResponse, *interfaces.Response, error) {
	if err := validateUpload(name, contentType, data); err != nil {
		return nil, nil, err
	}
	var result UpdateResponse
	resp, err := s.client.Put(ctx, Endpoint+"/"+string(name), data, map[string]string{"Content-Type": contentType}, &result)
	if err != nil {
		return nil, resp, err
	}
	if !result.Result.Success {
		return &result, resp, fmt.Errorf("appearance update failed: %s", result.ResultStatus.Description)
	}
	return &result, resp, nil
}

type AppearanceServiceInterface interface {
	GetLegacyAsset(context.Context, *auth.PortalSession, AssetName) (*LegacyAsset, *interfaces.Response, error)
	SaveLegacyAsset(context.Context, *auth.PortalSession, *SaveLegacyAssetRequest) (*SaveLegacyAssetResponse, *interfaces.Response, error)
	Get(context.Context, AssetName) ([]byte, *interfaces.Response, error)
	Update(context.Context, AssetName, string, []byte) (*UpdateResponse, *interfaces.Response, error)
}

var _ AppearanceServiceInterface = (*Service)(nil)
