package campaigns

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/campaigns/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestGetFeatures(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+EndpointFeatures, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
		return mocks.Responder(200, "Features_success")(r)
	})
	result, response, err := NewService(transport).GetFeatures(context.Background())
	require.NoError(t, err)
	require.NotNil(t, response)
	encoded, e := json.Marshal(result)
	require.NoError(t, e)
	assert.JSONEq(t, string(mocks.Fixture("Features_success")), string(encoded))
	assert.EqualValues(t, 100, result.LicenseVolumeData.CampaignVolume)
}
func TestFeaturesErrors(t *testing.T) {
	for _, status := range []int{401, 403, 500} {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("GET", testutil.BaseURL+EndpointFeatures, mocks.Responder(status, "error_unauthorized"))
		result, response, err := NewService(transport).GetFeatures(context.Background())
		require.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, status, response.StatusCode)
	}
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+EndpointFeatures, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
	_, _, err := NewService(transport).GetFeatures(context.Background())
	require.Error(t, err)
}
