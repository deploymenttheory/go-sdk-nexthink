// Configure NEXTHINK_INSTANCE, NEXTHINK_REGION, NEXTHINK_API=web and web authentication.
// NEXTHINK_REQUEST_FILE must contain the API request JSON.
// This call writes the explicitly supplied resource/configuration or event.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/device_configuration"
	"go.uber.org/zap"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web or both")
	}
	var request device_configuration.SaveProfilesRequest
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, response, err := c.WebAPI.DeviceConfiguration.SetProfiles(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Profiles: %d\n", len(result.Profiles))
	fmt.Printf("HTTP %d\n", response.StatusCode)
}
