package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// qunitFile is one file of the original QUnit suites (src/test/moment/*.js or src/test/locale/*.js) with each of
// its tests traced into a case program: the test's calls of the moment namespace and of moment, duration and locale
// objects, in order, and the lines the original printed when it ran that program.
type qunitFile struct {
	Source string      `json:"source"`
	Tests  []qunitTest `json:"tests"`
}

// qunitTest is one QUnit test of a qunitFile.
type qunitTest struct {
	Module string          `json:"module"`
	Name   string          `json:"name"`
	TZ     string          `json:"tz"`
	Case   json.RawMessage `json:"case"`
	Expect []string        `json:"expect"`
}

// readQUnitFile reads a gzip-compressed qunitFile; reading the whole stream verifies its checksum.
func readQUnitFile(path string) (*qunitFile, error) {
	compressed, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if err := zr.Close(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	var q qunitFile
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &q, nil
}

// TestQUnit replays the original QUnit suites: one subtest per QUnit module and test, named as in the original,
// whose traced program must print what the original printed.
func TestQUnit(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("testdata", "qunit", "*", "*.json.gz"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no traced QUnit tests: %v", err)
	}
	sort.Strings(files)
	for _, file := range files {
		q, err := readQUnitFile(file)
		if err != nil {
			t.Fatal(err)
		}
		for _, test := range q.Tests {
			t.Run(strings.ReplaceAll(test.Module, "/", "_")+"/"+test.Name, func(t *testing.T) {
				t.Parallel()
				compareRun(t, test.TZ, test.Case, test.Expect)
			})
		}
	}
}
