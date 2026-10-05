// Run with NEXTHINK_AUTH=chrome to explicitly opt into reading a signed-in Chrome tab.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth/chrome"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/experimental"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func main() {
	instance, region := os.Getenv("NEXTHINK_INSTANCE"), os.Getenv("NEXTHINK_REGION")
	var provider auth.TokenProvider
	var err error
	switch os.Getenv("NEXTHINK_AUTH") {
	case "chrome":
		provider, err = chrome.New(fmt.Sprintf("https://%s.%s.nexthink.cloud", instance, region))
	case "token":
		provider = auth.StaticToken(os.Getenv("NEXTHINK_ACCESS_TOKEN"), time.Time{})
	case "client":
		var c *nexthink.Client
		c, err = nexthink.NewClientFromEnv(client.WithLogger(zap.NewNop()))
		if err == nil {
			provider = c.GetTokenManager()
		}
	default:
		log.Fatal("NEXTHINK_AUTH must be chrome, token, or client")
	}
	if err != nil {
		log.Fatal(err)
	}
	c, err := experimental.NewClient(instance, region, provider, client.WithLogger(zap.NewNop()))
	if err != nil {
		log.Fatal(err)
	}
	operation := os.Getenv("NEXTHINK_EXPERIMENTAL_OPERATION")
	if operation == "" {
		operation = "shell.modules"
	}
	var request experimental.Request
	if path := os.Getenv("NEXTHINK_EXPERIMENTAL_REQUEST_FILE"); path != "" {
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
	body, resp, err := c.Do(ctx, operation, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d\n%s\n", resp.StatusCode, body)
}
