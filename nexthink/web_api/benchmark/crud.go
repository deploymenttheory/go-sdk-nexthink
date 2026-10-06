package benchmark

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// GetBinaryProductMaps executes the first-party UI operation getBinaryProductMaps.
func (s *Service) GetBinaryProductMaps(ctx context.Context, request *GetBinaryProductMapsRequest) (*GetBinaryProductMapsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetBinaryProductMapsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetBinaryProductMaps, OperationName: "getBinaryProductMaps", Variables: variables})
}

// GetProfile executes the first-party UI operation getProfile.
func (s *Service) GetProfile(ctx context.Context, request *GetProfileRequest) (*GetProfileResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetProfileResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetProfile, OperationName: "getProfile", Variables: variables})
}

// GetProfileSearchItems executes the first-party UI operation getProfileSearchItems.
func (s *Service) GetProfileSearchItems(ctx context.Context, request *GetProfileSearchItemsRequest) (*GetProfileSearchItemsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetProfileSearchItemsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetProfileSearchItems, OperationName: "getProfileSearchItems", Variables: variables})
}

// GetProfileVersions executes the first-party UI operation getProfileVersions.
func (s *Service) GetProfileVersions(ctx context.Context, request *GetProfileVersionsRequest) (*GetProfileVersionsResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[GetProfileVersionsResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryGetProfileVersions, OperationName: "getProfileVersions", Variables: variables})
}

// LookupBenchmark executes the first-party UI operation lookupBenchmark.
func (s *Service) LookupBenchmark(ctx context.Context, request *LookupBenchmarkRequest) (*LookupBenchmarkResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[LookupBenchmarkResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryLookupBenchmark, OperationName: "lookupBenchmark", Variables: variables})
}

// ProductProperties executes the first-party UI operation productProperties.
func (s *Service) ProductProperties(ctx context.Context, request *ProductPropertiesRequest) (*ProductPropertiesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ProductPropertiesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryProductProperties, OperationName: "productProperties", Variables: variables})
}

// ProfileProperties executes the first-party UI operation profileProperties.
func (s *Service) ProfileProperties(ctx context.Context, request *ProfilePropertiesRequest) (*ProfilePropertiesResponse, *interfaces.Response, error) {
	variables, err := requestVariables(request)
	if err != nil {
		return nil, nil, err
	}
	return graphql.ExecuteData[ProfilePropertiesResponse](ctx, graphql.NewService(s.client), operationID, graphql.GraphQLRequest{Query: queryProfileProperties, OperationName: "profileProperties", Variables: variables})
}

type BenchmarkServiceInterface interface {
	GetBinaryProductMaps(context.Context, *GetBinaryProductMapsRequest) (*GetBinaryProductMapsResponse, *interfaces.Response, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, *interfaces.Response, error)
	GetProfileSearchItems(context.Context, *GetProfileSearchItemsRequest) (*GetProfileSearchItemsResponse, *interfaces.Response, error)
	GetProfileVersions(context.Context, *GetProfileVersionsRequest) (*GetProfileVersionsResponse, *interfaces.Response, error)
	LookupBenchmark(context.Context, *LookupBenchmarkRequest) (*LookupBenchmarkResponse, *interfaces.Response, error)
	ProductProperties(context.Context, *ProductPropertiesRequest) (*ProductPropertiesResponse, *interfaces.Response, error)
	ProfileProperties(context.Context, *ProfilePropertiesRequest) (*ProfilePropertiesResponse, *interfaces.Response, error)
}

var _ BenchmarkServiceInterface = (*Service)(nil)
