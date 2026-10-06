package application_experience

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct {
	graphql *graphql.Service
	client  interfaces.HTTPClient
}

func NewService(c interfaces.HTTPClient) *Service {
	return &Service{graphql: graphql.NewService(c), client: c}
}

// GetApplicationsOverviewDesktopInvestigations executes the observed GetApplicationsOverviewDesktopInvestigations GraphQL operation.
func (s *Service) GetApplicationsOverviewDesktopInvestigations(ctx context.Context, request *GetApplicationsOverviewDesktopInvestigationsRequest) (*GetApplicationsOverviewDesktopInvestigationsResponse, *interfaces.Response, error) {
	if err := validateGetApplicationsOverviewDesktopInvestigations(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[GetApplicationsOverviewDesktopInvestigationsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetApplicationsOverviewDesktopInvestigations, OperationName: "GetApplicationsOverviewDesktopInvestigations", Variables: variables})
}

// GetAvgNetworkResponseTime executes the observed GetAvgNetworkResponseTime GraphQL operation.
func (s *Service) GetAvgNetworkResponseTime(ctx context.Context, request *GetAvgNetworkResponseTimeRequest) (*GetAvgNetworkResponseTimeResponse, *interfaces.Response, error) {
	if err := validateGetAvgNetworkResponseTime(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetAvgNetworkResponseTimeResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetAvgNetworkResponseTime, OperationName: "GetAvgNetworkResponseTime", Variables: variables})
}

// GetBinarySuggestion executes the observed GetBinarySuggestion GraphQL operation.
func (s *Service) GetBinarySuggestion(ctx context.Context, request *GetBinarySuggestionRequest) (*GetBinarySuggestionResponse, *interfaces.Response, error) {
	if err := validateGetBinarySuggestion(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"input": request.Input}

	return graphql.ExecuteData[GetBinarySuggestionResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetBinarySuggestion, OperationName: "GetBinarySuggestion", Variables: variables})
}

// GetDeviceCentricMetricBreakdown executes the observed GetDeviceCentricMetricBreakdown GraphQL operation.
func (s *Service) GetDeviceCentricMetricBreakdown(ctx context.Context, request *GetDeviceCentricMetricBreakdownRequest) (*GetDeviceCentricMetricBreakdownResponse, *interfaces.Response, error) {
	if err := validateGetDeviceCentricMetricBreakdown(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"desktopAppExperienceId": request.DesktopAppExperienceID, "dimensionToGroupBy": request.DimensionToGroupBy, "timeFrameFilter": request.TimeFrameFilter, "dimensionsFilterExcludingCurrentBreakdown": request.DimensionsFilterExcludingCurrentBreakdown, "limit": request.Limit}
	if request.OrderBy != nil {
		variables["orderBy"] = request.OrderBy
	}
	return graphql.ExecuteData[GetDeviceCentricMetricBreakdownResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetDeviceCentricMetricBreakdown, OperationName: "GetDeviceCentricMetricBreakdown", Variables: variables})
}

// GetFailedConnectionsRatio executes the observed GetFailedConnectionsRatio GraphQL operation.
func (s *Service) GetFailedConnectionsRatio(ctx context.Context, request *GetFailedConnectionsRatioRequest) (*GetFailedConnectionsRatioResponse, *interfaces.Response, error) {
	if err := validateGetFailedConnectionsRatio(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetFailedConnectionsRatioResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetFailedConnectionsRatio, OperationName: "GetFailedConnectionsRatio", Variables: variables})
}

// GetInsights executes the observed GetInsights GraphQL operation.
func (s *Service) GetInsights(ctx context.Context, request *GetInsightsRequest) (*GetInsightsResponse, *interfaces.Response, error) {
	if err := validateGetInsights(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"applicationId": request.ApplicationID, "applicationName": request.ApplicationName, "dimensionsFilter": request.DimensionsFilter, "insightType": request.InsightType, "timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[GetInsightsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetInsights, OperationName: "GetInsights", Variables: variables})
}

// GetMetricBreakdown executes the observed GetMetricBreakdown GraphQL operation.
func (s *Service) GetMetricBreakdown(ctx context.Context, request *GetMetricBreakdownRequest) (*GetMetricBreakdownResponse, *interfaces.Response, error) {
	if err := validateGetMetricBreakdown(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"desktopAppExperienceId": request.DesktopAppExperienceID, "dimensionToGroupBy": request.DimensionToGroupBy, "timeFrameFilter": request.TimeFrameFilter, "dimensionsFilterExcludingCurrentBreakdown": request.DimensionsFilterExcludingCurrentBreakdown, "limit": request.Limit}
	if request.OrderBy != nil {
		variables["orderBy"] = request.OrderBy
	}
	return graphql.ExecuteData[GetMetricBreakdownResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetMetricBreakdown, OperationName: "GetMetricBreakdown", Variables: variables})
}

