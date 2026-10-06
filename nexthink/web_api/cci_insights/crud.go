package cci_insights

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetBinaryInsights executes the first-party UI operation getBinaryInsights.
func (s *Service) GetBinaryInsights(ctx context.Context, request *GetBinaryInsightsRequest) (*GetBinaryInsightsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetBinaryInsightsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetBinaryInsights, OperationName: "getBinaryInsights", Variables: variables})
}

// GetDiagnosticBinaryInsights executes the first-party UI operation getBinaryInsights.
func (s *Service) GetDiagnosticBinaryInsights(ctx context.Context, request *GetDiagnosticBinaryInsightsRequest) (*GetDiagnosticBinaryInsightsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDiagnosticBinaryInsightsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDiagnosticBinaryInsights, OperationName: "getBinaryInsights", Variables: variables})
}

// GetDataset executes the first-party UI operation getDataset.
func (s *Service) GetDataset(ctx context.Context, request *GetDatasetRequest) (*GetDatasetResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetDatasetResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetDataset, OperationName: "getDataset", Variables: variables})
}

type CCIInsightsServiceInterface interface {
	GetBinaryInsights(context.Context, *GetBinaryInsightsRequest) (*GetBinaryInsightsResponse, *interfaces.Response, error)
	GetDiagnosticBinaryInsights(context.Context, *GetDiagnosticBinaryInsightsRequest) (*GetDiagnosticBinaryInsightsResponse, *interfaces.Response, error)
	GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, *interfaces.Response, error)
}

var _ CCIInsightsServiceInterface = (*Service)(nil)
