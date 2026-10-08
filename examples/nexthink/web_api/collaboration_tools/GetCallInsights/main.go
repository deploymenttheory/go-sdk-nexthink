package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collaboration_tools"
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
	deviceID := os.Getenv("NEXTHINK_COLLECTOR_ID")
	if deviceID == "" {
		deviceID = os.Getenv("NEXTHINK_DEVICE_ID")
	}
	if deviceID == "" {
		log.Fatal("set NEXTHINK_COLLECTOR_ID to the Collector UID used in Device View")
	}
	data, err := os.ReadFile(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	var request collaboration_tools.CallInsightsRequest
	if err = json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.CollaborationTools.GetCallInsights(ctx, deviceID, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
