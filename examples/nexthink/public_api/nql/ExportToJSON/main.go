package main

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	client, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()))
	if err != nil {
		log.Fatal(err)
	}
	if client.PublicAPI == nil {
		log.Fatal("set NEXTHINK_API=public")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Minute)
	defer cancel()
	queryID := strings.TrimSpace(os.Getenv("NEXTHINK_QUERY_ID"))
	if queryID == "" {
		log.Fatal("NEXTHINK_QUERY_ID is required")
	}
	output := strings.TrimSpace(os.Getenv("NEXTHINK_OUTPUT_FILE"))
	if output == "" {
		log.Fatal("NEXTHINK_OUTPUT_FILE is required")
	}
	result, err := client.PublicAPI.NQL.ExportToJSON(ctx, queryID)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		log.Fatal(err)
	}
	_, writeErr := file.Write(result.Data)
	closeErr := file.Close()
	if writeErr != nil {
		log.Fatal(writeErr)
	}
	if closeErr != nil {
		log.Fatal(closeErr)
	}
	fmt.Printf("Saved %d bytes (%s) to %s after %d polls\n", result.Size(), result.Format, output, result.PollCount)
}
