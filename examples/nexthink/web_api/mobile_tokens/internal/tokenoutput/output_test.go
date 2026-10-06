package tokenoutput

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/mobile_tokens"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultOutputOmitsCredential(t *testing.T) {
	token := "synthetic-enrollment-credential"
	result := &mobile_tokens.Token{Name: "fixture", JWTToken: &token}
	var stdout bytes.Buffer
	require.NoError(t, Write(result, "", &stdout))
	assert.NotContains(t, stdout.String(), token)
	assert.NotContains(t, stdout.String(), "jwtToken")
	assert.Contains(t, stdout.String(), "fixture")
	require.NotNil(t, result.JWTToken)
	assert.Equal(t, token, *result.JWTToken)
}
func TestPrivateOutputAndExclusiveCreation(t *testing.T) {
	token := "synthetic-enrollment-credential"
	result := &mobile_tokens.Token{Name: "fixture", JWTToken: &token}
	path := filepath.Join(t.TempDir(), "enrollment.json")
	var stdout bytes.Buffer
	require.NoError(t, Write(result, path, &stdout))
	assert.NotContains(t, stdout.String(), token)
	info, err := os.Stat(path)
	require.NoError(t, err)
	assert.Zero(t, info.Mode().Perm()&0o077)
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	var saved mobile_tokens.Token
	require.NoError(t, json.Unmarshal(data, &saved))
	require.NotNil(t, saved.JWTToken)
	assert.Equal(t, token, *saved.JWTToken)
	result.Name = "replacement"
	require.Error(t, Write(result, path, &stdout))
	after, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, data, after)
}
