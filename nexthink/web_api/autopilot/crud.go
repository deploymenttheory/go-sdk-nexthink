package autopilot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"mime"
	"net/url"
)

type Service struct {
	client  interfaces.HTTPClient
	graphql *graphql.Service
}

func NewService(c interfaces.HTTPClient) *Service {
	return &Service{client: c, graphql: graphql.NewService(c)}
}

// GetConfiguration calls the observed Autopilot UI operation.
func (s *Service) GetConfiguration(ctx context.Context) (*Configuration, *interfaces.Response, error) {
	var result Configuration
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/configuration", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateWebSearch calls the observed Autopilot UI operation.
func (s *Service) UpdateWebSearch(ctx context.Context, request *BooleanValue) (*interfaces.Response, error) {
	if err := validateUpdateWebSearch(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/cockpit/configuration/websearch", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// UpdateFilesystemTool calls the observed Autopilot UI operation.
func (s *Service) UpdateFilesystemTool(ctx context.Context, request *BooleanValue) (*interfaces.Response, error) {
	if err := validateUpdateFilesystemTool(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/cockpit/configuration/filesystemtool", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// UpdateAgentName calls the observed Autopilot UI operation.
func (s *Service) UpdateAgentName(ctx context.Context, request *StringValue) (*interfaces.Response, error) {
	if err := validateUpdateAgentName(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, Endpoint+"/cockpit/configuration/agentname", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// GetSettings calls the observed Autopilot UI operation.
func (s *Service) GetSettings(ctx context.Context) (*Settings, *interfaces.Response, error) {
	var result Settings
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/settings", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// SaveSettings calls the observed Autopilot UI operation.
func (s *Service) SaveSettings(ctx context.Context, request *Settings) (*Settings, *interfaces.Response, error) {
	if err := validateSaveSettings(request); err != nil {
		return nil, nil, err
	}
	var result Settings
	response, err := s.client.Post(ctx, Endpoint+"/cockpit/settings", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetWebSearchDomains calls the observed Autopilot UI operation.
func (s *Service) GetWebSearchDomains(ctx context.Context) (*WebSearchDomains, *interfaces.Response, error) {
	var result WebSearchDomains
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/settings/websearch/domains", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ReplaceWebSearchDomains calls the observed Autopilot UI operation.
func (s *Service) ReplaceWebSearchDomains(ctx context.Context, request *WebSearchDomainsRequest) (*WebSearchDomains, *interfaces.Response, error) {
	if err := validateReplaceWebSearchDomains(request); err != nil {
		return nil, nil, err
	}
	var result WebSearchDomains
	response, err := s.client.Put(ctx, Endpoint+"/cockpit/settings/websearch/domains", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetApproval calls the observed Autopilot UI operation.
func (s *Service) GetApproval(ctx context.Context, id string) (*Approval, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	var result Approval
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/approval/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateApproval calls the observed Autopilot UI operation.
func (s *Service) CreateApproval(ctx context.Context, request *Approval) (*Approval, *interfaces.Response, error) {
	if err := validateCreateApproval(request); err != nil {
		return nil, nil, err
	}
	var result Approval
	response, err := s.client.Post(ctx, Endpoint+"/cockpit/approval", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateApproval calls the observed Autopilot UI operation.
func (s *Service) UpdateApproval(ctx context.Context, id string, request *ApprovalInput) (*Approval, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	if err := validateUpdateApproval(request); err != nil {
		return nil, nil, err
	}
	var result Approval
	response, err := s.client.Patch(ctx, Endpoint+"/cockpit/approval/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListCalls calls the observed Autopilot UI operation.
func (s *Service) ListCalls(ctx context.Context) (*CallsResponse, *interfaces.Response, error) {
	var result CallsResponse
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/calls", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetKnowledgeArticleCount calls the observed Autopilot UI operation.
func (s *Service) GetKnowledgeArticleCount(ctx context.Context) (*ArticleCount, *interfaces.Response, error) {
	var result ArticleCount
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/knowledge-base/articles/count", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListKnowledgeConnectors calls the observed Autopilot UI operation.
func (s *Service) ListKnowledgeConnectors(ctx context.Context) (*KnowledgeConnectors, *interfaces.Response, error) {
	var result KnowledgeConnectors
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/knowledge-base/connectors", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// CreateTicket calls the observed Autopilot UI operation.
func (s *Service) CreateTicket(ctx context.Context, request *TicketRequest) (*TicketResult, *interfaces.Response, error) {
	if err := validateCreateTicket(request); err != nil {
		return nil, nil, err
	}
	var result TicketResult
	response, err := s.client.Post(ctx, ITSMEndpoint+"/cockpit/tickets", request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetConversation calls the observed Autopilot UI operation.
func (s *Service) GetConversation(ctx context.Context, id string) (*Conversation, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	var result Conversation
	response, err := s.client.Get(ctx, Endpoint+"/cockpit/conversations/"+url.PathEscape(id), nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// GetRecommendationConversationIDs calls the observed Autopilot UI operation.
func (s *Service) GetRecommendationConversationIDs(ctx context.Context, id string) (*ConversationIDs, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	var result ConversationIDs
	response, err := s.client.Get(ctx, RecommendationsEndpoint+"/knowledge-recommendations/"+url.PathEscape(id)+"/conversation-ids", nil, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// UpdateAgentActionInputs calls the observed Autopilot UI operation.
func (s *Service) UpdateAgentActionInputs(ctx context.Context, id string, request *AgentActionInputsRequest) (*interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := validateUpdateAgentActionInputs(request); err != nil {
		return nil, err
	}
	return s.client.Patch(ctx, AgentActionsEndpoint+"/"+url.PathEscape(id), request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, nil)
}

// DownloadCategorization downloads the original CSV bytes and server filename.
func (s *Service) DownloadCategorization(ctx context.Context) (*CategorizationDownload, *interfaces.Response, error) {
	response, data, err := s.client.GetBytes(ctx, Endpoint+"/cockpit/settings/categorization/csv", nil, map[string]string{"Accept": "text/csv"})
	if err != nil {
		return nil, response, err
	}
	result := &CategorizationDownload{Data: data}
	if response != nil {
		_, params, _ := mime.ParseMediaType(response.Headers.Get("Content-Disposition"))
		result.Filename = params["filename"]
	}
	return result, response, nil
}

// GetAgentActionInputs returns the agent action selected by the UI GraphQL query, retaining partial data.
func (s *Service) GetAgentActionInputs(ctx context.Context, id string) (*AgentAction, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	result, response, err := s.graphql.Execute(ctx, "graphql.remote_actions", graphql.GraphQLRequest{OperationName: "GetRemoteActionQuery", Query: `
query GetRemoteActionQuery($uuid: String!) {
  agentActionByUid(uid: $uuid) {
    scriptInfo {
      inputs {
        allowCustomValue
        description
        id
        name
        options
        usedByMacOs
        usedByWindows
        type
      }
    }
  }
}
`, Variables: map[string]any{"uuid": id}})
	if result == nil {
		return nil, response, err
	}
	var data struct {
		AgentAction *AgentAction `json:"agentActionByUid"`
	}
	if len(result.Data) > 0 && string(result.Data) != "null" {
		if decodeErr := json.Unmarshal(result.Data, &data); decodeErr != nil {
			return nil, response, fmt.Errorf("decode agent action: %w", decodeErr)
		}
	}
	return data.AgentAction, response, err
}

// GetAgentAction returns the agent action selected by the UI GraphQL query, retaining partial data.
func (s *Service) GetAgentAction(ctx context.Context, id string) (*AgentAction, *interfaces.Response, error) {
	if err := validateID(id); err != nil {
		return nil, nil, err
	}
	result, response, err := s.graphql.Execute(ctx, "graphql.remote_actions", graphql.GraphQLRequest{OperationName: "GetAgentActionDetailsQuery", Query: `
query GetAgentActionDetailsQuery($uuid: String!) {
  agentActionByUid(uid: $uuid) {
    contentId
    title
    name
    description
    purpose
    scriptInfo {
      runAs
      timeoutSeconds
      inputs {
        allowCustomValue
        description
        id
        name
        options
        type
        usedByMacOs
        usedByWindows
      }
      scriptWindows {
        name
      }
      scriptMacOs {
        name
      }
      impact {
        windows {
          behavior
          expectedImpact
        }
        macos: macOs {
          behavior
          expectedImpact
        }
      }
    }
  }
}
`, Variables: map[string]any{"uuid": id}})
	if result == nil {
		return nil, response, err
	}
	var data struct {
		AgentAction *AgentAction `json:"agentActionByUid"`
	}
	if len(result.Data) > 0 && string(result.Data) != "null" {
		if decodeErr := json.Unmarshal(result.Data, &data); decodeErr != nil {
			return nil, response, fmt.Errorf("decode agent action: %w", decodeErr)
		}
	}
	return data.AgentAction, response, err
}
