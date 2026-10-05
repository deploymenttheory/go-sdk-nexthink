//go:build !darwin

package chrome

import (
	"context"
	"fmt"
)

func readSession(context.Context, string) ([]byte, error) {
	return nil, fmt.Errorf("Chrome session provider requires macOS; supply an auth.TokenProvider on this platform")
}
