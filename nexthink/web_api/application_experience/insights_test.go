package application_experience

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/application_experience/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestApplicationInsightsProxy(t *testing.T) {
	var request ApplicationInsightsRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("GetApplicationInsights_request"), &request))
	for _, status := range []int{200, 403} {
		transport, mock := testutil.NewTransport(t)
		mock.RegisterResponder("POST", testutil.BaseURL+EndpointApplicationInsights, func(r *http.Request) (*http.Response, error) {
			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture("GetApplicationInsights_request")), string(body))
			for name, value := range map[string]string{"Authorization": "Bearer fixture-token", "X-NX-MaxAge": "300", "X-NX-Originator": "appex-dashboard", "X-NX-Timeout": "20", "X-NX-Priority": "1"} {
				assert.Equal(t, value, r.Header.Get(name))
			}
			if status == 200 {
				return mocks.Responder(status, "GetApplicationInsights_success")(r)
			}
			return mocks.Responder(status, "error_unauthorized")(r)
		})
		result, response, err := NewService(transport).GetApplicationInsights(context.Background(), &request)
		require.NotNil(t, response)
		assert.Equal(t, status, response.StatusCode)
		if status == 200 {
			require.NoError(t, err)
			b, e := json.Marshal(result)
			require.NoError(t, e)
			assert.JSONEq(t, string(mocks.Fixture("GetApplicationInsights_success")), string(b))
		} else {
			require.Error(t, err)
			assert.Nil(t, result)
		}
	}
}
func TestApplicationInsightsValidation(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	for _, request := range []*ApplicationInsightsRequest{nil, {}, {ApplicationID: "fixture", CurrentTimeframe: "from a to b", PreviousTimeframe: "from a to b", Timezone: "not-a-zone", Breakdowns: map[InsightCategory][]string{}}} {
		_, response, err := NewService(transport).GetApplicationInsights(context.Background(), request)
		require.Error(t, err)
		assert.Nil(t, response)
	}
	assert.Zero(t, mock.GetTotalCallCount())
}
