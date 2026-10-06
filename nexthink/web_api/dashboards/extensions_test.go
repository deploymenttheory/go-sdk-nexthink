package dashboards

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/dashboards/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func extensionLoad[T any](t *testing.T, name string) *T {
	t.Helper()
	var value T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &value))
	return &value
}

type extensionCase struct {
	name, verb, path, mode string
	status                 int
	hasBody, raw, graph    bool
	call                   func(*Service) (any, *interfaces.Response, error)
}

func extensionCases(t *testing.T) []extensionCase {
	t.Helper()
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	executionID := "22222222-2222-4333-8444-555555555555"
	_ = id
	_ = executionID
	return []extensionCase{{name: "ExtensionCreateWidget", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
		return s.CreateWidget(ctx, extensionLoad[WidgetRequest](t, "ExtensionCreateWidget_input"))
	}},
		{name: "ExtensionUpdateWidget", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateWidget(ctx, extensionLoad[WidgetRequest](t, "ExtensionUpdateWidget_input"))
		}},
		{name: "ExtensionDeleteWidget", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteWidget(ctx, extensionLoad[DeleteWidgetRequest](t, "ExtensionDeleteWidget_input"))
		}},
		{name: "ExtensionCreateFilter", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateFilter(ctx, extensionLoad[FilterRequest](t, "ExtensionCreateFilter_input"))
		}},
		{name: "ExtensionUpdateFilter", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateFilter(ctx, extensionLoad[FilterRequest](t, "ExtensionUpdateFilter_input"))
		}},
		{name: "ExtensionDeleteFilter", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteFilter(ctx, extensionLoad[DeleteFilterRequest](t, "ExtensionDeleteFilter_input"))
		}},
		{name: "ExtensionCreateTab", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.CreateTab(ctx, extensionLoad[MutationContext](t, "ExtensionCreateTab_input"))
		}},
		{name: "ExtensionUpdateTab", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateTab(ctx, extensionLoad[UpdateTabRequest](t, "ExtensionUpdateTab_input"))
		}},
		{name: "ExtensionUpdateTabs", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateTabs(ctx, extensionLoad[UpdateTabsRequest](t, "ExtensionUpdateTabs_input"))
		}},
		{name: "ExtensionDeleteTab", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.DeleteTab(ctx, extensionLoad[MutationContext](t, "ExtensionDeleteTab_input"))
		}},
		{name: "ExtensionUpdateLayout", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.UpdateLayout(ctx, extensionLoad[UpdateLayoutRequest](t, "ExtensionUpdateLayout_input"))
		}},
		{name: "ExtensionExport", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Export(ctx, extensionLoad[MutationContext](t, "ExtensionExport_input"))
		}},
		{name: "ExtensionDuplicate", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Duplicate(ctx, extensionLoad[DuplicateRequest](t, "ExtensionDuplicate_input"))
		}},
		{name: "ExtensionImport", verb: "POST", path: "/apigateway/dash/graphql", status: 200, mode: "json", hasBody: true, raw: false, graph: true, call: func(s *Service) (any, *interfaces.Response, error) {
			return s.Import(ctx, extensionLoad[ImportRequest](t, "ExtensionImport_input"))
		}}}
}
func TestExtensionWireContracts(t *testing.T) {
	for _, tt := range extensionCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				if tt.hasBody {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					if tt.raw {
						var expected string
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_request"), &expected))
						assert.Equal(t, expected, string(body))
						assert.Equal(t, "application/octet-stream", r.Header.Get("Content-Type"))
					} else {
						assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
						assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
					}
				}
				if tt.mode != "json" {
					body := ""
					if tt.mode == "text" {
						require.NoError(t, json.Unmarshal(mocks.Fixture(tt.name+"_success"), &body))
					}
					res := httpmock.NewStringResponse(tt.status, body)
					res.Header.Set("Content-Type", "text/plain")
					res.Header.Set("X-Request-ID", "fixture-request")
					return res, nil
				}
				return mocks.Responder(tt.status, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, tt.status, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			assert.Equal(t, 1, mock.GetTotalCallCount())
			if tt.mode == "empty" {
				assert.Nil(t, result)
				return
			}
			expected := mocks.Fixture(tt.name + "_success")
			if tt.mode == "text" {
				var message string
				require.NoError(t, json.Unmarshal(expected, &message))
				expected, _ = json.Marshal(map[string]string{"message": message})
			}
			if tt.graph {
				var envelope struct {
					Data json.RawMessage `json:"data"`
				}
				require.NoError(t, json.Unmarshal(expected, &envelope))
				expected = envelope.Data
			}
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
func TestExtensionFailures(t *testing.T) {
	for _, tt := range extensionCases(t) {
		t.Run(tt.name, func(t *testing.T) {
			for code, name := range map[int]string{400: "validation", 401: "unauthorized", 403: "forbidden"} {
				transport, mock := testutil.NewTransport(t)
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, mocks.Responder(code, "extension_error_"+name))
				_, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				assert.Equal(t, code, response.StatusCode)
				assert.JSONEq(t, string(mocks.Fixture("extension_error_"+name)), string(response.Body))
			}
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewErrorResponder(io.ErrUnexpectedEOF))
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
			if tt.mode == "json" {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
					res := httpmock.NewStringResponse(200, "{broken")
					res.Header.Set("Content-Type", "application/json")
					return res, nil
				})
				_, _, err = tt.call(NewService(transport))
				require.Error(t, err)
			}
			if tt.graph {
				mock.RegisterResponder(tt.verb, testutil.BaseURL+tt.path, httpmock.NewJsonResponderOrPanic(200, map[string]any{"data": map[string]any{}, "errors": []map[string]any{{"message": "fixture conflict"}}}))
				result, response, err := tt.call(NewService(transport))
				require.Error(t, err)
				require.NotNil(t, response)
				require.NotNil(t, result)
				assert.Equal(t, 200, response.StatusCode)
			}
		})
	}
}
func TestExtensionValidationBeforeTransport(t *testing.T) {
	ctx := context.Background()
	id := "11111111-2222-4333-8444-555555555555"
	_ = ctx
	_ = id
	cases := []struct {
		name string
		call func(*Service) (any, *interfaces.Response, error)
	}{{name: "CreateWidget", call: func(s *Service) (any, *interfaces.Response, error) { return s.CreateWidget(ctx, nil) }},
		{name: "UpdateWidget", call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateWidget(ctx, nil) }},
		{name: "DeleteWidget", call: func(s *Service) (any, *interfaces.Response, error) { return s.DeleteWidget(ctx, nil) }},
		{name: "CreateFilter", call: func(s *Service) (any, *interfaces.Response, error) { return s.CreateFilter(ctx, nil) }},
		{name: "UpdateFilter", call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateFilter(ctx, nil) }},
		{name: "DeleteFilter", call: func(s *Service) (any, *interfaces.Response, error) { return s.DeleteFilter(ctx, nil) }},
		{name: "CreateTab", call: func(s *Service) (any, *interfaces.Response, error) { return s.CreateTab(ctx, nil) }},
		{name: "UpdateTab", call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateTab(ctx, nil) }},
		{name: "UpdateTabs", call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateTabs(ctx, nil) }},
		{name: "DeleteTab", call: func(s *Service) (any, *interfaces.Response, error) { return s.DeleteTab(ctx, nil) }},
		{name: "UpdateLayout", call: func(s *Service) (any, *interfaces.Response, error) { return s.UpdateLayout(ctx, nil) }},
		{name: "Export", call: func(s *Service) (any, *interfaces.Response, error) { return s.Export(ctx, nil) }},
		{name: "Duplicate", call: func(s *Service) (any, *interfaces.Response, error) { return s.Duplicate(ctx, nil) }},
		{name: "Import", call: func(s *Service) (any, *interfaces.Response, error) { return s.Import(ctx, nil) }}}
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
