package appearance

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/appearance/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestImageContracts(t *testing.T) {
	var asset struct {
		Name        AssetName `json:"name"`
		ContentType string    `json:"contentType"`
		Data        []byte    `json:"data"`
	}
	require.NoError(t, json.Unmarshal(mocks.Fixture("asset"), &asset))
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	path := testutil.BaseURL + Endpoint + "/menu-logo"
	mock.RegisterResponder("GET", path, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "image/*", r.Header.Get("Accept"))
		assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
		resp := httpmock.NewBytesResponse(200, asset.Data)
		resp.Header.Set("Content-Type", asset.ContentType)
		return resp, nil
	})
	data, resp, err := s.Get(context.Background(), asset.Name)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, asset.Data, data)
	assert.Equal(t, asset.ContentType, resp.Headers.Get("Content-Type"))
	mock.RegisterResponder("PUT", path, func(r *http.Request) (*http.Response, error) {
		data, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		assert.Equal(t, asset.Data, data)
		assert.Equal(t, asset.ContentType, r.Header.Get("Content-Type"))
		return mocks.Responder(200, "Update_success")(r)
	})
	result, _, err := s.Update(context.Background(), asset.Name, asset.ContentType, asset.Data)
	require.NoError(t, err)
	assert.True(t, result.Result.Success)
	mock.RegisterResponder("PUT", path, mocks.Responder(200, "Update_failure"))
	result, resp, err = s.Update(context.Background(), asset.Name, asset.ContentType, asset.Data)
	require.Error(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.False(t, result.Result.Success)
	for _, code := range []int{400, 401, 403, 413} {
		mock.RegisterResponder("GET", path, mocks.Responder(code, "error"))
		_, res, err := s.Get(context.Background(), asset.Name)
		require.Error(t, err)
		assert.Equal(t, code, res.StatusCode)
		mock.RegisterResponder("PUT", path, mocks.Responder(code, "error"))
		_, res, err = s.Update(context.Background(), asset.Name, asset.ContentType, asset.Data)
		require.Error(t, err)
		assert.Equal(t, code, res.StatusCode)
	}
	mock.RegisterResponder("PUT", path, func(_ *http.Request) (*http.Response, error) {
		r := httpmock.NewStringResponse(200, "{broken")
		r.Header.Set("Content-Type", "application/json")
		return r, nil
	})
	_, _, err = s.Update(context.Background(), asset.Name, asset.ContentType, asset.Data)
	require.Error(t, err)
	for _, verb := range []string{"GET", "PUT"} {
		mock.RegisterResponder(verb, path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
	}
	_, _, err = s.Get(context.Background(), asset.Name)
	require.Error(t, err)
	_, _, err = s.Update(context.Background(), asset.Name, asset.ContentType, asset.Data)
	require.Error(t, err)
	before := mock.GetTotalCallCount()
	_, _, err = s.Get(context.Background(), "../bad")
	require.Error(t, err)
	_, _, err = s.Update(context.Background(), MenuLogo, "text/plain", asset.Data)
	require.Error(t, err)
	_, _, err = s.Update(context.Background(), MenuLogo, asset.ContentType, nil)
	require.Error(t, err)
	assert.Equal(t, before, mock.GetTotalCallCount())
}
