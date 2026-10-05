// Set NEXTHINK_API=web, web authentication, and NEXTHINK_REQUEST_FILE.
// This example changes the specific management object described in that file.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
)

func main() {
	c, err := nexthink.NewClientFromEnv(
		nexthink.WithLogger(zap.NewNop()),
		nexthink.WithRetryCount(0),
	)
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	var request struct {
		UUID string `json:"uuid"`
	}
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.Workflows.Delete(ctx, request.UUID)
	if err != nil {
		log.Fatal(err)
	}
	if result == nil || !result.Deleted {
		log.Fatal("server did not confirm deletion")
	}
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
