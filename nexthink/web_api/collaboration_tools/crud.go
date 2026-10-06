package collaboration_tools

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

// GetCallInsights reads the browser API contract.
func (s *Service) GetCallInsights(ctx context.Context, deviceID string, request *CallInsightsRequest) (*GetCallInsightsResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateCallInsightsRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetCallInsights
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetCallInsightsResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type CollaborationToolsServiceInterface interface {
	GetCallInsights(ctx context.Context, deviceID string, request *CallInsightsRequest) (*GetCallInsightsResponse, *interfaces.Response, error)
}

var _ CollaborationToolsServiceInterface = (*Service)(nil)
