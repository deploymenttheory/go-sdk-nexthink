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
	AutoConfigureMetering(ctx context.Context, request *AutoConfigureMeteringRequest) (*AutoConfigureMeteringResponse, *interfaces.Response, error)
	GetApplications(ctx context.Context, request *GetApplicationsRequest) (*GetApplicationsResponse, *interfaces.Response, error)
	GetConfigurationByApplicationUUID(ctx context.Context, request *GetConfigurationByApplicationUUIDRequest) (*GetConfigurationByApplicationUUIDResponse, *interfaces.Response, error)
	GetConfigurationDetails(ctx context.Context, request *GetConfigurationDetailsRequest) (*GetConfigurationDetailsResponse, *interfaces.Response, error)
	GetEmployeesTable(ctx context.Context, request *GetEmployeesTableRequest) (*GetEmployeesTableResponse, *interfaces.Response, error)
	GetPackages(ctx context.Context, request *GetPackagesRequest) (*GetPackagesResponse, *interfaces.Response, error)
	GetUsageBreakdown(ctx context.Context, request *GetUsageBreakdownRequest) (*GetUsageBreakdownResponse, *interfaces.Response, error)
	GetUsageByLicenseEndpoint(ctx context.Context, request *GetUsageByLicenseEndpointRequest) (*GetUsageByLicenseEndpointResponse, *interfaces.Response, error)
	GetUsageByLicenseEndpointCount(ctx context.Context, request *GetUsageByLicenseEndpointCountRequest) (*GetUsageByLicenseEndpointCountResponse, *interfaces.Response, error)
	GetUsageDistribution(ctx context.Context, request *GetUsageDistributionRequest) (*GetUsageDistributionResponse, *interfaces.Response, error)
	GetUsageDistributionByCategory(ctx context.Context, request *GetUsageDistributionByCategoryRequest) (*GetUsageDistributionByCategoryResponse, *interfaces.Response, error)
	GetUsageOverview(ctx context.Context, request *GetUsageOverviewRequest) (*GetUsageOverviewResponse, *interfaces.Response, error)
	GetConfigurationUsageOverview(ctx context.Context, request *GetConfigurationUsageOverviewRequest) (*GetConfigurationUsageOverviewResponse, *interfaces.Response, error)
	List(ctx context.Context) (*ListResponse, *interfaces.Response, error)
	Get(ctx context.Context, uuid string) (*GetResponse, *interfaces.Response, error)
	Create(ctx context.Context, request *ConfigurationInput) (*CreateResponse, *interfaces.Response, error)
	Update(ctx context.Context, uuid string, request *ConfigurationInput) (*UpdateResponse, *interfaces.Response, error)
	Delete(ctx context.Context, uuid string) (*DeleteResponse, *interfaces.Response, error)
}

var _ SoftwareMeteringServiceInterface = (*Service)(nil)

// AutoConfigureMetering executes the observed autoConfigureMetering GraphQL operation.
func (s *Service) AutoConfigureMetering(ctx context.Context, request *AutoConfigureMeteringRequest) (*AutoConfigureMeteringResponse, *interfaces.Response, error) {
	if err := validateAutoConfigureMetering(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"applicationId": request.ApplicationID}

	return graphql.ExecuteData[AutoConfigureMeteringResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryAutoConfigureMetering, OperationName: "autoConfigureMetering", Variables: variables})
}

// GetApplications executes the observed getApplications GraphQL operation.
func (s *Service) GetApplications(ctx context.Context, request *GetApplicationsRequest) (*GetApplicationsResponse, *interfaces.Response, error) {
	if err := validateGetApplications(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{}
	if request.UUID != nil {
		variables["uuid"] = request.UUID
	}
	return graphql.ExecuteData[GetApplicationsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetApplications, OperationName: "getApplications", Variables: variables})
}

// GetConfigurationByApplicationUUID executes the observed getConfigurationByApplicationUuid GraphQL operation.
func (s *Service) GetConfigurationByApplicationUUID(ctx context.Context, request *GetConfigurationByApplicationUUIDRequest) (*GetConfigurationByApplicationUUIDResponse, *interfaces.Response, error) {
	if err := validateGetConfigurationByApplicationUUID(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"applicationUuid": request.ApplicationUUID}

	return graphql.ExecuteData[GetConfigurationByApplicationUUIDResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetConfigurationByApplicationUUID, OperationName: "getConfigurationByApplicationUuid", Variables: variables})
}

// GetConfigurationDetails executes the observed getConfigurationDetails GraphQL operation.
func (s *Service) GetConfigurationDetails(ctx context.Context, request *GetConfigurationDetailsRequest) (*GetConfigurationDetailsResponse, *interfaces.Response, error) {
	if err := validateGetConfigurationDetails(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID}

	return graphql.ExecuteData[GetConfigurationDetailsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetConfigurationDetails, OperationName: "getConfigurationDetails", Variables: variables})
}

