// Package labconfig loads explicit fixtures for examples that change lab data.
package labconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadRequest reads the API request body from NEXTHINK_REQUEST_FILE.
func LoadRequest(destination any) error {
	path := os.Getenv("NEXTHINK_REQUEST_FILE")
	if path == "" {
		return fmt.Errorf("NEXTHINK_REQUEST_FILE must name a JSON request containing your intended lab targets")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, destination)
}
