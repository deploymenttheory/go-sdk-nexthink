package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/amplify"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	path := os.Getenv("NEXTHINK_REQUEST_FILE")
	if path == "" {
		log.Fatal("set NEXTHINK_REQUEST_FILE to a JSON request file")
	}
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var request amplify.ConfigurationRequest
	if err := json.Unmarshal(body, &request); err != nil {
		log.Fatal(err)
	}
	id := os.Getenv("NEXTHINK_CONTENT_ID")
	if id == "" {
		log.Fatal("set NEXTHINK_CONTENT_ID to the configuration id")
	}
	revision, err := strconv.Atoi(os.Getenv("NEXTHINK_REVISION_NUMBER"))
	if err != nil {
		log.Fatal("set NEXTHINK_REVISION_NUMBER from the latest GetConfiguration response")
	}
	result, _, err := c.WebAPI.Amplify.UpdateConfiguration(ctx, id, revision, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
