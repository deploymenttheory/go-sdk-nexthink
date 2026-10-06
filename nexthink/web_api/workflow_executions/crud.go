package workflow_executions

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

// GetTimeline reads the browser API contract.
func (s *Service) GetTimeline(ctx context.Context, workflowID string, executionID string, request *TimelineOptions) (*GetTimelineResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateTimelineOptions(request); err != nil {
		return nil, nil, err
	}
	path := EndpointGetTimeline
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	query["executionStatus"] = strings.ToUpper(query["executionStatus"])
	var result GetTimelineResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetTimelineV2 reads the browser API contract.
func (s *Service) GetTimelineV2(ctx context.Context, workflowID string, executionID string) (*GetTimelineV2Response, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetTimelineV2
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	var result GetTimelineV2Response
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListActivities reads the browser API contract.
func (s *Service) ListActivities(ctx context.Context, workflowID string, executionID string) (*ListActivitiesResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	path := EndpointListActivities
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	var result ListActivitiesResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetRemoteActionDetails reads the browser API contract.
func (s *Service) GetRemoteActionDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetRemoteActionDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetRemoteActionDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetRemoteActionDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCustomFieldsDetails reads the browser API contract.
func (s *Service) GetCustomFieldsDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetCustomFieldsDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetCustomFieldsDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetCustomFieldsDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetCampaignDetails reads the browser API contract.
func (s *Service) GetCampaignDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetCampaignDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetCampaignDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetCampaignDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetFunctionDetails reads the browser API contract.
func (s *Service) GetFunctionDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetFunctionDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetFunctionDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetFunctionDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetMessageDetails reads the browser API contract.
func (s *Service) GetMessageDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetMessageDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetMessageDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetMessageDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetSAPIDetails reads the browser API contract.
func (s *Service) GetSAPIDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetSAPIDetailsResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	if err := validateID(thinkletID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetSAPIDetails
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	path = strings.ReplaceAll(path, "{thinkletID}", url.PathEscape(thinkletID))
	var result GetSAPIDetailsResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListWorkflows reads the browser API contract.
func (s *Service) ListWorkflows(ctx context.Context, request *ListOptions) (*ListWorkflowsResponse, *interfaces.Response, error) {
	if err := validateListOptions(request); err != nil {
		return nil, nil, err
	}
	path := EndpointListWorkflows
	query, err := queryParameters(request)
	if err != nil {
		return nil, nil, err
	}
	var result ListWorkflowsResponse
	response, err := s.client.Get(ctx, path, query, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetWorkflow reads the browser API contract.
func (s *Service) GetWorkflow(ctx context.Context, workflowID string) (*GetWorkflowResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetWorkflow
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	var result GetWorkflowResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Get reads the browser API contract.
func (s *Service) Get(ctx context.Context, workflowID string, executionID string) (*GetResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	path := EndpointGet
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	var result GetResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetHistory reads the browser API contract.
func (s *Service) GetHistory(ctx context.Context, workflowID string, executionID string) (*GetHistoryResponse, *interfaces.Response, error) {
	if err := validateID(workflowID); err != nil {
		return nil, nil, err
	}
	if err := validateID(executionID); err != nil {
		return nil, nil, err
	}
	path := EndpointGetHistory
	path = strings.ReplaceAll(path, "{workflowID}", url.PathEscape(workflowID))
	path = strings.ReplaceAll(path, "{executionID}", url.PathEscape(executionID))
	var result GetHistoryResponse
	response, err := s.client.Get(ctx, path, nil, s.headers(), &result)
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

// ExecuteNQL submits an execution request. This can run work on real targets.
func (s *Service) ExecuteNQL(ctx context.Context, request *ExecuteNQLRequest) (*ExecuteNQLResponse, *interfaces.Response, error) {
	if err := validateExecuteNQLRequest(request); err != nil {
		return nil, nil, err
	}
	path := EndpointExecuteNQL
	var result ExecuteNQLResponse
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

type WorkflowExecutionsServiceInterface interface {
	GetTimeline(ctx context.Context, workflowID string, executionID string, request *TimelineOptions) (*GetTimelineResponse, *interfaces.Response, error)
	GetTimelineV2(ctx context.Context, workflowID string, executionID string) (*GetTimelineV2Response, *interfaces.Response, error)
	ListActivities(ctx context.Context, workflowID string, executionID string) (*ListActivitiesResponse, *interfaces.Response, error)
	GetRemoteActionDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetRemoteActionDetailsResponse, *interfaces.Response, error)
	GetCustomFieldsDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetCustomFieldsDetailsResponse, *interfaces.Response, error)
	GetCampaignDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetCampaignDetailsResponse, *interfaces.Response, error)
	GetFunctionDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetFunctionDetailsResponse, *interfaces.Response, error)
	GetMessageDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetMessageDetailsResponse, *interfaces.Response, error)
	GetSAPIDetails(ctx context.Context, workflowID string, executionID string, thinkletID string) (*GetSAPIDetailsResponse, *interfaces.Response, error)
	ListWorkflows(ctx context.Context, request *ListOptions) (*ListWorkflowsResponse, *interfaces.Response, error)
	GetWorkflow(ctx context.Context, workflowID string) (*GetWorkflowResponse, *interfaces.Response, error)
	Get(ctx context.Context, workflowID string, executionID string) (*GetResponse, *interfaces.Response, error)
	GetHistory(ctx context.Context, workflowID string, executionID string) (*GetHistoryResponse, *interfaces.Response, error)
	Execute(ctx context.Context, request *ExecuteRequest) (*ExecuteResponse, *interfaces.Response, error)
	ExecuteNQL(ctx context.Context, request *ExecuteNQLRequest) (*ExecuteNQLResponse, *interfaces.Response, error)
}

var _ WorkflowExecutionsServiceInterface = (*Service)(nil)