// GetNumOfCrashesAndDevices executes the observed GetNumOfCrashesAndDevices GraphQL operation.
func (s *Service) GetNumOfCrashesAndDevices(ctx context.Context, request *GetNumOfCrashesAndDevicesRequest) (*GetNumOfCrashesAndDevicesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfCrashesAndDevices(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfCrashesAndDevicesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfCrashesAndDevices, OperationName: "GetNumOfCrashesAndDevices", Variables: variables})
}

// GetNumOfCrashesAndEmployees executes the observed GetNumOfCrashesAndEmployees GraphQL operation.
func (s *Service) GetNumOfCrashesAndEmployees(ctx context.Context, request *GetNumOfCrashesAndEmployeesRequest) (*GetNumOfCrashesAndEmployeesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfCrashesAndEmployees(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfCrashesAndEmployeesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfCrashesAndEmployees, OperationName: "GetNumOfCrashesAndEmployees", Variables: variables})
}

// GetNumOfDevices executes the observed GetNumOfDevices GraphQL operation.
func (s *Service) GetNumOfDevices(ctx context.Context, request *GetNumOfDevicesRequest) (*GetNumOfDevicesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfDevices(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfDevicesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfDevices, OperationName: "GetNumOfDevices", Variables: variables})
}

// GetNumOfDevicesWithCrashes executes the observed GetNumOfDevicesWithCrashes GraphQL operation.
func (s *Service) GetNumOfDevicesWithCrashes(ctx context.Context, request *GetNumOfDevicesWithCrashesRequest) (*GetNumOfDevicesWithCrashesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfDevicesWithCrashes(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfDevicesWithCrashesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfDevicesWithCrashes, OperationName: "GetNumOfDevicesWithCrashes", Variables: variables})
}

// GetNumOfEmployees executes the observed GetNumOfEmployees GraphQL operation.
func (s *Service) GetNumOfEmployees(ctx context.Context, request *GetNumOfEmployeesRequest) (*GetNumOfEmployeesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfEmployees(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfEmployeesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfEmployees, OperationName: "GetNumOfEmployees", Variables: variables})
}

// GetNumOfEmployeesWithCrashes executes the observed GetNumOfEmployeesWithCrashes GraphQL operation.
func (s *Service) GetNumOfEmployeesWithCrashes(ctx context.Context, request *GetNumOfEmployeesWithCrashesRequest) (*GetNumOfEmployeesWithCrashesResponse, *interfaces.Response, error) {
	if err := validateGetNumOfEmployeesWithCrashes(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"id": request.ID, "timeFrameFilter": request.TimeFrameFilter}
	if request.DimensionsFilter != nil {
		variables["dimensionsFilter"] = request.DimensionsFilter
	}
	return graphql.ExecuteData[GetNumOfEmployeesWithCrashesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryGetNumOfEmployeesWithCrashes, OperationName: "GetNumOfEmployeesWithCrashes", Variables: variables})
}

// OverviewDesktopTooltips executes the observed OverviewDesktopTooltips GraphQL operation.
func (s *Service) OverviewDesktopTooltips(ctx context.Context, request *OverviewDesktopTooltipsRequest) (*OverviewDesktopTooltipsResponse, *interfaces.Response, error) {
	if err := validateOverviewDesktopTooltips(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"desktopTimeFrameFilter": request.DesktopTimeFrameFilter}

	return graphql.ExecuteData[OverviewDesktopTooltipsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryOverviewDesktopTooltips, OperationName: "OverviewDesktopTooltips", Variables: variables})
}

// TilesDesktopCrashesPerEmployee executes the observed TilesDesktopCrashesPerEmployee GraphQL operation.
func (s *Service) TilesDesktopCrashesPerEmployee(ctx context.Context, request *TilesDesktopCrashesPerEmployeeRequest) (*TilesDesktopCrashesPerEmployeeResponse, *interfaces.Response, error) {
	if err := validateTilesDesktopCrashesPerEmployee(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesDesktopCrashesPerEmployeeResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesDesktopCrashesPerEmployee, OperationName: "TilesDesktopCrashesPerEmployee", Variables: variables})
}

