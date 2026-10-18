package main

import (
	"os"
	"regexp"
	"sort"
	"testing"
)

// TestMembersMatchTheManifest: what is compiled in, at which version, is exactly what component.yaml declares, and the
// shell's own version is the manifest's (the pitfall "the shell's member list, the compiled members and the lock
// disagree").
func TestMembersMatchTheManifest(t *testing.T) {
	raw, err := os.ReadFile("component.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var declared []string
	for _, m := range regexp.MustCompile(`(?m)^    - (\S+@\S+)$`).FindAllStringSubmatch(string(raw), -1) {
		declared = append(declared, m[1])
	}
	var compiled []string
	for _, d := range members() {
		compiled = append(compiled, d.ID+"@"+d.Version)
	}
	sort.Strings(declared)
	sort.Strings(compiled)
	if len(declared) == 0 || len(declared) != len(compiled) {
		t.Fatalf("declared %v, compiled %v", declared, compiled)
	}
	for i := range declared {
		if declared[i] != compiled[i] {
			t.Errorf("declared %s, compiled %s", declared[i], compiled[i])
		}
	}
	if !regexp.MustCompile(`(?m)^  version: ` + regexp.QuoteMeta(Version) + `$`).Match(raw) {
		t.Errorf("component.yaml does not say version %s", Version)
	}
}
