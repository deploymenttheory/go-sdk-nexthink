// See README.md for authentication and inputs.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/global_search"
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	session := &global_search.PortalSession{Cookie: os.Getenv("NEXTHINK_PORTAL_COOKIE"), XAuthToken: os.Getenv("NEXTHINK_PORTAL_X_AUTH_TOKEN")}
	result, _, err := c.WebAPI.GlobalSearch.GetLegacyAuthToken(ctx, session)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(map[string]bool{"tokenReceived": result != nil && result.Result != nil && result.Result.Token != ""}); err != nil {
		log.Fatal(err)
	}
}
