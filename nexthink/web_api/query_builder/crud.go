package query_builder

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Transform calls the browser /apigateway/query-builder/api/v1/transform contract.
func (s *Service) Transform(ctx context.Context, request *TransformRequest) (*QueryResponse, *interfaces.Response, error) {
	if err := validateTransform(request); err != nil {
		return nil, nil, err
	}
	var result QueryResponse
	response, err := s.client.Post(ctx, EndpointTransform, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// ListDrilldownDestinations calls the browser /apigateway/query-builder/api/v1/drill-down/destinations contract.
func (s *Service) ListDrilldownDestinations(ctx context.Context, request *DestinationsRequest) (*DestinationsResponse, *interfaces.Response, error) {
	if err := validateListDrilldownDestinations(request); err != nil {
		return nil, nil, err
	}
	var result DestinationsResponse
	response, err := s.client.Post(ctx, EndpointListDrilldownDestinations, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

// TransformDrilldown calls the browser /apigateway/query-builder/api/v1/drill-down/transform contract.
func (s *Service) TransformDrilldown(ctx context.Context, request *DrilldownRequest) (*QueryResponse, *interfaces.Response, error) {
	if err := validateTransformDrilldown(request); err != nil {
		return nil, nil, err
	}
	var result QueryResponse
	response, err := s.client.Post(ctx, EndpointTransformDrilldown, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json"}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type QueryBuilderServiceInterface interface {
	Transform(ctx context.Context, request *TransformRequest) (*QueryResponse, *interfaces.Response, error)
	ListDrilldownDestinations(ctx context.Context, request *DestinationsRequest) (*DestinationsResponse, *interfaces.Response, error)
	TransformDrilldown(ctx context.Context, request *DrilldownRequest) (*QueryResponse, *interfaces.Response, error)
}

var _ QueryBuilderServiceInterface = (*Service)(nil)
