package dex_configuration

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ graphql *graphql.Service }

func NewService(c interfaces.HTTPClient) *Service { return &Service{graphql: graphql.NewService(c)} }

// GetAccount calls the observed getAccount UI operation.
func (s *Service) GetAccount(ctx context.Context) (*GetAccountResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetAccountResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetAccount, OperationName: "getAccount"})
}

// UpdateApplications calls the observed mutateApplications UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) UpdateApplications(ctx context.Context, request *UpdateApplicationsRequest) (*UpdateApplicationsResponse, *interfaces.Response, error) {
	if err := validateManagementUpdateApplications(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateApplicationsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdateApplications, OperationName: "mutateApplications", Variables: variables})
}

// GetScoreMetrics calls the observed getEcScoreMetrics UI operation.
func (s *Service) GetScoreMetrics(ctx context.Context, request *GetScoreMetricsRequest) (*GetScoreMetricsResponse, *interfaces.Response, error) {
	if err := validateManagementGetScoreMetrics(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetScoreMetricsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetScoreMetrics, OperationName: "getEcScoreMetrics", Variables: variables})
}

// UpdateScoreMetrics calls the observed mutateEcScoreMetrics UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) UpdateScoreMetrics(ctx context.Context, request *UpdateScoreMetricsRequest) (*UpdateScoreMetricsResponse, *interfaces.Response, error) {
	if err := validateManagementUpdateScoreMetrics(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[UpdateScoreMetricsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryUpdateScoreMetrics, OperationName: "mutateEcScoreMetrics", Variables: variables})
}

// GetVDIOptIn calls the observed getOptInVdi UI operation.
func (s *Service) GetVDIOptIn(ctx context.Context) (*GetVDIOptInResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetVDIOptInResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetVDIOptIn, OperationName: "getOptInVdi"})
}

// OptInVDI calls the observed optInVdi UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) OptInVDI(ctx context.Context) (*OptInVDIResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[OptInVDIResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryOptInVDI, OperationName: "optInVdi"})
}

// GetMemoryMetricsOptIn calls the observed getOptInMemoryMetrics UI operation.
func (s *Service) GetMemoryMetricsOptIn(ctx context.Context) (*GetMemoryMetricsOptInResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetMemoryMetricsOptInResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetMemoryMetricsOptIn, OperationName: "getOptInMemoryMetrics"})
}

// OptInMemoryMetrics calls the observed optInMemoryMetrics UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) OptInMemoryMetrics(ctx context.Context) (*OptInMemoryMetricsResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[OptInMemoryMetricsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryOptInMemoryMetrics, OperationName: "optInMemoryMetrics"})
}

// GetCampaign calls the observed getConfigCampaign UI operation.
func (s *Service) GetCampaign(ctx context.Context) (*GetCampaignResponse, *interfaces.Response, error) {

	return graphql.ExecuteData[GetCampaignResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetCampaign, OperationName: "getConfigCampaign"})
}

// EnableCampaign calls the observed enableCampaign UI operation. It can change tenant configuration; inspect partial GraphQL results even when an error is returned.
func (s *Service) EnableCampaign(ctx context.Context, request *EnableCampaignRequest) (*EnableCampaignResponse, *interfaces.Response, error) {
	if err := validateManagementEnableCampaign(request); err != nil {
		return nil, nil, err
	}
	variables, err := managementVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[EnableCampaignResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryEnableCampaign, OperationName: "enableCampaign", Variables: variables})
}

type DexConfigurationServiceInterface interface {
	GetApplications(context.Context) (*GetApplicationsResponse, *interfaces.Response, error)
	GetAccount(ctx context.Context) (*GetAccountResponse, *interfaces.Response, error)
	UpdateApplications(ctx context.Context, request *UpdateApplicationsRequest) (*UpdateApplicationsResponse, *interfaces.Response, error)
	GetScoreMetrics(ctx context.Context, request *GetScoreMetricsRequest) (*GetScoreMetricsResponse, *interfaces.Response, error)
	UpdateScoreMetrics(ctx context.Context, request *UpdateScoreMetricsRequest) (*UpdateScoreMetricsResponse, *interfaces.Response, error)
	GetVDIOptIn(ctx context.Context) (*GetVDIOptInResponse, *interfaces.Response, error)
	OptInVDI(ctx context.Context) (*OptInVDIResponse, *interfaces.Response, error)
	GetMemoryMetricsOptIn(ctx context.Context) (*GetMemoryMetricsOptInResponse, *interfaces.Response, error)
	OptInMemoryMetrics(ctx context.Context) (*OptInMemoryMetricsResponse, *interfaces.Response, error)
	GetCampaign(ctx context.Context) (*GetCampaignResponse, *interfaces.Response, error)
	EnableCampaign(ctx context.Context, request *EnableCampaignRequest) (*EnableCampaignResponse, *interfaces.Response, error)
}

var _ DexConfigurationServiceInterface = (*Service)(nil)

// GetApplications reads DEX application selections; it is distinct from the metering application query.
func (s *Service) GetApplications(ctx context.Context) (*GetApplicationsResponse, *interfaces.Response, error) {
	return graphql.ExecuteData[GetApplicationsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetApplications, OperationName: "getApplications"})
}
