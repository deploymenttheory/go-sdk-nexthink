package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflow_executions"
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
	if os.Getenv("NEXTHINK_ALLOW_EXECUTION") != "true" {
		log.Fatal("set NEXTHINK_ALLOW_EXECUTION=true only for approved execution targets")
	}
	data, err := os.ReadFile(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	var request workflow_executions.ExecuteRequest
	if err = json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.WorkflowExecutions.Execute(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
