package content_sharing

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
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/content_sharing/mocks"
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
	id := "11111111-2222-4333-8444-555555555555"
	return []contractCase{{name: "GetActions", verb: "GET", path: "/apigateway/content-administration/api/v3/share/actions?contentKey=custom-fields&resourceName=manual-custom-field", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetActions(ctx, loadContractFixture[ActionsOptions](t, "GetActions_input"))
	}},
		{name: "GetProfiles", verb: "GET", path: "/apigateway/content-administration/api/v3/share/profiles?bcsName=Fixture+field&contentId=11111111-2222-4333-8444-555555555555&contentKey=custom-fields&resourceName=manual-custom-field&shared=false", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetProfiles(ctx, loadContractFixture[ProfilesOptions](t, "GetProfiles_input"))
		}},
		{name: "GetUser", verb: "GET", path: "/apigateway/content-administration/api/v2/user/11111111-2222-4333-8444-555555555555", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetUser(ctx, id) }},
		{name: "GetLegacyOwner", verb: "GET", path: "/apigateway/coad/v1/sharingcontent/content-owner/11111111-2222-4333-8444-555555555555", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLegacyOwner(ctx, id) }},
		{name: "SetProfiles", verb: "POST", path: "/apigateway/content-administration/api/v3/share/profiles?contentKey=custom-fields", status: 200, hasBody: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.SetProfiles(ctx, "custom-fields", *loadContractFixture[[]ShareContent](t, "SetProfiles_input"))
			return nil, response, err
		}},
		{name: "GetLegacyActions", verb: "GET", path: "/apigateway/coad/v1/sharingcontent/permissions?service=appex", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLegacyActions(ctx, "appex") }},
		{name: "GetLegacyProfiles", verb: "GET", path: "/apigateway/coad/v1/sharingcontent/profiles?bcsName=Fixture+app&contentId=11111111-2222-4333-8444-555555555555&service=appex&shared=false", status: 200, hasBody: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetLegacyProfiles(ctx, loadContractFixture[LegacyOptions](t, "GetLegacyProfiles_input"), false)
		}},
		{name: "SetLegacyProfiles", verb: "POST", path: "/apigateway/coad/v1/sharingcontent/profiles?bcsName=Fixture+app&contentId=11111111-2222-4333-8444-555555555555&service=appex", status: 200, hasBody: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.SetLegacyProfiles(ctx, &LegacyOptions{Service: "appex", ContentID: id, BCSName: "Fixture app"}, loadContractFixture[LegacyUpdateRequest](t, "SetLegacyProfiles_input"))
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
	}{{name: "GetActions", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetActions(ctx, nil) }},
		{name: "GetProfiles", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetProfiles(ctx, nil) }},
		{name: "GetUser", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetUser(ctx, "../bad") }},
		{name: "GetLegacyOwner", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLegacyOwner(ctx, "../bad") }},
		{name: "SetProfiles", call: func(s *Service) (any, *interfaces.Response, error) {
			resp, err := s.SetProfiles(ctx, "custom-fields", nil)
			return nil, resp, err
		}},
		{name: "GetLegacyActions", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLegacyActions(ctx, "") }},
		{name: "GetLegacyProfiles", call: func(s *Service) (any, *interfaces.Response, error) { return s.GetLegacyProfiles(ctx, nil, false) }},
		{name: "SetLegacyProfiles", call: func(s *Service) (any, *interfaces.Response, error) { return s.SetLegacyProfiles(ctx, nil, nil) }}}
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
