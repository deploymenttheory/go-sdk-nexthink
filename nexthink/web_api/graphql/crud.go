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
	// The browser gateway negotiates binary gzip correctly for */*. With an
	// application/json Accept header, some routes incorrectly send base64 gzip
	// text while retaining Content-Encoding: gzip. Keep JSON as the request type.
	headers := map[string]string{"Accept": "*/*", "Content-Type": "application/json"}
	for key, value := range req.Headers {
		headers[key] = value
	}
	var result GraphQLResponse
	resp, err := s.client.Post(
		ctx,
		endpoints[operationID],
		req,
		headers,
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
