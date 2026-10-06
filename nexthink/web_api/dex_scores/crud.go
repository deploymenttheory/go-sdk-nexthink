package dex_scores

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetCampaign executes the first-party UI operation getCampaign.
func (s *Service) GetCampaign(ctx context.Context, request *GetCampaignRequest) (*GetCampaignResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetCampaignResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetCampaign, OperationName: "getCampaign", Variables: variables})
}

// GetDeviceExperience executes the first-party UI operation getDeviceExperience.
func (s *Service) GetDeviceExperience(ctx context.Context, request *GetDeviceExperienceRequest) (*GetDeviceExperienceResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDeviceExperienceResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDeviceExperience, OperationName: "getDeviceExperience", Variables: variables})
}

// GetDimensionBreakdowns executes the first-party UI operation getDimensionBreakdowns.
func (s *Service) GetDimensionBreakdowns(ctx context.Context, request *GetDimensionBreakdownsRequest) (*GetDimensionBreakdownsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDimensionBreakdownsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDimensionBreakdowns, OperationName: "getDimensionBreakdowns", Variables: variables})
}

// GetDimensionsV2 executes the first-party UI operation getDimensionsV2.
func (s *Service) GetDimensionsV2(ctx context.Context, request *GetDimensionsV2Request) (*GetDimensionsV2Response, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDimensionsV2Response](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDimensionsV2, OperationName: "getDimensionsV2", Variables: variables})
}

// GetInvestigationUrl executes the first-party UI operation getInvestigationUrl.
func (s *Service) GetInvestigationUrl(ctx context.Context, request *GetInvestigationUrlRequest) (*GetInvestigationUrlResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetInvestigationUrlResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetInvestigationUrl, OperationName: "getInvestigationUrl", Variables: variables})
}

// GetLeaves executes the first-party UI operation getLeaves.
func (s *Service) GetLeaves(ctx context.Context, request *GetLeavesRequest) (*GetLeavesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetLeavesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetLeaves, OperationName: "getLeaves", Variables: variables})
}

// GetMetricThreshold executes the first-party UI operation getMetricThreshold.
func (s *Service) GetMetricThreshold(ctx context.Context, request *GetMetricThresholdRequest) (*GetMetricThresholdResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetMetricThresholdResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetMetricThreshold, OperationName: "getMetricThreshold", Variables: variables})
}

// GetScores executes the first-party UI operation getScores.
func (s *Service) GetScores(ctx context.Context, request *GetScoresRequest) (*GetScoresResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetScoresResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetScores, OperationName: "getScores", Variables: variables})
}

// GetTrend executes the first-party UI operation getTrend.
func (s *Service) GetTrend(ctx context.Context, request *GetTrendRequest) (*GetTrendResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrend, OperationName: "getTrend", Variables: variables})
}

// GetTrendDevices executes the first-party UI operation getTrendDevices.
func (s *Service) GetTrendDevices(ctx context.Context, request *GetTrendDevicesRequest) (*GetTrendDevicesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendDevicesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendDevices, OperationName: "getTrendDevices", Variables: variables})
}

// GetTrendEmployeesWithIssues executes the first-party UI operation getTrendEmployeesWithIssues.
func (s *Service) GetTrendEmployeesWithIssues(ctx context.Context, request *GetTrendEmployeesWithIssuesRequest) (*GetTrendEmployeesWithIssuesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendEmployeesWithIssuesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendEmployeesWithIssues, OperationName: "getTrendEmployeesWithIssues", Variables: variables})
}

// GetTrendImprovement executes the first-party UI operation getTrendImprovement.
func (s *Service) GetTrendImprovement(ctx context.Context, request *GetTrendImprovementRequest) (*GetTrendImprovementResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendImprovementResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendImprovement, OperationName: "getTrendImprovement", Variables: variables})
}

// GetTrendScore executes the first-party UI operation getTrendScore.
func (s *Service) GetTrendScore(ctx context.Context, request *GetTrendScoreRequest) (*GetTrendScoreResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendScoreResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendScore, OperationName: "getTrendScore", Variables: variables})
}

// GetTrendTimeLost executes the first-party UI operation getTrendTimeLost.
func (s *Service) GetTrendTimeLost(ctx context.Context, request *GetTrendTimeLostRequest) (*GetTrendTimeLostResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendTimeLostResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendTimeLost, OperationName: "getTrendTimeLost", Variables: variables})
}

// GetTrendWithRange executes the first-party UI operation getTrendWithRange.
func (s *Service) GetTrendWithRange(ctx context.Context, request *GetTrendWithRangeRequest) (*GetTrendWithRangeResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetTrendWithRangeResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetTrendWithRange, OperationName: "getTrendWithRange", Variables: variables})
}

// GetWhatsChanged executes the first-party UI operation getWhatsChanged.
func (s *Service) GetWhatsChanged(ctx context.Context, request *GetWhatsChangedRequest) (*GetWhatsChangedResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetWhatsChangedResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetWhatsChanged, OperationName: "getWhatsChanged", Variables: variables})
}

type DexScoresServiceInterface interface {
	GetCampaign(context.Context, *GetCampaignRequest) (*GetCampaignResponse, *interfaces.Response, error)
	GetDeviceExperience(context.Context, *GetDeviceExperienceRequest) (*GetDeviceExperienceResponse, *interfaces.Response, error)
	GetDimensionBreakdowns(context.Context, *GetDimensionBreakdownsRequest) (*GetDimensionBreakdownsResponse, *interfaces.Response, error)
	GetDimensionsV2(context.Context, *GetDimensionsV2Request) (*GetDimensionsV2Response, *interfaces.Response, error)
	GetInvestigationUrl(context.Context, *GetInvestigationUrlRequest) (*GetInvestigationUrlResponse, *interfaces.Response, error)
	GetLeaves(context.Context, *GetLeavesRequest) (*GetLeavesResponse, *interfaces.Response, error)
	GetMetricThreshold(context.Context, *GetMetricThresholdRequest) (*GetMetricThresholdResponse, *interfaces.Response, error)
	GetScores(context.Context, *GetScoresRequest) (*GetScoresResponse, *interfaces.Response, error)
	GetTrend(context.Context, *GetTrendRequest) (*GetTrendResponse, *interfaces.Response, error)
	GetTrendDevices(context.Context, *GetTrendDevicesRequest) (*GetTrendDevicesResponse, *interfaces.Response, error)
	GetTrendEmployeesWithIssues(context.Context, *GetTrendEmployeesWithIssuesRequest) (*GetTrendEmployeesWithIssuesResponse, *interfaces.Response, error)
	GetTrendImprovement(context.Context, *GetTrendImprovementRequest) (*GetTrendImprovementResponse, *interfaces.Response, error)
	GetTrendScore(context.Context, *GetTrendScoreRequest) (*GetTrendScoreResponse, *interfaces.Response, error)
	GetTrendTimeLost(context.Context, *GetTrendTimeLostRequest) (*GetTrendTimeLostResponse, *interfaces.Response, error)
	GetTrendWithRange(context.Context, *GetTrendWithRangeRequest) (*GetTrendWithRangeResponse, *interfaces.Response, error)
	GetWhatsChanged(context.Context, *GetWhatsChangedRequest) (*GetWhatsChangedResponse, *interfaces.Response, error)
}

var _ DexScoresServiceInterface = (*Service)(nil)
