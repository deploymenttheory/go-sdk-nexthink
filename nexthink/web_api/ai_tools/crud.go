package ai_tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"net/url"
	"strconv"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(client interfaces.HTTPClient) *Service { return &Service{client: client} }

// List calls the observed GET /apigateway/aidex/config/v2/aitools/application browser contract.
func (s *Service) List(ctx context.Context) (*[]Tool, *interfaces.Response, error) {
	endpoint := EndpointList
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result []Tool
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Get calls the observed GET /apigateway/aidex/config/v2/aitools/application browser contract.
func (s *Service) Get(ctx context.Context, id string) (*Tool, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGet
	endpoint += "/" + url.PathEscape(id)
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Create calls the observed POST /apigateway/aidex/config/v2/aitools/application browser contract.
func (s *Service) Create(ctx context.Context, request *ToolRequest) (*Tool, *interfaces.Response, error) {
	if err := validateToolRequest(request); err != nil {
		return nil, nil, err
	}
	if request.NQLID[0] != '#' {
		return nil, nil, fmt.Errorf("custom tool nqlId must start with #")
	}
	endpoint := EndpointCreate
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Update calls the observed PUT /apigateway/aidex/config/v2/aitools/application browser contract.
func (s *Service) Update(ctx context.Context, id string, revision int, request *ToolRequest) (*Tool, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := validateToolRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointUpdate
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// Delete calls the observed DELETE /apigateway/aidex/config/v2/aitools/application browser contract.
func (s *Service) Delete(ctx context.Context, id string, revision int) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, err
	}
	endpoint := EndpointDelete
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	return s.client.Delete(ctx, endpoint, nil, headers, nil)
}

// GetCopilot calls the observed GET /apigateway/aidex/config/v2/aitools/ms-copilot browser contract.
func (s *Service) GetCopilot(ctx context.Context, id string) (*Tool, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetCopilot
	endpoint += "/" + url.PathEscape(id)
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateCopilot calls the observed POST /apigateway/aidex/config/v2/aitools/ms-copilot browser contract.
func (s *Service) CreateCopilot(ctx context.Context, request *CopilotRequest) (*Tool, *interfaces.Response, error) {
	if err := validateCopilotRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointCreateCopilot
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateCopilot calls the observed PUT /apigateway/aidex/config/v2/aitools/ms-copilot browser contract.
func (s *Service) UpdateCopilot(ctx context.Context, id string, revision int, request *CopilotRequest) (*Tool, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := validateCopilotRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointUpdateCopilot
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Tool
	response, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteCopilot calls the observed DELETE /apigateway/aidex/config/v2/aitools/ms-copilot browser contract.
func (s *Service) DeleteCopilot(ctx context.Context, id string, revision int) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, err
	}
	endpoint := EndpointDeleteCopilot
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	return s.client.Delete(ctx, endpoint, nil, headers, nil)
}

// ListGoals calls the observed GET /apigateway/aidex/config/v1/goal browser contract.
func (s *Service) ListGoals(ctx context.Context) (*[]Goal, *interfaces.Response, error) {
	endpoint := EndpointListGoals
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result []Goal
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGoal calls the observed GET /apigateway/aidex/config/v1/goal browser contract.
func (s *Service) GetGoal(ctx context.Context, id string) (*Goal, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetGoal
	endpoint += "/" + url.PathEscape(id)
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Goal
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateGoal calls the observed POST /apigateway/aidex/config/v1/goal browser contract.
func (s *Service) CreateGoal(ctx context.Context, request *GoalRequest) (*Goal, *interfaces.Response, error) {
	if err := validateGoalRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointCreateGoal
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Goal
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateGoal calls the observed PUT /apigateway/aidex/config/v1/goal browser contract.
func (s *Service) UpdateGoal(ctx context.Context, id string, revision int, request *GoalRequest) (*Goal, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := validateGoalRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointUpdateGoal
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Goal
	response, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// DeleteGoal calls the observed DELETE /apigateway/aidex/config/v1/goal browser contract.
func (s *Service) DeleteGoal(ctx context.Context, id string, revision int) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, err
	}
	endpoint := EndpointDeleteGoal
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	return s.client.Delete(ctx, endpoint, nil, headers, nil)
}

// GetLegacyTool calls the observed GET /apigateway/aidex/config/v1/tool browser contract.
func (s *Service) GetLegacyTool(ctx context.Context, id string) (*LegacyTool, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetLegacyTool
	endpoint += "/" + url.PathEscape(id)
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result LegacyTool
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListSystemTools calls the observed GET /apigateway/aidex/config/v1/tool/system-tools browser contract.
func (s *Service) ListSystemTools(ctx context.Context) (*[]LegacyTool, *interfaces.Response, error) {
	endpoint := EndpointListSystemTools
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result []LegacyTool
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetLicense calls the observed GET /apigateway/aidex/config/v1/aitools/license browser contract.
func (s *Service) GetLicense(ctx context.Context) (*License, *interfaces.Response, error) {
	endpoint := EndpointGetLicense
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result License
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetRedirectURLs calls the observed GET /apigateway/aidex/config/v1/tool/redirect-urls browser contract.
func (s *Service) GetRedirectURLs(ctx context.Context) (*[]RedirectURL, *interfaces.Response, error) {
	endpoint := EndpointGetRedirectURLs
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result []RedirectURL
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetModule calls the observed GET /apigateway/aidex/config/v1/module browser contract.
func (s *Service) GetModule(ctx context.Context) (*Module, *interfaces.Response, error) {
	endpoint := EndpointGetModule
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Module
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateModule creates the shared experience-campaign settings. This changes tenant configuration.
func (s *Service) CreateModule(ctx context.Context, request *Module) (*Module, *interfaces.Response, error) {
	if err := validateModule(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointCreateModule
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Module
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateModule replaces shared experience-campaign settings at the supplied revision.
func (s *Service) UpdateModule(ctx context.Context, id string, revision int, request *Module) (*Module, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateRevision(revision); err != nil {
		return nil, nil, err
	}
	if err := validateModule(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointUpdateModule
	endpoint += "/" + url.PathEscape(id)
	endpoint += "?" + url.Values{"rev": {strconv.Itoa(revision)}}.Encode()
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Module
	response, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CheckCopilotCredentials validates an existing credential reference with the Copilot integration.
func (s *Service) CheckCopilotCredentials(ctx context.Context, request *CredentialsRequest) (*json.RawMessage, *interfaces.Response, error) {
	if err := validateCredentialsRequest(request); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointCheckCopilotCredentials
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result json.RawMessage
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetOverviewInsights calls the observed GET /apigateway/aidex/insights/overview browser contract.
func (s *Service) GetOverviewInsights(ctx context.Context, language string) (*Insights, *interfaces.Response, error) {
	if err := validateLanguage(language); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetOverviewInsights
	if language != "" {
		endpoint += "?" + url.Values{"lang": {language}}.Encode()
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Insights
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetToolInsights calls the observed POST /apigateway/aidex/insights/tool browser contract.
func (s *Service) GetToolInsights(ctx context.Context, request *ToolInsightsRequest, language string) (*Insights, *interfaces.Response, error) {
	if err := validateToolInsightsRequest(request); err != nil {
		return nil, nil, err
	}
	if err := validateLanguage(language); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetToolInsights
	if language != "" {
		endpoint += "?" + url.Values{"lang": {language}}.Encode()
	}
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Insights
	response, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGovernanceTrends calls the observed GET /apigateway/aidex/insights/observability/governance/tool-governance-trends browser contract.
func (s *Service) GetGovernanceTrends(ctx context.Context) (*GovernanceTrends, *interfaces.Response, error) {
	endpoint := EndpointGetGovernanceTrends
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result GovernanceTrends
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGovernanceActiveUsers calls the observed GET /apigateway/aidex/insights/observability/governance/tool-governance-active-users browser contract.
func (s *Service) GetGovernanceActiveUsers(ctx context.Context) (*GovernanceActiveUsers, *interfaces.Response, error) {
	endpoint := EndpointGetGovernanceActiveUsers
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result GovernanceActiveUsers
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGovernanceDashboard calls the observed GET /apigateway/aidex/insights/observability/governance/tool-governance-dashboard browser contract.
func (s *Service) GetGovernanceDashboard(ctx context.Context) (*GovernanceDashboard, *interfaces.Response, error) {
	endpoint := EndpointGetGovernanceDashboard
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result GovernanceDashboard
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetGoalInsights calls the observed GET /apigateway/aidex/goals/v1/goals browser contract.
func (s *Service) GetGoalInsights(ctx context.Context, id string) (*Insights, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	endpoint := EndpointGetGoalInsights
	endpoint += "/" + url.PathEscape(id)
	endpoint += "/insights"
	headers := map[string]string{"Accept": "application/json", "Content-Type": "application/json"}
	var result Insights
	response, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type AIToolsServiceInterface interface {
	List(ctx context.Context) (*[]Tool, *interfaces.Response, error)
	Get(ctx context.Context, id string) (*Tool, *interfaces.Response, error)
	Create(ctx context.Context, request *ToolRequest) (*Tool, *interfaces.Response, error)
	Update(ctx context.Context, id string, revision int, request *ToolRequest) (*Tool, *interfaces.Response, error)
	Delete(ctx context.Context, id string, revision int) (*interfaces.Response, error)
	GetCopilot(ctx context.Context, id string) (*Tool, *interfaces.Response, error)
	CreateCopilot(ctx context.Context, request *CopilotRequest) (*Tool, *interfaces.Response, error)
	UpdateCopilot(ctx context.Context, id string, revision int, request *CopilotRequest) (*Tool, *interfaces.Response, error)
	DeleteCopilot(ctx context.Context, id string, revision int) (*interfaces.Response, error)
	ListGoals(ctx context.Context) (*[]Goal, *interfaces.Response, error)
	GetGoal(ctx context.Context, id string) (*Goal, *interfaces.Response, error)
	CreateGoal(ctx context.Context, request *GoalRequest) (*Goal, *interfaces.Response, error)
	UpdateGoal(ctx context.Context, id string, revision int, request *GoalRequest) (*Goal, *interfaces.Response, error)
	DeleteGoal(ctx context.Context, id string, revision int) (*interfaces.Response, error)
	GetLegacyTool(ctx context.Context, id string) (*LegacyTool, *interfaces.Response, error)
	ListSystemTools(ctx context.Context) (*[]LegacyTool, *interfaces.Response, error)
	GetLicense(ctx context.Context) (*License, *interfaces.Response, error)
	GetRedirectURLs(ctx context.Context) (*[]RedirectURL, *interfaces.Response, error)
	GetModule(ctx context.Context) (*Module, *interfaces.Response, error)
	CreateModule(ctx context.Context, request *Module) (*Module, *interfaces.Response, error)
	UpdateModule(ctx context.Context, id string, revision int, request *Module) (*Module, *interfaces.Response, error)
	CheckCopilotCredentials(ctx context.Context, request *CredentialsRequest) (*json.RawMessage, *interfaces.Response, error)
	GetOverviewInsights(ctx context.Context, language string) (*Insights, *interfaces.Response, error)
	GetToolInsights(ctx context.Context, request *ToolInsightsRequest, language string) (*Insights, *interfaces.Response, error)
	GetGovernanceTrends(ctx context.Context) (*GovernanceTrends, *interfaces.Response, error)
	GetGovernanceActiveUsers(ctx context.Context) (*GovernanceActiveUsers, *interfaces.Response, error)
	GetGovernanceDashboard(ctx context.Context) (*GovernanceDashboard, *interfaces.Response, error)
	GetGoalInsights(ctx context.Context, id string) (*Insights, *interfaces.Response, error)
}

var _ AIToolsServiceInterface = (*Service)(nil)
