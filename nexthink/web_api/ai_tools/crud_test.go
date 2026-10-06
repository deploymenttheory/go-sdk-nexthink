package ai_tools

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/ai_tools/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func TestContracts(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/application", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "List_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).List(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("List_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("Get", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/application/11111111-1111-4111-8111-111111111111", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "Get_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).Get(context.Background(), "11111111-1111-4111-8111-111111111111")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Get_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("Create", func(t *testing.T) {
		var request ToolRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Create_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/application", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Create_request")), string(b))
				fixture := "Create_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).Create(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Create_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("Update", func(t *testing.T) {
		var request ToolRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("Update_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("PUT", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/application/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Update_request")), string(b))
				fixture := "Update_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).Update(context.Background(), "11111111-1111-4111-8111-111111111111", 3, &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("Update_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("Delete", func(t *testing.T) {
		for _, status := range []int{204, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("DELETE", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/application/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				if status == 403 {
					return mocks.Responder(status, "error")(r)
				}
				return &http.Response{StatusCode: 204, Header: make(http.Header), Body: http.NoBody}, nil
			})
			response, err := NewService(transport).Delete(context.Background(), "11111111-1111-4111-8111-111111111111", 3)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetCopilot", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/ms-copilot/11111111-1111-4111-8111-111111111111", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetCopilot_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetCopilot(context.Background(), "11111111-1111-4111-8111-111111111111")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetCopilot_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("CreateCopilot", func(t *testing.T) {
		var request CopilotRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("CreateCopilot_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/ms-copilot", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateCopilot_request")), string(b))
				fixture := "CreateCopilot_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).CreateCopilot(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateCopilot_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("UpdateCopilot", func(t *testing.T) {
		var request CopilotRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("UpdateCopilot_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("PUT", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/ms-copilot/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateCopilot_request")), string(b))
				fixture := "UpdateCopilot_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).UpdateCopilot(context.Background(), "11111111-1111-4111-8111-111111111111", 3, &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateCopilot_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("DeleteCopilot", func(t *testing.T) {
		for _, status := range []int{204, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("DELETE", testutil.BaseURL+"/apigateway/aidex/config/v2/aitools/ms-copilot/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				if status == 403 {
					return mocks.Responder(status, "error")(r)
				}
				return &http.Response{StatusCode: 204, Header: make(http.Header), Body: http.NoBody}, nil
			})
			response, err := NewService(transport).DeleteCopilot(context.Background(), "11111111-1111-4111-8111-111111111111", 3)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("ListGoals", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/goal", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "ListGoals_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).ListGoals(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("ListGoals_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetGoal", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/goal/11111111-1111-4111-8111-111111111111", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetGoal_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetGoal(context.Background(), "11111111-1111-4111-8111-111111111111")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetGoal_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("CreateGoal", func(t *testing.T) {
		var request GoalRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("CreateGoal_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/config/v1/goal", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateGoal_request")), string(b))
				fixture := "CreateGoal_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).CreateGoal(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateGoal_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("UpdateGoal", func(t *testing.T) {
		var request GoalRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("UpdateGoal_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("PUT", testutil.BaseURL+"/apigateway/aidex/config/v1/goal/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateGoal_request")), string(b))
				fixture := "UpdateGoal_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).UpdateGoal(context.Background(), "11111111-1111-4111-8111-111111111111", 3, &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateGoal_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("DeleteGoal", func(t *testing.T) {
		for _, status := range []int{204, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("DELETE", testutil.BaseURL+"/apigateway/aidex/config/v1/goal/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				if status == 403 {
					return mocks.Responder(status, "error")(r)
				}
				return &http.Response{StatusCode: 204, Header: make(http.Header), Body: http.NoBody}, nil
			})
			response, err := NewService(transport).DeleteGoal(context.Background(), "11111111-1111-4111-8111-111111111111", 3)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetLegacyTool", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/tool/11111111-1111-4111-8111-111111111111", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetLegacyTool_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetLegacyTool(context.Background(), "11111111-1111-4111-8111-111111111111")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetLegacyTool_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("ListSystemTools", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/tool/system-tools", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "ListSystemTools_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).ListSystemTools(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("ListSystemTools_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetLicense", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/aitools/license", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetLicense_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetLicense(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetLicense_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetRedirectURLs", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/tool/redirect-urls", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetRedirectURLs_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetRedirectURLs(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetRedirectURLs_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetModule", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/config/v1/module", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetModule_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetModule(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetModule_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("CreateModule", func(t *testing.T) {
		var request Module
		require.NoError(t, json.Unmarshal(mocks.Fixture("CreateModule_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/config/v1/module", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateModule_request")), string(b))
				fixture := "CreateModule_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).CreateModule(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CreateModule_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("UpdateModule", func(t *testing.T) {
		var request Module
		require.NoError(t, json.Unmarshal(mocks.Fixture("UpdateModule_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("PUT", testutil.BaseURL+"/apigateway/aidex/config/v1/module/11111111-1111-4111-8111-111111111111?rev=3", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateModule_request")), string(b))
				fixture := "UpdateModule_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).UpdateModule(context.Background(), "11111111-1111-4111-8111-111111111111", 3, &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("UpdateModule_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("CheckCopilotCredentials", func(t *testing.T) {
		var request CredentialsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("CheckCopilotCredentials_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/check-credentials/ms-copilot", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CheckCopilotCredentials_request")), string(b))
				fixture := "CheckCopilotCredentials_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).CheckCopilotCredentials(context.Background(), &request)
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("CheckCopilotCredentials_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetOverviewInsights", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/insights/overview?lang=ja", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetOverviewInsights_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetOverviewInsights(context.Background(), "ja")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetOverviewInsights_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetToolInsights", func(t *testing.T) {
		var request ToolInsightsRequest
		require.NoError(t, json.Unmarshal(mocks.Fixture("GetToolInsights_request"), &request))
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("POST", testutil.BaseURL+"/apigateway/aidex/insights/tool?lang=ja", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				b, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetToolInsights_request")), string(b))
				fixture := "GetToolInsights_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetToolInsights(context.Background(), &request, "ja")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetToolInsights_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetGovernanceTrends", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/insights/observability/governance/tool-governance-trends", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetGovernanceTrends_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetGovernanceTrends(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetGovernanceTrends_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetGovernanceActiveUsers", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/insights/observability/governance/tool-governance-active-users", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetGovernanceActiveUsers_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetGovernanceActiveUsers(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetGovernanceActiveUsers_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetGovernanceDashboard", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/insights/observability/governance/tool-governance-dashboard", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetGovernanceDashboard_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetGovernanceDashboard(context.Background())
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetGovernanceDashboard_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
	t.Run("GetGoalInsights", func(t *testing.T) {
		for _, status := range []int{200, 403} {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder("GET", testutil.BaseURL+"/apigateway/aidex/goals/v1/goals/11111111-1111-4111-8111-111111111111/insights", func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "application/json", r.Header.Get("Accept"))
				fixture := "GetGoalInsights_success"
				if status == 403 {
					fixture = "error"
				}
				return mocks.Responder(status, fixture)(r)
			})
			result, response, err := NewService(transport).GetGoalInsights(context.Background(), "11111111-1111-4111-8111-111111111111")
			require.NotNil(t, response)
			assert.Equal(t, status, response.StatusCode)
			if status == 403 {
				require.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				b, err := json.Marshal(result)
				require.NoError(t, err)
				assert.JSONEq(t, string(mocks.Fixture("GetGoalInsights_success")), string(b))
			}
			assert.Equal(t, 1, mock.GetTotalCallCount())
		}
	})
}
func TestValidationPreventsRequests(t *testing.T) {
	transport, mock := testutil.NewTransport(t)
	s := NewService(transport)
	var err error
	var response *interfaces.Response
	_, response, err = s.Get(context.Background(), "../invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.Update(context.Background(), "../invalid", -1, nil)
	require.Error(t, err)
	assert.Nil(t, response)
	response, err = s.Delete(context.Background(), "../invalid", -1)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetCopilot(context.Background(), "../invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.CreateCopilot(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.UpdateCopilot(context.Background(), "../invalid", -1, nil)
	require.Error(t, err)
	assert.Nil(t, response)
	response, err = s.DeleteCopilot(context.Background(), "../invalid", -1)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetGoal(context.Background(), "../invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.CreateGoal(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.UpdateGoal(context.Background(), "../invalid", -1, nil)
	require.Error(t, err)
	assert.Nil(t, response)
	response, err = s.DeleteGoal(context.Background(), "../invalid", -1)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetLegacyTool(context.Background(), "../invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.CreateModule(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.UpdateModule(context.Background(), "../invalid", -1, nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.CheckCopilotCredentials(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetOverviewInsights(context.Background(), "invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetToolInsights(context.Background(), nil, "invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	_, response, err = s.GetGoalInsights(context.Background(), "../invalid")
	require.Error(t, err)
	assert.Nil(t, response)
	assert.Zero(t, mock.GetTotalCallCount())
}
