package nexthink

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/auth"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/client"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/graphql"
	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/web_api/nql_queries"
)

func TestCatalogWireContracts(t *testing.T) {
	seen := map[string]bool{}
	for _, op := range web_api.Operations() {
		if seen[op.ID] || op.Evidence == "" || op.Authentication == "" {
			t.Fatalf("invalid catalog entry %s", op.ID)
		}
		seen[op.ID] = true
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer user-token" {
			t.Error("missing user token")
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/apigateway/api/v1/product-shell/menu":
			if r.Method != "POST" {
				t.Error("menu requires POST")
			}
			w.Write([]byte(`{"result":[],"unknownFutureField":true}`))
		case "/apigateway/nqlapi/store/api/v1":
			var req map[string]any
			json.NewDecoder(r.Body).Decode(&req)
			if req["nql"] != "devices | limit 1" || req["nqlQuery"] != nil {
				t.Error("wrong query write representation")
			}
			w.WriteHeader(201)
			w.Write([]byte(`{"contentId":"created","nqlQuery":"devices | limit 1"}`))
		case "/apigateway/workflows/manage/graphql":
			w.Write([]byte(`{"data":{"partial":true},"errors":[{"message":"fixture failure"}]}`))
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
		}
	}))
	defer server.Close()
	c, err := NewClient(
		&AuthConfig{
			Instance: "test",
			Region:   "eu",
			WebAPI: &BrowserCredentials{
				TokenProvider: auth.StaticToken("user-token", time.Time{}),
			},
		},
		WithWebAPIOptions(client.WithBaseURL(server.URL)),
		WithLogger(zap.NewNop()),
	)
	if err != nil {
		t.Fatal(err)
	}
	raw, _, err := c.WebAPI.Do(
		context.Background(),
		"shell.menu",
		web_api.Request{Body: map[string]any{}},
	)
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]any
	json.Unmarshal(raw, &body)
	if body["unknownFutureField"] != true {
		t.Error("raw API lost unknown field")
	}
	if _, _, err := c.WebAPI.NQLQueries.Create(
		context.Background(),
		&nql_queries.SaveQueryRequest{
			NQLAPIID: "#lab_query",
			Name:     "lab",
			NQL:      "devices | limit 1",
		},
	); err != nil {
		t.Fatal(err)
	}
	result, resp, err := c.WebAPI.GraphQL.Execute(
		context.Background(),
		"graphql.workflows",
		graphql.GraphQLRequest{Query: "{ partial }"},
	)
	var gqlErrors graphql.GraphQLErrors
	if !errors.As(err, &gqlErrors) || resp.StatusCode != 200 ||
		string(result.Data) != `{"partial":true}` {
		t.Fatal("lost GraphQL errors or partial data")
	}
	for _, id := range []string{"missing-operation", "https://other.example"} {
		if _, _, err := c.WebAPI.Do(context.Background(), id, web_api.Request{}); err == nil {
			t.Error("unknown operation accepted")
		}
	}
	if _, _, err := c.WebAPI.Do(
		context.Background(),
		"nql_queries.get",
		web_api.Request{},
	); err == nil {
		t.Error("missing path parameter accepted")
	}
	for _, value := range []string{".", "..", "../other", "a/b", `a\b`} {
		if _, _, err := c.WebAPI.Do(
			context.Background(),
			"nql_queries.get",
			web_api.Request{PathParams: map[string]string{"contentId": value}},
		); err == nil {
			t.Errorf("path segment escape accepted: %q", value)
		}
	}
}
