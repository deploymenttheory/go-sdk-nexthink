package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/access_management"
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
	var request access_management.GrantRoleContentPermissionsRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	service := c.WebAPI.AccessManagement
	if os.Getenv("NEXTHINK_ACCESS_LEGACY") == "true" {
		service = c.WebAPI.LegacyAccessManagement
	}
	result, _, err := service.GrantRoleContentPermissions(ctx, &request)
	if result != nil {
		if e := json.NewEncoder(os.Stdout).Encode(result); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
