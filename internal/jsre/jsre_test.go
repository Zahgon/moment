package jsre

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type fixture struct {
	Source, Flags, Input, Src string
	Err                       bool
	Exec                      []any
	Match                     []*string
	Rep                       string
}

// TestRecordedMatches checks exec, String#match, String#replace and the source property against results
// recorded from Node.js for patterns of moment and its locales.
func TestRecordedMatches(t *testing.T) {
	data, err := os.ReadFile("testdata/cases.json")
	if err != nil {
		t.Fatal(err)
	}
	var items []fixture
	if err := json.Unmarshal(data, &items); err != nil {
		t.Fatal(err)
	}
	for _, it := range items {
		re, err := Compile(it.Source, it.Flags)
		if it.Err != (err != nil) {
			t.Errorf("Compile(%q, %q) error = %v, want error %v", it.Source, it.Flags, err, it.Err)
			continue
		}
		if err != nil {
			continue
		}
		var exec []any
		if m := re.Exec(it.Input); m != nil {
			exec = append(exec, float64(m.Index))
			for i, g := range m.Groups {
				if m.Defined[i] {
					exec = append(exec, g)
				} else {
					exec = append(exec, nil)
				}
			}
		}
		if !reflect.DeepEqual(exec, it.Exec) {
			t.Errorf("/%s/%s.exec(%q) = %v, want %v", it.Source, it.Flags, it.Input, exec, it.Exec)
		}
		re.LastIndex = 0
		var match []*string
		if m := re.Match(it.Input); m != nil {
			for i := range m.Groups {
				g := m.Groups[i]
				if m.Defined[i] {
					match = append(match, &g)
				} else {
					match = append(match, nil)
				}
			}
		}
		if !reflect.DeepEqual(match, it.Match) {
			t.Errorf("%q.match(/%s/%s) differs", it.Input, it.Source, it.Flags)
		}
		if got := re.Replace(it.Input, "<$&|$1|$`>"); got != it.Rep {
			t.Errorf("%q.replace(/%s/%s) = %q, want %q", it.Input, it.Source, it.Flags, got, it.Rep)
		}
		if got := re.Source(); got != it.Src {
			t.Errorf("source of /%s/ = %q, want %q", it.Source, got, it.Src)
		}
	}
}
