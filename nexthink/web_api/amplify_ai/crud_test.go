package amplify_ai

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/amplify_ai/mocks"
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
	cases := []struct {
		name, verb, path string
		body, empty      bool
		call             func(*Service) (any, *interfaces.Response, error)
		invalid          func(*Service) (any, *interfaces.Response, error)
	}{{name: "GenerateOrFetchAnalysis", verb: "POST", path: EndpointGenerateOrFetchAnalysis, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GenerateOrFetchAnalysis(ctx, load[AnalysisRequest](t, "GenerateOrFetchAnalysis_request"))
	}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GenerateOrFetchAnalysis(ctx, nil) }},
		{name: "SubmitFeedback", verb: "POST", path: EndpointSubmitFeedback, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.SubmitFeedback(ctx, load[FeedbackRequest](t, "SubmitFeedback_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.SubmitFeedback(ctx, nil) }},
		{name: "PostMetric", verb: "POST", path: EndpointPostMetric, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.PostMetric(ctx, load[MetricRequest](t, "PostMetric_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { r, e := s.PostMetric(ctx, nil); return nil, r, e }},
		{name: "ReportTicketRetrievalDuration", verb: "POST", path: EndpointReportTicketRetrievalDuration, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ReportTicketRetrievalDuration(ctx, load[TicketRetrievalDurationRequest](t, "ReportTicketRetrievalDuration_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ReportTicketRetrievalDuration(ctx, nil)
			return nil, r, e
		}},
		{name: "ExecuteAction", verb: "POST", path: EndpointExecuteAction, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ExecuteAction(ctx, load[ExecuteActionRequest](t, "ExecuteAction_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ExecuteAction(ctx, nil)
			return nil, r, e
		}},
		{name: "ExecuteUserAction", verb: "POST", path: EndpointExecuteUserAction, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ExecuteUserAction(ctx, load[ExecuteUserActionRequest](t, "ExecuteUserAction_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ExecuteUserAction(ctx, nil)
			return nil, r, e
		}},
		{name: "GetResolutionPlan", verb: "GET", path: EndpointGetResolutionPlan + "?resolution_plan_id=" + id, body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetResolutionPlan(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetResolutionPlan(ctx, "") }},
		{name: "RefreshResolutionStep", verb: "POST", path: EndpointRefreshResolutionStep, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.RefreshResolutionStep(ctx, load[RefreshResolutionStepRequest](t, "RefreshResolutionStep_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.RefreshResolutionStep(ctx, nil) }},
		{name: "UpdateResolutionStep", verb: "POST", path: EndpointUpdateResolutionStep, body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateResolutionStep(ctx, load[UpdateResolutionStepRequest](t, "UpdateResolutionStep_request"))
		}, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateResolutionStep(ctx, nil) }},
		{name: "GetMockMetadata", verb: "GET", path: EndpointGetMockMetadata + "?ticket_sys_id=" + id, body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetMockMetadata(ctx, id) }, invalid: func(s *Service) (any, *interfaces.Response, error) { return s.GetMockMetadata(ctx, "") }},
		{name: "ResolveTicket", verb: "POST", path: EndpointResolveTicket, body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ResolveTicket(ctx, load[ResolveTicketRequest](t, "ResolveTicket_request"))
			return nil, r, e
		}, invalid: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.ResolveTicket(ctx, nil)
			return nil, r, e
		}}}
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
				assert.JSONEq(t, string(expected), string(encoded))
			}
			for _, status := range []int{400, 401, 403, 409, 422, 500} {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(status, "error"))
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
