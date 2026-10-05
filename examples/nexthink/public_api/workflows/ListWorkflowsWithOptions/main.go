// Configure NEXTHINK_API=public, instance, region and client credentials.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/workflows"
	"go.uber.org/zap"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.PublicAPI == nil {
		log.Fatal("set NEXTHINK_API=public or both")
	}
	active := true
	options := workflows.ListOptions{Dependency: os.Getenv("NEXTHINK_WORKFLOW_DEPENDENCY"), TriggerMethod: os.Getenv("NEXTHINK_TRIGGER_METHOD"), FetchOnlyActiveWorkflows: &active}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.PublicAPI.Workflows.ListWorkflowsWithOptions(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Workflows: %d\n", len(result))
}
