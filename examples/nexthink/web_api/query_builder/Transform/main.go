package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/query_builder"
	"go.uber.org/zap"
	"os"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		panic(err)
	}
	var request query_builder.TransformRequest
	b, err := os.ReadFile(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, &request); err != nil {
		panic(err)
	}
	if c.WebAPI == nil {
		panic("set NEXTHINK_API=web")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.QueryBuilder.Transform(ctx, &request)
	if err != nil {
		panic(err)
	}
	b, err = json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
