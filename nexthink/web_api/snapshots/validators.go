package snapshots

import (
	"fmt"
	"regexp"
	"strings"
)

var nqlName = regexp.MustCompile(`^#[a-z0-9_]+$`)

func validateDefinition(r *Definition) error {
	if r == nil || strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.NQLQuery) == "" || strings.TrimSpace(r.Origin) == "" {
		return fmt.Errorf("name, NQL query and origin are required")
	}
	if !nqlName.MatchString(r.NQLName) {
		return fmt.Errorf("nqlName must start with # and contain lowercase letters, digits or underscores")
	}
	if r.SnapshotVersion != 1 {
		return fmt.Errorf("the observed snapshot format version is 1")
	}
	return nil
}
