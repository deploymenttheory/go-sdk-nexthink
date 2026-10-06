package azure_ad_credentials

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateCheckCredentials(t *testing.T) {
	good := &CheckCredentialsRequest{TenantID: "tenant", ClientID: "client", ClientSecret: "secret", NationalCloud: "GLOBAL"}
	require.NoError(t, ValidateCheckCredentials(good))
	cases := []*CheckCredentialsRequest{nil, {}}
	{
		r := *good
		r.TenantID = " "
		cases = append(cases, &r)
	}
	{
		r := *good
		r.ClientID = " "
		cases = append(cases, &r)
	}
	{
		r := *good
		r.NationalCloud = " "
		cases = append(cases, &r)
	}
	for _, r := range cases {
		transport, mock := testutil.NewTransport(t)
		response, err := NewService(transport).CheckCredentials(context.Background(), r)
		require.Error(t, err)
		assert.Nil(t, response)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}
