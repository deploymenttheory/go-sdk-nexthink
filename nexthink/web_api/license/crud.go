package license

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type LicenseServiceInterface interface {
	GetFeatureStatus(
		ctx context.Context,
		feature string,
	) (*FeatureStatusResponse, *interfaces.Response, error)
}

var _ LicenseServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetFeatureStatus calls the observed GET endpoint.
func (s *Service) GetFeatureStatus(
	ctx context.Context,
	feature string,
) (*FeatureStatusResponse, *interfaces.Response, error) {
	if err := ValidateFeature(feature); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result FeatureStatusResponse
	resp, err := s.client.Get(
		ctx,
		fmt.Sprintf(EndpointGetFeatureStatus, url.PathEscape(feature)),
		nil,
		headers,
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
