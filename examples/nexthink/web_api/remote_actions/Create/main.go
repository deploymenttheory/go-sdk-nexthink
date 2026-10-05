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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/remote_actions"
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
	var request remote_actions.RemoteActionInput
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.RemoteActions.Create(ctx, &request)
	// Preserve a created identifier even if GraphQL also reports an error.
	if result != nil && result.RemoteAction != nil {
		fmt.Printf("Created content ID: %s\n", result.RemoteAction.ContentID)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
