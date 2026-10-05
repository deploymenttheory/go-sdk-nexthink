package graphql

import (
	"fmt"
	"strings"
)

func ValidateRequest(operationID string, req GraphQLRequest) error {
	if _, ok := endpoints[operationID]; !ok || strings.TrimSpace(req.Query) == "" {
		return fmt.Errorf("a known GraphQL operation and nonempty query are required")
	}
	return nil
}
