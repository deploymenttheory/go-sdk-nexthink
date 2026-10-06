// CSV upload inputs and side effects are described in ../README.md.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	path := os.Getenv("NEXTHINK_CSV_FILE")
	if path == "" {
		log.Fatal("NEXTHINK_CSV_FILE required")
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("NEXTHINK_API=web or both required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.CustomFieldValues.ImportCSV(ctx, filepath.Base(path), f, info.Size())
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(map[string]any{"statusCode": response.StatusCode, "result": result}); err != nil {
		log.Fatal(err)
	}
}
