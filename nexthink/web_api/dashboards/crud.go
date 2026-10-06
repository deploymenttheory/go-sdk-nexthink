package dashboards

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	GetProductShellMenu(ctx context.Context) (*ProductShellMenuResponse, *interfaces.Response, error)
	ListFields(ctx context.Context, request *FieldsRequest) (*ListFieldsResponse, *interfaces.Response, error)
	ListCollections(ctx context.Context) (*ListCollectionsResponse, *interfaces.Response, error)
	GetConfiguration(ctx context.Context) (*GetConfigurationResponse, *interfaces.Response, error)
	Import(context.Context, *ImportRequest) (*ImportResponse, *interfaces.Response, error)
	Duplicate(context.Context, *DuplicateRequest) (*DuplicateResponse, *interfaces.Response, error)
	Export(context.Context, *MutationContext) (*ExportResponse, *interfaces.Response, error)
	UpdateLayout(context.Context, *UpdateLayoutRequest) (*UpdateLayoutResponse, *interfaces.Response, error)
	DeleteTab(context.Context, *MutationContext) (*DeleteTabResponse, *interfaces.Response, error)
	UpdateTabs(context.Context, *UpdateTabsRequest) (*UpdateTabsResponse, *interfaces.Response, error)
	UpdateTab(context.Context, *UpdateTabRequest) (*UpdateTabResponse, *interfaces.Response, error)
	CreateTab(context.Context, *MutationContext) (*CreateTabResponse, *interfaces.Response, error)
	DeleteFilter(context.Context, *DeleteFilterRequest) (*DeleteFilterResponse, *interfaces.Response, error)
	UpdateFilter(context.Context, *FilterRequest) (*UpdateFilterResponse, *interfaces.Response, error)
	CreateFilter(context.Context, *FilterRequest) (*CreateFilterResponse, *interfaces.Response, error)
	DeleteWidget(context.Context, *DeleteWidgetRequest) (*DeleteWidgetResponse, *interfaces.Response, error)
	UpdateWidget(context.Context, *WidgetRequest) (*UpdateWidgetResponse, *interfaces.Response, error)
	CreateWidget(context.Context, *WidgetRequest) (*CreateWidgetResponse, *interfaces.Response, error)
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

func (s *Service) CreateWidget(ctx context.Context, request *WidgetRequest) (*CreateWidgetResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateWidgetResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryCreateWidget, OperationName: "CreateWidget", Variables: variables})
}

func (s *Service) UpdateWidget(ctx context.Context, request *WidgetRequest) (*UpdateWidgetResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateWidgetResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdateWidget, OperationName: "UpdateWidget", Variables: variables})
}

func (s *Service) DeleteWidget(ctx context.Context, request *DeleteWidgetRequest) (*DeleteWidgetResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteWidgetResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryDeleteWidget, OperationName: "DeleteWidget", Variables: variables})
}

func (s *Service) CreateFilter(ctx context.Context, request *FilterRequest) (*CreateFilterResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateFilterResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryCreateFilter, OperationName: "CreateFilter", Variables: variables})
}

func (s *Service) UpdateFilter(ctx context.Context, request *FilterRequest) (*UpdateFilterResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateFilterResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdateFilter, OperationName: "UpdateFilter", Variables: variables})
}

func (s *Service) DeleteFilter(ctx context.Context, request *DeleteFilterRequest) (*DeleteFilterResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteFilterResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryDeleteFilter, OperationName: "DeleteFilter", Variables: variables})
}

func (s *Service) CreateTab(ctx context.Context, request *MutationContext) (*CreateTabResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[CreateTabResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryCreateTab, OperationName: "CreateTab", Variables: variables})
}

func (s *Service) UpdateTab(ctx context.Context, request *UpdateTabRequest) (*UpdateTabResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateTabResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdateTab, OperationName: "UpdateTab", Variables: variables})
}

func (s *Service) UpdateTabs(ctx context.Context, request *UpdateTabsRequest) (*UpdateTabsResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateTabsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdateTabs, OperationName: "UpdateTabs", Variables: variables})
}

func (s *Service) DeleteTab(ctx context.Context, request *MutationContext) (*DeleteTabResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	if err := ValidateID(request.TabID); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DeleteTabResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryDeleteTab, OperationName: "DeleteTab", Variables: variables})
}

func (s *Service) UpdateLayout(ctx context.Context, request *UpdateLayoutRequest) (*UpdateLayoutResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateLayoutResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryUpdateLayout, OperationName: "UpdateLayout", Variables: variables})
}

func (s *Service) Export(ctx context.Context, request *MutationContext) (*ExportResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ExportResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryExport, OperationName: "ExportDashboard", Variables: variables})
}

func (s *Service) Duplicate(ctx context.Context, request *DuplicateRequest) (*DuplicateResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[DuplicateResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryDuplicate, OperationName: "DuplicateDashboard", Variables: variables})
}

func (s *Service) Import(ctx context.Context, request *ImportRequest) (*ImportResponse, *interfaces.Response, error) {
	if err := validateMutationRequest(request); err != nil {
		return nil, nil, err
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ImportResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryImport, OperationName: "DashboardImport", Variables: variables})
}

func mutationVariables(request any) (map[string]any, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	err = json.Unmarshal(data, &result)
	return result, err
}

func (s *Service) GetConfiguration(ctx context.Context) (*GetConfigurationResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[GetConfigurationResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetConfiguration, OperationName: "Configurations", Variables: map[string]any{}})
}

func (s *Service) ListCollections(ctx context.Context) (*ListCollectionsResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[ListCollectionsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Headers: metadataHeaders(), Query: queryListCollections, OperationName: "Collections", Variables: map[string]any{}})
}

func (s *Service) ListFields(ctx context.Context, request *FieldsRequest) (*ListFieldsResponse, *interfaces.Response, error) {
	if request == nil || strings.TrimSpace(request.Collection) == "" {
		return nil, nil, fmt.Errorf("collection is required")
	}
	variables, err := mutationVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ListFieldsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Headers: metadataHeaders(), Query: queryListFields, OperationName: "Fields", Variables: variables})
}

// Metadata resolvers require the same time context as the dashboard UI.
func metadataHeaders() map[string]string {
	return map[string]string{"x-nxt-waas-iso-date-time": time.Now().UTC().Format(time.RFC3339Nano), "x-nxt-waas-timezone": "UTC", "x-nxt-waas-utc-offset": "0"}
}

// GetProductShellMenu reads dashboard navigation entries exposed to the product shell.
func (s *Service) GetProductShellMenu(ctx context.Context) (*ProductShellMenuResponse, *interfaces.Response, error) {
	var result ProductShellMenuResponse
	response, err := s.client.Get(ctx, EndpointProductShellMenu, nil, map[string]string{"Accept": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
