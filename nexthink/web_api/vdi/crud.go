package vdi

import (
	"context"
	"encoding/json"
	"fmt"
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

// GetSessionTimeline reads the browser API contract.
func (s *Service) GetSessionTimeline(ctx context.Context, sessionID string, request *TimelineRequest) (*GetSessionTimelineResponse, *interfaces.Response, error) {
	if err := validateID(sessionID); err != nil {
		return nil, nil, err
	}
	if err := validateTimelineRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetSessionTimeline
	path = strings.ReplaceAll(path, "{sessionID}", url.PathEscape(sessionID))
	var result GetSessionTimelineResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetHypervisorTimeline reads the browser API contract.
func (s *Service) GetHypervisorTimeline(ctx context.Context, sessionID string, request *TimelineRequest) (*GetHypervisorTimelineResponse, *interfaces.Response, error) {
	if err := validateID(sessionID); err != nil {
		return nil, nil, err
	}
	if err := validateTimelineRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetHypervisorTimeline
	path = strings.ReplaceAll(path, "{sessionID}", url.PathEscape(sessionID))
	var result GetHypervisorTimelineResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGlobalHealth reads the browser API contract.
func (s *Service) GetGlobalHealth(ctx context.Context, sessionID string, request *HealthRequest) (*GetGlobalHealthResponse, *interfaces.Response, error) {
	if err := validateID(sessionID); err != nil {
		return nil, nil, err
	}
	if err := validateHealthRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetGlobalHealth
	path = strings.ReplaceAll(path, "{sessionID}", url.PathEscape(sessionID))
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result GetGlobalHealthResponse
	headers := s.headers()
	headers["Accept"] = "application/json, text/plain"
	response, err := s.client.Get(ctx, path, query, headers, nil)
	if err != nil {
		return nil, response, err
	}
	if response == nil {
		return nil, nil, fmt.Errorf("global health returned no response")
	}
	if strings.Contains(response.Headers.Get("Content-Type"), "json") {
		if err := json.Unmarshal(response.Body, &result); err != nil {
			return nil, response, err
		}
	} else {
		result = GetGlobalHealthResponse(string(response.Body))
	}
	return &result, response, nil
}

// ValidateHostname reads the browser API contract.
func (s *Service) ValidateHostname(ctx context.Context, request *HostnameRequest) (*ValidateHostnameResponse, *interfaces.Response, error) {
	if err := validateHostnameRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointValidateHostname
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result ValidateHostnameResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type VDIServiceInterface interface {
	GetSessionTimeline(ctx context.Context, sessionID string, request *TimelineRequest) (*GetSessionTimelineResponse, *interfaces.Response, error)
	GetHypervisorTimeline(ctx context.Context, sessionID string, request *TimelineRequest) (*GetHypervisorTimelineResponse, *interfaces.Response, error)
	GetGlobalHealth(ctx context.Context, sessionID string, request *HealthRequest) (*GetGlobalHealthResponse, *interfaces.Response, error)
	ValidateHostname(ctx context.Context, request *HostnameRequest) (*ValidateHostnameResponse, *interfaces.Response, error)
}

var _ VDIServiceInterface = (*Service)(nil)
