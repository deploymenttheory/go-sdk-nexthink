package visual_editor

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/visual_editor/mocks"
)

func loadContractFixture[T any](t *testing.T, name string) *T {
	t.Helper()
	var v T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &v))
	return &v
}

type contractCase struct {
	name, verb, path string
	status           int
	hasBody, empty   bool
	call             func(*Service) (any, *interfaces.Response, error)
}

func contractCases(t *testing.T) []contractCase {
	t.Helper()
	ctx := context.Background()
	return []contractCase{{name: "ListBCOs", verb: "GET", path: "/apigateway/visual-editor/api/collections/bco?uri=device%2Fdevice", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListBCOs(ctx, "device/device") }},
		{name: "GetDefaultColumns", verb: "GET", path: "/apigateway/visual-editor/api/columns?uri=device%2Fdevice", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetDefaultColumns(ctx, "device/device") }},
		{name: "ListFilterCollections", verb: "GET", path: "/apigateway/visual-editor/api/filters/collections?uri=device%2Fdevice", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ListFilterCollections(ctx, "device/device")
		}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
				}
				if tt.empty {
					return httpmock.NewStringResponse(tt.status, ""), nil
				}
				return mocks.Responder(tt.status, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, tt.status, response.StatusCode)
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if !tt.empty {
				body, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(body))
			}
		})
	}
}
func TestErrors(t *testing.T) {
	for _, tt := range contractCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, code := range []int{400, 401, 403, 409} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "error"))
				_, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, code, response.StatusCode)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			if !tt.empty {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
					r := httpmock.NewStringResponse(200, "{broken")
					r.Header.Set("Content-Type", "application/json")
					return r, nil
				})
				_, _, err = tt.call(NewService(transport))
				require.Error(t, err)
			}
		})
	}
}
func TestValidation(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name string
		call func(*Service) (any, *interfaces.Response, error)
	}{{name: "ListBCOs", call: func(s *Service) (any, *interfaces.Response, error) { return s.ListBCOs(ctx, "") }},
		{name: "GetDefaultColumns", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetDefaultColumns(ctx, "") }},
		{name: "ListFilterCollections", call: func(s *Service) (any, *interfaces.Response, error) { return s.ListFilterCollections(ctx, "") }}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, response)
			assert.Zero(t, mock.GetTotalCallCount())
		})
	}
}
