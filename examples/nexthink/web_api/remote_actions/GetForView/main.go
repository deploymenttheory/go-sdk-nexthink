// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
// Set NEXTHINK_REMOTE_ACTION_UUID to the intended resource identifier.
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
	argument := os.Getenv("NEXTHINK_REMOTE_ACTION_UUID")
	if argument == "" {
		log.Fatal("NEXTHINK_REMOTE_ACTION_UUID is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.RemoteActions.GetForView(ctx, argument)
	if err != nil {
		log.Fatal(err)
	}
	if result.RemoteAction == nil {
		log.Fatal("remote action not returned")
	}
	fmt.Printf("Action: %s\n", result.RemoteAction.ID)
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
