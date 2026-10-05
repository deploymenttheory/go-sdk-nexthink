package dashboards

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	response, err := s.client.Get(ctx, EndpointList, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
func (s *Service) Create(ctx context.Context, request *DashboardInput) (*CreateResponse, *interfaces.Response, error) {
	if err := ValidateInput(request); err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryCreate, OperationName: "CreateDashboard", Variables: map[string]any{"dashboard": request}})
}
func (s *Service) Get(ctx context.Context, id string, options *GetOptions) (*GetResponse, *interfaces.Response, error) {
	if err := ValidateID(id); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": id}
	if options != nil && options.ProductArea != "" {
		variables["productArea"] = options.ProductArea
	}
	return graphql.ExecuteData[GetResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGet, OperationName: "GetDashboard", Variables: variables})
}
func (s *Service) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, *interfaces.Response, error) {
	if err := ValidateUpdateRequest(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "revision": request.Revision, "dashboard": request.Dashboard}
	addContext(variables, request.ProductArea, request.DashboardType)
	return graphql.ExecuteData[UpdateResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdate, OperationName: "UpdateDashboard", Variables: variables})
}
func (s *Service) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, *interfaces.Response, error) {
	if err := ValidateDeleteRequest(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"dashboardId": request.DashboardID, "revision": request.Revision}
	addContext(variables, request.ProductArea, request.DashboardType)
	return graphql.ExecuteData[DeleteResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryDelete, OperationName: "DeleteDashboard", Variables: variables})
}

type DashboardsServiceInterface interface {
	List(context.Context) (*ListResponse, *interfaces.Response, error)
	Get(context.Context, string, *GetOptions) (*GetResponse, *interfaces.Response, error)
	Create(context.Context, *DashboardInput) (*CreateResponse, *interfaces.Response, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, *interfaces.Response, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, *interfaces.Response, error)
}

var _ DashboardsServiceInterface = (*Service)(nil)

func addContext(variables map[string]any, area, kind string) {
	if area != "" {
		variables["productArea"] = area
	}
	if kind != "" {
		variables["dashboardType"] = kind
	}
}
