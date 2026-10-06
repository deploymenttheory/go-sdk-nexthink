package support_timeline

import (
	"context"
	"encoding/json"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/interfaces"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/testutil"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/support_timeline/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func load[T any](t *testing.T, name string) *T {
	t.Helper()
	var v T
	require.NoError(t, json.Unmarshal(mocks.Fixture(name), &v))
	return &v
}

type contract struct {
	name, method, path, query string
	call                      func(*Service) (any, *interfaces.Response, error)
}

func contracts(t *testing.T) []contract {
	t.Helper()
	ctx := context.Background()
	return []contract{{name: "GetAlertsAndErrors", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/timeline/alertsAndErrors", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetAlertsAndErrors(ctx, "fixture-deviceID", load[TimeRange](t, "GetAlertsAndErrors_request"))
	}}, {name: "GetPerformance", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/timeline/performance", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetPerformance(ctx, "fixture-deviceID", load[TimeRange](t, "GetPerformance_request"))
	}}, {name: "GetConnectivity", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v5/device/fixture-deviceID/timeline/connectivity", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetConnectivity(ctx, "fixture-deviceID", load[TimeRange](t, "GetConnectivity_request"))
	}}, {name: "GetApplicationConnectivity", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v5/device/fixture-deviceID/timeline/connectivityApplications", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetApplicationConnectivity(ctx, "fixture-deviceID", load[TimeRange](t, "GetApplicationConnectivity_request"))
	}}, {name: "GetActivity", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v5/device/fixture-deviceID/timeline/activity", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetActivity(ctx, "fixture-deviceID", load[TimeRange](t, "GetActivity_request"))
	}}, {name: "GetApplications", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v5/device/fixture-deviceID/timeline/appex", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetApplications(ctx, "fixture-deviceID", load[TimeRange](t, "GetApplications_request"))
	}}, {name: "GetUserInteractions", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/timeline/usersInteractions", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetUserInteractions(ctx, "fixture-deviceID", load[TimeRange](t, "GetUserInteractions_request"))
	}}, {name: "GetCollaboration", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/timeline/collaboration", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetCollaboration(ctx, "fixture-deviceID", load[TimeRange](t, "GetCollaboration_request"))
	}}, {name: "GetErrorsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/errors", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetErrorsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetErrorsDrilldown_request"))
	}}, {name: "GetFreezesDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/freezes", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetFreezesDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetFreezesDrilldown_request"))
	}}, {name: "GetAlertsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/alerts", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetAlertsDrilldown(ctx, "fixture-deviceID", load[AlertsRequest](t, "GetAlertsDrilldown_request"))
	}}, {name: "GetActionsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/actions", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetActionsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetActionsDrilldown_request"))
	}}, {name: "GetSystemBootsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/systemboots", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetSystemBootsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetSystemBootsDrilldown_request"))
	}}, {name: "GetSystemBootsAndSuspendsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/systemboots", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetSystemBootsAndSuspendsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetSystemBootsAndSuspendsDrilldown_request"))
	}}, {name: "GetTeamsCallsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/msteams", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetTeamsCallsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetTeamsCallsDrilldown_request"))
	}}, {name: "GetZoomCallsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/zoom", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetZoomCallsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetZoomCallsDrilldown_request"))
	}}, {name: "GetDesktopConnectivityDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/connectivity/applications/desktop", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetDesktopConnectivityDrilldown(ctx, "fixture-deviceID", load[ApplicationRequest](t, "GetDesktopConnectivityDrilldown_request"))
	}}, {name: "GetWebConnectivityDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/connectivity/applications/web", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetWebConnectivityDrilldown(ctx, "fixture-deviceID", load[ApplicationRequest](t, "GetWebConnectivityDrilldown_request"))
	}}, {name: "GetDesktopApplicationsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/desktopapplications", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetDesktopApplicationsDrilldown(ctx, "fixture-deviceID", load[ApplicationRequest](t, "GetDesktopApplicationsDrilldown_request"))
	}}, {name: "GetWebApplicationsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v3/device/fixture-deviceID/drilldowns/webapplications", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetWebApplicationsDrilldown(ctx, "fixture-deviceID", load[ApplicationRequest](t, "GetWebApplicationsDrilldown_request"))
	}}, {name: "GetInstallationsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/installationevents", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetInstallationsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetInstallationsDrilldown_request"))
	}}, {name: "GetEthernetDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/ethernet", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetEthernetDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetEthernetDrilldown_request"))
	}}, {name: "GetWiFiDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v2/device/fixture-deviceID/drilldowns/wifi", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetWiFiDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetWiFiDrilldown_request"))
	}}, {name: "GetConnectionsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/connectionevents", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetConnectionsDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetConnectionsDrilldown_request"))
	}}, {name: "GetNetworkApplicationDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/networkapp", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetNetworkApplicationDrilldown(ctx, "fixture-deviceID", load[ApplicationRequest](t, "GetNetworkApplicationDrilldown_request"))
	}}, {name: "GetCPUDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/cpu", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetCPUDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetCPUDrilldown_request"))
	}}, {name: "GetMemoryDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/memory", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetMemoryDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetMemoryDrilldown_request"))
	}}, {name: "GetDiskPerformanceDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/diskperformance", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetDiskPerformanceDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetDiskPerformanceDrilldown_request"))
	}}, {name: "GetDriveSpaceDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/drivespace", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetDriveSpaceDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetDriveSpaceDrilldown_request"))
	}}, {name: "GetGPUDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/gpu/fixture-slot", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetGPUDrilldown(ctx, "fixture-deviceID", "fixture-slot", load[TimeRange](t, "GetGPUDrilldown_request"))
	}}, {name: "GetNPUDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/drilldowns/npu", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetNPUDrilldown(ctx, "fixture-deviceID", load[TimeRange](t, "GetNPUDrilldown_request"))
	}}, {name: "GetUserInteractionsDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v3/device/fixture-deviceID/drilldowns/user/fixture-userID/userinteractions", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetUserInteractionsDrilldown(ctx, "fixture-deviceID", "fixture-userID", load[TimeRange](t, "GetUserInteractionsDrilldown_request"))
	}}, {name: "GetRoundTripTimeDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/user/fixture-userID/drilldowns/rtt", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetRoundTripTimeDrilldown(ctx, "fixture-deviceID", "fixture-userID", load[TimeRange](t, "GetRoundTripTimeDrilldown_request"))
	}}, {name: "GetNetworkLatencyDrilldown", method: "POST", path: "/apigateway/atl/support-device-timeline-be/api/v1/device/fixture-deviceID/user/fixture-userID/drilldowns/networklatency", query: "{}", call: func(s *Service) (any, *interfaces.Response, error) {
		return s.GetNetworkLatencyDrilldown(ctx, "fixture-deviceID", "fixture-userID", load[TimeRange](t, "GetNetworkLatencyDrilldown_request"))
	}}}
}
func TestWireContracts(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(r *http.Request) (*http.Response, error) {
				assert.Equal(t, "Bearer fixture-token", r.Header.Get("Authorization"))
				assert.Equal(t, "Europe/London", r.Header.Get("time-zone"))
				assert.Equal(t, "-60", r.Header.Get("utc-offset"))
				var query map[string]string
				require.NoError(t, json.Unmarshal([]byte(tt.query), &query))
				actual := map[string]string{}
				for k := range r.URL.Query() {
					actual[k] = r.URL.Query().Get(k)
				}
				assert.Equal(t, query, actual)
				if tt.method == "POST" {
					body, err := io.ReadAll(r.Body)
					require.NoError(t, err)
					assert.JSONEq(t, string(mocks.Fixture(tt.name+"_request")), string(body))
				}
				return mocks.Responder(200, tt.name+"_success")(r)
			})
			result, response, err := tt.call(NewService(transport, WithTimeZone("Europe/London", -60)))
			require.NoError(t, err)
			require.NotNil(t, response)
			assert.Equal(t, 200, response.StatusCode)
			assert.Equal(t, "fixture-request", response.Headers.Get("X-Request-ID"))
			actual, err := json.Marshal(result)
			require.NoError(t, err)
			assert.JSONEq(t, string(mocks.Fixture(tt.name+"_success")), string(actual))
			assert.Equal(t, 1, mock.GetTotalCallCount())
		})
	}
}
func TestErrors(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, mocks.Responder(403, "error"))
			result, response, err := tt.call(NewService(transport))
			require.Error(t, err)
			assert.Nil(t, result)
			require.NotNil(t, response)
			assert.Equal(t, 403, response.StatusCode)
		})
	}
}
func TestMalformedSuccess(t *testing.T) {
	for _, tt := range contracts(t) {
		t.Run(tt.name, func(t *testing.T) {
			transport, mock := testutil.NewTransport(t)
			mock.RegisterResponder(tt.method, testutil.BaseURL+tt.path, func(_ *http.Request) (*http.Response, error) {
				r := httpmock.NewStringResponse(200, "{broken")
				r.Header.Set("Content-Type", "application/json")
				return r, nil
			})
			_, _, err := tt.call(NewService(transport))
			require.Error(t, err)
		})
	}
}
func TestInvalidIdentifier(t *testing.T) {
	for _, id := range []string{"", " ", "..", "a/b", "a?b", "a#b", "a\nb"} {
		assert.Error(t, validateID(id))
	}
	assert.NoError(t, validateID("fixture-id"))
}
