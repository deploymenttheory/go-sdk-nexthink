package network_insights

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetInsights executes the first-party UI operation GetInsights.
func (s *Service) GetInsights(ctx context.Context, request *GetInsightsRequest) (*GetInsightsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetInsightsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetInsights, OperationName: "GetInsights", Variables: variables})
}

type NetworkInsightsServiceInterface interface {
	GetInsights(context.Context, *GetInsightsRequest) (*GetInsightsResponse, *interfaces.Response, error)
}

var _ NetworkInsightsServiceInterface = (*Service)(nil)
