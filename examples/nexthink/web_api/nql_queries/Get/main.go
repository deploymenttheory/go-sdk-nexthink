// Set NEXTHINK_API=web and NEXTHINK_WEB_AUTH=chrome or token.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
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
	id := os.Getenv("NEXTHINK_QUERY_CONTENT_ID")
	if id == "" {
		log.Fatal("NEXTHINK_QUERY_CONTENT_ID is required")
	}
	result, resp, err := c.WebAPI.NQLQueries.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d; query=%s\n", resp.StatusCode, result.NQLAPIID)
}
