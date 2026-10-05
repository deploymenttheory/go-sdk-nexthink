// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
// NEXTHINK_SCRIPT_FILE contains PowerShell UTF-8 source including its BOM.
// Inspection does not execute the script; the SDK encodes its bytes.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
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
	path := os.Getenv("NEXTHINK_SCRIPT_FILE")
	if path == "" {
		log.Fatal("NEXTHINK_SCRIPT_FILE is required")
	}
	script, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.RemoteActions.InspectPowerShellScript(ctx, script)
	if err != nil {
		log.Fatal(err)
	}
	if result.Parameters == nil {
		log.Fatal("script parameters not returned")
	}
	fmt.Printf("Inputs: %d; outputs: %d\n", len(result.Parameters.Inputs), len(result.Parameters.Outputs))
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
