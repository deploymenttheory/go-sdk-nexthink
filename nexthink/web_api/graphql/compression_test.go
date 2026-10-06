package graphql

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
)

// The lab visual-editor gateway returns real gzip for Accept: */*, but base64
// gzip text mislabeled as gzip for Accept: application/json. Exercise Resty's
// actual decoder so the test catches negotiation failures before JSON parsing.
func TestBrowserGatewayCompressionNegotiation(t *testing.T) {
	payload := []byte(`{"data":{"filterFields":[{"uri":"device/device/name"}]}}`)
	var compressed bytes.Buffer
	writer := gzip.NewWriter(&compressed)
	_, err := writer.Write(payload)
	require.NoError(t, err)
	require.NoError(t, writer.Close())
	for _, override := range []bool{false, true} {
		name := "DefaultBrowserAccept"
		if override {
			name = "ExplicitJSONAcceptRetainsDecodeError"
		}
		t.Run(name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+EndpointVisualEditor, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "gzip", r.Header.Get("Accept-Encoding"))
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				body := compressed.Bytes()
				if r.Header.Get("Accept") == "application/json" {
					body = []byte(base64.StdEncoding.EncodeToString(body))
				} else {
					assert.Equal(t, "*/*", r.Header.Get("Accept"))
				}
				response := httpmock.NewBytesResponse(200, body)
				response.Header.Set("Content-Type", "application/json")
				response.Header.Set("Content-Encoding", "gzip")
				return response, nil
			})
			request := GraphQLRequest{Query: `query Fields { filterFields { uri } }`, OperationName: "Fields"}
			if override {
				request.Headers = map[string]string{"Accept": "application/json"}
			}
			result, response, err := NewService(transport).Execute(context.Background(), "graphql.visual_editor", request)
			require.NotNil(t, response)
			if override {
				require.ErrorContains(t, err, "gzip: invalid header")
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.JSONEq(t, `{"filterFields":[{"uri":"device/device/name"}]}`, string(result.Data))
			}
		})
	}
}

func TestRequestHeadersForwardWithoutSerializing(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+EndpointVisualEditor, func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "application/custom+json", r.Header.Get("Accept"))
		assert.Equal(t, "fixture-client-context", r.Header.Get("Nx-Client-Context"))
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		data, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var body map[string]json.RawMessage
		require.NoError(t, json.Unmarshal(data, &body))
		assert.NotContains(t, body, "Headers")
		assert.NotContains(t, body, "headers")
		assert.NotContains(t, string(data), "fixture-client-context")
		assert.JSONEq(t, `{"query":"query Fields { filterFields { uri } }","operationName":"Fields"}`, string(data))
		response := httpmock.NewStringResponse(200, `{"data":{"filterFields":[]}}`)
		response.Header.Set("Content-Type", "application/json")
		return response, nil
	})
	result, _, err := NewService(transport).Execute(context.Background(), "graphql.visual_editor", GraphQLRequest{Query: `query Fields { filterFields { uri } }`, OperationName: "Fields", Headers: map[string]string{"Accept": "application/custom+json", "Nx-Client-Context": "fixture-client-context"}})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 1, mock.GetTotalCallCount())
}
