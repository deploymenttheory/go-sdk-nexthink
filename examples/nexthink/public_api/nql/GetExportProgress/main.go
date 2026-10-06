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
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	exportID := strings.TrimSpace(os.Getenv("NEXTHINK_EXPORT_ID"))
	if exportID == "" {
		log.Fatal("NEXTHINK_EXPORT_ID is required")
	}
	result, err := client.PublicAPI.NQL.GetExportProgress(ctx, exportID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
