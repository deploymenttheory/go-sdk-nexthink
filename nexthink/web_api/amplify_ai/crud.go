package amplify_ai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"strings"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

type AmplifyAIServiceInterface interface {
	GenerateOrFetchAnalysis(ctx context.Context, request *AnalysisRequest) (json.RawMessage, *interfaces.Response, error)
	SubmitFeedback(ctx context.Context, request *FeedbackRequest) (json.RawMessage, *interfaces.Response, error)
	PostMetric(ctx context.Context, request *MetricRequest) (*interfaces.Response, error)
	ReportTicketRetrievalDuration(ctx context.Context, request *TicketRetrievalDurationRequest) (*interfaces.Response, error)
	ExecuteAction(ctx context.Context, request *ExecuteActionRequest) (*interfaces.Response, error)
	ExecuteUserAction(ctx context.Context, request *ExecuteUserActionRequest) (*interfaces.Response, error)
	GetResolutionPlan(ctx context.Context, id string) (json.RawMessage, *interfaces.Response, error)
	RefreshResolutionStep(ctx context.Context, request *RefreshResolutionStepRequest) (json.RawMessage, *interfaces.Response, error)
	UpdateResolutionStep(ctx context.Context, request *UpdateResolutionStepRequest) (json.RawMessage, *interfaces.Response, error)
	GetMockMetadata(ctx context.Context, id string) (json.RawMessage, *interfaces.Response, error)
	ResolveTicket(ctx context.Context, request *ResolveTicketRequest) (*interfaces.Response, error)
}

var _ AmplifyAIServiceInterface = (*Service)(nil)

// GenerateOrFetchAnalysis starts or retrieves an analysis; generation can consume AI capacity.
func (s *Service) GenerateOrFetchAnalysis(ctx context.Context, request *AnalysisRequest) (json.RawMessage, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateGenerateOrFetchAnalysis(request); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	response, err := s.client.Post(ctx, EndpointGenerateOrFetchAnalysis, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// SubmitFeedback calls the corresponding Amplify AI extension operation.
func (s *Service) SubmitFeedback(ctx context.Context, request *FeedbackRequest) (json.RawMessage, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateSubmitFeedback(request); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	response, err := s.client.Post(ctx, EndpointSubmitFeedback, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// PostMetric calls the corresponding Amplify AI extension operation.
func (s *Service) PostMetric(ctx context.Context, request *MetricRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validatePostMetric(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, EndpointPostMetric, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// ReportTicketRetrievalDuration calls the corresponding Amplify AI extension operation.
func (s *Service) ReportTicketRetrievalDuration(ctx context.Context, request *TicketRetrievalDurationRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validateReportTicketRetrievalDuration(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, EndpointReportTicketRetrievalDuration, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// ExecuteAction executes an action associated with a resolution step.
func (s *Service) ExecuteAction(ctx context.Context, request *ExecuteActionRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validateExecuteAction(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, EndpointExecuteAction, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// ExecuteUserAction executes the user action associated with a resolution step.
func (s *Service) ExecuteUserAction(ctx context.Context, request *ExecuteUserActionRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validateExecuteUserAction(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, EndpointExecuteUserAction, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// GetResolutionPlan calls the corresponding Amplify AI extension operation.
func (s *Service) GetResolutionPlan(ctx context.Context, id string) (json.RawMessage, *interfaces.Response, error) {
	if strings.TrimSpace(id) == "" {
		return nil, nil, fmt.Errorf("identifier is required")
	}
	var result json.RawMessage
	response, err := s.client.Get(ctx, EndpointGetResolutionPlan, map[string]string{"resolution_plan_id": id}, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// RefreshResolutionStep refreshes a resolution step from its current execution state.
func (s *Service) RefreshResolutionStep(ctx context.Context, request *RefreshResolutionStepRequest) (json.RawMessage, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateRefreshResolutionStep(request); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	response, err := s.client.Post(ctx, EndpointRefreshResolutionStep, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// UpdateResolutionStep calls the corresponding Amplify AI extension operation.
func (s *Service) UpdateResolutionStep(ctx context.Context, request *UpdateResolutionStepRequest) (json.RawMessage, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateUpdateResolutionStep(request); err != nil {
		return nil, nil, err
	}
	var result json.RawMessage
	response, err := s.client.Post(ctx, EndpointUpdateResolutionStep, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// GetMockMetadata calls the corresponding Amplify AI extension operation.
func (s *Service) GetMockMetadata(ctx context.Context, id string) (json.RawMessage, *interfaces.Response, error) {
	if strings.TrimSpace(id) == "" {
		return nil, nil, fmt.Errorf("identifier is required")
	}
	var result json.RawMessage
	response, err := s.client.Get(ctx, EndpointGetMockMetadata, map[string]string{"ticket_sys_id": id}, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return result, response, nil
}

// ResolveTicket resolves a connected ticket; this changes the external ticket.
func (s *Service) ResolveTicket(ctx context.Context, request *ResolveTicketRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if err := validateResolveTicket(request); err != nil {
		return nil, err
	}
	return s.client.Post(ctx, EndpointResolveTicket, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}
