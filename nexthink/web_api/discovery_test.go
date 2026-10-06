package web_api

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
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
				source, err := parser.ParseFile(token.NewFileSet(), filepath.Join(impl.Package, "crud.go"), nil, 0)
				require.NoError(t, err)
				found := false
				for _, decl := range source.Decls {
					if method, ok := decl.(*ast.FuncDecl); ok && method.Recv != nil && method.Name.Name == impl.Method {
						found = true
					}
				}
				assert.True(t, found, "%s.%s must exist", impl.Package, impl.Method)
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
