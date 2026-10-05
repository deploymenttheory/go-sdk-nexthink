package monitors

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
func (s *Service) Create(ctx context.Context, request *MonitorInput) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateCreateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "AddNqlMonitor", Variables: map[string]any{"monitor": request}})
}
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "UpdateNqlMonitor", Variables: map[string]any{"docUuid": request.DocUUID, "revision": request.Revision, "monitor": request.Monitor}})
}
func (s *Service) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateDeleteRequest(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteMonitor", Variables: map[string]any{"deleteMonitorInput": request}})
}
func (s *Service) Get(ctx context.Context, docUUID string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(docUUID); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "NqlMonitor", Variables: map[string]any{"docUuid": docUUID}})
}
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type MonitorsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *MonitorInput) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error)
	Get(ctx context.Context, docUUID string) (*GetResponse, *interfaces.Response, error)
}

var _ MonitorsServiceInterface = (*Service)(nil)
