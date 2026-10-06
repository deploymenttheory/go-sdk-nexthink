package diagnostics

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// BinaryInfo executes the first-party UI operation BinaryInfo.
func (s *Service) BinaryInfo(ctx context.Context, request *BinaryInfoRequest) (*BinaryInfoResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[BinaryInfoResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryBinaryInfo, OperationName: "BinaryInfo", Variables: variables})
}

// GetDiagnosticContexts executes the first-party UI operation GetDiagnosticContexts.
func (s *Service) GetDiagnosticContexts(ctx context.Context, request *GetDiagnosticContextsRequest) (*GetDiagnosticContextsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDiagnosticContextsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDiagnosticContexts, OperationName: "GetDiagnosticContexts", Variables: variables})
}

// GetDiagnosticOverview executes the first-party UI operation GetDiagnosticOverview.
func (s *Service) GetDiagnosticOverview(ctx context.Context, request *GetDiagnosticOverviewRequest) (*GetDiagnosticOverviewResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDiagnosticOverviewResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDiagnosticOverview, OperationName: "GetDiagnosticOverview", Variables: variables})
}

// GetImpactedAndTotalObjects executes the first-party UI operation GetImpactedAndTotalObjects.
func (s *Service) GetImpactedAndTotalObjects(ctx context.Context, request *GetImpactedAndTotalObjectsRequest) (*GetImpactedAndTotalObjectsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetImpactedAndTotalObjectsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetImpactedAndTotalObjects, OperationName: "GetImpactedAndTotalObjects", Variables: variables})
}

// GetStandaloneDashboard executes the first-party UI operation GetStandaloneDashboard.
func (s *Service) GetStandaloneDashboard(ctx context.Context, request *GetStandaloneDashboardRequest) (*GetStandaloneDashboardResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetStandaloneDashboardResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetStandaloneDashboard, OperationName: "GetStandaloneDashboard", Variables: variables})
}

// HierarchyBreakdownDimensions executes the first-party UI operation HierarchyBreakdownDimensions.
func (s *Service) HierarchyBreakdownDimensions(ctx context.Context, request *HierarchyBreakdownDimensionsRequest) (*HierarchyBreakdownDimensionsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[HierarchyBreakdownDimensionsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryHierarchyBreakdownDimensions, OperationName: "HierarchyBreakdownDimensions", Variables: variables})
}

// IssueTimeseries executes the first-party UI operation IssueTimeseries.
func (s *Service) IssueTimeseries(ctx context.Context, request *IssueTimeseriesRequest) (*IssueTimeseriesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssueTimeseriesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssueTimeseries, OperationName: "IssueTimeseries", Variables: variables})
}

// TroubleshootingInsights executes the first-party UI operation TroubleshootingInsights.
func (s *Service) TroubleshootingInsights(ctx context.Context, request *TroubleshootingInsightsRequest) (*TroubleshootingInsightsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[TroubleshootingInsightsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryTroubleshootingInsights, OperationName: "TroubleshootingInsights", Variables: variables})
}

// GetConfiguration executes the first-party UI operation configuration.
func (s *Service) GetConfiguration(ctx context.Context, request *GetConfigurationRequest) (*GetConfigurationResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetConfigurationResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetConfiguration, OperationName: "configuration", Variables: variables})
}

// IssueEventsAndAssociatedObjectsByHierarchyBreakdown executes the first-party UI operation issueEventsAndAssociatedObjectsByHierarchyBreakdown.
func (s *Service) IssueEventsAndAssociatedObjectsByHierarchyBreakdown(ctx context.Context, request *IssueEventsAndAssociatedObjectsByHierarchyBreakdownRequest) (*IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssueEventsAndAssociatedObjectsByHierarchyBreakdown, OperationName: "issueEventsAndAssociatedObjectsByHierarchyBreakdown", Variables: variables})
}

// IssueEventsAndAssociatedObjectsByLocationBreakdown executes the first-party UI operation issueEventsAndAssociatedObjectsByLocationBreakdown.
func (s *Service) IssueEventsAndAssociatedObjectsByLocationBreakdown(ctx context.Context, request *IssueEventsAndAssociatedObjectsByLocationBreakdownRequest) (*IssueEventsAndAssociatedObjectsByLocationBreakdownResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssueEventsAndAssociatedObjectsByLocationBreakdownResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssueEventsAndAssociatedObjectsByLocationBreakdown, OperationName: "issueEventsAndAssociatedObjectsByLocationBreakdown", Variables: variables})
}

// IssueEventsAndAssociatedObjectsByTechnicalBreakdown executes the first-party UI operation issueEventsAndAssociatedObjectsByTechnicalBreakdown.
func (s *Service) IssueEventsAndAssociatedObjectsByTechnicalBreakdown(ctx context.Context, request *IssueEventsAndAssociatedObjectsByTechnicalBreakdownRequest) (*IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryIssueEventsAndAssociatedObjectsByTechnicalBreakdown, OperationName: "issueEventsAndAssociatedObjectsByTechnicalBreakdown", Variables: variables})
}

// GetDashboard executes the first-party UI operation GetDashboard.
func (s *Service) GetDashboard(ctx context.Context, request *GetDashboardRequest) (*GetDashboardResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDashboardResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDashboard, OperationName: "GetDashboard", Variables: variables})
}

type DiagnosticsServiceInterface interface {
	BinaryInfo(context.Context, *BinaryInfoRequest) (*BinaryInfoResponse, *interfaces.Response, error)
	GetDiagnosticContexts(context.Context, *GetDiagnosticContextsRequest) (*GetDiagnosticContextsResponse, *interfaces.Response, error)
	GetDiagnosticOverview(context.Context, *GetDiagnosticOverviewRequest) (*GetDiagnosticOverviewResponse, *interfaces.Response, error)
	GetImpactedAndTotalObjects(context.Context, *GetImpactedAndTotalObjectsRequest) (*GetImpactedAndTotalObjectsResponse, *interfaces.Response, error)
	GetStandaloneDashboard(context.Context, *GetStandaloneDashboardRequest) (*GetStandaloneDashboardResponse, *interfaces.Response, error)
	HierarchyBreakdownDimensions(context.Context, *HierarchyBreakdownDimensionsRequest) (*HierarchyBreakdownDimensionsResponse, *interfaces.Response, error)
	IssueTimeseries(context.Context, *IssueTimeseriesRequest) (*IssueTimeseriesResponse, *interfaces.Response, error)
	TroubleshootingInsights(context.Context, *TroubleshootingInsightsRequest) (*TroubleshootingInsightsResponse, *interfaces.Response, error)
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, *interfaces.Response, error)
	IssueEventsAndAssociatedObjectsByHierarchyBreakdown(context.Context, *IssueEventsAndAssociatedObjectsByHierarchyBreakdownRequest) (*IssueEventsAndAssociatedObjectsByHierarchyBreakdownResponse, *interfaces.Response, error)
	IssueEventsAndAssociatedObjectsByLocationBreakdown(context.Context, *IssueEventsAndAssociatedObjectsByLocationBreakdownRequest) (*IssueEventsAndAssociatedObjectsByLocationBreakdownResponse, *interfaces.Response, error)
	IssueEventsAndAssociatedObjectsByTechnicalBreakdown(context.Context, *IssueEventsAndAssociatedObjectsByTechnicalBreakdownRequest) (*IssueEventsAndAssociatedObjectsByTechnicalBreakdownResponse, *interfaces.Response, error)
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, *interfaces.Response, error)
}

var _ DiagnosticsServiceInterface = (*Service)(nil)
