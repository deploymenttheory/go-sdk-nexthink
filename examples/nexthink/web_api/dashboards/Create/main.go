// Configure NEXTHINK_API=web and browser authentication. See ../README.md for inputs.
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dashboards"
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
	var request dashboards.DashboardInput
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.Dashboards.Create(ctx, &request)
	// Retain partial GraphQL results so IDs remain available for cleanup.
	if result != nil {
		if encodeErr := json.NewEncoder(os.Stdout).Encode(result); encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
