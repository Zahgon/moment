package locale

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// TestMonthsParseSettings is the original's months-parse lint: every locale sets all three of monthsParse,
// longMonthsParse and shortMonthsParse, or none of them.
func TestMonthsParseSettings(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	setting := regexp.MustCompile(`(?i)monthsParse",`)
	counts := map[int][]string{}
	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		src, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		n := len(setting.FindAll(src, -1))
		counts[n] = append(counts[n], file)
	}
	if len(counts) != 2 || counts[0] == nil || counts[3] == nil {
		var partial []string
		for n, names := range counts {
			if n != 0 && n != 3 {
				partial = append(partial, names...)
			}
		}
		sort.Strings(partial)
		t.Fatalf("Months parse issue: each locale must have all three monthsParse settings or none; got %d "+
			"kinds of locale, partial in %v", len(counts), partial)
	}
}
