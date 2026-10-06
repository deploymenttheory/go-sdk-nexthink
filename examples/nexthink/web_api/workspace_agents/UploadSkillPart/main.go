// See README.md for authentication, inputs and side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace_agents"
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
	defer c.Close()
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	id := os.Getenv("NEXTHINK_RESOURCE_ID")
	if id == "" {
		log.Fatal("set NEXTHINK_RESOURCE_ID")
	}
	input := os.Getenv("NEXTHINK_REQUEST_FILE")
	if input == "" {
		log.Fatal("set NEXTHINK_REQUEST_FILE to reviewed request JSON")
	}
	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	var request workspace_agents.UploadPartRequest
	if err := json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.WorkspaceAgents.UploadSkillPart(ctx, id, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