// GetEmployeesTable executes the observed getEmployeesTable GraphQL operation.
func (s *Service) GetEmployeesTable(ctx context.Context, request *GetEmployeesTableRequest) (*GetEmployeesTableResponse, *interfaces.Response, error) {
	if err := validateGetEmployeesTable(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "applicationUuid": request.ApplicationUUID, "pagination": request.Pagination, "orderInput": request.OrderInput, "timeRange": request.TimeRange}

	return graphql.ExecuteData[GetEmployeesTableResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetEmployeesTable, OperationName: "getEmployeesTable", Variables: variables})
}

// GetPackages executes the observed getPackages GraphQL operation.
func (s *Service) GetPackages(ctx context.Context, request *GetPackagesRequest) (*GetPackagesResponse, *interfaces.Response, error) {
	if err := validateGetPackages(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"search": request.Search}

	return graphql.ExecuteData[GetPackagesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetPackages, OperationName: "getPackages", Variables: variables})
}

// GetUsageBreakdown executes the observed getUsageBreakdown GraphQL operation.
func (s *Service) GetUsageBreakdown(ctx context.Context, request *GetUsageBreakdownRequest) (*GetUsageBreakdownResponse, *interfaces.Response, error) {
	if err := validateGetUsageBreakdown(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "applicationUuid": request.ApplicationUUID, "timeRange": request.TimeRange}

	return graphql.ExecuteData[GetUsageBreakdownResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageBreakdown, OperationName: "getUsageBreakdown", Variables: variables})
}

// GetUsageByLicenseEndpoint executes the observed getUsageByLicenseEndpoint GraphQL operation.
func (s *Service) GetUsageByLicenseEndpoint(ctx context.Context, request *GetUsageByLicenseEndpointRequest) (*GetUsageByLicenseEndpointResponse, *interfaces.Response, error) {
	if err := validateGetUsageByLicenseEndpoint(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "pagination": request.Pagination, "orderInput": request.OrderInput, "timeRange": request.TimeRange}
	if request.Filter != nil {
		variables["filter"] = request.Filter
	}
	return graphql.ExecuteData[GetUsageByLicenseEndpointResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageByLicenseEndpoint, OperationName: "getUsageByLicenseEndpoint", Variables: variables})
}

// GetUsageByLicenseEndpointCount executes the observed getUsageByLicenseEndpointCount GraphQL operation.
func (s *Service) GetUsageByLicenseEndpointCount(ctx context.Context, request *GetUsageByLicenseEndpointCountRequest) (*GetUsageByLicenseEndpointCountResponse, *interfaces.Response, error) {
	if err := validateGetUsageByLicenseEndpointCount(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"configurationUuid": request.ConfigurationUUID, "timeRange": request.TimeRange}
	if request.Filter != nil {
		variables["filter"] = request.Filter
	}
	return graphql.ExecuteData[GetUsageByLicenseEndpointCountResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageByLicenseEndpointCount, OperationName: "getUsageByLicenseEndpointCount", Variables: variables})
}

// GetUsageDistribution executes the observed getUsageDistribution GraphQL operation.
func (s *Service) GetUsageDistribution(ctx context.Context, request *GetUsageDistributionRequest) (*GetUsageDistributionResponse, *interfaces.Response, error) {
	if err := validateGetUsageDistribution(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "timeRange": request.TimeRange}
	if request.Filter != nil {
		variables["filter"] = request.Filter
	}
	return graphql.ExecuteData[GetUsageDistributionResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageDistribution, OperationName: "getUsageDistribution", Variables: variables})
}

// GetUsageDistributionByCategory executes the observed getUsageDistributionByCategory GraphQL operation.
func (s *Service) GetUsageDistributionByCategory(ctx context.Context, request *GetUsageDistributionByCategoryRequest) (*GetUsageDistributionByCategoryResponse, *interfaces.Response, error) {
	if err := validateGetUsageDistributionByCategory(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "categoryKey": request.CategoryKey, "timeRange": request.TimeRange}
	if request.Filter != nil {
		variables["filter"] = request.Filter
	}
	return graphql.ExecuteData[GetUsageDistributionByCategoryResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageDistributionByCategory, OperationName: "getUsageDistributionByCategory", Variables: variables})
}

// GetUsageOverview executes the observed getUsageOverview GraphQL operation.
func (s *Service) GetUsageOverview(ctx context.Context, request *GetUsageOverviewRequest) (*GetUsageOverviewResponse, *interfaces.Response, error) {
	if err := validateGetUsageOverview(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "applicationUuid": request.ApplicationUUID, "timeRange": request.TimeRange}

	return graphql.ExecuteData[GetUsageOverviewResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetUsageOverview, OperationName: "getUsageOverview", Variables: variables})
}

// GetConfigurationUsageOverview retrieves the complete metering configuration usage breakdown.
func (s *Service) GetConfigurationUsageOverview(ctx context.Context, request *GetConfigurationUsageOverviewRequest) (*GetConfigurationUsageOverviewResponse, *interfaces.Response, error) {
	if err := validateGetConfigurationUsageOverview(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"uuid": request.UUID, "timeRange": request.TimeRange}
	if request.Filter != nil {
		variables["filter"] = request.Filter
	}
	return graphql.ExecuteData[GetConfigurationUsageOverviewResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetConfigurationUsageOverview, OperationName: "getUsageOverview", Variables: variables})
}
