package cci_benchmarks

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/cci_benchmarks/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestPartialResultsAndTimeZone(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+EndpointQuery, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "Europe/London", r.Header.Get("x-client-timezone"))
		return mocks.Responder(200, "Query_partial")(r)
	})
	request := &QueryRequest{TimeZone: "Europe/London", Queries: []BenchmarkQuery{{Source: Source{Name: "dex_login_duration"}, Metrics: []string{"avg_login_duration_per_device"}}}}
	result, _, err := NewService(transport).Query(context.Background(), request)
	require.NoError(t, err)
	require.Len(t, result.Results, 1)
	require.Len(t, result.Errors, 1)
	b, err := json.Marshal(result)
	require.NoError(t, err)
	assert.JSONEq(t, string(mocks.Fixture("Query_partial")), string(b))
}
func TestInvalidTimeZone(t *testing.T) {
	require.Error(t, validateQuery(&QueryRequest{TimeZone: "bad-zone", Queries: []BenchmarkQuery{{Source: Source{Name: "benchmark"}, Metrics: []string{"metric"}}}}))
}
