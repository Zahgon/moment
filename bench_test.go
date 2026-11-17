package moment_test

import (
	"testing"
	"time"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	_ "github.com/moment/moment/locale"
)

func mustNew(b *testing.B, args ...moment.Value) *moment.Moment {
	b.Helper()
	m, err := moment.New(args...)
	if err != nil {
		b.Fatal(err)
	}
	return m
}

// The benchmarks below mirror the files of the original benchmarks/ directory: each runs the same operation, with
// one sub-benchmark per unit where the original defines one test per unit.

var (
	addUnits      = []string{"milliseconds", "seconds", "minutes", "hours", "days", "weeks", "months", "quarters", "years"}
	boundaryUnits = []string{"second", "minute", "hour", "date", "day", "isoWeek", "week", "month", "quarter", "year"}
)

func BenchmarkAdd(b *testing.B) {
	for _, unit := range addUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				m.Add(8.0, unit)
			}
		})
	}
}

func BenchmarkSubtract(b *testing.B) {
	for _, unit := range addUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				m.Subtract(8.0, unit)
			}
		})
	}
}

func BenchmarkCompare(b *testing.B) {
	now := float64(time.Now().UnixMilli())
	const diff = -24 * 60 * 60 * 1e3
	this, input := mustNew(b, now), mustNew(b, now+diff)
	from, to := mustNew(b, now+diff), mustNew(b, now+2*diff)
	tests := []struct {
		name string
		fn   func() bool
	}{
		{"isAfter", func() bool { return this.IsAfter(input, "date") }},
		{"isBefore", func() bool { return this.IsBefore(input, "date") }},
		{"isSame", func() bool { return this.IsSame(input, "date") }},
		{"isSameOrAfter", func() bool { return this.IsSameOrAfter(input, "date") }},
		{"isSameOrBefore", func() bool { return this.IsSameOrBefore(input, "date") }},
		{"isBetween", func() bool { return this.IsBetween(from, to, "date", nil) }},
	}
	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.fn()
			}
		})
	}
}

func BenchmarkQuery(b *testing.B) {
	m := mustNew(b, "2013-05-25")
	tests := []struct {
		name string
		fn   func() bool
	}{
		{"isBefore_true", func() bool { return m.IsBefore("2013-06-25", nil) }},
		{"isBefore_self", func() bool { return m.IsBefore("2013-05-25", nil) }},
		{"isBefore_false", func() bool { return m.IsBefore("2013-04-25", nil) }},
		{"isAfter_true", func() bool { return m.IsAfter("2013-04-25", nil) }},
		{"isAfter_self", func() bool { return m.IsAfter("2013-05-25", nil) }},
		{"isAfter_false", func() bool { return m.IsAfter("2013-06-25", nil) }},
	}
	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.fn()
			}
		})
	}
}

func BenchmarkEndOf(b *testing.B) {
	for _, unit := range boundaryUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				m.EndOf(unit)
			}
		})
	}
}

func BenchmarkGet(b *testing.B) {
	for _, unit := range boundaryUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				if _, err := m.Get(unit); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	for _, unit := range boundaryUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				m.Set(unit, 2.0)
			}
		})
	}
}

func BenchmarkFromDate(b *testing.B) {
	base := js.NewDate(float64(time.Now().UnixMilli()))
	for i := 0; i < b.N; i++ {
		mustNew(b, base)
	}
}

func BenchmarkFromDateUTC(b *testing.B) {
	base := js.NewDate(float64(time.Now().UnixMilli()))
	for i := 0; i < b.N; i++ {
		if _, err := moment.CreateUTC(base); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMakeDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moment.CreateDuration(5.0, "years")
	}
}

// BenchmarkLoadMissingLocale selects a locale that does not exist, which falls back to the current one.
func BenchmarkLoadMissingLocale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		moment.GetSetGlobalLocale("en-US", nil)
	}
}

func BenchmarkParseCompactWeekday(b *testing.B) {
	tests := []struct {
		name   string
		input  string
		strict bool
	}{
		{"compact strict", "21530", true},
		{"compact non-strict", "21530", false},
		{"padded strict", "021530", true},
		{"padded non-strict", "021530", false},
	}
	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				mustNew(b, tc.input, "eHHmm", tc.strict)
			}
		})
	}
}

func BenchmarkClone(b *testing.B) {
	m := mustNew(b, 1700000000000.0)
	for i := 0; i < b.N; i++ {
		m.Clone()
	}
}

func BenchmarkFormat(b *testing.B) {
	m := mustNew(b, 1700000000000.0)
	for i := 0; i < b.N; i++ {
		m.Format("dddd, MMMM Do YYYY, h:mm:ss a")
	}
}

func BenchmarkFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := moment.New("2013-02-08T09:30:26.123+07:00"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFromStringAndFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := moment.New("2013-02-08 09:30", "YYYY-MM-DD HH:mm"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStartOf(b *testing.B) {
	for _, unit := range boundaryUnits {
		b.Run(unit, func(b *testing.B) {
			m := mustNew(b, "2013-05-25")
			for i := 0; i < b.N; i++ {
				m.StartOf(unit)
			}
		})
	}
}

func BenchmarkHumanize(b *testing.B) {
	d := moment.CreateDuration(123456789.0, nil)
	for i := 0; i < b.N; i++ {
		d.Humanize(true, nil)
	}
}

func BenchmarkLocaleFormat(b *testing.B) {
	m := mustNew(b, 1700000000000.0).SetLocale("ru")
	for i := 0; i < b.N; i++ {
		m.Format("LLLL")
	}
}

func BenchmarkZeroFillFormat(b *testing.B) {
	m := mustNew(b, 1700000000000.0)
	for i := 0; i < b.N; i++ {
		m.Format("YYYYYY-MM-DD HH:mm:ss.SSSSSS")
	}
}
