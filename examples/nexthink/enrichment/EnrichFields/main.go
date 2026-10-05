package main

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/services/enrichment"
	"go.uber.org/zap"
)

// This example demonstrates how to enrich fields for Nexthink objects.
//
// Enrichment can be done for:
// - Manual custom fields (any object)
// - Virtualization fields (devices only)
// - Configuration_tag (devices only)
// - Organization field (users only)
// - Entra ID fields (users only)
//
// The request can contain 1-10000 enrichment operations.
//
// Response types:
// - 200 OK: All objects processed successfully
// - 207 Multi-Status: Some objects processed, others failed
// - 400 Bad Request: All objects failed
//
// Use this to:
// - Populate custom fields with external data
// - Update device virtualization information
// - Sync user organization data
// - Tag devices with configuration information

func main() {
	clientID := os.Getenv("NEXTHINK_CLIENT_ID")
	clientSecret := os.Getenv("NEXTHINK_CLIENT_SECRET")
	instance := os.Getenv("NEXTHINK_INSTANCE")
	region := os.Getenv("NEXTHINK_REGION")

	if clientID == "" || clientSecret == "" || instance == "" || region == "" {
		log.Fatal("NEXTHINK_CLIENT_ID, NEXTHINK_CLIENT_SECRET, NEXTHINK_INSTANCE, and NEXTHINK_REGION environment variables are required")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	nxClient, err := nexthink.NewClient(
		clientID,
		clientSecret,
		instance,
		region,
		client.WithLogger(logger),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Enrich device and user fields
	request := new(enrichment.EnrichmentRequest)
	if err := labconfig.LoadRequest(request); err != nil {
		log.Fatal(err)
	}

	result, resp, err := nxClient.Enrichment.EnrichFields(ctx, request)
	if err != nil {
		log.Fatalf("Failed to enrich fields: %v", err)
	}

	// Display results
	fmt.Printf("\n=== Enrichment Results ===\n")
	fmt.Printf("HTTP Status: %d\n", resp.StatusCode)
	fmt.Printf("Duration: %v\n\n", resp.Duration)

	// Handle different response types based on status code
	switch resp.StatusCode {
	case 200:
		fmt.Printf("✓ All enrichments processed successfully!\n")
		fmt.Printf("Result: %+v\n", result)

	case 207:
		fmt.Printf("⚠ Partial success - some enrichments failed\n")
		fmt.Printf("Result: %+v\n", result)
		fmt.Printf("\n💡 Check the errors array for details on failed enrichments\n")

	case 400:
		fmt.Printf("✗ All enrichments failed\n")
		fmt.Printf("Result: %+v\n", result)

	default:
		fmt.Printf("Result: %+v\n", result)
	}

	logger.Info("Enrichment completed",
		zap.Int("status_code", resp.StatusCode),
		zap.Int("enrichment_count", len(request.Enrichments)))

	fmt.Printf("\n✓ Enrichment operation completed!\n")
	fmt.Printf("\n💡 Tips:\n")
	fmt.Printf("   - Use device/device/name, device/device/uid, or user/user/upn for identification\n")
	fmt.Printf("   - Custom fields must be prefixed with # (e.g., #cost_center)\n")
	fmt.Printf("   - Batch up to 10000 enrichments in a single request\n")
}
