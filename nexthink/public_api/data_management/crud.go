// Package data_management schedules asynchronous inventory device deletions.
// API reference: https://docs.nexthink.com/api/data-management/schedule-device-deletions
package data_management

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type DataManagementServiceInterface interface {
	DeleteDevices(
		ctx context.Context,
		req *DeleteDevicesRequest,
		requestID string,
	) (*DeleteDevicesResponse, *interfaces.Response, error)
}

var _ DataManagementServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// DeleteDevices returns scheduling outcomes, not proof that deletion has completed.
// requestID is an optional correlation UUID echoed in the response headers.
func (s *Service) DeleteDevices(
	ctx context.Context,
	req *DeleteDevicesRequest,
	requestID string,
) (*DeleteDevicesResponse, *interfaces.Response, error) {
	if err := ValidateDeleteDevicesRequest(req); err != nil {
		return nil, nil, err
	}
	headers := map[string]string{"Content-Type": "application/json", "Accept": "application/json"}
	if requestID != "" {
		headers["x-request-id"] = requestID
	}
	var result DeleteDevicesResponse
	resp, err := s.client.Post(ctx, EndpointDeviceDeletions, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
