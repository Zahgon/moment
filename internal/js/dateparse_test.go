package js

import (
	"encoding/json"
	"math"
	"os"
	"testing"
	"time"
)

// TestParseDate checks new Date(string) against time values recorded from Node.js with TZ=UTC.
func TestParseDate(t *testing.T) {
	saved := time.Local
	time.Local = time.UTC
	defer func() { time.Local = saved }()
	data, err := os.ReadFile("testdata/dateparse.json")
	if err != nil {
		t.Fatal(err)
	}
	var items []struct {
		S string
		T *float64
	}
	if err := json.Unmarshal(data, &items); err != nil {
		t.Fatal(err)
	}
	for _, it := range items {
		got := ParseDate(it.S).Time()
		if it.T == nil {
			if !math.IsNaN(got) {
				t.Errorf("new Date(%q) = %v, want Invalid Date", it.S, got)
			}
			continue
		}
		if got != *it.T {
			t.Errorf("new Date(%q) = %v, want %v", it.S, got, *it.T)
		}
	}
}

// TestLocalTimeGap checks that a local time skipped by a transition resolves with the offset in force before it,
// as Node.js does: 02:30 on 2014-03-09 in New York is 03:30 EDT.
func TestLocalTimeGap(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Skip("no time-zone database")
	}
	saved := time.Local
	time.Local = loc
	defer func() { time.Local = saved }()
	d := NewLocalDate(2014, 2, 9, 2, 30)
	if got, want := d.Time(), 1394350200000.0; got != want {
		t.Errorf("new Date(2014, 2, 9, 2, 30) = %v, want %v", got, want)
	}
	if got := d.Hours(false); got != 3 {
		t.Errorf("getHours() = %v, want 3", got)
	}
}
