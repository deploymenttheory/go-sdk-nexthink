package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
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
	workflowID := os.Getenv("NEXTHINK_WORKFLOW_ID")
	executionID := os.Getenv("NEXTHINK_EXECUTION_ID")
	thinkletID := os.Getenv("NEXTHINK_THINKLET_ID")

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.WorkflowExecutions.GetRemoteActionDetails(ctx, workflowID, executionID, thinkletID)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
