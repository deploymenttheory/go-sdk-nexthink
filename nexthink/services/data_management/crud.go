// Package data_management schedules asynchronous inventory device deletions.
// API reference: https://docs.nexthink.com/api/data-management/schedule-device-deletions
package data_management

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

const EndpointDeviceDeletions = "/api/v1/data-management/device/deletions"

type Device struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}
type DeleteDevicesRequest struct {
	Devices []Device `json:"devices"`
}
type DeviceStatus struct {
	UID    string `json:"uid"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type DeleteDevicesResponse struct {
	ScheduledCount int            `json:"scheduledCount"`
	Status         string         `json:"status"`
	Devices        []DeviceStatus `json:"devices"`
}

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

func ValidateDeleteDevicesRequest(req *DeleteDevicesRequest) error {
	if req == nil || len(req.Devices) == 0 || len(req.Devices) > 100 {
		return fmt.Errorf("devices must contain 1 to 100 entries")
	}
	for i, d := range req.Devices {
		if strings.TrimSpace(d.UID) == "" || strings.TrimSpace(d.Name) == "" {
			return fmt.Errorf("devices[%d]: uid and name are required", i)
		}
	}
	// Malformed nonempty UIDs are deliberately passed to the server: the API
	// reports INVALID per device while scheduling the rest of the batch.
	return nil
}

// DeleteDevices returns scheduling outcomes, not proof that deletion has completed.
// requestID is an optional correlation UUID echoed in the response headers.
func (s *Service) DeleteDevices(ctx context.Context, req *DeleteDevicesRequest, requestID string) (*DeleteDevicesResponse, *interfaces.Response, error) {
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
