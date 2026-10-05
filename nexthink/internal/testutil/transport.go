// Package testutil provides isolated, in-memory HTTP transports for contract tests.
package testutil

import (
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
)

const BaseURL = "https://test.eu.nexthink.cloud"

func NewTransport(t *testing.T) (*client.Transport, *httpmock.MockTransport) {
	t.Helper()
	mock := httpmock.NewMockTransport()
	transport, err := client.NewTransportWithTokenProvider(
		BaseURL,
		auth.StaticToken("fixture-token", time.Time{}),
		client.WithLogger(zap.NewNop()),
		client.WithTransport(mock),
		client.WithRetryCount(0),
	)
	require.NoError(t, err)
	return transport, mock
}
