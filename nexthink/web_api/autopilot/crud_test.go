package autopilot

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/autopilot/mocks"
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

type contract struct {
	name, method, path   string
	body, empty, graphql bool
	call                 func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	ctx := context.Background()
	return []contract{
		{name: "GetConfiguration", method: "GET", path: Endpoint + "/cockpit/configuration", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetConfiguration(ctx) }},
		{name: "UpdateWebSearch", method: "PATCH", path: Endpoint + "/cockpit/configuration/websearch", body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateWebSearch(ctx, load[BooleanValue](t, "UpdateWebSearch_request"))
			return nil, response, err
		}},
		{name: "UpdateFilesystemTool", method: "PATCH", path: Endpoint + "/cockpit/configuration/filesystemtool", body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateFilesystemTool(ctx, load[BooleanValue](t, "UpdateFilesystemTool_request"))
			return nil, response, err
		}},
		{name: "UpdateAgentName", method: "PATCH", path: Endpoint + "/cockpit/configuration/agentname", body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateAgentName(ctx, load[StringValue](t, "UpdateAgentName_request"))
			return nil, response, err
		}},
		{name: "GetSettings", method: "GET", path: Endpoint + "/cockpit/settings", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetSettings(ctx) }},
		{name: "SaveSettings", method: "POST", path: Endpoint + "/cockpit/settings", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.SaveSettings(ctx, load[Settings](t, "SaveSettings_request"))
		}},
		{name: "GetWebSearchDomains", method: "GET", path: Endpoint + "/cockpit/settings/websearch/domains", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetWebSearchDomains(ctx) }},
		{name: "ReplaceWebSearchDomains", method: "PUT", path: Endpoint + "/cockpit/settings/websearch/domains", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.ReplaceWebSearchDomains(ctx, load[WebSearchDomainsRequest](t, "ReplaceWebSearchDomains_request"))
		}},
		{name: "GetApproval", method: "GET", path: Endpoint + "/cockpit/approval/" + "fixture-id", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetApproval(ctx, "fixture-id") }},
		{name: "CreateApproval", method: "POST", path: Endpoint + "/cockpit/approval", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateApproval(ctx, load[Approval](t, "CreateApproval_request"))
		}},
		{name: "UpdateApproval", method: "PATCH", path: Endpoint + "/cockpit/approval/" + "fixture-id", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateApproval(ctx, "fixture-id", load[ApprovalInput](t, "UpdateApproval_request"))
		}},
		{name: "ListCalls", method: "GET", path: Endpoint + "/cockpit/calls", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListCalls(ctx) }},
		{name: "GetKnowledgeArticleCount", method: "GET", path: Endpoint + "/cockpit/knowledge-base/articles/count", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetKnowledgeArticleCount(ctx) }},
		{name: "ListKnowledgeConnectors", method: "GET", path: Endpoint + "/cockpit/knowledge-base/connectors", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListKnowledgeConnectors(ctx) }},
		{name: "CreateTicket", method: "POST", path: ITSMEndpoint + "/cockpit/tickets", body: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateTicket(ctx, load[TicketRequest](t, "CreateTicket_request"))
		}},
		{name: "GetConversation", method: "GET", path: Endpoint + "/cockpit/conversations/" + "fixture-id", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetConversation(ctx, "fixture-id") }},
		{name: "GetRecommendationConversationIDs", method: "GET", path: RecommendationsEndpoint + "/knowledge-recommendations/" + "fixture-id" + "/conversation-ids", body: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.GetRecommendationConversationIDs(ctx, "fixture-id")
		}},
		{name: "UpdateAgentActionInputs", method: "PATCH", path: AgentActionsEndpoint + "/" + "fixture-id", body: true, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			response, err := s.UpdateAgentActionInputs(ctx, "fixture-id", load[AgentActionInputsRequest](t, "UpdateAgentActionInputs_request"))
			return nil, response, err
		}},
		{name: "GetAgentAction", method: "POST", path: "/apigateway/act/manage/graphql", body: true, graphql: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetAgentAction(ctx, "fixture-id") }},
		{name: "GetAgentActionInputs", method: "POST", path: "/apigateway/act/manage/graphql", body: true, graphql: true, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetAgentActionInputs(ctx, "fixture-id") }},
	}
}
func TestContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.body {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				}
				status, body := 204, ""
				if !tt.empty {
					status = 200
					body = string(mocks.Fixture(tt.name + "_success"))
				}
				response := httpmock.NewStringResponse(status, body)
				response.Header.Set("Content-Type", "application/json")
				response.Header.Set("X-Request-ID", "fixture-request")
				return response, nil
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			if tt.empty {
				assert.Nil(t, result)
			} else {
				require.NotNil(t, result)
				expected := mocks.Fixture(tt.name + "_success")
				if tt.graphql {
					var envelope struct {
						Data struct {
							AgentAction json.RawMessage `json:"agentActionByUid"`
						} `json:"data"`
					}
					require.NoError(t, json.Unmarshal(expected, &envelope))
					expected = envelope.Data.AgentAction
				}
				actual, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(expected), string(actual))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestFailures(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			for _, status := range []int{401, 403, 409, 429, 500} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, httpmock.NewStringResponder(status, string(mocks.Fixture("error"))))
				_, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, status, response.StatusCode)
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
		})
	}
}
func TestDownloadPreservesCSVAndFilename(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	var fixture struct{ CSV, Filename string }
	require.NoError(t, json.Unmarshal(mocks.Fixture("DownloadCategorization_success"), &fixture))
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/cockpit/settings/categorization/csv", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, "text/csv", r.Header.Get("Accept"))
		response := httpmock.NewStringResponse(200, fixture.CSV)
		response.Header.Set("Content-Type", "text/csv")
		response.Header.Set("Content-Disposition", "attachment; filename*=UTF-8''fixture%20categories.csv")
		return response, nil
	})
	result, response, err := NewService(transport).DownloadCategorization(context.Background())
	require.NoError(t, err)
	assert.Equal(t, fixture.Filename, result.Filename)
	assert.Equal(t, fixture.CSV, string(result.Data))
	assert.Equal(t, result.Data, response.Body)
}
func TestGraphQLPartialData(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/act/manage/graphql", func(r *http.Request) (*http.Response, error) {
		response := httpmock.NewStringResponse(200, `{"data":{"agentActionByUid":{"contentId":"fixture-id","scriptInfo":null}},"errors":[{"message":"partial permission denied"}]}`)
		response.Header.Set("Content-Type", "application/json")
		return response, nil
	})
	result, response, err := NewService(transport).GetAgentAction(context.Background(), "fixture-id")
	require.Error(t, err)
	require.NotNil(t, response)
	require.NotNil(t, result)
	assert.Equal(t, "fixture-id", *result.ContentID)
}
func TestEscaping(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	mock.RegisterResponder("GET", testutil.BaseURL+Endpoint+"/cockpit/approval/fixture%2Fid%3Fx", func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, Endpoint+"/cockpit/approval/fixture%2Fid%3Fx", r.URL.EscapedPath())
		response := httpmock.NewStringResponse(200, string(mocks.Fixture("GetApproval_success")))
		response.Header.Set("Content-Type", "application/json")
		return response, nil
	})
	_, _, err := NewService(transport).GetApproval(context.Background(), "fixture/id?x")
	require.NoError(t, err)
}
