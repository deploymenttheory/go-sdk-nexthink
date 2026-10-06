// See README.md for inputs and tenant-wide side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/user_classification"
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
	path := os.Getenv("NEXTHINK_REQUEST_FILE")
	if path == "" {
		log.Fatal("set NEXTHINK_REQUEST_FILE")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var request user_classification.ReplaceRequest
	if err := json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.WebAPI.UserClassification.Replace(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(map[string]int{"status": response.StatusCode}); err != nil {
		log.Fatal(err)
	}
}
