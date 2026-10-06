package workspace_agents

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/workspace_agents/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, n string) *T {
	t.Helper()
	var r T
	require.NoError(t, json.Unmarshal(mocks.Fixture(n), &r))
	return &r
}

type contract struct {
	name, method, path string
	body, raw, empty   bool
	call               func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	ctx := context.Background()
	return []contract{{name: "ListSkills", method: "GET", path: "/apigateway/nlp/assist-skills/api/v1?source=user", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.ListSkills(ctx, "user") }},
		{name: "CheckSkillAvailability", method: "GET", path: "/apigateway/nlp/assist-skills/api/v1?check-availability=system", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.CheckSkillAvailability(ctx, "system") }},
		{name: "GetSkill", method: "GET", path: "/apigateway/nlp/assist-skills/api/v1/fixture-id", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) { return s.GetSkill(ctx, "fixture-id") }},
		{name: "CreateSkill", method: "POST", path: "/apigateway/nlp/assist-skills/api/v1", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateSkill(ctx, load[SkillRequest](t, "CreateSkill_request"))
		}},
		{name: "UpdateSkill", method: "PUT", path: "/apigateway/nlp/assist-skills/api/v1/fixture-id", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateSkill(ctx, "fixture-id", load[SkillRequest](t, "UpdateSkill_request"))
		}},
		{name: "DeleteSkill", method: "DELETE", path: "/apigateway/nlp/assist-skills/api/v1/fixture-id", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.DeleteSkill(ctx, "fixture-id")
			return nil, r, e
		}},
		{name: "UploadSkillFile", method: "POST", path: "/apigateway/nlp/assist-skills/api/v1/skills/fixture-id/knowledgebase/files?filename=fixture.txt&mime_type=text%2Fplain", body: true, raw: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UploadSkillFile(ctx, "fixture-id", load[FileUploadRequest](t, "UploadSkillFile_input"))
		}},
		{name: "DeleteSkillFile", method: "DELETE", path: "/apigateway/nlp/assist-skills/api/v1/skills/fixture-id/knowledgebase/files/file-id", body: false, raw: false, empty: true, call: func(s *Service) (any, *interfaces.Response, error) {
			r, e := s.DeleteSkillFile(ctx, "fixture-id", "file-id")
			return nil, r, e
		}},
		{name: "StartSkillMultipartUpload", method: "POST", path: "/apigateway/nlp/assist-skills/api/v1/skills/fixture-id/knowledgebase/files/multipart?filename=fixture.txt&mime_type=text%2Fplain", body: false, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.StartSkillMultipartUpload(ctx, "fixture-id", load[StartMultipartRequest](t, "StartSkillMultipartUpload_input"))
		}},
		{name: "UploadSkillPart", method: "PUT", path: "/apigateway/nlp/assist-skills/api/v1/skills/fixture-id/knowledgebase/files/multipart?fileId=file-id&partNumber=1&uploadId=upload-id", body: true, raw: true, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UploadSkillPart(ctx, "fixture-id", load[UploadPartRequest](t, "UploadSkillPart_input"))
		}},
		{name: "CompleteSkillMultipartUpload", method: "POST", path: "/apigateway/nlp/assist-skills/api/v1/skills/fixture-id/knowledgebase/files/multipart/complete?fileId=file-id&uploadId=upload-id", body: true, raw: false, empty: false, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CompleteSkillMultipartUpload(ctx, "fixture-id", load[CompleteMultipartRequest](t, "CompleteSkillMultipartUpload_input"))
		}}}
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
					if tt.raw {
						var expected string
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_request"), &expected))
						assert.Equal(t, expected, string(body))
						assert.Equal(t, "application/octet-stream", r.Header.Get("Content-Type"))
					} else {
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
					}
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
				actual, e := json.Marshal(result)
				require.NoError(t, e)
				assert.JSONEq(t, string(expected), string(actual))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}

func TestHTTPErrorsPreserveResponse(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, httpmock.NewStringResponder(403, string(mocks.Fixture("error"))))
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 403, response.StatusCode)
			assert.JSONEq(t, string(mocks.Fixture("error")), string(response.Body))
		})
	}
}
func TestMalformedJSONResponses(t *testing.T) {
	for _, tt := range contracts(t) {
		if tt.empty {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				response := httpmock.NewStringResponse(200, "{broken")
				response.Header.Set("Content-Type", "application/json")
				return response, nil
			})
			_, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
		})
	}
}
