package graphql

import (
	"encoding/json"
	"fmt"
)

type GraphQLRequest struct {
	// Headers carries per-request UI context and is never serialized into the GraphQL body.
	Headers       map[string]string `json:"-"`
	Query         string            `json:"query"`
	Variables     map[string]any    `json:"variables,omitempty"`
	OperationName string            `json:"operationName,omitempty"`
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
