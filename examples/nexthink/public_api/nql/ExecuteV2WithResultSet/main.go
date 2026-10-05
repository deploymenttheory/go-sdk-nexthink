// Configure NEXTHINK_API=public, instance, region and client credentials.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
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
	var request nql.ExecuteRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.PublicAPI.NQL.ExecuteV2WithResultSet(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Rows: %d\n", result.Rows())
}
