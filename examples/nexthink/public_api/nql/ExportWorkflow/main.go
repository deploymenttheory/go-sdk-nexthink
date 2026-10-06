package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/nql"
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
	requestPath := strings.TrimSpace(os.Getenv("NEXTHINK_REQUEST_FILE"))
	if requestPath == "" {
		log.Fatal("NEXTHINK_REQUEST_FILE is required")
	}
	data, err := os.ReadFile(requestPath)
	if err != nil {
		log.Fatal(err)
	}
	var request nql.ExportRequest
	if err = json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}
	format := strings.ToLower(strings.TrimSpace(os.Getenv("NEXTHINK_EXPORT_FORMAT")))
	if format == "" {
		format = nql.ExportFormatCSV
	}
	options := nql.DefaultExportOptions().WithFormat(format).WithPollInterval(5 * time.Second).WithTimeout(10 * time.Minute).WithOnStatusChange(func(oldStatus, newStatus string, elapsed time.Duration) {
		fmt.Fprintf(os.Stderr, "%s -> %s (%s)\n", oldStatus, newStatus, elapsed.Round(time.Second))
	})
	output := strings.TrimSpace(os.Getenv("NEXTHINK_OUTPUT_FILE"))
	if output == "" {
		log.Fatal("NEXTHINK_OUTPUT_FILE is required")
	}
	result, err := client.PublicAPI.NQL.ExportWorkflow(ctx, &request, options)
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
