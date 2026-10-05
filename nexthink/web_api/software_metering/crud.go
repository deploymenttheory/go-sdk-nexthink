package software_metering

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

// List calls the management operation observed in the Nexthink UI.
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[ListResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryList, OperationName: "getConfigurations", Variables: nil})
}

// Get calls the management operation observed in the Nexthink UI.
func (s *Service) Get(ctx context.Context, uuid string) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "getConfiguration", Variables: map[string]any{"uuid": uuid}})
}

// Create calls the management operation observed in the Nexthink UI.
func (s *Service) Create(ctx context.Context, request *ConfigurationInput) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateConfiguration(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "createConfiguration", Variables: map[string]any{"configuration": request}})
}

// Update calls the management operation observed in the Nexthink UI.
func (s *Service) Update(ctx context.Context, uuid string, request *ConfigurationInput) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(uuid, request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "updateConfiguration", Variables: map[string]any{"uuid": uuid, "configuration": request}})
}

// Delete calls the management operation observed in the Nexthink UI.
func (s *Service) Delete(ctx context.Context, uuid string) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateID(uuid); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "deleteConfiguration", Variables: map[string]any{"uuid": uuid}})
}

type SoftwareMeteringServiceInterface interface {
	List(ctx context.Context) (*ListResponse, *interfaces.Response, error)
	Get(ctx context.Context, uuid string) (*GetResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *ConfigurationInput) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, uuid string, request *ConfigurationInput) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, uuid string) (*DeleteResponse, *interfaces.Response, error)
}

var _ SoftwareMeteringServiceInterface = (*Service)(nil)
