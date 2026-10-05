package spark

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

func TestValidateHandoff(t *testing.T) {
	empty := ""
	for _, req := range []*HandoffRequest{nil, {}, {Message: Message{Parts: []Part{{Type: "OTHER"}}}}, {Message: Message{Parts: []Part{{Type: "TEXT"}}}}, {Message: Message{Parts: []Part{{Type: "TEXT", Text: "x", MIMEType: "text/plain"}}}}, {Message: Message{Parts: []Part{{Type: "TEXT", Text: "x", FileContent: &empty}}}}, {Message: Message{Parts: []Part{{Type: "FILE", MIMEType: "text/plain"}}}}, {Message: Message{Parts: []Part{{Type: "FILE", FileContent: &empty}}}}, {Message: Message{Parts: []Part{{Type: "FILE", MIMEType: "text/plain", FileContent: &empty, Text: "x"}}}}} {
		transport, mock := testutil.NewTransport(t)
		resp, err := NewService(
			transport,
		).Handoff(context.Background(), "fixture@example.test", "", req)
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
}

func TestValidateHeaders(t *testing.T) {
	for _, tc := range []struct{ upn, zone string }{{"", ""}, {" ", ""}, {"user\r\nInjected: true", ""}, {"fixture@example.test", "UTC\nInjected: true"}} {
		transport, mock := testutil.NewTransport(t)
		resp, err := NewService(
			transport,
		).Handoff(context.Background(), tc.upn, tc.zone, handoffRequest())
		require.Error(t, err)
		assert.Nil(t, resp)
		assert.Zero(t, mock.GetTotalCallCount())
	}
	require.NoError(t, ValidateHandoffHeaders("fixture@example.test", ""))
}
