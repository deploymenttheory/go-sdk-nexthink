// Use browser authentication. See ../README.md for inputs and side effects.
package main

import (
	"context"
	"encoding/json"
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
	id := os.Getenv("NEXTHINK_CONTENT_ID")
	if id == "" {
		log.Fatal("NEXTHINK_CONTENT_ID is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.WebAPI.Webhooks.Delete(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	result := map[string]int{"status": response.StatusCode}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
