// Configure NEXTHINK_API=public, instance, region and client credentials.
// The builder only validates the local structure; execution uses the saved NEXTHINK_QUERY_ID.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/nql"
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
	queryID := os.Getenv("NEXTHINK_QUERY_ID")
	builder := nql.NewQueryBuilder().FromDevices().List("device.name")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.PublicAPI.NQL.ExecuteQueryBuilder(ctx, queryID, builder)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rows: %d\n", result.Rows())
}
