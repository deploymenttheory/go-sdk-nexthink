package auth

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPortalSession(t *testing.T) {
	for _, s := range []*PortalSession{nil, {}, {Cookie: "x\r\ny"}, {XAuthToken: "x\x00y"}} {
		require.Error(t, s.Validate())
	}
	for _, s := range []*PortalSession{{Cookie: "synthetic-secret"}, {XAuthToken: "synthetic-secret"}} {
		require.NoError(t, s.Validate())
		b, err := json.Marshal(s)
		require.NoError(t, err)
		require.NotContains(t, string(b), "synthetic-secret")
		require.NotContains(t, fmt.Sprint(s), "synthetic-secret")
	}
}
