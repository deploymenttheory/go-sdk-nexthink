package main

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/nexthink/web_api/mobile_tokens/internal/tokenoutput"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/mobile_tokens"
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
		log.Fatal("set NEXTHINK_API=web or both")
	}
	var request mobile_tokens.CreateRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.MobileTokens.Create(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err := tokenoutput.Write(result, os.Getenv("NEXTHINK_OUTPUT_FILE"), os.Stdout); err != nil {
		log.Fatal(err)
	}
}
