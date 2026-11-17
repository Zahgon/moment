package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var binary string

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

// testMain builds momentcase into a temporary directory, runs the tests against it and removes the directory; a
// failure to build or to clean up fails the run.
func testMain(m *testing.M) (code int) {
	dir, err := os.MkdirTemp("", "momentcase-test")
	if err != nil {
		log.Printf("creating the build directory: %v", err)
		return 1
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			log.Printf("removing the build directory: %v", err)
			code = 1
		}
	}()
	binary = filepath.Join(dir, "momentcase")
	build := exec.Command("go", "build", "-o", binary, ".")
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		log.Printf("building momentcase: %v", err)
		return 1
	}
	return m.Run()
}

// recordedCase is a case with the lines the original Moment.js printed for it.
type recordedCase struct {
	ID     string          `json:"id"`
	TZ     string          `json:"tz"`
	Case   json.RawMessage `json:"case"`
	Expect []string        `json:"expect"`
}

// decodeLine parses an output line keeping object key order, so that key order is compared too.
func decodeLine(t *testing.T, line string) any {
	t.Helper()
	v, err := readJSON(strings.NewReader(line))
	if err != nil {
		t.Fatalf("bad output line %q: %v", line, err)
	}
	return normalize(v)
}

func normalize(v any) any {
	switch x := v.(type) {
	case []any:
		out := make([]any, len(x))
		for i, e := range x {
			out[i] = normalize(e)
		}
		return out
	case interface{ Keys() []string }:
		type pair struct {
			Key   string
			Value any
		}
		obj := x.(interface {
			Keys() []string
			Get(string) any
		})
		var pairs []pair
		for _, k := range obj.Keys() {
			pairs = append(pairs, pair{k, normalize(obj.Get(k))})
		}
		return pairs
	}
	return v
}

// TestRecordedCases replays the recorded cases of every group in testdata and compares each output line with
// the line Moment.js printed.
func TestRecordedCases(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("testdata", "*.json"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no recorded cases: %v", err)
	}
	sort.Strings(files)
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		var cases []recordedCase
		if err := json.Unmarshal(data, &cases); err != nil {
			t.Fatalf("%s: %v", file, err)
		}
		group := strings.TrimSuffix(filepath.Base(file), ".json")
		for _, c := range cases {
			t.Run(group+"/"+c.ID, func(t *testing.T) {
				compareRun(t, c.TZ, c.Case, c.Expect)
			})
		}
	}
}

// compareRun runs momentcase on one case in the time zone tz and compares every output line with the line the
// original Moment.js printed.
func compareRun(t *testing.T, tz string, body []byte, expect []string) {
	t.Helper()
	cmd := exec.Command(binary)
	cmd.Stdin = bytes.NewReader(body)
	cmd.Env = append(os.Environ(), "TZ="+tz, "LANG=C.UTF-8")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("momentcase failed: %v", err)
	}
	got := strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")
	if len(expect) == 0 && len(got) == 1 && got[0] == "" {
		return
	}
	if len(got) != len(expect) {
		t.Fatalf("got %d lines, want %d:\n got: %q\nwant: %q", len(got), len(expect), got, expect)
	}
	for i := range got {
		if !reflect.DeepEqual(decodeLine(t, got[i]), decodeLine(t, expect[i])) {
			t.Errorf("line %d:\n got: %s\nwant: %s", i, got[i], expect[i])
		}
	}
}
