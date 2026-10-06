package web_api

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Every resource method must have a runnable example that actually calls it.
// This guards against adding API methods without extending the example suite.
func TestEveryResourceMethodHasExample(t *testing.T) {
	files, err := filepath.Glob("*/*.go")
	publicFiles, publicErr := filepath.Glob("../public_api/*/*.go")
	require.NoError(t, publicErr)
	files = append(files, publicFiles...)
	require.NoError(t, err)
	require.NotEmpty(t, files)
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		source, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
		require.NoError(t, err)
		resource := filepath.Base(filepath.Dir(file))
		family := "web_api"
		if filepath.Dir(filepath.Dir(file)) == "../public_api" {
			family = "public_api"
		}
		for _, declaration := range source.Decls {
			method, ok := declaration.(*ast.FuncDecl)
			if !ok || method.Recv == nil || !method.Name.IsExported() {
				continue
			}
			receiver := method.Recv.List[0].Type
			if pointer, ok := receiver.(*ast.StarExpr); ok {
				receiver = pointer.X
			}
			name, ok := receiver.(*ast.Ident)
			if !ok || name.Name != "Service" {
				continue
			}
			t.Run(family+"/"+resource+"/"+method.Name.Name, func(t *testing.T) {
				path := filepath.Join("..", "..", "examples", "nexthink", family, resource, method.Name.Name, "main.go")
				example, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
				require.NoError(t, err, "add a runnable example for every resource method")
				called := false
				ast.Inspect(example, func(node ast.Node) bool {
					call, ok := node.(*ast.CallExpr)
					if !ok {
						return true
					}
					selector, ok := call.Fun.(*ast.SelectorExpr)
					if ok && selector.Sel.Name == method.Name.Name {
						called = true
					}
					return true
				})
				require.True(t, called, "example must call the SDK method")
			})
		}
	}
}
