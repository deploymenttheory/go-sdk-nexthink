package main

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/appearance"
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
	name := appearance.AssetName(os.Getenv("NEXTHINK_ASSET_NAME"))
	if name == "" {
		name = appearance.MenuLogo
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	file := os.Getenv("NEXTHINK_OUTPUT_FILE")
	if file == "" {
		log.Fatal("NEXTHINK_OUTPUT_FILE is required")
	}
	data, _, err := c.WebAPI.Appearance.Get(ctx, name)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(file, data, 0600); err != nil {
		log.Fatal(err)
	}
}
