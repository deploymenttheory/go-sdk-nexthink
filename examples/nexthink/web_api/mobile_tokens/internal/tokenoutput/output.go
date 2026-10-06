// Package tokenoutput keeps enrollment credentials out of example stdout.
package tokenoutput

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/mobile_tokens"
)

// Write prints token metadata without JWTToken. An explicit outputPath also saves
// the full response in a new owner-only file; an existing path is never replaced.
func Write(result *mobile_tokens.Token, outputPath string, stdout io.Writer) error {
	if result == nil {
		return fmt.Errorf("token response is nil")
	}
	if outputPath != "" {
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("encode full token response: %w", err)
		}
		file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
		if err != nil {
			return fmt.Errorf("create token output file: %w", err)
		}
		_, writeErr := file.Write(append(data, '\n'))
		closeErr := file.Close()
		if writeErr != nil {
			return fmt.Errorf("write token output file: %w", writeErr)
		}
		if closeErr != nil {
			return fmt.Errorf("close token output file: %w", closeErr)
		}
	}
	metadata := *result
	metadata.JWTToken = nil
	if err := json.NewEncoder(stdout).Encode(metadata); err != nil {
		return fmt.Errorf("write token metadata: %w", err)
	}
	return nil
}
