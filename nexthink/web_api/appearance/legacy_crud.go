package appearance

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/portalsession"
)

const EndpointLegacyGetAsset = portalsession.GetAsset
const EndpointLegacySaveAsset = portalsession.SaveAsset

// GetLegacyAsset reads the current portal asset with defaultImage=false, as the UI does.
func (s *Service) GetLegacyAsset(ctx context.Context, session *auth.PortalSession, name AssetName) (*LegacyAsset, *interfaces.Response, error) {
	if err := session.Validate(); err != nil {
		return nil, nil, err
	}
	if err := validateName(name); err != nil {
		return nil, nil, err
	}
	var result LegacyAsset
	resp, err := s.client.Post(portalsession.Context(ctx), EndpointLegacyGetAsset, struct {
		DefaultImage bool      `json:"defaultImage"`
		Name         AssetName `json:"name"`
	}{false, name}, legacyHeaders(session), &result)
	if err != nil {
		return nil, resp, err
	}
	if len(result.ID) == 0 || len(result.Version) == 0 || result.MIMEType == "" {
		return &result, resp, fmt.Errorf("legacy appearance returned incomplete asset metadata")
	}
	return &result, resp, nil
}

// SaveLegacyAsset replaces a portal asset. Read its current ID/version first.
// This is one save request; call GetLegacyAsset separately to obtain fresh metadata.
func (s *Service) SaveLegacyAsset(ctx context.Context, session *auth.PortalSession, request *SaveLegacyAssetRequest) (*SaveLegacyAssetResponse, *interfaces.Response, error) {
	if err := session.Validate(); err != nil {
		return nil, nil, err
	}
	if err := validateLegacyAsset(request); err != nil {
		return nil, nil, err
	}
	var result SaveLegacyAssetResponse
	resp, err := s.client.Post(portalsession.Context(ctx), EndpointLegacySaveAsset, request, legacyHeaders(session), &result)
	if err != nil {
		return nil, resp, err
	}
	if result.ResultStatus != nil && result.ResultStatus.Code != 0 {
		return &result, resp, fmt.Errorf("legacy appearance save failed: %s", result.ResultStatus.Description)
	}
	if result.Result != nil && result.Result.Error != nil && legacyCodeTruthy(result.Result.Error.Code) {
		return &result, resp, fmt.Errorf("legacy appearance save failed: %s", result.Result.Error.Code)
	}
	return &result, resp, nil
}
func legacyHeaders(s *auth.PortalSession) map[string]string {
	return map[string]string{"Accept": "application/json", "Content-Type": "application/json", "Cookie": s.Cookie, "x-auth-token": s.XAuthToken}
}
