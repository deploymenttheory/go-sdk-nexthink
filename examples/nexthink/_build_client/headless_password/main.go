// Package main demonstrates unattended local-password authentication for Web APIs.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := nexthink.NewClient(&nexthink.AuthConfig{
		Instance: os.Getenv("NEXTHINK_INSTANCE"),
		Region:   os.Getenv("NEXTHINK_REGION"),
		WebAPI: &nexthink.BrowserCredentials{
			UsernamePassword: &nexthink.UsernamePasswordCredentials{
				Username:              os.Getenv("NEXTHINK_USERNAME"),
				Password:              os.Getenv("NEXTHINK_PASSWORD"),
				LoginTimeout:          90 * time.Second,
				BrowserProxy:          os.Getenv("NEXTHINK_BROWSER_PROXY"),
				BrowserExecutablePath: os.Getenv("NEXTHINK_BROWSER_EXECUTABLE_PATH"),
			},
		},
	}, nexthink.WithTimeout(30*time.Second), nexthink.WithRetryCount(0))
	if err != nil {
		return err
	}
	defer c.Close()

	// This deadline includes the initial login as well as the read-only API call.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	result, response, err := c.WebAPI.Applications.List(ctx, nil)
	if err != nil {
		return err
	}
	fmt.Printf("HTTP %d: %d applications on this page, %d total\n",
		response.StatusCode, len(result.Items), result.Total)

	return nil
}
