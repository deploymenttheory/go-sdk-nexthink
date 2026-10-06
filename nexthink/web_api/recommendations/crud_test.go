package recommendations

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/recommendations/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type contractCase struct {
	name, method, path string
	call               func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	var request UpdateStatusRequest
	require.NoError(t, json.Unmarshal(mocks.Fixture("UpdateStatus_request"), &request))
	return []contractCase{{"List", "GET", Endpoint, func(s *Service) (any, *interfaces.Response, error) { return s.List(context.Background()) }}, {"UpdateStatus", "PUT", Endpoint + "/fixture%2Fid/status", func(s *Service) (any, *interfaces.Response, error) {
		return s.UpdateStatus(context.Background(), "fixture/id", &request)
	}}}
}
func TestContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, tt.path, r.URL.EscapedPath())
				if tt.method == "PUT" {
					b, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture("UpdateStatus_request")), string(b))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 1, mock.GetTotalCallCount())
			data, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(data))
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contractCases(t) {
		for _, failure := range []struct {
			name      string
			responder httpmock.Responder
		}{{"400", mocks.Responder(400, "error_400")}, {"401", mocks.Responder(401, "error_401")}, {"403", mocks.Responder(403, "error_403")}, {"404", mocks.Responder(404, "error_404")}, {"Transport", httpmock.NewErrorResponder(io.ErrUnexpectedEOF)}, {"Malformed", func(r *http.Request) (*http.Response, error) {
			x := httpmock.NewStringResponse(200, "{broken")
			x.Header.Set("Content-Type", "application/json")
			return x, nil
		}}} {
			t.Run(tt.name+failure.name, func(t *testing.T) {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, failure.responder)
				result, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				assert.Nil(t, result)
				if failure.name != "Transport" {
					require.NotNil(t, response)
				}
			})
		}
	}
}
func TestEmptyListAndOptionalNote(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint, mocks.Responder(200, "List_empty"))
	result, _, err := NewService(transport).List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Empty(t, *result)
	b, err := json.Marshal(&UpdateStatusRequest{Status: StatusNew})
	require.NoError(t, err)
	assert.JSONEq(t, `{"status":"new"}`, string(b))
}
