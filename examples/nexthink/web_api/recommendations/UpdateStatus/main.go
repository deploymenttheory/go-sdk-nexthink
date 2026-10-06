package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/recommendations"
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
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	id := os.Getenv("NEXTHINK_ID")
	if id == "" {
		log.Fatal("NEXTHINK_ID is required")
	}
	var request recommendations.UpdateStatusRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	result, _, err := c.WebAPI.Recommendations.UpdateStatus(ctx, id, &request)
	if result != nil {
		if e := json.NewEncoder(os.Stdout).Encode(result); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
