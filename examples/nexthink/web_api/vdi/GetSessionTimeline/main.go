package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/vdi"
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
	sessionID := os.Getenv("NEXTHINK_SESSION_ID")
	data, err := os.ReadFile(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	var request vdi.TimelineRequest
	if err = json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.VDI.GetSessionTimeline(ctx, sessionID, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
