package writing_assistant

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct {
	client  interfaces.HTTPClient
	graphql *graphql.Service
}

func NewService(c interfaces.HTTPClient) *Service {
	return &Service{client: c, graphql: graphql.NewService(c)}
}

// Get calls the management operation observed in the Nexthink UI.
func (s *Service) Get(ctx context.Context, id string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "GetWritingAssistant", Variables: map[string]any{"id": id}})
}

// Create calls the management operation observed in the Nexthink UI.
func (s *Service) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "CreateWritingAssistant", Variables: map[string]any{"input": request}})
}

// Update calls the management operation observed in the Nexthink UI.
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "UpdateWritingAssistant", Variables: map[string]any{"input": request}})
}

// Delete calls the management operation observed in the Nexthink UI.
func (s *Service) Delete(ctx context.Context, id string) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteWritingAssistant", Variables: map[string]any{"id": id}})
}

// List returns the Writing Assistant content listing, including tool and revision metadata.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type WritingAssistantServiceInterface interface {
	Get(ctx context.Context, id string) (*GetResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *CreateRequest) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, id string) (*DeleteResponse, *interfaces.Response, error)
	List(context.Context) (*ListResponse, *interfaces.Response, error)
}

var _ WritingAssistantServiceInterface = (*Service)(nil)