// TilesDesktopNumberOfEmployees executes the observed TilesDesktopNumberOfEmployees GraphQL operation.
func (s *Service) TilesDesktopNumberOfEmployees(ctx context.Context, request *TilesDesktopNumberOfEmployeesRequest) (*TilesDesktopNumberOfEmployeesResponse, *interfaces.Response, error) {
	if err := validateTilesDesktopNumberOfEmployees(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesDesktopNumberOfEmployeesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesDesktopNumberOfEmployees, OperationName: "TilesDesktopNumberOfEmployees", Variables: variables})
}

// TilesErrorCount executes the observed TilesErrorCount GraphQL operation.
func (s *Service) TilesErrorCount(ctx context.Context, request *TilesErrorCountRequest) (*TilesErrorCountResponse, *interfaces.Response, error) {
	if err := validateTilesErrorCount(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesErrorCountResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesErrorCount, OperationName: "TilesErrorCount", Variables: variables})
}

// TilesFrustratingPageLoads executes the observed TilesFrustratingPageLoads GraphQL operation.
func (s *Service) TilesFrustratingPageLoads(ctx context.Context, request *TilesFrustratingPageLoadsRequest) (*TilesFrustratingPageLoadsResponse, *interfaces.Response, error) {
	if err := validateTilesFrustratingPageLoads(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesFrustratingPageLoadsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesFrustratingPageLoads, OperationName: "TilesFrustratingPageLoads", Variables: variables})
}

// TilesNumberOfEmployees executes the observed TilesNumberOfEmployees GraphQL operation.
func (s *Service) TilesNumberOfEmployees(ctx context.Context, request *TilesNumberOfEmployeesRequest) (*TilesNumberOfEmployeesResponse, *interfaces.Response, error) {
	if err := validateTilesNumberOfEmployees(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesNumberOfEmployeesResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesNumberOfEmployees, OperationName: "TilesNumberOfEmployees", Variables: variables})
}

// TilesPageLoadTime executes the observed TilesPageLoadTime GraphQL operation.
func (s *Service) TilesPageLoadTime(ctx context.Context, request *TilesPageLoadTimeRequest) (*TilesPageLoadTimeResponse, *interfaces.Response, error) {
	if err := validateTilesPageLoadTime(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesPageLoadTimeResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesPageLoadTime, OperationName: "TilesPageLoadTime", Variables: variables})
}

// TilesTransactionTime executes the observed TilesTransactionTime GraphQL operation.
func (s *Service) TilesTransactionTime(ctx context.Context, request *TilesTransactionTimeRequest) (*TilesTransactionTimeResponse, *interfaces.Response, error) {
	if err := validateTilesTransactionTime(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesTransactionTimeResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesTransactionTime, OperationName: "TilesTransactionTime", Variables: variables})
}

// TilesUsageTime executes the observed TilesUsageTime GraphQL operation.
func (s *Service) TilesUsageTime(ctx context.Context, request *TilesUsageTimeRequest) (*TilesUsageTimeResponse, *interfaces.Response, error) {
	if err := validateTilesUsageTime(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[TilesUsageTimeResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryTilesUsageTime, OperationName: "TilesUsageTime", Variables: variables})
}

// WebOverviewTooltips executes the observed WebOverviewTooltips GraphQL operation.
func (s *Service) WebOverviewTooltips(ctx context.Context, request *WebOverviewTooltipsRequest) (*WebOverviewTooltipsResponse, *interfaces.Response, error) {
	if err := validateWebOverviewTooltips(request); err != nil {
		return nil, nil, err
	}
	variables := map[string]any{"timeFrameFilter": request.TimeFrameFilter}

	return graphql.ExecuteData[WebOverviewTooltipsResponse](ctx, s.graphql, operationID, graphql.GraphQLRequest{Query: queryWebOverviewTooltips, OperationName: "WebOverviewTooltips", Variables: variables})
}

