package action_executions

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

// ListRemoteActions reads the browser API contract.
func (s *Service) ListRemoteActions(ctx context.Context, request *ListOptions) (*ListRemoteActionsResponse, *interfaces.Response, error) {
	if err := validateListOptions(request); err != nil {
		return nil, nil, err
	}
	path := EndpointListRemoteActions
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result ListRemoteActionsResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetRemoteAction reads the browser API contract.
func (s *Service) GetRemoteAction(ctx context.Context, request *DetailsRequest) (*GetRemoteActionResponse, *interfaces.Response, error) {
	if err := validateDetailsRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetRemoteAction
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result GetRemoteActionResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListRemoteActionsForQuery reads the browser API contract.
func (s *Service) ListRemoteActionsForQuery(ctx context.Context, request *NQLRequest) (*ListRemoteActionsForQueryResponse, *interfaces.Response, error) {
	if err := validateNQLRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointListRemoteActionsForQuery
	var result ListRemoteActionsForQueryResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Execute submits an execution request. This can run work on real targets.
func (s *Service) Execute(ctx context.Context, request *ExecuteRequest) (*ExecuteResponse, *interfaces.Response, error) {
	if err := validateExecuteRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointExecute
	var result ExecuteResponse
	headers := s.headers()
	if request.Source != "" {
		headers["nx-source"] = request.Source
	}
	response, err := s.client.Post(ctx, path, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListActions reads the browser API contract.
func (s *Service) ListActions(ctx context.Context) (*ListActionsResponse, *interfaces.Response, error) {
	path := EndpointListActions
	var result ListActionsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetDeviceHistory reads the browser API contract.
func (s *Service) GetDeviceHistory(ctx context.Context, deviceID string, request *HistoryRequest) (*GetDeviceHistoryResponse, *interfaces.Response, error) {
	if err := validateID(deviceID); err != nil {
		return nil, nil, err
	}
	if err := validateHistoryRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetDeviceHistory
	path = strings.ReplaceAll(path, "{deviceID}", url.PathEscape(deviceID))
	var result GetDeviceHistoryResponse
	response, err := s.client.Post(ctx, path, request, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type ActionExecutionsServiceInterface interface {
	ListRemoteActions(ctx context.Context, request *ListOptions) (*ListRemoteActionsResponse, *interfaces.Response, error)
	GetRemoteAction(ctx context.Context, request *DetailsRequest) (*GetRemoteActionResponse, *interfaces.Response, error)
	ListRemoteActionsForQuery(ctx context.Context, request *NQLRequest) (*ListRemoteActionsForQueryResponse, *interfaces.Response, error)
	Execute(ctx context.Context, request *ExecuteRequest) (*ExecuteResponse, *interfaces.Response, error)
	ListActions(ctx context.Context) (*ListActionsResponse, *interfaces.Response, error)
	GetDeviceHistory(ctx context.Context, deviceID string, request *HistoryRequest) (*GetDeviceHistoryResponse, *interfaces.Response, error)
}

var _ ActionExecutionsServiceInterface = (*Service)(nil)
