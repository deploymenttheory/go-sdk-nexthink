package support_timeline

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strconv"
	"strings"
)

type Service struct {
	client    interfaces.HTTPClient
	timeZone  string
	utcOffset int
}

// Option configures the browser's local time context. UTC is the default.
type Option func(*Service)

// WithTimeZone supplies an IANA timezone and JavaScript getTimezoneOffset minutes (UTC minus local time).
func WithTimeZone(name string, offsetMinutes int) Option {
	return func(s *Service) { s.timeZone = name; s.utcOffset = offsetMinutes }
}
func NewService(c interfaces.HTTPClient, options ...Option) *Service {
	s := &Service{client: c, timeZone: "UTC"}
	for _, o := range options {
		o(s)
	}
	return s
}
func (s *Service) headers() map[string]string {
	return map[string]string{"Accept": "application/json", "Content-Type": "application/json", "time-zone": s.timeZone, "utc-offset": strconv.Itoa(s.utcOffset)}
}

// GetAlertsAndErrors reads the browser API contract.
func (s *Service) GetAlertsAndErrors(ctx context.Context, deviceID string, request *TimeRange) (*GetAlertsAndErrorsResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetAlertsAndErrors
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetAlertsAndErrorsResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetPerformance reads the browser API contract.
func (s *Service) GetPerformance(ctx context.Context, deviceID string, request *TimeRange) (*GetPerformanceResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetPerformance
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetPerformanceResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetConnectivity reads the browser API contract.
func (s *Service) GetConnectivity(ctx context.Context, deviceID string, request *TimeRange) (*GetConnectivityResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetConnectivity
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetConnectivityResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetApplicationConnectivity reads the browser API contract.
func (s *Service) GetApplicationConnectivity(ctx context.Context, deviceID string, request *TimeRange) (*GetApplicationConnectivityResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetApplicationConnectivity
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetApplicationConnectivityResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetActivity reads the browser API contract.
func (s *Service) GetActivity(ctx context.Context, deviceID string, request *TimeRange) (*GetActivityResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetActivity
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetActivityResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetApplications reads the browser API contract.
func (s *Service) GetApplications(ctx context.Context, deviceID string, request *TimeRange) (*GetApplicationsResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetApplications
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetApplicationsResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetUserInteractions reads the browser API contract.
func (s *Service) GetUserInteractions(ctx context.Context, deviceID string, request *TimeRange) (*GetUserInteractionsResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetUserInteractions
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetUserInteractionsResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCollaboration reads the browser API contract.
func (s *Service) GetCollaboration(ctx context.Context, deviceID string, request *TimeRange) (*GetCollaborationResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetCollaboration
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetCollaborationResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetErrorsDrilldown reads the browser API contract.
func (s *Service) GetErrorsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetErrorsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetErrorsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetErrorsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetFreezesDrilldown reads the browser API contract.
func (s *Service) GetFreezesDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetFreezesDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetFreezesDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetFreezesDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetAlertsDrilldown reads the browser API contract.
func (s *Service) GetAlertsDrilldown(ctx context.Context, deviceID string, request *AlertsRequest) (*GetAlertsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateAlertsRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetAlertsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetAlertsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetActionsDrilldown reads the browser API contract.
func (s *Service) GetActionsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetActionsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetActionsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetActionsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetSystemBootsDrilldown reads the browser API contract.
func (s *Service) GetSystemBootsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetSystemBootsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetSystemBootsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetSystemBootsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetSystemBootsAndSuspendsDrilldown reads the browser API contract.
func (s *Service) GetSystemBootsAndSuspendsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetSystemBootsAndSuspendsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetSystemBootsAndSuspendsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetSystemBootsAndSuspendsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetTeamsCallsDrilldown reads the browser API contract.
func (s *Service) GetTeamsCallsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetTeamsCallsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetTeamsCallsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetTeamsCallsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetZoomCallsDrilldown reads the browser API contract.
func (s *Service) GetZoomCallsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetZoomCallsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetZoomCallsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetZoomCallsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDesktopConnectivityDrilldown reads the browser API contract.
func (s *Service) GetDesktopConnectivityDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetDesktopConnectivityDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateApplicationRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetDesktopConnectivityDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetDesktopConnectivityDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetWebConnectivityDrilldown reads the browser API contract.
func (s *Service) GetWebConnectivityDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetWebConnectivityDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateApplicationRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetWebConnectivityDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetWebConnectivityDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDesktopApplicationsDrilldown reads the browser API contract.
func (s *Service) GetDesktopApplicationsDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetDesktopApplicationsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateApplicationRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetDesktopApplicationsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetDesktopApplicationsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetWebApplicationsDrilldown reads the browser API contract.
func (s *Service) GetWebApplicationsDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetWebApplicationsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateApplicationRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetWebApplicationsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetWebApplicationsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetInstallationsDrilldown reads the browser API contract.
func (s *Service) GetInstallationsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetInstallationsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetInstallationsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetInstallationsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetEthernetDrilldown reads the browser API contract.
func (s *Service) GetEthernetDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetEthernetDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetEthernetDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetEthernetDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetWiFiDrilldown reads the browser API contract.
func (s *Service) GetWiFiDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetWiFiDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetWiFiDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetWiFiDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetConnectionsDrilldown reads the browser API contract.
func (s *Service) GetConnectionsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetConnectionsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetConnectionsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetConnectionsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetNetworkApplicationDrilldown reads the browser API contract.
func (s *Service) GetNetworkApplicationDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetNetworkApplicationDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateApplicationRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetNetworkApplicationDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetNetworkApplicationDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCPUDrilldown reads the browser API contract.
func (s *Service) GetCPUDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetCPUDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetCPUDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetCPUDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetMemoryDrilldown reads the browser API contract.
func (s *Service) GetMemoryDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetMemoryDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetMemoryDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetMemoryDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDiskPerformanceDrilldown reads the browser API contract.
func (s *Service) GetDiskPerformanceDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetDiskPerformanceDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetDiskPerformanceDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetDiskPerformanceDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDriveSpaceDrilldown reads the browser API contract.
func (s *Service) GetDriveSpaceDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetDriveSpaceDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetDriveSpaceDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetDriveSpaceDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGPUDrilldown reads the browser API contract.
func (s *Service) GetGPUDrilldown(ctx context.Context, deviceID string, slot string, request *TimeRange) (*GetGPUDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateID(slot); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetGPUDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	path = strings.ReplaceAll(path, "{slot}", url.PathEscape(slot))
	var result GetGPUDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetNPUDrilldown reads the browser API contract.
func (s *Service) GetNPUDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetNPUDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetNPUDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetNPUDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetUserInteractionsDrilldown reads the browser API contract.
func (s *Service) GetUserInteractionsDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetUserInteractionsDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateID(userID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetUserInteractionsDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	path = strings.ReplaceAll(path, "{userID}", url.PathEscape(userID))
	var result GetUserInteractionsDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetRoundTripTimeDrilldown reads the browser API contract.
func (s *Service) GetRoundTripTimeDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetRoundTripTimeDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateID(userID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetRoundTripTimeDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	path = strings.ReplaceAll(path, "{userID}", url.PathEscape(userID))
	var result GetRoundTripTimeDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetNetworkLatencyDrilldown reads the browser API contract.
func (s *Service) GetNetworkLatencyDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetNetworkLatencyDrilldownResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateID(userID); err != nil {
		return nil, nil, err
	}
	if err := validateTimeRange(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetNetworkLatencyDrilldown
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	path = strings.ReplaceAll(path, "{userID}", url.PathEscape(userID))
	var result GetNetworkLatencyDrilldownResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type SupportTimelineServiceInterface interface {
	GetAlertsAndErrors(ctx context.Context, deviceID string, request *TimeRange) (*GetAlertsAndErrorsResponse, *interfaces.Response, error)
	GetPerformance(ctx context.Context, deviceID string, request *TimeRange) (*GetPerformanceResponse, *interfaces.Response, error)
	GetConnectivity(ctx context.Context, deviceID string, request *TimeRange) (*GetConnectivityResponse, *interfaces.Response, error)
	GetApplicationConnectivity(ctx context.Context, deviceID string, request *TimeRange) (*GetApplicationConnectivityResponse, *interfaces.Response, error)
	GetActivity(ctx context.Context, deviceID string, request *TimeRange) (*GetActivityResponse, *interfaces.Response, error)
	GetApplications(ctx context.Context, deviceID string, request *TimeRange) (*GetApplicationsResponse, *interfaces.Response, error)
	GetUserInteractions(ctx context.Context, deviceID string, request *TimeRange) (*GetUserInteractionsResponse, *interfaces.Response, error)
	GetCollaboration(ctx context.Context, deviceID string, request *TimeRange) (*GetCollaborationResponse, *interfaces.Response, error)
	GetErrorsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetErrorsDrilldownResponse, *interfaces.Response, error)
	GetFreezesDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetFreezesDrilldownResponse, *interfaces.Response, error)
	GetAlertsDrilldown(ctx context.Context, deviceID string, request *AlertsRequest) (*GetAlertsDrilldownResponse, *interfaces.Response, error)
	GetActionsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetActionsDrilldownResponse, *interfaces.Response, error)
	GetSystemBootsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetSystemBootsDrilldownResponse, *interfaces.Response, error)
	GetSystemBootsAndSuspendsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetSystemBootsAndSuspendsDrilldownResponse, *interfaces.Response, error)
	GetTeamsCallsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetTeamsCallsDrilldownResponse, *interfaces.Response, error)
	GetZoomCallsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetZoomCallsDrilldownResponse, *interfaces.Response, error)
	GetDesktopConnectivityDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetDesktopConnectivityDrilldownResponse, *interfaces.Response, error)
	GetWebConnectivityDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetWebConnectivityDrilldownResponse, *interfaces.Response, error)
	GetDesktopApplicationsDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetDesktopApplicationsDrilldownResponse, *interfaces.Response, error)
	GetWebApplicationsDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetWebApplicationsDrilldownResponse, *interfaces.Response, error)
	GetInstallationsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetInstallationsDrilldownResponse, *interfaces.Response, error)
	GetEthernetDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetEthernetDrilldownResponse, *interfaces.Response, error)
	GetWiFiDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetWiFiDrilldownResponse, *interfaces.Response, error)
	GetConnectionsDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetConnectionsDrilldownResponse, *interfaces.Response, error)
	GetNetworkApplicationDrilldown(ctx context.Context, deviceID string, request *ApplicationRequest) (*GetNetworkApplicationDrilldownResponse, *interfaces.Response, error)
	GetCPUDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetCPUDrilldownResponse, *interfaces.Response, error)
	GetMemoryDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetMemoryDrilldownResponse, *interfaces.Response, error)
	GetDiskPerformanceDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetDiskPerformanceDrilldownResponse, *interfaces.Response, error)
	GetDriveSpaceDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetDriveSpaceDrilldownResponse, *interfaces.Response, error)
	GetGPUDrilldown(ctx context.Context, deviceID string, slot string, request *TimeRange) (*GetGPUDrilldownResponse, *interfaces.Response, error)
	GetNPUDrilldown(ctx context.Context, deviceID string, request *TimeRange) (*GetNPUDrilldownResponse, *interfaces.Response, error)
	GetUserInteractionsDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetUserInteractionsDrilldownResponse, *interfaces.Response, error)
	GetRoundTripTimeDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetRoundTripTimeDrilldownResponse, *interfaces.Response, error)
	GetNetworkLatencyDrilldown(ctx context.Context, deviceID string, userID string, request *TimeRange) (*GetNetworkLatencyDrilldownResponse, *interfaces.Response, error)
}

var _ SupportTimelineServiceInterface = (*Service)(nil)
