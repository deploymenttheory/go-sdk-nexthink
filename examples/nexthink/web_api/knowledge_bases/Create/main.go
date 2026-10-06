// Use browser authentication. See ../README.md for inputs and side effects.
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/knowledge_bases"
	"go.uber.org/zap"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	var request knowledge_bases.CreateRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.WebAPI.KnowledgeBases.Create(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	result := map[string]int{"status": response.StatusCode}
	if err := json.NewEncoder(os.Stdout).Encode(result); err != nil {
		log.Fatal(err)
	}
}
