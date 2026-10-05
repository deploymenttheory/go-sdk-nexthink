package nexthink_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/data_management"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/nql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/spark"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/public_api/workflows"
	"go.uber.org/zap"
)

func TestNewPublicServiceWireContracts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Header.Get("Authorization") != "Bearer test" {
			t.Errorf("unexpected request method or authentication")
		}
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case data_management.EndpointDeviceDeletions:
			devices := body["devices"].([]any)
			if len(devices) != 1 || devices[0].(map[string]any)["uid"] != "malformed" {
				t.Error("lost per-device validation input")
			}
			w.WriteHeader(202)
			w.Write([]byte(`{"status":"ACCEPTED","scheduledCount":0,"devices":[{"uid":"malformed","name":"lab","status":"INVALID"}]}`))
		case spark.EndpointHandoff:
			if r.Header.Get("User-Principal-Name") != "lab@example.test" || r.Header.Get("Timezone") != "Europe/London" {
				t.Error("missing Spark headers")
			}
			part := body["message"].(map[string]any)["parts"].([]any)[0].(map[string]any)
			if part["type"] != "TEXT" || part["text"] != "test" || len(part) != 2 {
				t.Error("incorrect message union encoding")
			}
			w.WriteHeader(204)
		case "/api/v2/nql/execute":
			if len(body) != 2 || body["parameters"].(map[string]any)["platform"] != "macOS" || r.Header.Get("Accept") != "application/json" {
				t.Error("incorrect NQL execute contract")
			}
			w.Write([]byte(`{"queryId":"#lab_query","rows":0,"data":[]}`))
		case "/api/v1/nql/export":
			if body["format"] != nil || body["platform"] != nil || body["compression"] != "GZIP" {
				t.Error("incorrect NQL export contract")
			}
			w.Write([]byte(`{"exportId":"opaque=="}`))
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
		}
	}))
	defer server.Close()
	c, err := client.NewTransportWithTokenProvider(server.URL, auth.StaticToken("test", time.Time{}), client.WithLogger(zap.NewNop()))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	d, resp, err := data_management.NewService(c).DeleteDevices(ctx, &data_management.DeleteDevicesRequest{Devices: []data_management.Device{{UID: "malformed", Name: "lab"}}}, "")
	if err != nil || resp.StatusCode != 202 || d.ScheduledCount != 0 || d.Devices[0].Status != "INVALID" {
		t.Fatalf("deletion scheduling contract: %v", err)
	}
	resp, err = spark.NewService(c).Handoff(ctx, "lab@example.test", "Europe/London", &spark.HandoffRequest{Message: spark.Message{Parts: []spark.Part{{Type: "TEXT", Text: "test"}}}})
	if err != nil || resp.StatusCode != 204 {
		t.Fatalf("empty handoff response: %v", err)
	}
	n := nql.NewService(c)
	if _, _, err = n.ExecuteNQLV2(ctx, &nql.ExecuteRequest{QueryID: "#lab_query", Parameters: map[string]string{"platform": "macOS"}}); err != nil {
		t.Fatal(err)
	}
	if _, _, err = n.StartNQLExport(ctx, &nql.ExportRequest{QueryID: "#lab_query", Format: "json", Compression: "GZIP"}); err != nil {
		t.Fatal(err)
	}
}

func TestWorkflowDetailsUsesLiveNQLIDContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" || r.URL.Path != "/api/v1/workflows/details" || r.URL.Query().Get("nqlId") != "lab_workflow" || r.URL.Query().Has("nql-id") {
			t.Error("incorrect workflow detail request")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"lab_workflow","triggerMethods":{"apiEnabled":true,"manualEnabled":false},"versions":[{"version":1,"status":"ACTIVE","parameters":[{"id":"mode","allowCustomValue":true,"options":["test"]}],"valid":true,"hasDevice":true,"hasUser":false}]}`))
	}))
	defer server.Close()
	c, err := client.NewTransportWithTokenProvider(server.URL, auth.StaticToken("test", time.Time{}), client.WithLogger(zap.NewNop()))
	if err != nil {
		t.Fatal(err)
	}
	result, _, err := workflows.NewService(c).GetWorkflowDetails(context.Background(), "lab_workflow")
	if err != nil {
		t.Fatal(err)
	}
	if !result.TriggerMethods.APIEnabled || result.Versions[0].Parameters[0].ID != "mode" {
		t.Fatal("lost live workflow fields")
	}
}
