package graphql

import (
	"context"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type GraphQLServiceInterface interface {
	Execute(context.Context, string, GraphQLRequest) (*GraphQLResponse, *interfaces.Response, error)
}

var _ GraphQLServiceInterface = (*Service)(nil)

type Service struct{ client interfaces.HTTPClient }

func NewService(c interfaces.HTTPClient) *Service { return &Service{client: c} }

// Execute retains partial data and reports GraphQL errors even on HTTP 200.
func (s *Service) Execute(
	ctx context.Context,
	operationID string,
	req GraphQLRequest,
) (*GraphQLResponse, *interfaces.Response, error) {
	if err := ValidateRequest(operationID, req); err != nil {
		return nil, nil, err
	}
	var result GraphQLResponse
	resp, err := s.client.Post(
		ctx,
		endpoints[operationID],
		req,
		map[string]string{"Accept": "application/json", "Content-Type": "application/json"},
		&result,
	)
	if err != nil {
		return nil, resp, err
	}
	if len(result.Errors) > 0 {
		return &result, resp, GraphQLErrors(result.Errors)
	}
	return &result, resp, nil
}
