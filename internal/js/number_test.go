package js

import (
	"encoding/json"
	"math"
	"os"
	"testing"
)

type numberFixture struct {
	Numbers []struct {
		X, S, Fixed, Round string
	}
	Strings []struct {
		S, Num, Int, Float string
	}
}

func loadNumbers(t *testing.T) numberFixture {
	t.Helper()
	data, err := os.ReadFile("testdata/numbers.json")
	if err != nil {
		t.Fatal(err)
	}
	var f numberFixture
	if err := json.Unmarshal(data, &f); err != nil {
		t.Fatal(err)
	}
	return f
}

// numberString prints a number as the fixture does: String(x), with -0 spelled out.
func numberString(x float64) string {
	if x == 0 && math.Signbit(x) {
		return "-0"
	}
	return NumberToString(x)
}

// TestNumberFormatting checks Number#toString, toFixed(3) and Math.round against values recorded from Node.js.
func TestNumberFormatting(t *testing.T) {
	for _, c := range loadNumbers(t).Numbers {
		x := StringToNumber(c.S)
		if c.X == "-0" {
			x = math.Copysign(0, -1)
		}
		if got := NumberToString(x); got != c.S {
			t.Errorf("String(%s) = %q, want %q", c.X, got, c.S)
		}
		fixed := NumberToString(x)
		if !math.IsInf(x, 0) && !math.IsNaN(x) && math.Abs(x) < 1e21 {
			fixed = ToFixed(x, 3)
		}
		if fixed != c.Fixed {
			t.Errorf("(%s).toFixed(3) = %q, want %q", c.X, fixed, c.Fixed)
		}
		if got := numberString(Round(x)); got != c.Round {
			t.Errorf("Math.round(%s) = %q, want %q", c.X, got, c.Round)
		}
	}
}

// TestStringConversion checks Number(s), parseInt(s, 10) and parseFloat(s) against values recorded from Node.js.
func TestStringConversion(t *testing.T) {
	for _, c := range loadNumbers(t).Strings {
		if got := NumberToString(StringToNumber(c.S)); got != c.Num {
			t.Errorf("Number(%q) = %s, want %s", c.S, got, c.Num)
		}
		if got := NumberToString(ParseInt(c.S)); got != c.Int {
			t.Errorf("parseInt(%q) = %s, want %s", c.S, got, c.Int)
		}
		if got := NumberToString(ParseFloat(c.S)); got != c.Float {
			t.Errorf("parseFloat(%q) = %s, want %s", c.S, got, c.Float)
		}
	}
}
