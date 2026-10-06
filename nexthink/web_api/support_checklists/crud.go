package support_checklists

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

// ListProperties reads the browser API contract.
func (s *Service) ListProperties(ctx context.Context, deviceID string) (*ListPropertiesResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	path := EndpointListProperties
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result ListPropertiesResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// List reads the browser API contract.
func (s *Service) List(ctx context.Context, deviceID string) (*ListResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	path := EndpointList
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result ListResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Get reads the browser API contract.
func (s *Service) Get(ctx context.Context, deviceID string, checklistID string) (*GetResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateID(checklistID); err != nil {
		return nil, nil, err
	}
	path := EndpointGet
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	path = strings.ReplaceAll(path, "{checklistID}", url.PathEscape(checklistID))
	var result GetResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type SupportChecklistsServiceInterface interface {
	ListProperties(ctx context.Context, deviceID string) (*ListPropertiesResponse, *interfaces.Response, error)
	List(ctx context.Context, deviceID string) (*ListResponse, *interfaces.Response, error)
	Get(ctx context.Context, deviceID string, checklistID string) (*GetResponse, *interfaces.Response, error)
}

var _ SupportChecklistsServiceInterface = (*Service)(nil)
