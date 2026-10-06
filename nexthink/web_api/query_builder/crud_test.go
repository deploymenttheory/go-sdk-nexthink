package query_builder

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/query_builder/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestContracts(t *testing.T) {
	t.Run("Transform", func(t *testing.T) {
		var request TransformRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Transform_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+EndpointTransform, func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Transform_request")), string(body))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				fixture := "Transform_success"
				if status != 200 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).Transform(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 200 {
				require.NoError(t, err)
				b, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture("Transform_success")), string(b))
			} else {
				require.Error(t, err)
				assert.Nil(t, result)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("ListDrilldownDestinations", func(t *testing.T) {
		var request DestinationsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("ListDrilldownDestinations_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+EndpointListDrilldownDestinations, func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("ListDrilldownDestinations_request")), string(body))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				fixture := "ListDrilldownDestinations_success"
				if status != 200 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).ListDrilldownDestinations(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 200 {
				require.NoError(t, err)
				b, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture("ListDrilldownDestinations_success")), string(b))
			} else {
				require.Error(t, err)
				assert.Nil(t, result)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("TransformDrilldown", func(t *testing.T) {
		var request DrilldownRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("TransformDrilldown_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+EndpointTransformDrilldown, func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("TransformDrilldown_request")), string(body))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				fixture := "TransformDrilldown_success"
				if status != 200 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).TransformDrilldown(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 200 {
				require.NoError(t, err)
				b, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture("TransformDrilldown_success")), string(b))
			} else {
				require.Error(t, err)
				assert.Nil(t, result)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
}
func TestValidationPreventsRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	var response *interfaces.Response
	var err error
	_, response, err = s.Transform(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.Transform(context.Background(), &TransformRequest{})
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.ListDrilldownDestinations(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.ListDrilldownDestinations(context.Background(), &DestinationsRequest{})
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.TransformDrilldown(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.TransformDrilldown(context.Background(), &DrilldownRequest{})
	require.Error(t, err)
	assert.Nil(t, response)
	assert.Zero(t, mock.GetTotalCallCount())
}
