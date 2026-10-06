package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
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
	search := os.Getenv("NEXTHINK_SEARCH")
	result, _, err := c.WebAPI.CollaborationComments.ListUserMentions(ctx, search)
	if result != nil {
		if e := json.NewEncoder(os.Stdout).Encode(result); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
