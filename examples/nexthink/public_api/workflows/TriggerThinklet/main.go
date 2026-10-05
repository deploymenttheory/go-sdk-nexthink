// Configure NEXTHINK_API=public, instance, region and client credentials.
// This resumes the supplied workflow execution; use a dedicated lab execution.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
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
	var request workflows.ThinkletTriggerRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	workflowID := os.Getenv("NEXTHINK_WORKFLOW_UUID")
	executionID := os.Getenv("NEXTHINK_EXECUTION_UUID")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.PublicAPI.Workflows.TriggerThinklet(ctx, workflowID, executionID, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request ID: %s\n", result.RequestUUID)
}
