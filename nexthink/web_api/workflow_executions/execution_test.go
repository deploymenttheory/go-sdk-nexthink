package workflow_executions

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workflow_executions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestExecutionSourceHeader(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	request := load[ExecuteRequest](t, "Execute_request")
	request.Source = "DEVICE_VIEW"
	mock.RegisterResponder("POST", testutil.BaseURL+EndpointExecute, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "DEVICE_VIEW", r.Header.Get("nx-source"))
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var value map[string]json.RawMessage
		require.NoError(t, json.Unmarshal(body, &value))
		assert.NotContains(t, value, "Source")
		assert.NotContains(t, value, "source")
		return mocks.Responder(200, "Execute_success")(r)
	})
	_, _, err := NewService(transport).Execute(context.Background(), request)
	require.NoError(t, err)
}
func TestExecutionRequiresTargets(t *testing.T) {
	r := load[ExecuteRequest](t, "Execute_request")
	r.Targets = nil
	require.Error(t, validateExecuteRequest(r))
	r.Targets = []Target{{}}
	require.Error(t, validateExecuteRequest(r))
}
