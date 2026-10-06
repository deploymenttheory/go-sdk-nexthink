package snapshots

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/snapshots/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var r T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &r))
	return &r
}
func TestDiscoveryContracts(t *testing.T) {
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	_ = id
	cases := []struct {
		name, verb, path string
		body, empty      bool
		call             func(*Service) (any, *interfaces.Response, error)
		invalid          func(*Service) (any, *interfaces.Response, error)
	}{{name: "List", verb: "GET", path: "/apigateway/content-administration/api/v2/contents/trends", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.List(ctx) }},
		{name: "Get", verb: "GET", path: "/apigateway/api/v1/custom-trends/definition/11111111-2222-4333-8444-555555555555", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.Get(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Get(ctx, "../bad") }},
		{name: "Export", verb: "GET", path: "/apigateway/api/v1/custom-trends/definition/11111111-2222-4333-8444-555555555555", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.Export(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Export(ctx, "../bad") }},
		{name: "Delete", verb: "DELETE", path: "/apigateway/api/v1/custom-trends/definition/11111111-2222-4333-8444-555555555555", body: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) { r, e := s.Delete(ctx, id); return nil, r, e }, invalid: func(s *Service) (any, *interfaces.Response, error) { r, e := s.Delete(ctx, "../bad"); return nil, r, e }},
		{name: "Create", verb: "POST", path: "/apigateway/api/v1/custom-trends/definition", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Create(ctx, load[Definition](t, "Create_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Create(ctx, nil) }},
		{name: "Update", verb: "PUT", path: "/apigateway/api/v1/custom-trends/definition/11111111-2222-4333-8444-555555555555", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Update(ctx, id, load[Definition](t, "Update_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Update(ctx, id, nil) }},
		{name: "Import", verb: "POST", path: "/apigateway/api/v1/custom-trends/definition", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Import(ctx, load[Definition](t, "Import_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.Import(ctx, nil) }}}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				}
				if tt.empty {
					return httpmock.NewStringResponse(200, ""), nil
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, resp, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode)
			if !tt.empty {
				encoded, err := json.Marshal(result)
				require.NoError(t, err)
				expected := mocks.Fixture(tt.name + "_success")
				if tt.name == "Export" {
					var raw map[string]json.RawMessage
					require.NoError(t, json.Unmarshal(expected, &raw))
					delete(raw, "contentId")
					delete(raw, "bcsRevisionNumber")
					expected, err = json.Marshal(raw)
					require.NoError(t, err)
				}
				assert.JSONEq(t, string(expected), string(encoded))
			}
			for _, status := range []int{400, 401, 403, 409, 422, 500} {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(status, "leads_error"))
				_, resp, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, resp)
				assert.Equal(t, status, resp.StatusCode)
			}
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err = tt.call(NewService(transport))
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
			if tt.invalid != nil {
				before := mock.GetTotalCallCount()
				_, resp, err := tt.invalid(NewService(transport))
				require.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, before, mock.GetTotalCallCount())
			}
		})
	}
}
