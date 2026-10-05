package experimental

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
)

type GraphQLRequest struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables,omitempty"`
	OperationName string         `json:"operationName,omitempty"`
}
type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations,omitempty"`
	Path       []any          `json:"path,omitempty"`
	Extensions map[string]any `json:"extensions,omitempty"`
}
type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []GraphQLError  `json:"errors,omitempty"`
}

type GraphQLErrors []GraphQLError

func (e GraphQLErrors) Error() string {
	return fmt.Sprintf("GraphQL returned %d error(s); inspect response.Errors", len(e))
}

// GraphQL retains partial data and reports GraphQL errors even on HTTP 200.
// The operation ID selects a catalog endpoint; Query supplies its query or mutation.
func (c *Client) GraphQL(ctx context.Context, operationID string, req GraphQLRequest) (*GraphQLResponse, *interfaces.Response, error) {
	if !strings.HasPrefix(operationID, "graphql.") || strings.TrimSpace(req.Query) == "" {
		return nil, nil, fmt.Errorf("a GraphQL catalog operation and nonempty query are required")
	}
	result, resp, err := decode[GraphQLResponse](c.Do(ctx, operationID, Request{Body: req}))
	if err != nil {
		return nil, resp, err
	}
	if len(result.Errors) > 0 {
		return result, resp, GraphQLErrors(result.Errors)
	}
	return result, resp, nil
}
