// moment.js locale configuration
// locale : Frisian [fy]
// author : Robin van der Vliet : https://github.com/robin0van0der0v

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fy", func() {
		moment.DefineLocale("fy", js.Obj(
			"months", []string{"jannewaris", "febrewaris", "maart", "april", "maaie", "juny", "july", "augustus", "septimber", "oktober", "novimber", "desimber"},
			"monthsShort", fyMonthsShort,
			"monthsParseExact", true,
			"weekdays", []string{"snein", "moandei", "tiisdei", "woansdei", "tongersdei", "freed", "sneon"},
			"weekdaysShort", []string{"si.", "mo.", "ti.", "wo.", "to.", "fr.", "so."},
			"weekdaysMin", []string{"Si", "Mo", "Ti", "Wo", "To", "Fr", "So"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD-MM-YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[hjoed om] LT",
				"nextDay", "[moarn om] LT",
				"nextWeek", "dddd [om] LT",
				"lastDay", "[juster om] LT",
				"lastWeek", "[ôfrûne] dddd [om] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "oer %s",
				"past", "%s lyn",
				"s", "in pear sekonden",
				"ss", "%d sekonden",
				"m", "ien minút",
				"mm", "%d minuten",
				"h", "ien oere",
				"hh", "%d oeren",
				"d", "ien dei",
				"dd", "%d dagen",
				"M", "ien moanne",
				"MM", "%d moannen",
				"y", "ien jier",
				"yy", "%d jierren",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(ste|de)", ""),
			"ordinal", fyOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	fyMonthsShortWithDots    = []string{"jan.", "feb.", "mrt.", "apr.", "mai", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "des."}
	fyMonthsShortWithoutDots = []string{"jan", "feb", "mrt", "apr", "mai", "jun", "jul", "aug", "sep", "okt", "nov", "des"}
)

// fyMonthsShort drops the abbreviation dots in a -MMM- format.
func fyMonthsShort(_ *moment.Locale, m *moment.Moment, format any) any {
	return pickDottedMonth(m, format, fyMonthsShortWithDots, fyMonthsShortWithoutDots)
}

// fyOrdinal is the ordinal callback of fy.
func fyOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	if sameNumber(number, 1) || sameNumber(number, 8) || js.ToNumber(number) >= 20 {
		return str(number) + "ste"
	}
	return str(number) + "de"
}
