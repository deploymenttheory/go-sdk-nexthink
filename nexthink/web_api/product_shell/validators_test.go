package product_shell

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestGetFlagValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetFlag(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestGetDynamicMenuValidation(t *testing.T) {
	for _, value := range []string{"", " ", ".", "..", "a/b", `a\b`} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(transport).GetDynamicMenu(context.Background(), value)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestValidateClaimsValidation(t *testing.T) {
	for _, request := range []*ClaimsRequest{nil, {}, {Method: "hasAnyClaim"}, {Method: "hasPatternClaim"}, {Method: "hasClaimValue"}} {
		transport, mock := testutil.NewTransport(t)
		result, response, err := NewService(transport).ValidateClaims(context.Background(), request)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, response)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
func TestPostTelemetryValidation(t *testing.T) {
	for _, value := range []string{"", "null", "{broken"} {
		transport, mock := testutil.NewTransport(t)
		result, resp, err := NewService(
			transport,
		).PostTelemetry(context.Background(), json.RawMessage(value))
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestOpaqueRequestShapes(t *testing.T) {
	for _, raw := range []string{`{}`, `[]`, `["fixture-claim"]`} {
		require.NoError(t, ValidateRequest(json.RawMessage(raw)))
	}
}

func TestClaimsVariants(t *testing.T) {
	empty := []string{}
	pattern, claim := "nx_*", "nx_example"
	for _, request := range []*ClaimsRequest{
		{Method: "hasAllClaims", Claims: &empty},
		{Method: "hasAnyClaim", Claims: &empty},
		{Method: "hasPatternClaim", PatternClaim: &pattern},
		{Method: "hasClaimValue", Claim: &claim, Value: json.RawMessage(`false`)},
	} {
		require.NoError(t, ValidateClaimsRequest(request))
	}
	encoded, err := json.Marshal(&ClaimsRequest{Method: "hasAnyClaim", Claims: &empty})
	require.NoError(t, err)
	assert.JSONEq(t, `{"method":"hasAnyClaim","claims":[]}`, string(encoded))
}
