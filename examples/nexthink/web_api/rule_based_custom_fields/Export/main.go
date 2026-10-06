// Browser authentication and resource-specific inputs are documented in ../README.md.
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
	if c.WebAPI == nil {
		log.Fatal("NEXTHINK_API=web or both is required")
	}
	id := os.Getenv("NEXTHINK_CONTENT_ID")
	if id == "" {
		log.Fatal("NEXTHINK_CONTENT_ID required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.RuleBasedCustomFields.Export(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
