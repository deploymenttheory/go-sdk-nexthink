// Set NEXTHINK_API=web and NEXTHINK_WEB_AUTH=chrome or token.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id := os.Getenv("NEXTHINK_WORKFLOW_UUID")
	if id == "" {
		log.Fatal("NEXTHINK_WORKFLOW_UUID is required")
	}
	result, resp, err := c.WebAPI.Workflows.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if result.Workflow == nil {
		log.Fatal("workflow was not returned")
	}
	fmt.Printf("HTTP %d; action=%s\n", resp.StatusCode, result.Workflow.ID)
}
