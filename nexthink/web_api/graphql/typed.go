package graphql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

// ExecuteData decodes the data envelope for resource-specific GraphQL methods.
// Partial data is returned alongside GraphQLErrors; callers must inspect both.
// Missing data is an error unless the server already supplied GraphQL errors.
func ExecuteData[T any](
	ctx context.Context,
	service *Service,
	operationID string,
	request GraphQLRequest,
) (*T, *interfaces.Response, error) {
	result, response, requestErr := service.Execute(ctx, operationID, request)
	if result == nil {
		return nil, response, requestErr
	}
	var data *T
	if len(result.Data) > 0 {
		if err := json.Unmarshal(result.Data, &data); err != nil {
			return nil, response, errors.Join(
				requestErr,
				fmt.Errorf("decode GraphQL data: %w", err),
			)
		}
	}
	if data == nil && requestErr == nil {
		return nil, response, fmt.Errorf("GraphQL response has no data or errors")
	}
	return data, response, requestErr
}
