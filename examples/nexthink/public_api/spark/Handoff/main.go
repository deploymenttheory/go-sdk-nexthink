package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/spark"
)

func main() {
	req := new(spark.HandoffRequest)
	if err := labconfig.LoadRequest(req); err != nil {
		log.Fatal(err)
	}
	c, err := nexthink.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	resp, err := c.PublicAPI.Spark.Handoff(ctx, os.Getenv("NEXTHINK_USER_UPN"), os.Getenv("NEXTHINK_TIMEZONE"), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("HTTP %d\n", resp.StatusCode)
}
