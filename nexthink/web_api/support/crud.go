package support

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

// GetProfile reads the browser API contract.
func (s *Service) GetProfile(ctx context.Context, deviceID string) (*GetProfileResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetProfile
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetProfileResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetPlatform reads the browser API contract.
func (s *Service) GetPlatform(ctx context.Context, deviceID string) (*GetPlatformResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetPlatform
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetPlatformResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListUsers reads the browser API contract.
func (s *Service) ListUsers(ctx context.Context, deviceID string) (*ListUsersResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	path := EndpointListUsers
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result ListUsersResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Search reads the browser API contract.
func (s *Service) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, *interfaces.Response, error) {
	if err := validateSearchRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointSearch
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result SearchResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type SupportServiceInterface interface {
	GetProfile(ctx context.Context, deviceID string) (*GetProfileResponse, *interfaces.Response, error)
	GetPlatform(ctx context.Context, deviceID string) (*GetPlatformResponse, *interfaces.Response, error)
	ListUsers(ctx context.Context, deviceID string) (*ListUsersResponse, *interfaces.Response, error)
	Search(ctx context.Context, request *SearchRequest) (*SearchResponse, *interfaces.Response, error)
}

var _ SupportServiceInterface = (*Service)(nil)
