// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
package main

import (
	"context"
	"fmt"
	"log"
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
	result, response, err := c.WebAPI.CollectorManagement.GetGroups(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Values: %d\n", len(result))
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
