package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	if c.WebAPI == nil {
		panic("web API client is required")
	}
	id := os.Getenv("NEXTHINK_CONTENT_ID")
	if id == "" {
		panic("NEXTHINK_CONTENT_ID is required")
	}
	revision, err := strconv.Atoi(os.Getenv("NEXTHINK_REVISION"))
	if err != nil {
		panic("set NEXTHINK_REVISION to the latest _rev")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.WebAPI.AITools.DeleteGoal(ctx, id, revision)
	if err != nil {
		panic(err)
	}
	result := map[string]int{"status": response.StatusCode}
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
