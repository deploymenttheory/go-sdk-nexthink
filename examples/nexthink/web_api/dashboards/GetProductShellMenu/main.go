package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dashboards"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		panic(err)
	}
	defer c.Close()
	if c.WebAPI == nil {
		panic("set NEXTHINK_API=web or both")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	var options *dashboards.ProductShellMenuOptions
	if areas := os.Getenv("NEXTHINK_PRODUCT_AREAS"); areas != "" {
		options = &dashboards.ProductShellMenuOptions{ProductAreas: strings.Split(areas, ",")}
	}
	result, _, err := c.WebAPI.Dashboards.GetProductShellMenu(ctx, options)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
