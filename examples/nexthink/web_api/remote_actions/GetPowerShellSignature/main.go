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
	result, response, err := c.WebAPI.RemoteActions.GetPowerShellSignature(ctx, script)
	if err != nil {
		log.Fatal(err)
	}
	if result.Signature == nil {
		log.Fatal("signature information not returned")
	}
	fmt.Printf("Signature state: %s\n", result.Signature.State)
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
