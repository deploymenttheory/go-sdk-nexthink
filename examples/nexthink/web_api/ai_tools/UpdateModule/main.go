package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ai_tools"
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
	var request ai_tools.Module
	data, err := os.ReadFile(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(data, &request); err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.AITools.UpdateModule(ctx, id, revision, &request)
	if err != nil {
		panic(err)
	}
	output, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
