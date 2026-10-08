// See README.md for inputs and tenant-wide side effects.
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
	defer c.Close()
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	output := os.Getenv("NEXTHINK_OUTPUT_FILE")
	if output == "" {
		log.Fatal("set NEXTHINK_OUTPUT_FILE")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, _, err := c.WebAPI.DeviceClassification.DownloadOrganization(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(output, result, 0o600); err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(map[string]string{"output": output}); err != nil {
		log.Fatal(err)
	}
}
