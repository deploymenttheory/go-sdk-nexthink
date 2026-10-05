// Set NEXTHINK_API=web and NEXTHINK_WEB_AUTH=chrome or token.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, resp, err := c.WebAPI.ContentAdministration.List(
		ctx,
		"fe-nqlapi-nx-content-administration-config",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d; saved queries=%d\n", resp.StatusCode, len(result.Rows))
}
