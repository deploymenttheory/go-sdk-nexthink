// See README.md for inputs and tenant-wide side effects.
package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/device_classification"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
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
	path := os.Getenv("NEXTHINK_REQUEST_FILE")
	if path == "" {
		log.Fatal("set NEXTHINK_REQUEST_FILE")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var request device_classification.RulesetUpload
	if err := json.Unmarshal(data, &request); err != nil {
		log.Fatal(err)
	}
	if csvFile := os.Getenv("NEXTHINK_CSV_FILE"); csvFile != "" {
		request.CSV, err = os.ReadFile(csvFile)
		if err != nil {
			log.Fatal(err)
		}
		request.Filename = filepath.Base(csvFile)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	response, err := c.WebAPI.DeviceClassification.CreateVPNEgress(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.NewEncoder(os.Stdout).Encode(map[string]int{"status": response.StatusCode}); err != nil {
		log.Fatal(err)
	}
}
