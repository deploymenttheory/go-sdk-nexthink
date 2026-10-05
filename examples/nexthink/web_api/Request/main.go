// Set NEXTHINK_API=web and NEXTHINK_WEB_AUTH=chrome (or token).
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	operation := os.Getenv("NEXTHINK_WEB_OPERATION")
	if operation == "" {
		operation = "shell.modules"
	}
	var request web_api.Request
	if path := os.Getenv("NEXTHINK_WEB_REQUEST_FILE"); path != "" {
		b, e := os.ReadFile(path)
		if e != nil {
			log.Fatal(e)
		}
		if e = json.Unmarshal(b, &request); e != nil {
			log.Fatal(e)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	body, resp, err := c.WebAPI.Do(ctx, operation, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d; response bytes=%d\n", resp.StatusCode, len(body))
}
