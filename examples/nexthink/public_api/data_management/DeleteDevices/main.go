package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/data_management"
)

func main() {
	req := new(data_management.DeleteDevicesRequest)
	if err := labconfig.LoadRequest(req); err != nil {
		log.Fatal(err)
	}
	c, err := nexthink.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, resp, err := c.PublicAPI.DataManagement.DeleteDevices(ctx, req, os.Getenv("NEXTHINK_REQUEST_ID"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d: scheduled %d; status %s (asynchronous)\n", resp.StatusCode, result.ScheduledCount, result.Status)
}
