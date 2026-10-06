package main

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/examples/internal/labconfig"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

func main() {
	c, err := nexthink.NewClientFromEnv(nexthink.WithLogger(zap.NewNop()), nexthink.WithRetryCount(0))
	if err != nil {
		log.Fatal(err)
	}
	if c.WebAPI == nil {
		log.Fatal("set NEXTHINK_API=web")
	}
	var request struct {
		LibraryID string `json:"libraryID"`
	}
	if err := labconfig.LoadRequest(&request); err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	response, err := c.WebAPI.Library.DeleteCustomPack(ctx, request.LibraryID)
	if err != nil {
		log.Fatal(err)
	}
	if e := json.NewEncoder(os.Stdout).Encode(map[string]int{"status": response.StatusCode}); e != nil {
		log.Fatal(e)
	}
}
