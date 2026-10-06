// See README.md for browser authentication, inputs and side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/autopilot"
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
	var request autopilot.Approval
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.Autopilot.CreateApproval(ctx, &request)
	if result != nil {
		if e := json.NewEncoder(os.Stdout).Encode(result); e != nil {
			log.Fatal(e)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