type ApplicationExperienceServiceInterface interface {
	GetApplicationInsights(ctx context.Context, request *ApplicationInsightsRequest) (*ApplicationInsightsResponse, *interfaces.Response, error)
	GetApplicationsOverviewDesktopInvestigations(ctx context.Context, request *GetApplicationsOverviewDesktopInvestigationsRequest) (*GetApplicationsOverviewDesktopInvestigationsResponse, *interfaces.Response, error)
	GetAvgNetworkResponseTime(ctx context.Context, request *GetAvgNetworkResponseTimeRequest) (*GetAvgNetworkResponseTimeResponse, *interfaces.Response, error)
	GetBinarySuggestion(ctx context.Context, request *GetBinarySuggestionRequest) (*GetBinarySuggestionResponse, *interfaces.Response, error)
	GetDeviceCentricMetricBreakdown(ctx context.Context, request *GetDeviceCentricMetricBreakdownRequest) (*GetDeviceCentricMetricBreakdownResponse, *interfaces.Response, error)
	GetFailedConnectionsRatio(ctx context.Context, request *GetFailedConnectionsRatioRequest) (*GetFailedConnectionsRatioResponse, *interfaces.Response, error)
	GetInsights(ctx context.Context, request *GetInsightsRequest) (*GetInsightsResponse, *interfaces.Response, error)
	GetMetricBreakdown(ctx context.Context, request *GetMetricBreakdownRequest) (*GetMetricBreakdownResponse, *interfaces.Response, error)
	GetNumOfCrashesAndDevices(ctx context.Context, request *GetNumOfCrashesAndDevicesRequest) (*GetNumOfCrashesAndDevicesResponse, *interfaces.Response, error)
	GetNumOfCrashesAndEmployees(ctx context.Context, request *GetNumOfCrashesAndEmployeesRequest) (*GetNumOfCrashesAndEmployeesResponse, *interfaces.Response, error)
	GetNumOfDevices(ctx context.Context, request *GetNumOfDevicesRequest) (*GetNumOfDevicesResponse, *interfaces.Response, error)
	GetNumOfDevicesWithCrashes(ctx context.Context, request *GetNumOfDevicesWithCrashesRequest) (*GetNumOfDevicesWithCrashesResponse, *interfaces.Response, error)
	GetNumOfEmployees(ctx context.Context, request *GetNumOfEmployeesRequest) (*GetNumOfEmployeesResponse, *interfaces.Response, error)
	GetNumOfEmployeesWithCrashes(ctx context.Context, request *GetNumOfEmployeesWithCrashesRequest) (*GetNumOfEmployeesWithCrashesResponse, *interfaces.Response, error)
	OverviewDesktopTooltips(ctx context.Context, request *OverviewDesktopTooltipsRequest) (*OverviewDesktopTooltipsResponse, *interfaces.Response, error)
	TilesDesktopCrashesPerEmployee(ctx context.Context, request *TilesDesktopCrashesPerEmployeeRequest) (*TilesDesktopCrashesPerEmployeeResponse, *interfaces.Response, error)
	TilesDesktopNumberOfEmployees(ctx context.Context, request *TilesDesktopNumberOfEmployeesRequest) (*TilesDesktopNumberOfEmployeesResponse, *interfaces.Response, error)
	TilesErrorCount(ctx context.Context, request *TilesErrorCountRequest) (*TilesErrorCountResponse, *interfaces.Response, error)
	TilesFrustratingPageLoads(ctx context.Context, request *TilesFrustratingPageLoadsRequest) (*TilesFrustratingPageLoadsResponse, *interfaces.Response, error)
	TilesNumberOfEmployees(ctx context.Context, request *TilesNumberOfEmployeesRequest) (*TilesNumberOfEmployeesResponse, *interfaces.Response, error)
	TilesPageLoadTime(ctx context.Context, request *TilesPageLoadTimeRequest) (*TilesPageLoadTimeResponse, *interfaces.Response, error)
	TilesTransactionTime(ctx context.Context, request *TilesTransactionTimeRequest) (*TilesTransactionTimeResponse, *interfaces.Response, error)
	TilesUsageTime(ctx context.Context, request *TilesUsageTimeRequest) (*TilesUsageTimeResponse, *interfaces.Response, error)
	WebOverviewTooltips(ctx context.Context, request *WebOverviewTooltipsRequest) (*WebOverviewTooltipsResponse, *interfaces.Response, error)
}

var _ ApplicationExperienceServiceInterface = (*Service)(nil)

// GetApplicationInsights generates application narratives through the UI's insights proxy.
// A configured application ID is required. No device action is executed.
func (s *Service) GetApplicationInsights(ctx context.Context, request *ApplicationInsightsRequest) (*ApplicationInsightsResponse, *interfaces.Response, error) {
	if err := validateApplicationInsights(request); err != nil {
		return nil, nil, err
	}
	var result ApplicationInsightsResponse
	response, err := s.client.Post(ctx, EndpointApplicationInsights, request, map[string]string{"Content-Type": "application/json", "Accept": "application/json", "X-NX-MaxAge": "300", "X-NX-Originator": "appex-dashboard", "X-NX-Timeout": "20", "X-NX-Priority": "1"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}
