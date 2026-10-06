package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInventoryAndReport(t *testing.T) {
	root := t.TempDir()
	for _, family := range []string{"web_api", "public_api"} {
		if err := os.MkdirAll(filepath.Join(root, "nexthink", family, "sample"), 0700); err != nil {
			t.Fatal(err)
		}
	}
	source := []byte("package sample; type Service struct{}; func (*Service) Get() {}; func (*Service) hidden() {}; func Other() {}")
	if err := os.WriteFile(filepath.Join(root, "nexthink", "web_api", "sample", "crud.go"), source, 0600); err != nil {
		t.Fatal(err)
	}
	rows, err := inventory(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 1 || rows[0].key() != "web_api/sample/Get" {
		t.Fatalf("unexpected inventory: %+v", rows)
	}
	if validate(rows, nil) == nil {
		t.Fatal("accepted missing method")
	}
	pass := rows[0]
	pass.Status = "pass"
	pass.Reason = "curl matched"
	if validate(rows, []entry{pass}) == nil {
		t.Fatal("accepted unevidenced pass")
	}
	pass.Evidence = []string{"sanitized-run-reference"}
	if err := validate(rows, []entry{pass}); err != nil {
		t.Fatal(err)
	}
	if validate(rows, []entry{pass, pass}) == nil {
		t.Fatal("accepted duplicate")
	}
	blocked := rows[0]
	blocked.Status = "blocked"
	blocked.Reason = "requires populated fixture"
	if validate(rows, []entry{blocked}) == nil {
		t.Fatal("accepted uncategorized blocker")
	}
	blocked.Blocker = "fixture_required"
	if err := validate(rows, []entry{blocked}); err != nil {
		t.Fatal(err)
	}
	invalid := blocked
	invalid.Blocker = "mystery"
	if validate(rows, []entry{invalid}) == nil {
		t.Fatal("accepted unknown blocker category")
	}
	unknown := pass
	unknown.Method = "Unknown"
	if validate(rows, []entry{unknown}) == nil {
		t.Fatal("accepted unknown method")
	}
}
