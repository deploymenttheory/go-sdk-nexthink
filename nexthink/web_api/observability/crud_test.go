package observability

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/observability/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestSubmitProtocol(t *testing.T) {
	var req Submission
	require.NoError(t, json.Unmarshal(mocks.Fixture("Submit_request"), &req))
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, func(r *http.Request) (*http.Response, error) {
		q := r.URL.Query()
		assert.Equal(t, req.ClientToken, q.Get("dd-api-key"))
		assert.Equal(t, req.Source, q.Get("ddsource"))
		assert.Equal(t, req.Origin, q.Get("dd-evp-origin"))
		assert.Equal(t, req.OriginVersion, q.Get("dd-evp-origin-version"))
		assert.Equal(t, req.RequestID, q.Get("dd-request-id"))
		assert.Equal(t, req.BatchTime, q.Get("batch_time"))
		assert.Equal(t, req.API, q.Get("_dd.api"))
		assert.Equal(t, req.ContentType, r.Header.Get("Content-Type"))
		data, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		assert.Equal(t, req.Payload, data)
		return httpmock.NewStringResponse(202, ""), nil
	})
	resp, err := s.Submit(context.Background(), &req)
	require.NoError(t, err)
	assert.Equal(t, 202, resp.StatusCode)
	for _, code := range []int{400, 401, 403, 413, 429} {
		mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, mocks.Responder(code, "error"))
		resp, err := s.Submit(context.Background(), &req)
		require.Error(t, err)
		assert.Equal(t, code, resp.StatusCode)
	}
	mock.RegisterResponder("POST", testutil.BaseURL+Endpoint, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
	_, err = s.Submit(context.Background(), &req)
	require.Error(t, err)
	before := mock.GetTotalCallCount()
	_, err = s.Submit(context.Background(), nil)
	require.Error(t, err)
	req.ClientToken = ""
	_, err = s.Submit(context.Background(), &req)
	require.Error(t, err)
	assert.Equal(t, before, mock.GetTotalCallCount())
}
