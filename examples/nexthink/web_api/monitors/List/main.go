// Configure NEXTHINK_API=web and browser authentication. See ../README.md for inputs.
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.Monitors.List(ctx)
	// Print partial GraphQL data before reporting an error so created IDs remain available.
	if result != nil {
		if encodeErr := json.NewEncoder(os.Stdout).Encode(result); encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
