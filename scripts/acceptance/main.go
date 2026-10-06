// Command acceptance inventories exported resource methods and checks that a
// sanitized acceptance report accounts for every method exactly once.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type entry struct {
	Family   string   `json:"family"`
	Resource string   `json:"resource"`
	Method   string   `json:"method"`
	Status   string   `json:"status,omitempty"`
	Reason   string   `json:"reason,omitempty"`
	Evidence []string `json:"evidence,omitempty"`
}

func (e entry) key() string { return e.Family + "/" + e.Resource + "/" + e.Method }
func inventory(root string) ([]entry, error) {
	var rows []entry
	for _, family := range []string{"public_api", "web_api"} {
		base := filepath.Join(root, "nexthink", family)
		err := filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if path != base && strings.Count(strings.TrimPrefix(path, base+string(os.PathSeparator)), string(os.PathSeparator)) > 0 {
					return filepath.SkipDir
				}
				return nil
			}
			if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
			if err != nil {
				return err
			}
			for _, decl := range f.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok || fn.Recv == nil || !fn.Name.IsExported() {
					continue
				}
				t := fn.Recv.List[0].Type
				if ptr, ok := t.(*ast.StarExpr); ok {
					t = ptr.X
				}
				name, ok := t.(*ast.Ident)
				if !ok || name.Name != "Service" {
					continue
				}
				rows = append(rows, entry{Family: family, Resource: filepath.Base(filepath.Dir(path)), Method: fn.Name.Name})
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].key() < rows[j].key() })
	return rows, nil
}
func validate(want, got []entry) error {
	expected := map[string]bool{}
	for _, e := range want {
		expected[e.key()] = true
	}
	seen := map[string]bool{}
	for _, e := range got {
		k := e.key()
		if !expected[k] {
			return fmt.Errorf("unknown method %s", k)
		}
		if seen[k] {
			return fmt.Errorf("duplicate method %s", k)
		}
		seen[k] = true
		switch e.Status {
		case "pass", "failed", "blocked":
		default:
			return fmt.Errorf("invalid status for %s", k)
		}
		if strings.TrimSpace(e.Reason) == "" {
			return fmt.Errorf("missing reason for %s", k)
		}
		if e.Status == "pass" && len(e.Evidence) == 0 {
			return fmt.Errorf("pass without evidence for %s", k)
		}
	}
	var missing []string
	for k := range expected {
		if !seen[k] {
			missing = append(missing, k)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		return fmt.Errorf("%d unaccounted methods: %s", len(missing), strings.Join(missing, ", "))
	}
	return nil
}
func run() error {
	root := flag.String("root", ".", "repository root")
	report := flag.String("report", "", "validate a sanitized JSON array against current source")
	flag.Parse()
	rows, err := inventory(*root)
	if err != nil {
		return err
	}
	if *report != "" {
		data, err := os.ReadFile(*report)
		if err != nil {
			return err
		}
		var results []entry
		if err = json.Unmarshal(data, &results); err != nil {
			return err
		}
		if err = validate(rows, results); err != nil {
			return err
		}
		counts := map[string]int{}
		for _, r := range results {
			counts[r.Status]++
		}
		return json.NewEncoder(os.Stdout).Encode(map[string]any{"methods": len(rows), "results": counts})
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(rows)
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
