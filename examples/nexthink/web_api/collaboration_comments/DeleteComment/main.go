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
	documentID := os.Getenv("NEXTHINK_DOCUMENT_ID")
	if documentID == "" {
		log.Fatal("NEXTHINK_DOCUMENT_ID is required")
	}
	commentID := os.Getenv("NEXTHINK_COMMENT_ID")
	if commentID == "" {
		log.Fatal("NEXTHINK_COMMENT_ID is required")
	}
	response, err := c.WebAPI.CollaborationComments.DeleteComment(ctx, documentID, commentID)
	if response != nil {
		if e := json.NewEncoder(os.Stdout).Encode(map[string]int{"statusCode": response.StatusCode}); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
