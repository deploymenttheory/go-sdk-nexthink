package cci_benchmarks

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Query calls the browser /apigateway/cci/benchmarks/query contract.
func (s *Service) Query(ctx context.Context, request *QueryRequest) (*QueryResponse, *interfaces.Response, error) {
	if err := validateQuery(request); err != nil {
		return nil, nil, err
	}
	zone := request.TimeZone
	if zone == "" {
		zone = "UTC"
	}
	var result QueryResponse
	response, err := s.client.Post(ctx, EndpointQuery, request, map[string]string{"Accept": "application/json", "Content-Type": "application/json", "x-client-timezone": zone}, &result)
	if err != nil {
		return nil, response, err
	}
	return &result, response, nil
}

type CCIBenchmarksServiceInterface interface {
	Query(ctx context.Context, request *QueryRequest) (*QueryResponse, *interfaces.Response, error)
}

var _ CCIBenchmarksServiceInterface = (*Service)(nil)
