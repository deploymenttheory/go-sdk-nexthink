// Configure NEXTHINK_API=public, instance, region and client credentials.
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
	if c.PublicAPI == nil {
		log.Fatal("set NEXTHINK_API=public or both")
	}
	url := os.Getenv("NEXTHINK_DOWNLOAD_URL")
	output := os.Getenv("NEXTHINK_OUTPUT_FILE")
	if output == "" {
		log.Fatal("set NEXTHINK_OUTPUT_FILE")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := c.PublicAPI.NQL.DownloadNQLExport(ctx, url)
	if err != nil {
		log.Fatal("export download failed; inspect the error without logging a signed URL")
	}
	if err := os.WriteFile(output, result, 0600); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes\n", len(result))
}
