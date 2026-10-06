// See README.md for inputs and tenant-wide side effects.
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
	key := os.Getenv("NEXTHINK_CONFIGURATION_KEY")
	if key == "" {
		log.Fatal("set NEXTHINK_CONFIGURATION_KEY")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.ProductConfiguration.GetInstance(ctx, key)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
