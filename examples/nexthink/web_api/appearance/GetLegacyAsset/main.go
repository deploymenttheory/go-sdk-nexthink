// See README.md for credentials, inputs and side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
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
		log.Fatal("set NEXTHINK_API=web")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	session := &auth.PortalSession{Cookie: os.Getenv("NEXTHINK_PORTAL_COOKIE"), XAuthToken: os.Getenv("NEXTHINK_PORTAL_X_AUTH_TOKEN")}
	name := os.Getenv("NEXTHINK_ASSET_NAME")
	if name == "" {
		name = "menu-logo"
	}
	result, _, err := c.WebAPI.Appearance.GetLegacyAsset(ctx, session, appearance.AssetName(name))
	if result != nil {
		if e := json.NewEncoder(os.Stdout).Encode(result); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
