package collector_management

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

const EndpointQueryDevices = "/apigateway/api/v1/clt-updater/query/devices"

func (s *Service) QueryDevices(ctx context.Context, r *DeviceQuery) (*DevicesResponse, *interfaces.Response, error) {
	if r == nil || r.Limit < 1 || r.Platforms == nil || r.Versions == nil || r.Groups == nil || r.TargetVersions == nil {
		return nil, nil, fmt.Errorf("positive limit and explicit filter arrays are required")
	}
	var result DevicesResponse
	resp, err := s.client.Post(ctx, EndpointQueryDevices, r, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
