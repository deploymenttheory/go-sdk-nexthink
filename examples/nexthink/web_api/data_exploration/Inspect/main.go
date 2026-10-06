// Configure browser authentication and NEXTHINK_API=web. See ../README.md.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/data_exploration"
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
	path := os.Getenv("NEXTHINK_REQUEST_FILE")
	if path == "" {
		log.Fatal("set NEXTHINK_REQUEST_FILE; see request.example.json")
	}
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var request data_exploration.QueryInput
	if err = json.Unmarshal(body, &request); err != nil {
		log.Fatal(err)
	}
	var options *data_exploration.TimeContext
	if path := os.Getenv("NEXTHINK_TIME_CONTEXT_FILE"); path != "" {
		body, readErr := os.ReadFile(path)
		if readErr != nil {
			log.Fatal(readErr)
		}
		if err = json.Unmarshal(body, &options); err != nil {
			log.Fatal(err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.DataExploration.Inspect(ctx, &request, options)
	if result != nil {
		if encodeErr := json.NewEncoder(os.Stdout).Encode(result); encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
