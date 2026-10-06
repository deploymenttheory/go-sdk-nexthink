// Set NEXTHINK_API=web and NEXTHINK_REQUEST_FILE to an input JSON file.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dex_scores"
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
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var request dex_scores.GetScoresRequest
	if err = json.Unmarshal(b, &request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.DexScores.GetScores(ctx, &request)
	if result != nil {
		if encodeErr := json.NewEncoder(os.Stdout).Encode(result); encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
