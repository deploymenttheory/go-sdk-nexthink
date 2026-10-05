// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
// Set NEXTHINK_WORKFLOW_UUID to the intended resource identifier.
// The opaque export is written to stdout; redirect it to a file to retain it.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
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
	argument := os.Getenv("NEXTHINK_WORKFLOW_UUID")
	if argument == "" {
		log.Fatal("NEXTHINK_WORKFLOW_UUID is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.Workflows.Export(ctx, argument)
	if err != nil {
		log.Fatal(err)
	}
	if result.Content == nil {
		log.Fatal("export not returned")
	}
	fmt.Print(*result.Content)
	fmt.Fprintf(os.Stderr, "Export bytes: %d; HTTP %d\n", len(*result.Content), response.StatusCode)
}
