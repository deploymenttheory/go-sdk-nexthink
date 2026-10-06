package web_api

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Reconciliation must point to real resource methods and catalog gateways.
// Same-named documents may target different gateways, so validate every variant.
func TestDiscoveryGraphQLImplementationsExist(t *testing.T) {
	body, err := os.ReadFile("../../docs/web-api-discovery.json")
	require.NoError(t, err)
	var inventory struct {
		Documents []struct {
			Kind            string `json:"kind"`
			Name            string `json:"name"`
			Hash            string `json:"document_sha256"`
			Status          string `json:"status"`
			Reason          string `json:"reason"`
			Implementations []struct {
				Package     string `json:"package"`
				Method      string `json:"method"`
				Endpoint    string `json:"endpoint"`
				OperationID string `json:"operation_id"`
				Example     string `json:"example"`
			} `json:"implementations"`
		} `json:"graphql_documents"`
	}
	require.NoError(t, json.Unmarshal(body, &inventory))
	require.NotEmpty(t, inventory.Documents)
	seen := map[string]bool{}
	methods := map[string]map[string]bool{}
	for _, doc := range inventory.Documents {
		t.Run(doc.Kind+"/"+doc.Name+"/"+doc.Hash[:12], func(t *testing.T) {
			key := doc.Kind + ":" + doc.Name + ":" + doc.Hash
			require.False(t, seen[key], "duplicate document variant")
			seen[key] = true
			if doc.Status == "unsupported_live_schema" {
				require.NotEmpty(t, doc.Reason)
				require.Empty(t, doc.Implementations)
				return
			}
			require.Contains(t, []string{"implemented", "equivalent_selection"}, doc.Status)
			require.NotEmpty(t, doc.Implementations)
			for _, impl := range doc.Implementations {
				assert.True(t, discoveryServiceMethods(t, methods, impl.Package)[impl.Method], "%s.%s must exist", impl.Package, impl.Method)

				op, exists := OperationByID(impl.OperationID)
				require.True(t, exists, "catalog operation %s", impl.OperationID)
				assert.Equal(t, impl.Endpoint, op.Path)
				assert.Equal(t, "POST", op.Method)
				_, err = os.Stat(filepath.Join("../..", impl.Example))
				assert.NoError(t, err)
			}
		})
	}
}

func TestCatalogOperationIDsAreUnique(t *testing.T) {
	seen := map[string]bool{}
	for _, op := range Operations() {
		require.NotEmpty(t, op.ID)
		assert.False(t, seen[op.ID], "duplicate operation ID %s", op.ID)
		seen[op.ID] = true
		assert.NotEmpty(t, op.Path)
		assert.NotEmpty(t, op.Method)
	}
}

// Resource methods may live in extension files, not only crud.go.
func discoveryServiceMethods(t *testing.T, cache map[string]map[string]bool, pkg string) map[string]bool {
	t.Helper()
	if methods, ok := cache[pkg]; ok {
		return methods
	}
	paths, err := filepath.Glob(filepath.Join(pkg, "*.go"))
	require.NoError(t, err)
	require.NotEmpty(t, paths)
	methods := map[string]bool{}
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		source, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		require.NoError(t, err)
		for _, decl := range source.Decls {
			method, ok := decl.(*ast.FuncDecl)
			if !ok || method.Recv == nil || len(method.Recv.List) == 0 {
				continue
			}
			receiver := method.Recv.List[0].Type
			if pointer, ok := receiver.(*ast.StarExpr); ok {
				receiver = pointer.X
			}
			if name, ok := receiver.(*ast.Ident); ok && name.Name == "Service" {
				methods[method.Name.Name] = true
			}
		}
	}
	cache[pkg] = methods
	return methods
}

