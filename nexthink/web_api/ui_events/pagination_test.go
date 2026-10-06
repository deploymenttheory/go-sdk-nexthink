package ui_events

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ui_events/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCursorAndScope(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/messages?cursor=a%2Bb", mocks.Responder(200, "Poll_success"))
	s := NewService(transport)
	_, _, err := s.Poll(context.Background(), "messages?cursor=a%2Bb")
	require.NoError(t, err)
	_, _, err = s.Poll(context.Background(), testutil.BaseURL+Endpoint+"/messages?cursor=a%2Bb")
	require.NoError(t, err)
	for _, bad := range []string{"//elsewhere.invalid/messages", "https://elsewhere.invalid" + Endpoint + "/messages", "../messages", "/different/path", "messages#fragment", "messages%2f..%2fsecrets"} {
		before := mock.GetTotalCallCount()
		_, resp, err := s.Poll(context.Background(), bad)
		require.Error(t, err)
		if resp != nil {
			assert.Zero(t, resp.StatusCode)
		}
		assert.Equal(t, before, mock.GetTotalCallCount())
	}
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/messages", httpmock.NewStringResponder(422, "disabled"))
	_, resp, err := s.Poll(context.Background(), "")
	require.Error(t, err)
	assert.Equal(t, 422, resp.StatusCode)
}
