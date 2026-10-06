// See README.md for authentication, inputs and side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	id := os.Getenv("NEXTHINK_RESOURCE_ID")
	if id == "" {
		log.Fatal("set NEXTHINK_RESOURCE_ID")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.Workspace.CreateConversationShare(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
