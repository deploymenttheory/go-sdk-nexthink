// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
// NEXTHINK_REQUEST_FILE must contain the API request JSON.
// The request file supplies query, variables, and operationName.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"go.uber.org/zap"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	var request graphql.GraphQLRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	operation := os.Getenv("NEXTHINK_GRAPHQL_OPERATION")
	if operation == "" {
		log.Fatal("NEXTHINK_GRAPHQL_OPERATION is required (for example graphql.workflows)")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.GraphQL.Execute(ctx, operation, request)
	if result != nil {
		fmt.Printf("GraphQL errors=%d; data bytes=%d\n", len(result.Errors), len(result.Data))
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
