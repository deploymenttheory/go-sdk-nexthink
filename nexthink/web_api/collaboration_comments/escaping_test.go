package collaboration_comments

import (
	"context"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/collaboration_comments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestReferencesAndSearchEscaping(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/documents/doc%2Fid/comments", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, Endpoint+"/documents/doc%2Fid/comments", r.URL.EscapedPath())
		return mocks.Responder(200, "ListComments_success")(r)
	})
	_, _, err := NewService(transport).ListComments(context.Background(), "doc/id")
	require.NoError(t, err)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/user-mentions?u=name%26admin%3Dtrue", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "name&admin=true", r.URL.Query().Get("u"))
		assert.Len(t, r.URL.Query(), 1)
		return mocks.Responder(200, "ListUserMentions_success")(r)
	})
	_, _, err = NewService(transport).ListUserMentions(context.Background(), "name&admin=true")
	require.NoError(t, err)
}

func TestAcknowledgmentsAllowEmptyBodies(t *testing.T) {
	for _, tt := range contractCases(t) {
		if !tt.empty {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(204, ""))
			_, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 204, response.StatusCode)
		})
	}
}
