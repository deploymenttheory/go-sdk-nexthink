package cci_benchmarks

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/cci_benchmarks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestContracts(t *testing.T) {
	t.Run("Query", func(t *testing.T) {
		var request QueryRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Query_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+EndpointQuery, func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Query_request")), string(body))
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "UTC", r.Header.Get("x-client-timezone"))
				fixture := "Query_success"
				if status != 200 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).Query(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 200 {
				require.NoError(t, err)
				b, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(mocks.Fixture("Query_success")), string(b))
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
	_, response, err = s.Query(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.Query(context.Background(), &QueryRequest{})
	require.Error(t, err)
	assert.Nil(t, response)
	assert.Zero(t, mock.GetTotalCallCount())
}
