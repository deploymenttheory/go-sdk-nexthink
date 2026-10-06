package main

import (
	"context"
	"encoding/json"
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
	file := os.Getenv("NEXTHINK_IMAGE_FILE")
	if file == "" {
		log.Fatal("NEXTHINK_IMAGE_FILE is required")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	result, _, err := c.WebAPI.Appearance.Update(ctx, name, os.Getenv("NEXTHINK_IMAGE_CONTENT_TYPE"), data)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