func TestDiscoveryHTTPContractsAndSourceAudit(t *testing.T) {
	body, err := os.ReadFile("../../docs/web-api-discovery.json")
	require.NoError(t, err)
	var inventory struct {
		Contracts []struct {
			Method      string `json:"method"`
			Path        string `json:"path"`
			Status      string `json:"status"`
			OperationID string `json:"operation_id"`
			Kind        string `json:"kind"`
			Round       string `json:"discovery_round"`
			Methods     []struct {
				Package string `json:"package"`
				Method  string `json:"method"`
				Example string `json:"example"`
			} `json:"sdk_methods"`
		} `json:"known_follow_up_http_contracts"`
		Literals []struct {
			Literal    string   `json:"literal"`
			Status     string   `json:"status"`
			Previous   string   `json:"previous_status"`
			Operations []string `json:"catalog_operation_ids"`
			Audit      *struct {
				Owner      string `json:"owner"`
				Resolution string `json:"resolution"`
				Scope      string `json:"scope_limit"`
			} `json:"source_audit"`
		} `json:"api_path_literals"`
		Reconciliation struct {
			HTTPContracts   int    `json:"follow_up_http_contracts"`
			HTTPImplemented int    `json:"follow_up_http_implemented"`
			Catalog         int    `json:"catalog_http_contracts_including_graphql_gateways"`
			UniqueRoutes    int    `json:"catalog_unique_http_method_path_pairs"`
			OriginalLeads   int    `json:"original_outstanding_discovery_leads"`
			AuditedOriginal int    `json:"original_outstanding_leads_source_audited"`
			Unresolved      int    `json:"unresolved_source_discovery_leads"`
			NewMethods      int    `json:"current_round_new_sdk_methods"`
			NewOptions      int    `json:"current_round_new_route_options"`
			Scope           string `json:"unknown_backend_completeness"`
		} `json:"reconciliation"`
	}
	require.NoError(t, json.Unmarshal(body, &inventory))
	require.NotEmpty(t, inventory.Contracts)
	methods := map[string]map[string]bool{}
	newMethods, newOptions := 0, 0
	for _, row := range inventory.Contracts {
		t.Run(row.Method+"/"+row.Path+"/"+row.Kind, func(t *testing.T) {
			assert.Equal(t, "implemented", row.Status)
			require.NotEmpty(t, row.Methods)
			op, ok := OperationByID(row.OperationID)
			require.True(t, ok, "catalog operation %s", row.OperationID)
			assert.Equal(t, row.Method, op.Method)
			assert.Equal(t, row.Path, op.Path)
			for _, impl := range row.Methods {
				assert.True(t, discoveryServiceMethods(t, methods, impl.Package)[impl.Method], "%s.%s must exist", impl.Package, impl.Method)
				_, err := os.Stat(filepath.Join("../..", impl.Example))
				assert.NoError(t, err)
			}
		})
		if row.Round == "remaining_leads" {
			switch row.Kind {
			case "api_method":
				newMethods++
			case "route_option":
				newOptions++
			default:
				t.Errorf("unknown contract kind %q", row.Kind)
			}
		}
	}
	original := 0
	for _, row := range inventory.Literals {
		assert.Contains(t, []string{"observed_contract_implemented", "audited_source_contracts_implemented"}, row.Status, row.Literal)
		require.NotEmpty(t, row.Operations, row.Literal)
		for _, id := range row.Operations {
			_, ok := OperationByID(id)
			assert.True(t, ok, "%s catalog mapping %s", row.Literal, id)
		}
		if row.Status == "audited_source_contracts_implemented" {
			require.NotNil(t, row.Audit, row.Literal)
			assert.NotEmpty(t, row.Audit.Owner)
			assert.NotEmpty(t, row.Audit.Resolution)
			assert.NotEmpty(t, row.Audit.Scope)
			if row.Previous != "" {
				original++
			}
		}
	}
	rec := inventory.Reconciliation
	assert.Len(t, inventory.Contracts, rec.HTTPContracts)
	assert.Equal(t, rec.HTTPContracts, rec.HTTPImplemented)
	assert.Len(t, Operations(), rec.Catalog)
	unique := map[string]bool{}
	for _, op := range Operations() {
		unique[op.Method+" "+op.Path] = true
	}
	assert.Len(t, unique, rec.UniqueRoutes)
	assert.Equal(t, rec.NewMethods, newMethods)
	assert.Equal(t, rec.NewOptions, newOptions)
	assert.Equal(t, 64, original, "all original 10 specific, 52 prefix and 2 proxy leads need source evidence")
	assert.Equal(t, rec.OriginalLeads, original)
	assert.Equal(t, rec.AuditedOriginal, original)
	assert.Zero(t, rec.Unresolved)
	assert.NotEmpty(t, rec.Scope)
}
